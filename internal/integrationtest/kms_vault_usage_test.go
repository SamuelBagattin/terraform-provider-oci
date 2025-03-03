// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	vaultUsageSingularDataSourceRepresentation = map[string]interface{}{
		"vault_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.id}`},
	}

	VaultUsageResourceConfig = KeyResourceDependencies
)

// issue-routing-tag: kms/default
func TestKmsVaultUsageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsVaultUsageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_kms_vault_usage.test_vault_usage"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_vault_usage", "test_vault_usage", acctest.Required, acctest.Create, vaultUsageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VaultUsageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_version_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_key_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_key_version_count"),
			),
		},
	})
}
