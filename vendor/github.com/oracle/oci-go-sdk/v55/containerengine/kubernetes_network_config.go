// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// KubernetesNetworkConfig The properties that define the network configuration for Kubernetes.
type KubernetesNetworkConfig struct {

	// The CIDR block for Kubernetes pods. Optional, defaults to 10.244.0.0/16.
	PodsCidr *string `mandatory:"false" json:"podsCidr"`

	// The CIDR block for Kubernetes services. Optional, defaults to 10.96.0.0/16.
	ServicesCidr *string `mandatory:"false" json:"servicesCidr"`
}

func (m KubernetesNetworkConfig) String() string {
	return common.PointerString(m)
}
