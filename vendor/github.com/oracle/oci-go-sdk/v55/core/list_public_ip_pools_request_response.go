// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/v55/common"
	"net/http"
)

// ListPublicIpPoolsRequest wrapper for the ListPublicIpPools operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListPublicIpPools.go.html to see an example of how to use ListPublicIpPoolsRequest.
type ListPublicIpPoolsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given BYOIP CIDR block.
	ByoipRangeId *string `mandatory:"false" contributesTo:"query" name:"byoipRangeId"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListPublicIpPoolsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListPublicIpPoolsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPublicIpPoolsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPublicIpPoolsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPublicIpPoolsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPublicIpPoolsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListPublicIpPoolsResponse wrapper for the ListPublicIpPools operation
type ListPublicIpPoolsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PublicIpPoolCollection instances
	PublicIpPoolCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListPublicIpPoolsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPublicIpPoolsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPublicIpPoolsSortByEnum Enum with underlying type: string
type ListPublicIpPoolsSortByEnum string

// Set of constants representing the allowable values for ListPublicIpPoolsSortByEnum
const (
	ListPublicIpPoolsSortByTimecreated ListPublicIpPoolsSortByEnum = "TIMECREATED"
	ListPublicIpPoolsSortByDisplayname ListPublicIpPoolsSortByEnum = "DISPLAYNAME"
)

var mappingListPublicIpPoolsSortBy = map[string]ListPublicIpPoolsSortByEnum{
	"TIMECREATED": ListPublicIpPoolsSortByTimecreated,
	"DISPLAYNAME": ListPublicIpPoolsSortByDisplayname,
}

// GetListPublicIpPoolsSortByEnumValues Enumerates the set of values for ListPublicIpPoolsSortByEnum
func GetListPublicIpPoolsSortByEnumValues() []ListPublicIpPoolsSortByEnum {
	values := make([]ListPublicIpPoolsSortByEnum, 0)
	for _, v := range mappingListPublicIpPoolsSortBy {
		values = append(values, v)
	}
	return values
}

// ListPublicIpPoolsSortOrderEnum Enum with underlying type: string
type ListPublicIpPoolsSortOrderEnum string

// Set of constants representing the allowable values for ListPublicIpPoolsSortOrderEnum
const (
	ListPublicIpPoolsSortOrderAsc  ListPublicIpPoolsSortOrderEnum = "ASC"
	ListPublicIpPoolsSortOrderDesc ListPublicIpPoolsSortOrderEnum = "DESC"
)

var mappingListPublicIpPoolsSortOrder = map[string]ListPublicIpPoolsSortOrderEnum{
	"ASC":  ListPublicIpPoolsSortOrderAsc,
	"DESC": ListPublicIpPoolsSortOrderDesc,
}

// GetListPublicIpPoolsSortOrderEnumValues Enumerates the set of values for ListPublicIpPoolsSortOrderEnum
func GetListPublicIpPoolsSortOrderEnumValues() []ListPublicIpPoolsSortOrderEnum {
	values := make([]ListPublicIpPoolsSortOrderEnum, 0)
	for _, v := range mappingListPublicIpPoolsSortOrder {
		values = append(values, v)
	}
	return values
}
