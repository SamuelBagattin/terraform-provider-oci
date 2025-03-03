// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// OperationsInsightsWarehouseUserSummary Summary of a Operations Insights Warehouse User.
type OperationsInsightsWarehouseUserSummary struct {

	// OPSI Warehouse OCID
	OperationsInsightsWarehouseId *string `mandatory:"true" json:"operationsInsightsWarehouseId"`

	// Hub User OCID
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Username for schema which would have access to AWR Data,  Enterprise Manager Data and Operations Insights OPSI Hub.
	Name *string `mandatory:"true" json:"name"`

	// Indicate whether user has access to AWR data.
	IsAwrDataAccess *bool `mandatory:"true" json:"isAwrDataAccess"`

	// The time at which the resource was first created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Possible lifecycle states
	LifecycleState OperationsInsightsWarehouseUserLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// User provided connection password for the AWR Data,  Enterprise Manager Data and Operations Insights OPSI Hub.
	ConnectionPassword *string `mandatory:"false" json:"connectionPassword"`

	// Indicate whether user has access to EM data.
	IsEmDataAccess *bool `mandatory:"false" json:"isEmDataAccess"`

	// Indicate whether user has access to OPSI data.
	IsOpsiDataAccess *bool `mandatory:"false" json:"isOpsiDataAccess"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The time at which the resource was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m OperationsInsightsWarehouseUserSummary) String() string {
	return common.PointerString(m)
}
