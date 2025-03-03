// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_apm_config "github.com/oracle/oci-go-sdk/v55/apmconfig"

	oci_common "github.com/oracle/oci-go-sdk/v55/common"
)

func init() {
	RegisterOracleClient("oci_apm_config.ConfigClient", &OracleClient{InitClientFn: initApmconfigConfigClient})
}

func initApmconfigConfigClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apm_config.NewConfigClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) ConfigClient() *oci_apm_config.ConfigClient {
	return m.GetClient("oci_apm_config.ConfigClient").(*oci_apm_config.ConfigClient)
}
