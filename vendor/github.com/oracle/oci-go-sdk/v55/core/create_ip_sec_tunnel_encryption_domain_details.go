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

// CreateIpSecTunnelEncryptionDomainDetails Request to enable a multi-encryption domain policy on the IPSec tunnel.
// There can't be more than 50 security associations in use at one time. See Encryption domain for policy-based
// tunnels (https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/ipsecencryptiondomains.htm#spi_policy_based_tunnel) for more.
type CreateIpSecTunnelEncryptionDomainDetails struct {

	// Lists IPv4 or IPv6-enabled subnets in your Oracle tenancy.
	OracleTrafficSelector []string `mandatory:"false" json:"oracleTrafficSelector"`

	// Lists IPv4 or IPv6-enabled subnets in your on-premises network.
	CpeTrafficSelector []string `mandatory:"false" json:"cpeTrafficSelector"`
}

func (m CreateIpSecTunnelEncryptionDomainDetails) String() string {
	return common.PointerString(m)
}
