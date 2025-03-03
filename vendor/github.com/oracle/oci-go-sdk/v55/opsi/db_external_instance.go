// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// DbExternalInstance Configuration parameters defined for external databases instance level.
type DbExternalInstance struct {

	// Name of the database instance.
	InstanceName *string `mandatory:"true" json:"instanceName"`

	// Host name of the database instance.
	HostName *string `mandatory:"true" json:"hostName"`

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"false" json:"timeCollected"`

	// Total number of CPUs allocated for the host.
	CpuCount *int `mandatory:"false" json:"cpuCount"`

	// Total amount of usable Physical RAM Memory available in gigabytes.
	HostMemoryCapacity *float64 `mandatory:"false" json:"hostMemoryCapacity"`

	// Database version.
	Version *string `mandatory:"false" json:"version"`

	// Indicates whether the instance is mounted in cluster database mode (YES) or not (NO).
	Parallel *string `mandatory:"false" json:"parallel"`

	// Role (permissions) of the database instance.
	InstanceRole *string `mandatory:"false" json:"instanceRole"`

	// Indicates if logins are allowed or restricted.
	Logins *string `mandatory:"false" json:"logins"`

	// Status of the database.
	DatabaseStatus *string `mandatory:"false" json:"databaseStatus"`

	// Status of the instance.
	Status *string `mandatory:"false" json:"status"`

	// The edition of the database.
	Edition *string `mandatory:"false" json:"edition"`

	// Start up time of the database instance.
	StartupTime *common.SDKTime `mandatory:"false" json:"startupTime"`
}

//GetTimeCollected returns TimeCollected
func (m DbExternalInstance) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m DbExternalInstance) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DbExternalInstance) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbExternalInstance DbExternalInstance
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeDbExternalInstance
	}{
		"DB_EXTERNAL_INSTANCE",
		(MarshalTypeDbExternalInstance)(m),
	}

	return json.Marshal(&s)
}
