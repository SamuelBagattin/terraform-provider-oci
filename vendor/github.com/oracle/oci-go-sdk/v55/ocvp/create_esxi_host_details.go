// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// CreateEsxiHostDetails Details of the ESXi host to add to the SDDC.
type CreateEsxiHostDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SDDC to add the
	// ESXi host to.
	SddcId *string `mandatory:"true" json:"sddcId"`

	// A descriptive name for the ESXi host. It's changeable.
	// Esxi Host name requirements are 1-16 character length limit, Must start with a letter, Must be English letters, numbers, - only, No repeating hyphens, Must be unique within the SDDC.
	// If this attribute is not specified, the SDDC's `instanceDisplayNamePrefix` attribute is used
	// to name and incrementally number the ESXi host. For example, if you're creating the fourth
	// ESXi host in the SDDC, and `instanceDisplayNamePrefix` is `MySDDC`, the host's display
	// name is `MySDDC-4`.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The billing option currently used by the ESXi host.
	// ListSupportedSkus.
	CurrentSku SkuEnum `mandatory:"false" json:"currentSku,omitempty"`

	// The billing option to switch to after the existing billing cycle ends.
	// If `nextSku` is null or empty, `currentSku` continues to the next billing cycle.
	// ListSupportedSkus.
	NextSku SkuEnum `mandatory:"false" json:"nextSku,omitempty"`

	// The availability domain to create the ESXi host in.
	// If keep empty, for AD-specific SDDC, new ESXi host will be created in the same availability domain;
	// for multi-AD SDDC, new ESXi host will be auto assigned to the next availability domain following evenly distribution strategy.
	ComputeAvailabilityDomain *string `mandatory:"false" json:"computeAvailabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the esxi host that
	// is failed. It is an optional param, when user supplies this param, new Esxi
	// Host will be created to replace the failed one, and failedEsxiHostId field
	// will be udpated in the newly created EsxiHost.
	FailedEsxiHostId *string `mandatory:"false" json:"failedEsxiHostId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateEsxiHostDetails) String() string {
	return common.PointerString(m)
}
