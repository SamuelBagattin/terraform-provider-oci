// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NoSQL Database API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// QueryResultCollection The result of a query.
type QueryResultCollection struct {

	// Array of objects representing query results.
	Items []map[string]interface{} `mandatory:"false" json:"items"`

	Usage *RequestUsage `mandatory:"false" json:"usage"`
}

func (m QueryResultCollection) String() string {
	return common.PointerString(m)
}
