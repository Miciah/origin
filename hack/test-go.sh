#!/bin/bash
#
# This script runs Go language unit tests for the repository. Arguments to this script
# are parsed as a list of packages to test until the first argument starting with '-' or '--' is
# found. That argument and all following arguments are interpreted as flags to be passed directly
# to `go test`. If no arguments are given, then "all" packages are tested.
#
# Coverage reports and jUnit XML reports can be generated by this script as well, but both cannot
# be generated at once.
#
# This script consumes the following parameters as environment variables:
#  - DRY_RUN:             prints all packages that would be tested with the args that would be used and exits
#  - TEST_KUBE:           toggles testing of non-essential Kubernetes unit tests
#  - TIMEOUT:             the timeout for any one unit test (default '60s')
#  - DETECT_RACES:        toggles the 'go test' race detector (defaults '-race')
#  - COVERAGE_OUTPUT_DIR: locates the directory in which coverage output files will be placed
#  - COVERAGE_SPEC:       a set of flags for 'go test' that specify the coverage behavior (default '-cover -covermode=atomic')
#  - GOTEST_FLAGS:        any other flags to be sent to 'go test'
#  - JUNIT_REPORT:        toggles the creation of jUnit XML from the test output and changes this script's output behavior
#                         to use the 'junitreport' tool for summarizing the tests.
#  - DLV_DEBUG            toggles running tests using delve debugger
function cleanup() {
    return_code=$?

    os::util::describe_return_code "${return_code}"
    exit "${return_code}"
}
trap "cleanup" EXIT

source "$(dirname "${BASH_SOURCE}")/lib/init.sh"
os::build::setup_env
os::cleanup::tmpdir

# Internalize environment variables we consume and default if they're not set
dry_run="${DRY_RUN:-}"
test_timeout="${TIMEOUT:-120s}"
detect_races="${DETECT_RACES:-true}"
coverage_output_dir="${COVERAGE_OUTPUT_DIR:-}"
coverage_spec="${COVERAGE_SPEC:--cover -covermode atomic}"
gotest_flags="${GOTEST_FLAGS:-}"
junit_report="${JUNIT_REPORT:-}"
dlv_debug="${DLV_DEBUG:-}"

if [[ -n "${junit_report}" && -n "${coverage_output_dir}" ]]; then
    echo "$0 cannot create jUnit XML reports and coverage reports at the same time."
    exit 1
fi

# determine if user wanted verbosity
verbose=
if [[ "${gotest_flags}" =~ -v( |$) ]]; then
    verbose=true
fi

# Build arguments for 'go test'
if [[ -z "${verbose}" && -n "${junit_report}" ]]; then
    # verbosity can be set explicitly by the user or set implicitly by asking for the jUnit
    # XML report, so we only want to add the flag if it hasn't been added by a user already
    # and is being implicitly set by jUnit report generation
    gotest_flags+=" -v"
fi

if [[ "${detect_races}" == "true" ]]; then
    gotest_flags+=" -race"
fi

# check to see if user has not disabled coverage mode
if [[ -n "${coverage_spec}" ]]; then
    # if we have a coverage spec set, we add it. '-race' implies '-cover -covermode atomic'
    # but specifying both at the same time does not lead to an error so we can add both specs
    gotest_flags+=" ${coverage_spec}"
fi

# check to see if user has not disabled test timeouts
if [[ -n "${test_timeout}" ]]; then
    gotest_flags+=" -timeout ${test_timeout}"
fi


# Break up the positional arguments into packages that need to be tested and arguments that need to be passed to `go test`
package_args=
for arg in "$@"; do
    if [[ "${arg}" =~ ^-.* ]]; then
        # we found an arg that begins with a dash, so we stop interpreting arguments
        # henceforth as packages and instead interpret them as flags to give to `go test`
        break
    fi
    # an arg found before the first flag is a package
    package_args+=" ${arg}"
    shift
done
gotest_flags+=" $*"

# Determine packages to test
godeps_package_prefix="vendor/"
test_packages=
if [[ -n "${package_args}" ]]; then
    for package in ${package_args}; do
        # If we're trying to recursively test a package under Godeps, strip the Godeps prefix so go test can find the packages correctly
        if [[ "${package}" == "${godeps_package_prefix}"*"/..." ]]; then
            test_packages="${test_packages} ${package:${#godeps_package_prefix}}"
        else
            test_packages="${test_packages} ${OS_GO_PACKAGE}/${package}"
        fi
    done
else
    # If no packages are given to test, we need to generate a list of all packages with unit tests
    test_packages="$(os::util::list_test_packages_under '*')"
fi

if [[ -n "${dry_run}" ]]; then
    echo "The following base flags for \`go test\` will be used by $0:"
    echo "go test ${gotest_flags}"
    echo "The following packages will be tested by $0:"
    for package in ${test_packages}; do
        echo "${package}"
    done
    exit 0
fi

# Run 'go test' with the accumulated arguments and packages:
if [[ -n "${junit_report}" ]]; then
    # we need to generate jUnit xml

    test_error_file="${LOG_DIR}/test-go-err.log"

    os::log::info "Running \`go test\`..."
    # we don't care if the `go test` fails in this pipe, as we want to generate the report and summarize the output anyway
    set +o pipefail

    os::util::ensure::built_binary_exists 'gotest2junit'
    report_file="$( mktemp "${ARTIFACT_DIR}/unit_report_XXXXX" ).xml"

    go test -json ${gotest_flags} ${test_packages} 2>"${test_error_file}" | tee "${JUNIT_REPORT_OUTPUT}" | gotest2junit > "${report_file}"
    test_return_code="${PIPESTATUS[0]}"

    gzip "${test_error_file}" -c > "${ARTIFACT_DIR}/unit-error.log.gz"
    gzip "${JUNIT_REPORT_OUTPUT}" -c > "${ARTIFACT_DIR}/unit.log.gz"

    set -o pipefail

    if [[ -s "${test_error_file}" ]]; then
        os::log::warning "\`go test\` had the following output to stderr:
$( cat "${test_error_file}") "
    fi

    if grep -q 'WARNING: DATA RACE' "${JUNIT_REPORT_OUTPUT}"; then
        locations=( $( sed -n '/WARNING: DATA RACE/=' "${JUNIT_REPORT_OUTPUT}") )
        if [[ "${#locations[@]}" -gt 1 ]]; then
            os::log::warning "\`go test\` detected data races."
            os::log::warning "Details can be found in the full output file at lines ${locations[*]}."
        else
            os::log::warning "\`go test\` detected a data race."
            os::log::warning "Details can be found in the full output file at line ${locations[*]}."
        fi
    fi

    exit "${test_return_code}"

elif [[ -n "${coverage_output_dir}" ]]; then
    # we need to generate coverage reports
    for test_package in ${test_packages}; do
        mkdir -p "${coverage_output_dir}/${test_package}"
        local_gotest_flags="${gotest_flags} -coverprofile=${coverage_output_dir}/${test_package}/profile.out"

        go test ${local_gotest_flags} ${test_package}
    done

    # assemble all profiles and generate a coverage report
    echo 'mode: atomic' > "${coverage_output_dir}/profiles.out"
    find "${coverage_output_dir}" -name profile.out | xargs sed '/^mode: atomic$/d' >> "${coverage_output_dir}/profiles.out"

    go tool cover "-html=${coverage_output_dir}/profiles.out" -o "${coverage_output_dir}/coverage.html"
    os::log::info "Coverage profile written to ${coverage_output_dir}/coverage.html"

    # clean up all of the individual coverage reports as they have been subsumed into the report at ${coverage_output_dir}/coverage.html
    # we can clean up all of the coverage reports at once as they all exist in subdirectories of ${coverage_output_dir}/${OS_GO_PACKAGE}
    # and they are the only files found in those subdirectories
    rm -rf "${coverage_output_dir:?}/${OS_GO_PACKAGE}"

elif [[ -n "${dlv_debug}" ]]; then
    # run tests using delve debugger
    dlv test ${test_packages}
else
    # we need to generate neither jUnit XML nor coverage reports
    go test ${gotest_flags} ${test_packages}
fi
