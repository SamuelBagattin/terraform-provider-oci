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

// ResourceTypesEnum Enum with underlying type: string
type ResourceTypesEnum string

// Set of constants representing the allowable values for ResourceTypesEnum
const (
	ResourceTypesExacc                 ResourceTypesEnum = "EXACC"
	ResourceTypesExadatainfrastructure ResourceTypesEnum = "EXADATAINFRASTRUCTURE"
	ResourceTypesAutonomousvmcluster   ResourceTypesEnum = "AUTONOMOUSVMCLUSTER"
)

var mappingResourceTypes = map[string]ResourceTypesEnum{
	"EXACC":                 ResourceTypesExacc,
	"EXADATAINFRASTRUCTURE": ResourceTypesExadatainfrastructure,
	"AUTONOMOUSVMCLUSTER":   ResourceTypesAutonomousvmcluster,
}

// GetResourceTypesEnumValues Enumerates the set of values for ResourceTypesEnum
func GetResourceTypesEnumValues() []ResourceTypesEnum {
	values := make([]ResourceTypesEnum, 0)
	for _, v := range mappingResourceTypes {
		values = append(values, v)
	}
	return values
}
