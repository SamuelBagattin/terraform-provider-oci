// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

// AccessRequestSeveritiesEnum Enum with underlying type: string
type AccessRequestSeveritiesEnum string

// Set of constants representing the allowable values for AccessRequestSeveritiesEnum
const (
	AccessRequestSeveritiesS1 AccessRequestSeveritiesEnum = "S1"
	AccessRequestSeveritiesS2 AccessRequestSeveritiesEnum = "S2"
	AccessRequestSeveritiesS3 AccessRequestSeveritiesEnum = "S3"
	AccessRequestSeveritiesS4 AccessRequestSeveritiesEnum = "S4"
)

var mappingAccessRequestSeverities = map[string]AccessRequestSeveritiesEnum{
	"S1": AccessRequestSeveritiesS1,
	"S2": AccessRequestSeveritiesS2,
	"S3": AccessRequestSeveritiesS3,
	"S4": AccessRequestSeveritiesS4,
}

// GetAccessRequestSeveritiesEnumValues Enumerates the set of values for AccessRequestSeveritiesEnum
func GetAccessRequestSeveritiesEnumValues() []AccessRequestSeveritiesEnum {
	values := make([]AccessRequestSeveritiesEnum, 0)
	for _, v := range mappingAccessRequestSeverities {
		values = append(values, v)
	}
	return values
}
