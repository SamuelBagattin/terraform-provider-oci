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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// HostInsight Host insight resource.
type HostInsight interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	GetId() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The host name. The host name is unique amongst the hosts managed by the same management agent.
	GetHostName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Indicates the status of a host insight in Operations Insights
	GetStatus() ResourceStatusEnum

	// The time the the host insight was first enabled. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The current state of the host.
	GetLifecycleState() LifecycleStateEnum

	// The user-friendly name for the host. The name does not have to be unique.
	GetHostDisplayName() *string

	// Operations Insights internal representation of the host type. Possible value is EXTERNAL-HOST.
	GetHostType() *string

	// Processor count. This is the OCPU count for Autonomous Database and CPU core count for other database types.
	GetProcessorCount() *int

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// The time the host insight was updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string
}

type hostinsight struct {
	JsonData         []byte
	Id               *string                           `mandatory:"true" json:"id"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	HostName         *string                           `mandatory:"true" json:"hostName"`
	FreeformTags     map[string]string                 `mandatory:"true" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
	Status           ResourceStatusEnum                `mandatory:"true" json:"status"`
	TimeCreated      *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState   LifecycleStateEnum                `mandatory:"true" json:"lifecycleState"`
	HostDisplayName  *string                           `mandatory:"false" json:"hostDisplayName"`
	HostType         *string                           `mandatory:"false" json:"hostType"`
	ProcessorCount   *int                              `mandatory:"false" json:"processorCount"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	TimeUpdated      *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleDetails *string                           `mandatory:"false" json:"lifecycleDetails"`
	EntitySource     string                            `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *hostinsight) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerhostinsight hostinsight
	s := struct {
		Model Unmarshalerhostinsight
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.HostName = s.Model.HostName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Status = s.Model.Status
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.HostDisplayName = s.Model.HostDisplayName
	m.HostType = s.Model.HostType
	m.ProcessorCount = s.Model.ProcessorCount
	m.SystemTags = s.Model.SystemTags
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *hostinsight) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "EM_MANAGED_EXTERNAL_HOST":
		mm := EmManagedExternalHostInsight{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MACS_MANAGED_EXTERNAL_HOST":
		mm := MacsManagedExternalHostInsight{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m hostinsight) GetId() *string {
	return m.Id
}

//GetCompartmentId returns CompartmentId
func (m hostinsight) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetHostName returns HostName
func (m hostinsight) GetHostName() *string {
	return m.HostName
}

//GetFreeformTags returns FreeformTags
func (m hostinsight) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m hostinsight) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetStatus returns Status
func (m hostinsight) GetStatus() ResourceStatusEnum {
	return m.Status
}

//GetTimeCreated returns TimeCreated
func (m hostinsight) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetLifecycleState returns LifecycleState
func (m hostinsight) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

//GetHostDisplayName returns HostDisplayName
func (m hostinsight) GetHostDisplayName() *string {
	return m.HostDisplayName
}

//GetHostType returns HostType
func (m hostinsight) GetHostType() *string {
	return m.HostType
}

//GetProcessorCount returns ProcessorCount
func (m hostinsight) GetProcessorCount() *int {
	return m.ProcessorCount
}

//GetSystemTags returns SystemTags
func (m hostinsight) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

//GetTimeUpdated returns TimeUpdated
func (m hostinsight) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleDetails returns LifecycleDetails
func (m hostinsight) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m hostinsight) String() string {
	return common.PointerString(m)
}
