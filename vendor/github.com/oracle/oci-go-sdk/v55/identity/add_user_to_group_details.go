// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// AddUserToGroupDetails The representation of AddUserToGroupDetails
type AddUserToGroupDetails struct {

	// The OCID of the user.
	UserId *string `mandatory:"true" json:"userId"`

	// The OCID of the group.
	GroupId *string `mandatory:"true" json:"groupId"`
}

func (m AddUserToGroupDetails) String() string {
	return common.PointerString(m)
}
