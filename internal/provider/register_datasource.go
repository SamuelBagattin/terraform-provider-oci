package provider

import (
	"github.com/terraform-providers/terraform-provider-oci/internal/service/audit"
	"github.com/terraform-providers/terraform-provider-oci/internal/service/budget"
	tf_identity "github.com/terraform-providers/terraform-provider-oci/internal/service/identity"
	tf_kms "github.com/terraform-providers/terraform-provider-oci/internal/service/kms"
)

func init() {
	RegisterDatasource("oci_identity_smtp_credentials", tf_identity.IdentitySmtpCredentialsDataSource())
	RegisterDatasource("oci_budget_alert_rule", budget.BudgetAlertRuleDataSource())
	RegisterDatasource("oci_budget_budget", budget.BudgetBudgetDataSource())
	RegisterDatasource("oci_budget_alert_rules", budget.BudgetAlertRulesDataSource())
	RegisterDatasource("oci_budget_budgets", budget.BudgetBudgetsDataSource())
	RegisterDatasource("oci_identity_compartment", tf_identity.IdentityCompartmentDataSource())
	RegisterDatasource("oci_identity_compartments", tf_identity.IdentityCompartmentsDataSource())
	RegisterDatasource("oci_identity_network_source", tf_identity.IdentityNetworkSourceDataSource())
	RegisterDatasource("oci_identity_region_subscriptions", tf_identity.IdentityRegionSubscriptionsDataSource())
	RegisterDatasource("oci_identity_iam_work_requests", tf_identity.IdentityIamWorkRequestsDataSource())
	RegisterDatasource("oci_identity_groups", tf_identity.IdentityGroupsDataSource())
	RegisterDatasource("oci_identity_cost_tracking_tags", tf_identity.IdentityCostTrackingTagsDataSource())
	RegisterDatasource("oci_identity_users", tf_identity.IdentityUsersDataSource())
	RegisterDatasource("oci_identity_availability_domains", tf_identity.IdentityAvailabilityDomainsDataSource())
	RegisterDatasource("oci_identity_iam_work_request_logs", tf_identity.IdentityIamWorkRequestLogsDataSource())
	RegisterDatasource("oci_identity_tag_defaults", tf_identity.IdentityTagDefaultsDataSource())
	RegisterDatasource("oci_identity_user_group_memberships", tf_identity.IdentityUserGroupMembershipsDataSource())
	RegisterDatasource("oci_identity_swift_passwords", tf_identity.IdentitySwiftPasswordsDataSource())
	RegisterDatasource("oci_identity_group", tf_identity.IdentityGroupDataSource())
	RegisterDatasource("oci_identity_api_keys", tf_identity.IdentityApiKeysDataSource())
	RegisterDatasource("oci_identity_domain", tf_identity.IdentityDomainDataSource())
	RegisterDatasource("oci_identity_tag_default", tf_identity.IdentityTagDefaultDataSource())
	RegisterDatasource("oci_identity_availability_domain", tf_identity.IdentityAvailabilityDomainDataSource())
	RegisterDatasource("oci_identity_customer_secret_keys", tf_identity.IdentityCustomerSecretKeysDataSource())
	RegisterDatasource("oci_identity_iam_work_request", tf_identity.IdentityIamWorkRequestDataSource())
	RegisterDatasource("oci_identity_iam_work_request_errors", tf_identity.IdentityIamWorkRequestErrorsDataSource())
	RegisterDatasource("oci_identity_allowed_domain_license_types", tf_identity.IdentityAllowedDomainLicenseTypesDataSource())
	RegisterDatasource("oci_identity_idp_group_mappings", tf_identity.IdentityIdpGroupMappingsDataSource())
	RegisterDatasource("oci_identity_dynamic_groups", tf_identity.IdentityDynamicGroupsDataSource())
	RegisterDatasource("oci_identity_tenancy", tf_identity.IdentityTenancyDataSource())
	RegisterDatasource("oci_identity_auth_tokens", tf_identity.IdentityAuthTokensDataSource())
	RegisterDatasource("oci_identity_user", tf_identity.IdentityUserDataSource())
	RegisterDatasource("oci_identity_network_sources", tf_identity.IdentityNetworkSourcesDataSource())
	RegisterDatasource("oci_identity_tags", tf_identity.IdentityTagsDataSource())
	RegisterDatasource("oci_identity_policies", tf_identity.IdentityPoliciesDataSource())
	RegisterDatasource("oci_identity_authentication_policy", tf_identity.IdentityAuthenticationPolicyDataSource())
	RegisterDatasource("oci_identity_regions", tf_identity.IdentityRegionsDataSource())
	RegisterDatasource("oci_identity_tag", tf_identity.IdentityTagDataSource())
	RegisterDatasource("oci_identity_db_credentials", tf_identity.IdentityDbCredentialsDataSource())
	RegisterDatasource("oci_identity_identity_providers", tf_identity.IdentityIdentityProvidersDataSource())
	RegisterDatasource("oci_identity_identity_provider_groups", tf_identity.IdentityIdentityProviderGroupsDataSource())
	RegisterDatasource("oci_identity_tag_namespaces", tf_identity.IdentityTagNamespacesDataSource())
	RegisterDatasource("oci_identity_domains", tf_identity.IdentityDomainsDataSource())
	RegisterDatasource("oci_identity_ui_password", tf_identity.IdentityUiPasswordDataSource())
	RegisterDatasource("oci_identity_fault_domains", tf_identity.IdentityFaultDomainsDataSource())
	RegisterDatasource("oci_audit_configuration", audit.AuditConfigurationDataSource())
	RegisterDatasource("oci_audit_events", audit.AuditAuditEventsDataSource())

	//kms service
	RegisterDatasource("oci_kms_key", tf_kms.KmsKeyDataSource())
	RegisterDatasource("oci_kms_replication_status", tf_kms.KmsReplicationStatusDataSource())
	RegisterDatasource("oci_kms_vault_usage", tf_kms.KmsVaultUsageDataSource())
	RegisterDatasource("oci_kms_keys", tf_kms.KmsKeysDataSource())
	RegisterDatasource("oci_kms_key_version", tf_kms.KmsKeyVersionDataSource())
	RegisterDatasource("oci_kms_key_versions", tf_kms.KmsKeyVersionsDataSource())
	RegisterDatasource("oci_kms_decrypted_data", tf_kms.KmsDecryptedDataDataSource())
	RegisterDatasource("oci_kms_vaults", tf_kms.KmsVaultsDataSource())
	RegisterDatasource("oci_kms_encrypted_data", tf_kms.KmsEncryptedDataDataSource())
	RegisterDatasource("oci_kms_vault", tf_kms.KmsVaultDataSource())
	RegisterDatasource("oci_kms_vault_replicas", tf_kms.KmsVaultReplicasDataSource())
}
