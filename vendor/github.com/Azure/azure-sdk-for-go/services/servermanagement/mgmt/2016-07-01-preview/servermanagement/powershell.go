package servermanagement

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// PowerShellClient is the REST API for Azure Server Management Service.
type PowerShellClient struct {
	BaseClient
}

// NewPowerShellClient creates an instance of the PowerShellClient client.
func NewPowerShellClient(subscriptionID string) PowerShellClient {
	return NewPowerShellClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPowerShellClientWithBaseURI creates an instance of the PowerShellClient client.
func NewPowerShellClientWithBaseURI(baseURI string, subscriptionID string) PowerShellClient {
	return PowerShellClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CancelCommand cancels a PowerShell command.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user
// subscriptionId. nodeName is the node name (256 characters maximum). session is the sessionId from the user.
// pssession is the PowerShell sessionId from the user.
func (client PowerShellClient) CancelCommand(ctx context.Context, resourceGroupName string, nodeName string, session string, pssession string) (result PowerShellCancelCommandFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servermanagement.PowerShellClient", "CancelCommand", err.Error())
	}

	req, err := client.CancelCommandPreparer(ctx, resourceGroupName, nodeName, session, pssession)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "CancelCommand", nil, "Failure preparing request")
		return
	}

	result, err = client.CancelCommandSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "CancelCommand", result.Response(), "Failure sending request")
		return
	}

	return
}

// CancelCommandPreparer prepares the CancelCommand request.
func (client PowerShellClient) CancelCommandPreparer(ctx context.Context, resourceGroupName string, nodeName string, session string, pssession string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"pssession":         autorest.Encode("path", pssession),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}/features/powerShellConsole/pssessions/{pssession}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelCommandSender sends the CancelCommand request. The method will close the
// http.Response Body if it receives an error.
func (client PowerShellClient) CancelCommandSender(req *http.Request) (future PowerShellCancelCommandFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// CancelCommandResponder handles the response to the CancelCommand request. The method always
// closes the http.Response Body.
func (client PowerShellClient) CancelCommandResponder(resp *http.Response) (result PowerShellCommandResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateSession creates a PowerShell session.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user
// subscriptionId. nodeName is the node name (256 characters maximum). session is the sessionId from the user.
// pssession is the PowerShell sessionId from the user.
func (client PowerShellClient) CreateSession(ctx context.Context, resourceGroupName string, nodeName string, session string, pssession string) (result PowerShellCreateSessionFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servermanagement.PowerShellClient", "CreateSession", err.Error())
	}

	req, err := client.CreateSessionPreparer(ctx, resourceGroupName, nodeName, session, pssession)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "CreateSession", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSessionSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "CreateSession", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateSessionPreparer prepares the CreateSession request.
func (client PowerShellClient) CreateSessionPreparer(ctx context.Context, resourceGroupName string, nodeName string, session string, pssession string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"pssession":         autorest.Encode("path", pssession),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}/features/powerShellConsole/pssessions/{pssession}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSessionSender sends the CreateSession request. The method will close the
// http.Response Body if it receives an error.
func (client PowerShellClient) CreateSessionSender(req *http.Request) (future PowerShellCreateSessionFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// CreateSessionResponder handles the response to the CreateSession request. The method always
// closes the http.Response Body.
func (client PowerShellClient) CreateSessionResponder(resp *http.Response) (result PowerShellSessionResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetCommandStatus gets the status of a command.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user
// subscriptionId. nodeName is the node name (256 characters maximum). session is the sessionId from the user.
// pssession is the PowerShell sessionId from the user. expand is gets current output from an ongoing call.
func (client PowerShellClient) GetCommandStatus(ctx context.Context, resourceGroupName string, nodeName string, session string, pssession string, expand PowerShellExpandOption) (result PowerShellCommandStatus, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servermanagement.PowerShellClient", "GetCommandStatus", err.Error())
	}

	req, err := client.GetCommandStatusPreparer(ctx, resourceGroupName, nodeName, session, pssession, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "GetCommandStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetCommandStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "GetCommandStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetCommandStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "GetCommandStatus", resp, "Failure responding to request")
	}

	return
}

// GetCommandStatusPreparer prepares the GetCommandStatus request.
func (client PowerShellClient) GetCommandStatusPreparer(ctx context.Context, resourceGroupName string, nodeName string, session string, pssession string, expand PowerShellExpandOption) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"pssession":         autorest.Encode("path", pssession),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}/features/powerShellConsole/pssessions/{pssession}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetCommandStatusSender sends the GetCommandStatus request. The method will close the
// http.Response Body if it receives an error.
func (client PowerShellClient) GetCommandStatusSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetCommandStatusResponder handles the response to the GetCommandStatus request. The method always
// closes the http.Response Body.
func (client PowerShellClient) GetCommandStatusResponder(resp *http.Response) (result PowerShellCommandStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// InvokeCommand creates a PowerShell script and invokes it.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user
// subscriptionId. nodeName is the node name (256 characters maximum). session is the sessionId from the user.
// pssession is the PowerShell sessionId from the user. powerShellCommandParameters is parameters supplied to the
// Invoke PowerShell Command operation.
func (client PowerShellClient) InvokeCommand(ctx context.Context, resourceGroupName string, nodeName string, session string, pssession string, powerShellCommandParameters PowerShellCommandParameters) (result PowerShellInvokeCommandFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servermanagement.PowerShellClient", "InvokeCommand", err.Error())
	}

	req, err := client.InvokeCommandPreparer(ctx, resourceGroupName, nodeName, session, pssession, powerShellCommandParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "InvokeCommand", nil, "Failure preparing request")
		return
	}

	result, err = client.InvokeCommandSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "InvokeCommand", result.Response(), "Failure sending request")
		return
	}

	return
}

// InvokeCommandPreparer prepares the InvokeCommand request.
func (client PowerShellClient) InvokeCommandPreparer(ctx context.Context, resourceGroupName string, nodeName string, session string, pssession string, powerShellCommandParameters PowerShellCommandParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"pssession":         autorest.Encode("path", pssession),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}/features/powerShellConsole/pssessions/{pssession}/invokeCommand", pathParameters),
		autorest.WithJSON(powerShellCommandParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// InvokeCommandSender sends the InvokeCommand request. The method will close the
// http.Response Body if it receives an error.
func (client PowerShellClient) InvokeCommandSender(req *http.Request) (future PowerShellInvokeCommandFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// InvokeCommandResponder handles the response to the InvokeCommand request. The method always
// closes the http.Response Body.
func (client PowerShellClient) InvokeCommandResponder(resp *http.Response) (result PowerShellCommandResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListSession gets a list of the active sessions.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user
// subscriptionId. nodeName is the node name (256 characters maximum). session is the sessionId from the user.
func (client PowerShellClient) ListSession(ctx context.Context, resourceGroupName string, nodeName string, session string) (result PowerShellSessionResources, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servermanagement.PowerShellClient", "ListSession", err.Error())
	}

	req, err := client.ListSessionPreparer(ctx, resourceGroupName, nodeName, session)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "ListSession", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSessionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "ListSession", resp, "Failure sending request")
		return
	}

	result, err = client.ListSessionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "ListSession", resp, "Failure responding to request")
	}

	return
}

// ListSessionPreparer prepares the ListSession request.
func (client PowerShellClient) ListSessionPreparer(ctx context.Context, resourceGroupName string, nodeName string, session string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}/features/powerShellConsole/pssessions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSessionSender sends the ListSession request. The method will close the
// http.Response Body if it receives an error.
func (client PowerShellClient) ListSessionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListSessionResponder handles the response to the ListSession request. The method always
// closes the http.Response Body.
func (client PowerShellClient) ListSessionResponder(resp *http.Response) (result PowerShellSessionResources, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// TabCompletion gets tab completion values for a command.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user
// subscriptionId. nodeName is the node name (256 characters maximum). session is the sessionId from the user.
// pssession is the PowerShell sessionId from the user. powerShellTabCompletionParamters is parameters supplied to
// the tab completion call.
func (client PowerShellClient) TabCompletion(ctx context.Context, resourceGroupName string, nodeName string, session string, pssession string, powerShellTabCompletionParamters PowerShellTabCompletionParameters) (result PowerShellTabCompletionResults, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servermanagement.PowerShellClient", "TabCompletion", err.Error())
	}

	req, err := client.TabCompletionPreparer(ctx, resourceGroupName, nodeName, session, pssession, powerShellTabCompletionParamters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "TabCompletion", nil, "Failure preparing request")
		return
	}

	resp, err := client.TabCompletionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "TabCompletion", resp, "Failure sending request")
		return
	}

	result, err = client.TabCompletionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "TabCompletion", resp, "Failure responding to request")
	}

	return
}

// TabCompletionPreparer prepares the TabCompletion request.
func (client PowerShellClient) TabCompletionPreparer(ctx context.Context, resourceGroupName string, nodeName string, session string, pssession string, powerShellTabCompletionParamters PowerShellTabCompletionParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"pssession":         autorest.Encode("path", pssession),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}/features/powerShellConsole/pssessions/{pssession}/tab", pathParameters),
		autorest.WithJSON(powerShellTabCompletionParamters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TabCompletionSender sends the TabCompletion request. The method will close the
// http.Response Body if it receives an error.
func (client PowerShellClient) TabCompletionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// TabCompletionResponder handles the response to the TabCompletion request. The method always
// closes the http.Response Body.
func (client PowerShellClient) TabCompletionResponder(resp *http.Response) (result PowerShellTabCompletionResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateCommand updates a running PowerShell command with more data.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user
// subscriptionId. nodeName is the node name (256 characters maximum). session is the sessionId from the user.
// pssession is the PowerShell sessionId from the user.
func (client PowerShellClient) UpdateCommand(ctx context.Context, resourceGroupName string, nodeName string, session string, pssession string) (result PowerShellUpdateCommandFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servermanagement.PowerShellClient", "UpdateCommand", err.Error())
	}

	req, err := client.UpdateCommandPreparer(ctx, resourceGroupName, nodeName, session, pssession)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "UpdateCommand", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateCommandSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "UpdateCommand", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdateCommandPreparer prepares the UpdateCommand request.
func (client PowerShellClient) UpdateCommandPreparer(ctx context.Context, resourceGroupName string, nodeName string, session string, pssession string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"pssession":         autorest.Encode("path", pssession),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}/features/powerShellConsole/pssessions/{pssession}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateCommandSender sends the UpdateCommand request. The method will close the
// http.Response Body if it receives an error.
func (client PowerShellClient) UpdateCommandSender(req *http.Request) (future PowerShellUpdateCommandFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// UpdateCommandResponder handles the response to the UpdateCommand request. The method always
// closes the http.Response Body.
func (client PowerShellClient) UpdateCommandResponder(resp *http.Response) (result PowerShellCommandResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
