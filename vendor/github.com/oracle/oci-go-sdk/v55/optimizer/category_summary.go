// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// CategorySummary The metadata associated with the category summary.
type CategorySummary struct {

	// The unique OCID of the category.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy. The tenancy is the root compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name assigned to the category.
	Name *string `mandatory:"true" json:"name"`

	// Text describing the category.
	Description *string `mandatory:"true" json:"description"`

	// An array of `RecommendationCount` objects grouped by the level of importance assigned to each recommendation.
	RecommendationCounts []RecommendationCount `mandatory:"true" json:"recommendationCounts"`

	// An array of `ResourceCount` objects grouped by the status of each recommendation.
	ResourceCounts []ResourceCount `mandatory:"true" json:"resourceCounts"`

	// The category's current state.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The estimated cost savings, in dollars, for the category.
	EstimatedCostSaving *float64 `mandatory:"true" json:"estimatedCostSaving"`

	// The date and time the category details were created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the category details were last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Additional metadata key/value pairs for the category summary.
	// For example:
	// `{"EstimatedSaving": "200"}`
	ExtendedMetadata map[string]string `mandatory:"false" json:"extendedMetadata"`
}

func (m CategorySummary) String() string {
	return common.PointerString(m)
}
