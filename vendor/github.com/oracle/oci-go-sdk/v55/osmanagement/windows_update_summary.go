// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// WindowsUpdateSummary An update available for a Windows managed instance.
type WindowsUpdateSummary struct {

	// Windows Update name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique identifier for the Windows update. NOTE - This is not an OCID,
	// but is a unique identifier assigned by Microsoft.
	// Example: `6981d463-cd91-4a26-b7c4-ea4ded9183ed`
	Name *string `mandatory:"true" json:"name"`

	// The purpose of this update.
	UpdateType UpdateTypesEnum `mandatory:"true" json:"updateType"`

	// Indicates whether the update can be installed using OSMS.
	Installable IsEligibleForInstallationEnum `mandatory:"false" json:"installable,omitempty"`

	// Indicates whether a reboot may be required to complete installation of this update.
	IsRebootRequiredForInstallation *bool `mandatory:"false" json:"isRebootRequiredForInstallation"`
}

func (m WindowsUpdateSummary) String() string {
	return common.PointerString(m)
}
