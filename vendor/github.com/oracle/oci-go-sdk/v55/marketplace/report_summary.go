// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// ReportSummary The model of a single report.
type ReportSummary struct {

	// The type of report.
	ReportType *string `mandatory:"true" json:"reportType"`

	// The date of the report.
	Date *common.SDKTime `mandatory:"true" json:"date"`

	// The columns in the report.
	Columns []string `mandatory:"true" json:"columns"`

	// The contents of the report in comma-separated values (CSV) file format.
	Content *string `mandatory:"true" json:"content"`
}

func (m ReportSummary) String() string {
	return common.PointerString(m)
}
