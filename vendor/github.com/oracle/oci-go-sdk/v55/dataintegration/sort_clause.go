// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// SortClause The information about the sort object.
type SortClause struct {
	Field *ShapeField `mandatory:"false" json:"field"`

	// The sort order.
	Order SortClauseOrderEnum `mandatory:"false" json:"order,omitempty"`
}

func (m SortClause) String() string {
	return common.PointerString(m)
}

// SortClauseOrderEnum Enum with underlying type: string
type SortClauseOrderEnum string

// Set of constants representing the allowable values for SortClauseOrderEnum
const (
	SortClauseOrderAsc  SortClauseOrderEnum = "ASC"
	SortClauseOrderDesc SortClauseOrderEnum = "DESC"
)

var mappingSortClauseOrder = map[string]SortClauseOrderEnum{
	"ASC":  SortClauseOrderAsc,
	"DESC": SortClauseOrderDesc,
}

// GetSortClauseOrderEnumValues Enumerates the set of values for SortClauseOrderEnum
func GetSortClauseOrderEnumValues() []SortClauseOrderEnum {
	values := make([]SortClauseOrderEnum, 0)
	for _, v := range mappingSortClauseOrder {
		values = append(values, v)
	}
	return values
}
