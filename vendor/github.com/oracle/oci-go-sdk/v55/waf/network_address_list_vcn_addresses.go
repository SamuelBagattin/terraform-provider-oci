// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// NetworkAddressListVcnAddresses A NetworkAddressList that contains VCN addresses.
type NetworkAddressListVcnAddresses struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the NetworkAddressList.
	Id *string `mandatory:"true" json:"id"`

	// NetworkAddressList display name, can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the NetworkAddressList was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// A list of private address prefixes, each associated with a particular VCN.
	// To specify all addresses in a VCN, use "0.0.0.0/0" for IPv4 and "::/0" for IPv6.
	VcnAddresses []PrivateAddresses `mandatory:"true" json:"vcnAddresses"`

	// The time the NetworkAddressList was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in FAILED state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the NetworkAddressList.
	LifecycleState NetworkAddressListLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m NetworkAddressListVcnAddresses) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m NetworkAddressListVcnAddresses) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m NetworkAddressListVcnAddresses) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m NetworkAddressListVcnAddresses) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m NetworkAddressListVcnAddresses) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m NetworkAddressListVcnAddresses) GetLifecycleState() NetworkAddressListLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m NetworkAddressListVcnAddresses) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetFreeformTags returns FreeformTags
func (m NetworkAddressListVcnAddresses) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m NetworkAddressListVcnAddresses) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m NetworkAddressListVcnAddresses) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m NetworkAddressListVcnAddresses) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m NetworkAddressListVcnAddresses) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNetworkAddressListVcnAddresses NetworkAddressListVcnAddresses
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeNetworkAddressListVcnAddresses
	}{
		"VCN_ADDRESSES",
		(MarshalTypeNetworkAddressListVcnAddresses)(m),
	}

	return json.Marshal(&s)
}
