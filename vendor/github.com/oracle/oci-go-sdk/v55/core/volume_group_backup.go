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

// VolumeGroupBackup A point-in-time copy of a volume group that can then be used to create a new volume group
// or restore a volume group. For more information, see Volume Groups (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroups.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you
// supply string values using the API.
type VolumeGroupBackup struct {

	// The OCID of the compartment that contains the volume group backup.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the volume group backup.
	Id *string `mandatory:"true" json:"id"`

	// The current state of a volume group backup.
	LifecycleState VolumeGroupBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the volume group backup was created. This is the time the actual point-in-time image
	// of the volume group data was taken. Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The type of backup.
	Type VolumeGroupBackupTypeEnum `mandatory:"true" json:"type"`

	// OCIDs for the volume backups in this volume group backup.
	VolumeBackupIds []string `mandatory:"true" json:"volumeBackupIds"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The date and time the volume group backup will expire and be automatically deleted.
	// Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339). This parameter will always be present for volume group
	// backups that were created automatically by a scheduled-backup policy. For manually
	// created volume group backups, it will be absent, signifying that there is no expiration
	// time and the backup will last forever until manually deleted.
	ExpirationTime *common.SDKTime `mandatory:"false" json:"expirationTime"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The aggregate size of the volume group backup, in MBs.
	SizeInMBs *int64 `mandatory:"false" json:"sizeInMBs"`

	// The aggregate size of the volume group backup, in GBs.
	SizeInGBs *int64 `mandatory:"false" json:"sizeInGBs"`

	// Specifies whether the volume group backup was created manually, or via scheduled
	// backup policy.
	SourceType VolumeGroupBackupSourceTypeEnum `mandatory:"false" json:"sourceType,omitempty"`

	// The date and time the request to create the volume group backup was received. Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeRequestReceived *common.SDKTime `mandatory:"false" json:"timeRequestReceived"`

	// The aggregate size used by the volume group backup, in MBs.
	// It is typically smaller than sizeInMBs, depending on the spaceconsumed
	// on the volume group and whether the volume backup is full or incremental.
	UniqueSizeInMbs *int64 `mandatory:"false" json:"uniqueSizeInMbs"`

	// The aggregate size used by the volume group backup, in GBs.
	// It is typically smaller than sizeInGBs, depending on the spaceconsumed
	// on the volume group and whether the volume backup is full or incremental.
	UniqueSizeInGbs *int64 `mandatory:"false" json:"uniqueSizeInGbs"`

	// The OCID of the source volume group.
	VolumeGroupId *string `mandatory:"false" json:"volumeGroupId"`

	// The OCID of the source volume group backup.
	SourceVolumeGroupBackupId *string `mandatory:"false" json:"sourceVolumeGroupBackupId"`
}

func (m VolumeGroupBackup) String() string {
	return common.PointerString(m)
}

// VolumeGroupBackupLifecycleStateEnum Enum with underlying type: string
type VolumeGroupBackupLifecycleStateEnum string

// Set of constants representing the allowable values for VolumeGroupBackupLifecycleStateEnum
const (
	VolumeGroupBackupLifecycleStateCreating        VolumeGroupBackupLifecycleStateEnum = "CREATING"
	VolumeGroupBackupLifecycleStateCommitted       VolumeGroupBackupLifecycleStateEnum = "COMMITTED"
	VolumeGroupBackupLifecycleStateAvailable       VolumeGroupBackupLifecycleStateEnum = "AVAILABLE"
	VolumeGroupBackupLifecycleStateTerminating     VolumeGroupBackupLifecycleStateEnum = "TERMINATING"
	VolumeGroupBackupLifecycleStateTerminated      VolumeGroupBackupLifecycleStateEnum = "TERMINATED"
	VolumeGroupBackupLifecycleStateFaulty          VolumeGroupBackupLifecycleStateEnum = "FAULTY"
	VolumeGroupBackupLifecycleStateRequestReceived VolumeGroupBackupLifecycleStateEnum = "REQUEST_RECEIVED"
)

var mappingVolumeGroupBackupLifecycleState = map[string]VolumeGroupBackupLifecycleStateEnum{
	"CREATING":         VolumeGroupBackupLifecycleStateCreating,
	"COMMITTED":        VolumeGroupBackupLifecycleStateCommitted,
	"AVAILABLE":        VolumeGroupBackupLifecycleStateAvailable,
	"TERMINATING":      VolumeGroupBackupLifecycleStateTerminating,
	"TERMINATED":       VolumeGroupBackupLifecycleStateTerminated,
	"FAULTY":           VolumeGroupBackupLifecycleStateFaulty,
	"REQUEST_RECEIVED": VolumeGroupBackupLifecycleStateRequestReceived,
}

// GetVolumeGroupBackupLifecycleStateEnumValues Enumerates the set of values for VolumeGroupBackupLifecycleStateEnum
func GetVolumeGroupBackupLifecycleStateEnumValues() []VolumeGroupBackupLifecycleStateEnum {
	values := make([]VolumeGroupBackupLifecycleStateEnum, 0)
	for _, v := range mappingVolumeGroupBackupLifecycleState {
		values = append(values, v)
	}
	return values
}

// VolumeGroupBackupSourceTypeEnum Enum with underlying type: string
type VolumeGroupBackupSourceTypeEnum string

// Set of constants representing the allowable values for VolumeGroupBackupSourceTypeEnum
const (
	VolumeGroupBackupSourceTypeManual    VolumeGroupBackupSourceTypeEnum = "MANUAL"
	VolumeGroupBackupSourceTypeScheduled VolumeGroupBackupSourceTypeEnum = "SCHEDULED"
)

var mappingVolumeGroupBackupSourceType = map[string]VolumeGroupBackupSourceTypeEnum{
	"MANUAL":    VolumeGroupBackupSourceTypeManual,
	"SCHEDULED": VolumeGroupBackupSourceTypeScheduled,
}

// GetVolumeGroupBackupSourceTypeEnumValues Enumerates the set of values for VolumeGroupBackupSourceTypeEnum
func GetVolumeGroupBackupSourceTypeEnumValues() []VolumeGroupBackupSourceTypeEnum {
	values := make([]VolumeGroupBackupSourceTypeEnum, 0)
	for _, v := range mappingVolumeGroupBackupSourceType {
		values = append(values, v)
	}
	return values
}

// VolumeGroupBackupTypeEnum Enum with underlying type: string
type VolumeGroupBackupTypeEnum string

// Set of constants representing the allowable values for VolumeGroupBackupTypeEnum
const (
	VolumeGroupBackupTypeFull        VolumeGroupBackupTypeEnum = "FULL"
	VolumeGroupBackupTypeIncremental VolumeGroupBackupTypeEnum = "INCREMENTAL"
)

var mappingVolumeGroupBackupType = map[string]VolumeGroupBackupTypeEnum{
	"FULL":        VolumeGroupBackupTypeFull,
	"INCREMENTAL": VolumeGroupBackupTypeIncremental,
}

// GetVolumeGroupBackupTypeEnumValues Enumerates the set of values for VolumeGroupBackupTypeEnum
func GetVolumeGroupBackupTypeEnumValues() []VolumeGroupBackupTypeEnum {
	values := make([]VolumeGroupBackupTypeEnum, 0)
	for _, v := range mappingVolumeGroupBackupType {
		values = append(values, v)
	}
	return values
}
