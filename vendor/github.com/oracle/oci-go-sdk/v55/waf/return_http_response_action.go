// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// ReturnHttpResponseAction An object that represents an action which returns a defined HTTP response.
type ReturnHttpResponseAction struct {

	// Action name. Can be used to reference the action.
	Name *string `mandatory:"true" json:"name"`

	// Response code.
	// The following response codes are valid values for this property:
	// * 2xx
	//   200 OK
	//   201 Created
	//   202 Accepted
	//   206 Partial Content
	// * 3xx
	//   300 Multiple Choices
	//   301 Moved Permanently
	//   302 Found
	//   303 See Other
	//   307 Temporary Redirect
	// * 4xx
	//   400 Bad Request
	//   401 Unauthorized
	//   403 Forbidden
	//   404 Not Found
	//   405 Method Not Allowed
	//   408 Request Timeout
	//   409 Conflict
	//   411 Length Required
	//   412 Precondition Failed
	//   413 Payload Too Large
	//   414 URI Too Long
	//   415 Unsupported Media Type
	//   416 Range Not Satisfiable
	//   422 Unprocessable Entity
	//   494 Request Header Too Large
	//   495 Cert Error
	//   496 No Cert
	//   497 HTTP to HTTPS
	// * 5xx
	//   500 Internal Server Error
	//   501 Not Implemented
	//   502 Bad Gateway
	//   503 Service Unavailable
	//   504 Gateway Timeout
	//   507 Insufficient Storage
	// Example: `200`
	Code *int `mandatory:"true" json:"code"`

	// Adds headers defined in this array for HTTP response.
	// Hop-by-hop headers are not allowed to be set:
	// * Connection
	// * Keep-Alive
	// * Proxy-Authenticate
	// * Proxy-Authorization
	// * TE
	// * Trailer
	// * Transfer-Encoding
	// * Upgrade
	Headers []ResponseHeader `mandatory:"false" json:"headers"`

	Body HttpResponseBody `mandatory:"false" json:"body"`
}

//GetName returns Name
func (m ReturnHttpResponseAction) GetName() *string {
	return m.Name
}

func (m ReturnHttpResponseAction) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ReturnHttpResponseAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeReturnHttpResponseAction ReturnHttpResponseAction
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeReturnHttpResponseAction
	}{
		"RETURN_HTTP_RESPONSE",
		(MarshalTypeReturnHttpResponseAction)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ReturnHttpResponseAction) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Headers []ResponseHeader `json:"headers"`
		Body    httpresponsebody `json:"body"`
		Name    *string          `json:"name"`
		Code    *int             `json:"code"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Headers = make([]ResponseHeader, len(model.Headers))
	for i, n := range model.Headers {
		m.Headers[i] = n
	}

	nn, e = model.Body.UnmarshalPolymorphicJSON(model.Body.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Body = nn.(HttpResponseBody)
	} else {
		m.Body = nil
	}

	m.Name = model.Name

	m.Code = model.Code

	return
}
