// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// TopologyAssociatedWithEntityRelationship Defines the `AssociatedWith` relationship between virtual network topology entities. An `AssociatedWith` relationship
// is defined when there is no obvious `contains` relationship but entities are still related.
// For example, a DRG is associated with a VCN because a DRG is not managed by VCN but can be
// attached to a VCN.
type TopologyAssociatedWithEntityRelationship struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the first entity in the relationship.
	Id1 *string `mandatory:"true" json:"id1"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the second entity in the relationship.
	Id2 *string `mandatory:"true" json:"id2"`

	AssociatedWithDetails *TopologyAssociatedWithRelationshipDetails `mandatory:"false" json:"associatedWithDetails"`
}

//GetId1 returns Id1
func (m TopologyAssociatedWithEntityRelationship) GetId1() *string {
	return m.Id1
}

//GetId2 returns Id2
func (m TopologyAssociatedWithEntityRelationship) GetId2() *string {
	return m.Id2
}

func (m TopologyAssociatedWithEntityRelationship) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m TopologyAssociatedWithEntityRelationship) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTopologyAssociatedWithEntityRelationship TopologyAssociatedWithEntityRelationship
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTopologyAssociatedWithEntityRelationship
	}{
		"ASSOCIATED_WITH",
		(MarshalTypeTopologyAssociatedWithEntityRelationship)(m),
	}

	return json.Marshal(&s)
}
