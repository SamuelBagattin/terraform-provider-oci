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

// MfaTotpToken Totp token for MFA
type MfaTotpToken struct {

	// The Totp token for MFA.
	TotpToken *string `mandatory:"false" json:"totpToken"`
}

func (m MfaTotpToken) String() string {
	return common.PointerString(m)
}
