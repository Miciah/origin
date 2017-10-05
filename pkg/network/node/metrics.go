package node

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/openshift/origin/pkg/util/ovs"
	"github.com/prometheus/client_golang/prometheus"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

const (
	SDNNamespace = "openshift"
	SDNSubsystem = "sdn"

	OVSFlowsKey                 = "ovs_flows"
	ARPCacheAvailableEntriesKey = "arp_cache_entries"
	PodIPsKey                   = "pod_ips"
	PodSetupErrorsKey           = "pod_setup_errors"
	PodSetupLatencyKey          = "pod_setup_latency"
	PodTeardownErrorsKey        = "pod_teardown_errors"
	PodTeardownLatencyKey       = "pod_teardown_latency"
)

var (
	OVSFlows = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: SDNNamespace,
			Subsystem: SDNSubsystem,
			Name:      OVSFlowsKey,
			Help:      "Number of Open vSwitch flows",
		},
	)

	ARPCacheAvailableEntries = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: SDNNamespace,
			Subsystem: SDNSubsystem,
			Name:      ARPCacheAvailableEntriesKey,
			Help:      "Number of available entries in the ARP cache",
		},
	)

	PodIPs = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: SDNNamespace,
			Subsystem: SDNSubsystem,
			Name:      PodIPsKey,
			Help:      "Number of allocated pod IPs",
		},
	)

	PodSetupErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: SDNNamespace,
			Subsystem: SDNSubsystem,
			Name:      PodSetupErrorsKey,
			Help:      "Number pod setup errors",
		},
	)

	PodSetupLatency = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: SDNNamespace,
			Subsystem: SDNSubsystem,
			Name:      PodSetupLatencyKey,
			Help:      "Latency of pod network setup in microseconds",
		},
		[]string{"pod_namespace", "pod_name", "sandbox_id"},
	)

	PodTeardownErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: SDNNamespace,
			Subsystem: SDNSubsystem,
			Name:      PodTeardownErrorsKey,
			Help:      "Number pod teardown errors",
		},
	)

	PodTeardownLatency = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: SDNNamespace,
			Subsystem: SDNSubsystem,
			Name:      PodTeardownLatencyKey,
			Help:      "Latency of pod network teardown in microseconds",
		},
		[]string{"pod_namespace", "pod_name", "sandbox_id"},
	)

	// num stale OVS flows (flows that reference non-existent ports)
	// num vnids (in the master)
	// num netnamespaces (in the master)
	// iptables call time (in upstream kube)
	// iptables call failures (in upstream kube)
	// iptables num rules (in upstream kube)
)

var registerMetrics sync.Once

// Register all node metrics.
func RegisterMetrics() {
	registerMetrics.Do(func() {
		prometheus.MustRegister(OVSFlows)
		prometheus.MustRegister(ARPCacheAvailableEntries)
		prometheus.MustRegister(PodIPs)
		prometheus.MustRegister(PodSetupErrors)
		prometheus.MustRegister(PodSetupLatency)
		prometheus.MustRegister(PodTeardownErrors)
		prometheus.MustRegister(PodTeardownLatency)
	})
}

// Gets the time since the specified start in microseconds.
func sinceInMicroseconds(start time.Time) float64 {
	return float64(time.Since(start).Nanoseconds() / time.Microsecond.Nanoseconds())
}

func gatherPeriodicMetrics(ovs ovs.Interface) {
	updateOVSMetrics(ovs)
	updateARPMetrics()
	updatePodIPMetrics()
}

func updateOVSMetrics(ovs ovs.Interface) {
	flows, err := ovs.DumpFlows()
	if err == nil {
		OVSFlows.Set(float64(len(flows)))
	} else {
		utilruntime.HandleError(fmt.Errorf("failed to dump OVS flows for metrics: %v", err))
	}
}

func updateARPMetrics() {
	var used int
	data, err := ioutil.ReadFile("/proc/net/arp")
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("failed to read ARP entries for metrics: %v", err))
		return
	}
	lines := strings.Split(string(data), "\n")
	// Skip the header line
	used = len(lines) - 1

	// gc_thresh2 isn't the absolute max, but it's the level at which
	// garbage collection (and thus problems) could start.
	data, err = ioutil.ReadFile("/proc/sys/net/ipv4/neigh/default/gc_thresh2")
	if pathErr, ok := err.(*os.PathError); ok && pathErr.Err == syscall.ENOENT {
		// gc_thresh* may not exist in some cases; don't log an error
		return
	} else if err != nil {
		utilruntime.HandleError(fmt.Errorf("failed to read max ARP entries for metrics: %T %v", err, err))
		return
	}

	max, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err == nil {
		available := max - used
		if available < 0 {
			available = 0
		}
		ARPCacheAvailableEntries.Set(float64(available))
	} else {
		utilruntime.HandleError(fmt.Errorf("failed to parse max ARP entries %q for metrics: %T %v", data, err, err))
	}
}

func updatePodIPMetrics() {
	numAddrs := 0
	items, err := ioutil.ReadDir("/var/lib/cni/networks/openshift-sdn/")
	if pathErr, ok := err.(*os.PathError); ok && pathErr.Err == syscall.ENOENT {
		// Don't log an error if the directory doesn't exist (eg, no pods started yet)
		return
	} else if err != nil {
		utilruntime.HandleError(fmt.Errorf("failed to read pod IPs for metrics: %v", err))
	}

	for _, i := range items {
		if net.ParseIP(i.Name()) != nil {
			numAddrs++
		}
	}
	OVSFlows.Set(float64(numAddrs))
}
