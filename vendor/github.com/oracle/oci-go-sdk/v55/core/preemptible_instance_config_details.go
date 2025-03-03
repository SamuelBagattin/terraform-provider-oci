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

// PreemptibleInstanceConfigDetails Configuration options for preemptible instances.
type PreemptibleInstanceConfigDetails struct {
	PreemptionAction PreemptionAction `mandatory:"true" json:"preemptionAction"`
}

func (m PreemptibleInstanceConfigDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *PreemptibleInstanceConfigDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		PreemptionAction preemptionaction `json:"preemptionAction"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.PreemptionAction.UnmarshalPolymorphicJSON(model.PreemptionAction.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PreemptionAction = nn.(PreemptionAction)
	} else {
		m.PreemptionAction = nil
	}

	return
}
