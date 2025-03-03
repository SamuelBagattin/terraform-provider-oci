// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package logging

import (
	"github.com/oracle/oci-go-sdk/v55/common"
	"net/http"
)

// DeleteUnifiedAgentConfigurationRequest wrapper for the DeleteUnifiedAgentConfiguration operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/logging/DeleteUnifiedAgentConfiguration.go.html to see an example of how to use DeleteUnifiedAgentConfigurationRequest.
type DeleteUnifiedAgentConfigurationRequest struct {

	// The OCID of the Unified Agent configuration.
	UnifiedAgentConfigurationId *string `mandatory:"true" contributesTo:"path" name:"unifiedAgentConfigurationId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call for a
	// resource, set the `if-match` parameter to the value of the etag from a
	// previous GET or POST response for that resource. The resource will be
	// updated or deleted only if the etag you provide matches the resource's
	// current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeleteUnifiedAgentConfigurationRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteUnifiedAgentConfigurationRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DeleteUnifiedAgentConfigurationRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteUnifiedAgentConfigurationRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DeleteUnifiedAgentConfigurationResponse wrapper for the DeleteUnifiedAgentConfiguration operation
type DeleteUnifiedAgentConfigurationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OCID of the work request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteUnifiedAgentConfigurationResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteUnifiedAgentConfigurationResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
