// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Visual Builder API
//
// Oracle Visual Builder enables developers to quickly build web and mobile applications. With a visual development environment that makes it easy to connect to Oracle data and third-party REST services, developers can build modern, consumer-grade applications in a fraction of the time it would take in other tools.
// The Visual Builder Instance Management API allows users to create and manage a Visual Builder instance.
//

package visualbuilder

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// VbInstanceSummaryCollection Result of a VbInstance Summary request. Contains VbInstanceSummary items.
type VbInstanceSummaryCollection struct {

	// The collection of VbInstanceSummary objects.
	Items []VbInstanceSummary `mandatory:"true" json:"items"`
}

func (m VbInstanceSummaryCollection) String() string {
	return common.PointerString(m)
}
