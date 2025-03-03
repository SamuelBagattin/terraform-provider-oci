// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// ShapeSummary The shape of the DB System. The shape determines resources to allocate
// to the DB System - CPU cores and memory for VM shapes; CPU cores, memory
// and storage for non-VM (or bare metal) shapes.  For a description of
// shapes, see DB System Shape Options (https://docs.cloud.oracle.com/mysql-database/doc/shapes.htm).
type ShapeSummary struct {

	// The name of the shape used for the DB System.
	Name *string `mandatory:"true" json:"name"`

	// The number of CPU Cores the Instance provides. These are "OCPU"s.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The amount of RAM the Instance provides. This is an IEC base-2 number.
	MemorySizeInGBs *int `mandatory:"true" json:"memorySizeInGBs"`

	// What service features the shape is supported for.
	IsSupportedFor []ShapeSummaryIsSupportedForEnum `mandatory:"false" json:"isSupportedFor,omitempty"`
}

func (m ShapeSummary) String() string {
	return common.PointerString(m)
}

// ShapeSummaryIsSupportedForEnum Enum with underlying type: string
type ShapeSummaryIsSupportedForEnum string

// Set of constants representing the allowable values for ShapeSummaryIsSupportedForEnum
const (
	ShapeSummaryIsSupportedForDbsystem         ShapeSummaryIsSupportedForEnum = "DBSYSTEM"
	ShapeSummaryIsSupportedForAnalyticscluster ShapeSummaryIsSupportedForEnum = "ANALYTICSCLUSTER"
	ShapeSummaryIsSupportedForHeatwavecluster  ShapeSummaryIsSupportedForEnum = "HEATWAVECLUSTER"
)

var mappingShapeSummaryIsSupportedFor = map[string]ShapeSummaryIsSupportedForEnum{
	"DBSYSTEM":         ShapeSummaryIsSupportedForDbsystem,
	"ANALYTICSCLUSTER": ShapeSummaryIsSupportedForAnalyticscluster,
	"HEATWAVECLUSTER":  ShapeSummaryIsSupportedForHeatwavecluster,
}

// GetShapeSummaryIsSupportedForEnumValues Enumerates the set of values for ShapeSummaryIsSupportedForEnum
func GetShapeSummaryIsSupportedForEnumValues() []ShapeSummaryIsSupportedForEnum {
	values := make([]ShapeSummaryIsSupportedForEnum, 0)
	for _, v := range mappingShapeSummaryIsSupportedFor {
		values = append(values, v)
	}
	return values
}
