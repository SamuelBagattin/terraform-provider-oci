// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v55/common"
	"net/http"
)

// ListReferencesRequest wrapper for the ListReferences operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListReferences.go.html to see an example of how to use ListReferencesRequest.
type ListReferencesRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The application key.
	ApplicationKey *string `mandatory:"true" contributesTo:"path" name:"applicationKey"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListReferencesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListReferencesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListReferencesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListReferencesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListReferencesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListReferencesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListReferencesResponse wrapper for the ListReferences operation
type ListReferencesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ReferenceSummaryCollection instances
	ReferenceSummaryCollection `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListReferencesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListReferencesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListReferencesSortOrderEnum Enum with underlying type: string
type ListReferencesSortOrderEnum string

// Set of constants representing the allowable values for ListReferencesSortOrderEnum
const (
	ListReferencesSortOrderAsc  ListReferencesSortOrderEnum = "ASC"
	ListReferencesSortOrderDesc ListReferencesSortOrderEnum = "DESC"
)

var mappingListReferencesSortOrder = map[string]ListReferencesSortOrderEnum{
	"ASC":  ListReferencesSortOrderAsc,
	"DESC": ListReferencesSortOrderDesc,
}

// GetListReferencesSortOrderEnumValues Enumerates the set of values for ListReferencesSortOrderEnum
func GetListReferencesSortOrderEnumValues() []ListReferencesSortOrderEnum {
	values := make([]ListReferencesSortOrderEnum, 0)
	for _, v := range mappingListReferencesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListReferencesSortByEnum Enum with underlying type: string
type ListReferencesSortByEnum string

// Set of constants representing the allowable values for ListReferencesSortByEnum
const (
	ListReferencesSortByTimeCreated ListReferencesSortByEnum = "TIME_CREATED"
	ListReferencesSortByDisplayName ListReferencesSortByEnum = "DISPLAY_NAME"
)

var mappingListReferencesSortBy = map[string]ListReferencesSortByEnum{
	"TIME_CREATED": ListReferencesSortByTimeCreated,
	"DISPLAY_NAME": ListReferencesSortByDisplayName,
}

// GetListReferencesSortByEnumValues Enumerates the set of values for ListReferencesSortByEnum
func GetListReferencesSortByEnumValues() []ListReferencesSortByEnum {
	values := make([]ListReferencesSortByEnum, 0)
	for _, v := range mappingListReferencesSortBy {
		values = append(values, v)
	}
	return values
}
