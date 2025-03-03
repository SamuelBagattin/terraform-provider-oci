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

// MultipathDevice Secondary multipath device, it uses the charUsername and chapSecret from primary volume attachment
type MultipathDevice struct {

	// The volume's iSCSI IP address.
	// Example: `169.254.2.2`
	Ipv4 *string `mandatory:"true" json:"ipv4"`

	// The target volume's iSCSI Qualified Name in the format defined
	// by RFC 3720 (https://tools.ietf.org/html/rfc3720#page-32).
	// Example: `iqn.2015-12.com.oracleiaas:40b7ee03-883f-46c6-a951-63d2841d2195`
	Iqn *string `mandatory:"true" json:"iqn"`

	// The volume's iSCSI port, usually port 860 or 3260.
	// Example: `3260`
	Port *int `mandatory:"false" json:"port"`
}

func (m MultipathDevice) String() string {
	return common.PointerString(m)
}
