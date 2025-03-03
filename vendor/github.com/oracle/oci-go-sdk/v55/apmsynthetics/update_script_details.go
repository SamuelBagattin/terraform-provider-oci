// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// UpdateScriptDetails Details of the request body used to update a script.
// Only Side or JavaScript content types are supported and content should be in Side or JavaScript formats only.
type UpdateScriptDetails struct {

	// Unique name that can be edited. The name should not contain any confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Content type of script.
	ContentType ContentTypesEnum `mandatory:"false" json:"contentType,omitempty"`

	// The content of the script. It may contain custom-defined tags that can be used for setting dynamic parameters.
	// The format to set dynamic parameters is: `<ORAP><ON>param name</ON><OV>param value</OV><OS>isParamValueSecret(true/false)</OS></ORAP>`.
	// Param value and isParamValueSecret are optional, the default value for isParamValueSecret is false.
	// Examples:
	// With mandatory param name : `<ORAP><ON>param name</ON></ORAP>`
	// With parameter name and value : `<ORAP><ON>param name</ON><OV>param value</OV></ORAP>`
	// Note that the content is valid if it matches the given content type. For example, if the content type is SIDE, then the content should be in Side script format. If the content type is JS, then the content should be in JavaScript format.
	Content *string `mandatory:"false" json:"content"`

	// File name of uploaded script content.
	ContentFileName *string `mandatory:"false" json:"contentFileName"`

	// List of script parameters. Example: `[{"paramName": "userid", "paramValue":"testuser", "isSecret": false}]`
	Parameters []ScriptParameter `mandatory:"false" json:"parameters"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateScriptDetails) String() string {
	return common.PointerString(m)
}
