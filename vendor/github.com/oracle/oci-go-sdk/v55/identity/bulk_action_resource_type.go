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

// BulkActionResourceType The representation of BulkActionResourceType
type BulkActionResourceType struct {

	// The unique name of the resource-type.
	Name *string `mandatory:"true" json:"name"`

	// List of metadata keys required to identify a specific resource. Some resource-types require information besides an OCID to identify
	// a specific resource. For example, the resource-type `buckets` requires metadataKeys DeleteBucket.
	MetadataKeys []string `mandatory:"false" json:"metadataKeys"`
}

func (m BulkActionResourceType) String() string {
	return common.PointerString(m)
}
