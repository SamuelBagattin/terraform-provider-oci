// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// UserSummary The summary of a specific User.
type UserSummary struct {

	// The name of the User.
	Name *string `mandatory:"true" json:"name"`

	// The account status of the User
	Status UserSummaryStatusEnum `mandatory:"true" json:"status"`

	// The default tablespace for data.
	DefaultTablespace *string `mandatory:"true" json:"defaultTablespace"`

	// The name of the default tablespace for temporary tables or the name of a tablespace group.
	TempTablespace *string `mandatory:"true" json:"tempTablespace"`

	// The User creation date.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The User resource profile name.
	Profile *string `mandatory:"true" json:"profile"`

	// The date of expiration of the account
	TimeExpiring *common.SDKTime `mandatory:"false" json:"timeExpiring"`
}

func (m UserSummary) String() string {
	return common.PointerString(m)
}

// UserSummaryStatusEnum Enum with underlying type: string
type UserSummaryStatusEnum string

// Set of constants representing the allowable values for UserSummaryStatusEnum
const (
	UserSummaryStatusOpen                          UserSummaryStatusEnum = "OPEN"
	UserSummaryStatusExpired                       UserSummaryStatusEnum = "EXPIRED"
	UserSummaryStatusExpiredGrace                  UserSummaryStatusEnum = "EXPIRED_GRACE"
	UserSummaryStatusLocked                        UserSummaryStatusEnum = "LOCKED"
	UserSummaryStatusLockedTimed                   UserSummaryStatusEnum = "LOCKED_TIMED"
	UserSummaryStatusExpiredAndLocked              UserSummaryStatusEnum = "EXPIRED_AND_LOCKED"
	UserSummaryStatusExpiredGraceAndLocked         UserSummaryStatusEnum = "EXPIRED_GRACE_AND_LOCKED"
	UserSummaryStatusExpiredAndLockedTimed         UserSummaryStatusEnum = "EXPIRED_AND_LOCKED_TIMED"
	UserSummaryStatusExpiredGraceAndLockedTimed    UserSummaryStatusEnum = "EXPIRED_GRACE_AND_LOCKED_TIMED"
	UserSummaryStatusOpenAndInRollover             UserSummaryStatusEnum = "OPEN_AND_IN_ROLLOVER"
	UserSummaryStatusExpiredAndInRollover          UserSummaryStatusEnum = "EXPIRED_AND_IN_ROLLOVER"
	UserSummaryStatusLockedAndInRollover           UserSummaryStatusEnum = "LOCKED_AND_IN_ROLLOVER"
	UserSummaryStatusExpiredAndLockedAndInRollover UserSummaryStatusEnum = "EXPIRED_AND_LOCKED_AND_IN_ROLLOVER"
	UserSummaryStatusLockedTimedAndInRollover      UserSummaryStatusEnum = "LOCKED_TIMED_AND_IN_ROLLOVER"
	UserSummaryStatusExpiredAndLockedTimedAndInRol UserSummaryStatusEnum = "EXPIRED_AND_LOCKED_TIMED_AND_IN_ROL"
)

var mappingUserSummaryStatus = map[string]UserSummaryStatusEnum{
	"OPEN":                                UserSummaryStatusOpen,
	"EXPIRED":                             UserSummaryStatusExpired,
	"EXPIRED_GRACE":                       UserSummaryStatusExpiredGrace,
	"LOCKED":                              UserSummaryStatusLocked,
	"LOCKED_TIMED":                        UserSummaryStatusLockedTimed,
	"EXPIRED_AND_LOCKED":                  UserSummaryStatusExpiredAndLocked,
	"EXPIRED_GRACE_AND_LOCKED":            UserSummaryStatusExpiredGraceAndLocked,
	"EXPIRED_AND_LOCKED_TIMED":            UserSummaryStatusExpiredAndLockedTimed,
	"EXPIRED_GRACE_AND_LOCKED_TIMED":      UserSummaryStatusExpiredGraceAndLockedTimed,
	"OPEN_AND_IN_ROLLOVER":                UserSummaryStatusOpenAndInRollover,
	"EXPIRED_AND_IN_ROLLOVER":             UserSummaryStatusExpiredAndInRollover,
	"LOCKED_AND_IN_ROLLOVER":              UserSummaryStatusLockedAndInRollover,
	"EXPIRED_AND_LOCKED_AND_IN_ROLLOVER":  UserSummaryStatusExpiredAndLockedAndInRollover,
	"LOCKED_TIMED_AND_IN_ROLLOVER":        UserSummaryStatusLockedTimedAndInRollover,
	"EXPIRED_AND_LOCKED_TIMED_AND_IN_ROL": UserSummaryStatusExpiredAndLockedTimedAndInRol,
}

// GetUserSummaryStatusEnumValues Enumerates the set of values for UserSummaryStatusEnum
func GetUserSummaryStatusEnumValues() []UserSummaryStatusEnum {
	values := make([]UserSummaryStatusEnum, 0)
	for _, v := range mappingUserSummaryStatus {
		values = append(values, v)
	}
	return values
}
