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

// EncryptionInTransitTypeEnum Enum with underlying type: string
type EncryptionInTransitTypeEnum string

// Set of constants representing the allowable values for EncryptionInTransitTypeEnum
const (
	EncryptionInTransitTypeNone                  EncryptionInTransitTypeEnum = "NONE"
	EncryptionInTransitTypeBmEncryptionInTransit EncryptionInTransitTypeEnum = "BM_ENCRYPTION_IN_TRANSIT"
)

var mappingEncryptionInTransitType = map[string]EncryptionInTransitTypeEnum{
	"NONE":                     EncryptionInTransitTypeNone,
	"BM_ENCRYPTION_IN_TRANSIT": EncryptionInTransitTypeBmEncryptionInTransit,
}

// GetEncryptionInTransitTypeEnumValues Enumerates the set of values for EncryptionInTransitTypeEnum
func GetEncryptionInTransitTypeEnumValues() []EncryptionInTransitTypeEnum {
	values := make([]EncryptionInTransitTypeEnum, 0)
	for _, v := range mappingEncryptionInTransitType {
		values = append(values, v)
	}
	return values
}
