// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// CreateRuleSetDetails A named set of rules to add to the load balancer.
type CreateRuleSetDetails struct {

	// The name for this set of rules. It must be unique and it cannot be changed. Avoid entering
	// confidential information.
	// Example: `example_rule_set`
	Name *string `mandatory:"true" json:"name"`

	// An array of rules that compose the rule set.
	Items []Rule `mandatory:"true" json:"items"`
}

func (m CreateRuleSetDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateRuleSetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name  *string `json:"name"`
		Items []rule  `json:"items"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Name = model.Name

	m.Items = make([]Rule, len(model.Items))
	for i, n := range model.Items {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Items[i] = nn.(Rule)
		} else {
			m.Items[i] = nil
		}
	}

	return
}
