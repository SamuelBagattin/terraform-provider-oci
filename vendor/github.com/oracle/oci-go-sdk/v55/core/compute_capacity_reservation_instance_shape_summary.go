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

// ComputeCapacityReservationInstanceShapeSummary An available shape used to launch instances in a compute capacity reservation.
type ComputeCapacityReservationInstanceShapeSummary struct {

	// The shape's availability domain.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The name of the available shape used to launch instances in a compute capacity reservation.
	InstanceShape *string `mandatory:"true" json:"instanceShape"`
}

func (m ComputeCapacityReservationInstanceShapeSummary) String() string {
	return common.PointerString(m)
}
