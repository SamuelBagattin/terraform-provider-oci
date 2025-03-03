// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// CaBundle Reference to the CA bundle that should be used on the gateway
type CaBundle interface {
}

type cabundle struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *cabundle) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercabundle cabundle
	s := struct {
		Model Unmarshalercabundle
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *cabundle) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "CA_BUNDLE":
		mm := CertificatesCaBundle{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CERTIFICATE_AUTHORITY":
		mm := CertificatesCertificateAuthority{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m cabundle) String() string {
	return common.PointerString(m)
}

// CaBundleTypeEnum Enum with underlying type: string
type CaBundleTypeEnum string

// Set of constants representing the allowable values for CaBundleTypeEnum
const (
	CaBundleTypeCaBundle             CaBundleTypeEnum = "CA_BUNDLE"
	CaBundleTypeCertificateAuthority CaBundleTypeEnum = "CERTIFICATE_AUTHORITY"
)

var mappingCaBundleType = map[string]CaBundleTypeEnum{
	"CA_BUNDLE":             CaBundleTypeCaBundle,
	"CERTIFICATE_AUTHORITY": CaBundleTypeCertificateAuthority,
}

// GetCaBundleTypeEnumValues Enumerates the set of values for CaBundleTypeEnum
func GetCaBundleTypeEnumValues() []CaBundleTypeEnum {
	values := make([]CaBundleTypeEnum, 0)
	for _, v := range mappingCaBundleType {
		values = append(values, v)
	}
	return values
}
