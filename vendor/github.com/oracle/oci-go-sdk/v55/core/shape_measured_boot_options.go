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

// ShapeMeasuredBootOptions Configuration options for the Measured Boot feature.
type ShapeMeasuredBootOptions struct {

	// Boolean values that indicate whether the Measured Boot feature can be enabled or disabled.
	AllowedValues []bool `mandatory:"false" json:"allowedValues"`

	// Whether the Measured Boot feature is enabled by default.
	IsDefaultEnabled *bool `mandatory:"false" json:"isDefaultEnabled"`
}

func (m ShapeMeasuredBootOptions) String() string {
	return common.PointerString(m)
}
