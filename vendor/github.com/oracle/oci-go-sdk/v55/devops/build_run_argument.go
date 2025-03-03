// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// BuildRunArgument Values for pipeline parameters to be supplied at the time of running the build.
type BuildRunArgument struct {

	// Name of the parameter (case-sensitive). Parameter name must be ^[a-zA-Z][a-zA-Z_0-9]*$.
	// Example: 'Build_Pipeline_param' is not same as 'build_pipeline_Param'
	Name *string `mandatory:"true" json:"name"`

	// Value of the argument.
	Value *string `mandatory:"true" json:"value"`
}

func (m BuildRunArgument) String() string {
	return common.PointerString(m)
}
