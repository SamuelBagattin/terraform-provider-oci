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
	"github.com/oracle/oci-go-sdk/v55/common"
)

// CopyVolumeGroupBackupDetails The representation of CopyVolumeGroupBackupDetails
type CopyVolumeGroupBackupDetails struct {

	// The name of the destination region.
	// Example: `us-ashburn-1`
	DestinationRegion *string `mandatory:"true" json:"destinationRegion"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the Key Management key in the destination region which will be the master encryption key
	// for the copied volume group backup.
	// If you do not specify this attribute the volume group backup will be encrypted with the Oracle-provided encryption
	// key when it is copied to the destination region.
	//
	// For more information about the Key Management service and encryption keys, see
	// Overview of Key Management (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm) and
	// Using Keys (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/usingkeys.htm).
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`
}

func (m CopyVolumeGroupBackupDetails) String() string {
	return common.PointerString(m)
}
