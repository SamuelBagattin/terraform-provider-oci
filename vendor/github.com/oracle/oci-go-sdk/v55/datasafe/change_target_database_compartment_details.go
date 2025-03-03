// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// ChangeTargetDatabaseCompartmentDetails The details used to change the compartment of a Data Safe target database.
type ChangeTargetDatabaseCompartmentDetails struct {

	// The OCID of the new compartment to where you want to move the Data Safe target database.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeTargetDatabaseCompartmentDetails) String() string {
	return common.PointerString(m)
}
