// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"github.com/oracle/oci-go-sdk/v55/common"
	"net/http"
)

// GetDeployEnvironmentRequest wrapper for the GetDeployEnvironment operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetDeployEnvironment.go.html to see an example of how to use GetDeployEnvironmentRequest.
type GetDeployEnvironmentRequest struct {

	// Unique environment identifier.
	DeployEnvironmentId *string `mandatory:"true" contributesTo:"path" name:"deployEnvironmentId"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDeployEnvironmentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDeployEnvironmentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetDeployEnvironmentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDeployEnvironmentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetDeployEnvironmentResponse wrapper for the GetDeployEnvironment operation
type GetDeployEnvironmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DeployEnvironment instance
	DeployEnvironment `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDeployEnvironmentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDeployEnvironmentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
