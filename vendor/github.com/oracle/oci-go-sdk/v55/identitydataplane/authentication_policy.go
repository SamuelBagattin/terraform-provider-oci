// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Service
//
// API for the Identity Dataplane
//

package identitydataplane

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// AuthenticationPolicy Authentication policy, currently set for the given compartment
type AuthenticationPolicy struct {

	// Password policy.
	PasswordPolicy *PasswordPolicy `mandatory:"false" json:"passwordPolicy"`

	// Compartment OCID.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m AuthenticationPolicy) String() string {
	return common.PointerString(m)
}
