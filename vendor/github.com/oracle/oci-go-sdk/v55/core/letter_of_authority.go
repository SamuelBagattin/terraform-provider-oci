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

// LetterOfAuthority The Letter of Authority for the cross-connect. You must submit this letter when
// requesting cabling for the cross-connect at the FastConnect location.
type LetterOfAuthority struct {

	// The name of the entity authorized by this Letter of Authority.
	AuthorizedEntityName *string `mandatory:"false" json:"authorizedEntityName"`

	// The type of cross-connect fiber, termination, and optical specification.
	CircuitType LetterOfAuthorityCircuitTypeEnum `mandatory:"false" json:"circuitType,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cross-connect.
	CrossConnectId *string `mandatory:"false" json:"crossConnectId"`

	// The address of the FastConnect location.
	FacilityLocation *string `mandatory:"false" json:"facilityLocation"`

	// The meet-me room port for this cross-connect.
	PortName *string `mandatory:"false" json:"portName"`

	// The date and time when the Letter of Authority expires, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`

	// The date and time the Letter of Authority was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeIssued *common.SDKTime `mandatory:"false" json:"timeIssued"`
}

func (m LetterOfAuthority) String() string {
	return common.PointerString(m)
}

// LetterOfAuthorityCircuitTypeEnum Enum with underlying type: string
type LetterOfAuthorityCircuitTypeEnum string

// Set of constants representing the allowable values for LetterOfAuthorityCircuitTypeEnum
const (
	LetterOfAuthorityCircuitTypeLc LetterOfAuthorityCircuitTypeEnum = "Single_mode_LC"
	LetterOfAuthorityCircuitTypeSc LetterOfAuthorityCircuitTypeEnum = "Single_mode_SC"
)

var mappingLetterOfAuthorityCircuitType = map[string]LetterOfAuthorityCircuitTypeEnum{
	"Single_mode_LC": LetterOfAuthorityCircuitTypeLc,
	"Single_mode_SC": LetterOfAuthorityCircuitTypeSc,
}

// GetLetterOfAuthorityCircuitTypeEnumValues Enumerates the set of values for LetterOfAuthorityCircuitTypeEnum
func GetLetterOfAuthorityCircuitTypeEnumValues() []LetterOfAuthorityCircuitTypeEnum {
	values := make([]LetterOfAuthorityCircuitTypeEnum, 0)
	for _, v := range mappingLetterOfAuthorityCircuitType {
		values = append(values, v)
	}
	return values
}
