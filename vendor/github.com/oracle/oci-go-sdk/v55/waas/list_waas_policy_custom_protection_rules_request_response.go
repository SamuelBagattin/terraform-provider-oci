// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waas

import (
	"github.com/oracle/oci-go-sdk/v55/common"
	"net/http"
)

// ListWaasPolicyCustomProtectionRulesRequest wrapper for the ListWaasPolicyCustomProtectionRules operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waas/ListWaasPolicyCustomProtectionRules.go.html to see an example of how to use ListWaasPolicyCustomProtectionRulesRequest.
type ListWaasPolicyCustomProtectionRulesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the WAAS policy.
	WaasPolicyId *string `mandatory:"true" contributesTo:"path" name:"waasPolicyId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated call. If unspecified, defaults to `10`.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous paginated call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Filter rules using a list of ModSecurity rule IDs.
	ModSecurityRuleId []string `contributesTo:"query" name:"modSecurityRuleId" collectionFormat:"multi"`

	// Filter rules using a list of actions.
	Action []ListWaasPolicyCustomProtectionRulesActionEnum `contributesTo:"query" name:"action" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWaasPolicyCustomProtectionRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWaasPolicyCustomProtectionRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWaasPolicyCustomProtectionRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWaasPolicyCustomProtectionRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListWaasPolicyCustomProtectionRulesResponse wrapper for the ListWaasPolicyCustomProtectionRules operation
type ListWaasPolicyCustomProtectionRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []WaasPolicyCustomProtectionRuleSummary instances
	Items []WaasPolicyCustomProtectionRuleSummary `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// For list pagination. When this header appears in the response, additional pages of results may remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListWaasPolicyCustomProtectionRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWaasPolicyCustomProtectionRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWaasPolicyCustomProtectionRulesActionEnum Enum with underlying type: string
type ListWaasPolicyCustomProtectionRulesActionEnum string

// Set of constants representing the allowable values for ListWaasPolicyCustomProtectionRulesActionEnum
const (
	ListWaasPolicyCustomProtectionRulesActionDetect ListWaasPolicyCustomProtectionRulesActionEnum = "DETECT"
	ListWaasPolicyCustomProtectionRulesActionBlock  ListWaasPolicyCustomProtectionRulesActionEnum = "BLOCK"
)

var mappingListWaasPolicyCustomProtectionRulesAction = map[string]ListWaasPolicyCustomProtectionRulesActionEnum{
	"DETECT": ListWaasPolicyCustomProtectionRulesActionDetect,
	"BLOCK":  ListWaasPolicyCustomProtectionRulesActionBlock,
}

// GetListWaasPolicyCustomProtectionRulesActionEnumValues Enumerates the set of values for ListWaasPolicyCustomProtectionRulesActionEnum
func GetListWaasPolicyCustomProtectionRulesActionEnumValues() []ListWaasPolicyCustomProtectionRulesActionEnum {
	values := make([]ListWaasPolicyCustomProtectionRulesActionEnum, 0)
	for _, v := range mappingListWaasPolicyCustomProtectionRulesAction {
		values = append(values, v)
	}
	return values
}
