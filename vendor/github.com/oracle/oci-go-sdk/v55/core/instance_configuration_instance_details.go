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

// InstanceConfigurationInstanceDetails The representation of InstanceConfigurationInstanceDetails
type InstanceConfigurationInstanceDetails interface {
}

type instanceconfigurationinstancedetails struct {
	JsonData     []byte
	InstanceType string `json:"instanceType"`
}

// UnmarshalJSON unmarshals json
func (m *instanceconfigurationinstancedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinstanceconfigurationinstancedetails instanceconfigurationinstancedetails
	s := struct {
		Model Unmarshalerinstanceconfigurationinstancedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.InstanceType = s.Model.InstanceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *instanceconfigurationinstancedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.InstanceType {
	case "compute":
		mm := ComputeInstanceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m instanceconfigurationinstancedetails) String() string {
	return common.PointerString(m)
}
