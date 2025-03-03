// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

// RecommendationResourceTypeEnum Enum with underlying type: string
type RecommendationResourceTypeEnum string

// Set of constants representing the allowable values for RecommendationResourceTypeEnum
const (
	RecommendationResourceTypeDataEntity RecommendationResourceTypeEnum = "DATA_ENTITY"
	RecommendationResourceTypeAttribute  RecommendationResourceTypeEnum = "ATTRIBUTE"
	RecommendationResourceTypeTerm       RecommendationResourceTypeEnum = "TERM"
	RecommendationResourceTypeCategory   RecommendationResourceTypeEnum = "CATEGORY"
)

var mappingRecommendationResourceType = map[string]RecommendationResourceTypeEnum{
	"DATA_ENTITY": RecommendationResourceTypeDataEntity,
	"ATTRIBUTE":   RecommendationResourceTypeAttribute,
	"TERM":        RecommendationResourceTypeTerm,
	"CATEGORY":    RecommendationResourceTypeCategory,
}

// GetRecommendationResourceTypeEnumValues Enumerates the set of values for RecommendationResourceTypeEnum
func GetRecommendationResourceTypeEnumValues() []RecommendationResourceTypeEnum {
	values := make([]RecommendationResourceTypeEnum, 0)
	for _, v := range mappingRecommendationResourceType {
		values = append(values, v)
	}
	return values
}
