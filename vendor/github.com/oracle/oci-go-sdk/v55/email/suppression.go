// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// API for the Email Delivery service. Use this API to send high-volume, application-generated
// emails. For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//
// **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
// If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// Suppression The full information representing an email suppression.
type Suppression struct {

	// The OCID of the compartment to contain the suppression. Since
	// suppressions are at the customer level, this must be the tenancy
	// OCID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Email address of the suppression.
	EmailAddress *string `mandatory:"true" json:"emailAddress"`

	// The unique ID of the suppression.
	Id *string `mandatory:"true" json:"id"`

	// The reason that the email address was suppressed. For more information on the types of bounces, see Suppression List (https://docs.cloud.oracle.com/Content/Email/Concepts/overview.htm#components).
	Reason SuppressionReasonEnum `mandatory:"false" json:"reason,omitempty"`

	// The date and time the suppression was added in "YYYY-MM-ddThh:mmZ"
	// format with a Z offset, as defined by RFC 3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last date and time the suppression prevented submission
	// in "YYYY-MM-ddThh:mmZ"
	// format with a Z offset, as defined by RFC 3339.
	TimeLastSuppressed *common.SDKTime `mandatory:"false" json:"timeLastSuppressed"`

	// The value of the Message-ID header from the email that triggered a suppression.
	// This value is as defined in RFC 5322 section 3.6.4, excluding angle-brackets.
	// Not provided for all types of suppressions.
	MessageId *string `mandatory:"false" json:"messageId"`

	// The specific error message returned by a system that resulted in the suppression.
	// This message is usually an SMTP error code with additional descriptive text.
	// Not provided for all types of suppressions.
	ErrorDetail *string `mandatory:"false" json:"errorDetail"`

	// DNS name of the source of the error that caused the suppression.
	// Will be set to either the remote-mta or reporting-mta field from a delivery status notification (RFC 3464) when available.
	// Not provided for all types of suppressions, and not always known.
	// Note: Most SMTP errors that cause suppressions come from software run by email receiving systems rather than from OCI email delivery itself.
	ErrorSource *string `mandatory:"false" json:"errorSource"`
}

func (m Suppression) String() string {
	return common.PointerString(m)
}

// SuppressionReasonEnum Enum with underlying type: string
type SuppressionReasonEnum string

// Set of constants representing the allowable values for SuppressionReasonEnum
const (
	SuppressionReasonUnknown     SuppressionReasonEnum = "UNKNOWN"
	SuppressionReasonHardbounce  SuppressionReasonEnum = "HARDBOUNCE"
	SuppressionReasonComplaint   SuppressionReasonEnum = "COMPLAINT"
	SuppressionReasonManual      SuppressionReasonEnum = "MANUAL"
	SuppressionReasonSoftbounce  SuppressionReasonEnum = "SOFTBOUNCE"
	SuppressionReasonUnsubscribe SuppressionReasonEnum = "UNSUBSCRIBE"
)

var mappingSuppressionReason = map[string]SuppressionReasonEnum{
	"UNKNOWN":     SuppressionReasonUnknown,
	"HARDBOUNCE":  SuppressionReasonHardbounce,
	"COMPLAINT":   SuppressionReasonComplaint,
	"MANUAL":      SuppressionReasonManual,
	"SOFTBOUNCE":  SuppressionReasonSoftbounce,
	"UNSUBSCRIBE": SuppressionReasonUnsubscribe,
}

// GetSuppressionReasonEnumValues Enumerates the set of values for SuppressionReasonEnum
func GetSuppressionReasonEnumValues() []SuppressionReasonEnum {
	values := make([]SuppressionReasonEnum, 0)
	for _, v := range mappingSuppressionReason {
		values = append(values, v)
	}
	return values
}
