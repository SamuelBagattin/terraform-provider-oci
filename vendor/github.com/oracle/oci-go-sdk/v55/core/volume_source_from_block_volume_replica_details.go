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

// VolumeSourceFromBlockVolumeReplicaDetails Specifies the source block volume replica which the block volume will be created from.
// The block volume replica shoulbe be in the same availability domain as the block volume.
// Only one volume can be created from a replica at the same time.
type VolumeSourceFromBlockVolumeReplicaDetails struct {

	// The OCID of the block volume replica.
	Id *string `mandatory:"true" json:"id"`
}

func (m VolumeSourceFromBlockVolumeReplicaDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m VolumeSourceFromBlockVolumeReplicaDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVolumeSourceFromBlockVolumeReplicaDetails VolumeSourceFromBlockVolumeReplicaDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeVolumeSourceFromBlockVolumeReplicaDetails
	}{
		"blockVolumeReplica",
		(MarshalTypeVolumeSourceFromBlockVolumeReplicaDetails)(m),
	}

	return json.Marshal(&s)
}
