CHANGELOG
=========

## HEAD (Unreleased)

#### Breaking Changes

- Remove env var support for auxiliary tenant IDs (ARM_AUXILIARY_TENANT_IDS). This option can still be set directly as a Provider argument or config.

#### Improvements

- Upgrade to Go 1.16
  [#630](https://github.com/pulumi/pulumi-azure-native/issues/630)
- New resource: `sql.DataMaskingPolicy`
  [#605](https://github.com/pulumi/pulumi-azure-native/issues/605)
- `Endpoint` and `CustomDomain` resources are auto-named with a random suffix.
  [#617](https://github.com/pulumi/pulumi-azure-native/issues/629)

#### Fixes

- Support migration from `azure-nextgen` with custom providers.
  [#617](https://github.com/pulumi/pulumi-azure-native/issues/617)

---

## 0.7.1 (2020-03-02)

Fixed the aliases for smooth migration from `azure-nextgen` to `azure-native`.

## 0.7.0 (2020-02-27)

### Azure NextGen is renamed to Azure-Native

See the [announcement](https://pulumi.com/blog/full-coverage-of-azure-resources-with-azure-native/) in the official Pulumi blog. 

#### Top-Level Resources

Top-Level Resources ([#169](https://github.com/pulumi/pulumi-azure-nextgen/issues/169)):

1. We pick a version for every Azure resource and place it to the top module of each resource provider. E.g., you will be able to create a storage account like

```
new Pulumi.AzureNextGen.Storage.StorageAccount("sa", ...);
// or
new azure_nextgen.storage.StorageAccount("sa", ...);
```

2. Resources that have preview versions but no stable versions are also available in the top module (they were not in Latest).

3. At 1.0, for a given resource, we will likely pick an API version that is the latest. If a later version is published by Microsoft and it contains breaking changes, we will NOT introduce this version to the top-level resource. It will be available under a versioned namespace only.

4. At 2.0, whenever it comes, we will bring all those postponed breaking changes by promoting top-level resources to their respective latest API versions.

5. We will encourage everyone to use top-level modules unless they have a reason not to do so (e.g., they need a newer version with its new features).

6. All "Latest" modules/namespaces become deprecated and will be removed before 1.0. You will be able to migrate to top-level resources without re-creating Azure resources.

Note: `storage.Blob` and `storage.StorageAccountStaticWebsite` are top-level resources only, their `latest` counterpart has been removed. 


#### Auto-naming ([#5](https://github.com/pulumi/pulumi-azure-nextgen/issues/5))

Auto-naming is applied for properties that are the last path parameters in their API endpoint path.

There is a twist compared to the normal Pulumi auto-naming:

- For top-level resources (e.g., resource groups, virtual networks, AKS), we append the same random suffix as in pulumi-azure.
- For child resources (e.g., subnets, databases, app slots), we simply copy the logical name to the physical name without any randomization.

Example:

```ts
const resourceGroup = new resources.ResourceGroup("rg");
// the actual resource would be named e.g. "rg8a43bc22"

const blob = new storage.Blob("wwwroot", {
    resourceGroupName: resourceGroup.name,
    accountName: storageAccount.name,
    containerName: storageContainer.name,
    source: new pulumi.asset.FileArchive("wwwroot"),
});
// blob name is simply "wwwroot"
```

#### Breaking Changes

- `web.WebApplicationSettings` renamed to `web.WebAppApplicationSettings` ([#282](https://github.com/pulumi/pulumi-azure-nextgen/issues/282))
- `authorization.GetClientConfig`, `authorization.GetClientToken`, `Storage.Blob`, `Storage.StorageAccountStaticWebsite` were removed from
  corresponding `latest` modules and are only available in the top-level modules
  
#### Bug Fixes

- Relax the enum check to accomodate irregularities like in [#294](https://github.com/pulumi/pulumi-azure-nextgen/issues/294)

## 0.6.1 (2020-02-11)

New features:

- Automatic location propagation ([#6](https://github.com/pulumi/pulumi-azure-nextgen/issues/6))
- New resource: `Storage.Blob` ([#13](https://github.com/pulumi/pulumi-azure-nextgen/issues/13))

## 0.6.0 (2020-02-02)

#### Breaking changes in SDKs

Limit the `latest` version to published versions of APIs ([#191](https://github.com/pulumi/pulumi-azure-nextgen/issues/191))

Several resources and functions were moved to another namespace to correctly match their Azure API classification.
([#230](https://github.com/pulumi/pulumi-azure-nextgen/issues/230)):

Moved from `Billing` to `CostManagement`:

- `GetCostAllocationRule` function
- `CostAllocationRule` resource
- `ReportByDepartment` resource
- `ReportByBillingAccount` resource
- `GetReportByDepartment` function
- `GetReportByBillingAccount` function
- `GetCostAllocationRule` function

Moved from `Compute` to `Automanage`:

- `ConfigurationProfileAssignment` resource
- `GetConfigurationProfileAssignment` function

Moved from `HybridCompute` to `GuestConfiguration`:

- `GuestConfigurationHCRPAssignment` resource
- `GuestConfigurationAssignment` resource
- `GetGuestConfigurationHCRPAssignment` function
- `GetGuestConfigurationAssignment` function

Moved from `Management` to `Authorization`:

- `GetPolicyDefinitionAtManagementGroup` function
- `GetPolicySetDefinitionAtManagementGroup` function
- `PolicyDefinitionAtManagementGroup` resource
- `PolicySetDefinitionAtManagementGroup` resource

Moved from `Management` to `Blueprint`:

- `Artifact` resource
- `Blueprint` resource
- `GetPublishedBlueprint` function
- `GetBlueprint` function
- `GetArtifact` function
- `PublishedBlueprint` resource

Moved from `Management` to `Insights`:

- `ManagementGroupDiagnosticSetting` resource
- `GetManagementGroupDiagnosticSetting` function

Moved from `Management` to `Resources`:

- `GetDeploymentAtManagementGroupScope` function
- `DeploymentAtManagementGroupScope` resource

Moved from `OperationalInsights` to `SecurityInsights`:

- `Incident` resource
- `GetIncident` function
- `GetDataConnector` function
- `GetBookmark` function
- `GetAlertRule` function
- `GetAction` function
- `DataConnector` resource
- `Bookmark` resource
- `AlertRule` resource
- `Action` resource

#### Fixes

- Fix unknown propagation in Update's preview ([#115](https://github.com/pulumi/pulumi-azure-nextgen/issues/115))
- Re-model Identity and Encryption properties of EventHubs/ServiceBus namespace ([#243](https://github.com/pulumi/pulumi-azure-nextgen/issues/243))
- Various fixes to Go SDK to address compilation issues due to name collisions. Conflicting types have been renamed for various resource providers but Go programs would not have been able to successfully compile with previous names anyway. ([#137](https://github.com/pulumi/pulumi-azure-nextgen/issues/137))
- Respect custom timeouts ([#252](https://github.com/pulumi/pulumi-azure-nextgen/issues/252))

#### New resources

- `Sql.DatabaseSecurityAlertPolicy` ([#257](https://github.com/pulumi/pulumi-azure-nextgen/issues/257))
- `Storage.StorageAccountStaticWebsite` ([#86](https://github.com/pulumi/pulumi-azure-nextgen/issues/86))
- `Web.WebAppAuthSettings`
- `Web.WebAppAuthSettingsSlot`
- `Web.WebAppAuthSettingsV2` ([#34](https://github.com/pulumi/pulumi-azure-nextgen/issues/34))
- `Web.WebAppAuthSettingsV2Slot`
- `Web.WebAppAzureStorageAccounts`
- `Web.WebAppAzureStorageAccountsSlot`
- `Web.WebAppBackupConfiguration`
- `Web.WebAppBackupConfigurationSlot`
- `Web.WebAppConnectionStrings`
- `Web.WebAppConnectionStringsSlot`
- `Web.WebAppMetadata`
- `Web.WebAppMetadataSlot`
- `Web.WebAppSitePushSettings`
- `Web.WebAppSitePushSettingsSlot`
- `Web.WebApplicationSettings`
- `Web.WebApplicationSettingsSlot`

#### New invoke functions

- `AppConfiguration.listConfigurationStoreKeys`
- `Relay.listHybridConnectionKeys`
- `ServiceBus.listQueueKeys`
- `ServiceBus.listTopicKeys` ([#248](https://github.com/pulumi/pulumi-azure-nextgen/issues/248))
- `Storage.listStorageAccountSAS`
- `Storage.listStorageAccountServiceSAS` ([#201](https://github.com/pulumi/pulumi-azure-nextgen/issues/201))

## 0.5.1 (2021-01-26)

Fixes:

- Empty arrays are now serialized to HTTP payloads as-is, fixing 'accessPolicies is not specified' ([#231](https://github.com/pulumi/pulumi-azure-nextgen/issues/231))
- Do not require replacements when updating to a version with default values ([#238](https://github.com/pulumi/pulumi-azure-nextgen/issues/238))

## 0.5.0 (2021-01-23)

Features:

- Respect default values as Azure API specs define. This is potentially breaking, as
  applying default values for a resource may require an update operation. You may
  need to adjust your code to avoid undesired updates.
  ([#183](https://github.com/pulumi/pulumi-azure-nextgen/issues/183))
- GetClientToken invoke to retrieve an OAuth token for the current auth context ([#207](https://github.com/pulumi/pulumi-azure-nextgen/issues/207))

Fixes:

- Generate all types along inheritance hierachy, not just one level ([#186](https://github.com/pulumi/pulumi-azure-nextgen/issues/186))
- Correct Resource ID ([#211](https://github.com/pulumi/pulumi-azure-nextgen/issues/211))
- Reduce the memory footprint ([#203](https://github.com/pulumi/pulumi-azure-nextgen/issues/203))
- Improved error messaging ([#213](https://github.com/pulumi/pulumi-azure-nextgen/issues/213))
- Fix failure when creating Certificate resource ([#212](https://github.com/pulumi/pulumi-azure-nextgen/issues/212))
- Creation of PrivateDnsZoneGroup ([#227](https://github.com/pulumi/pulumi-azure-nextgen/issues/227))

---

## 0.4.0 (2021-01-05)

Breaking changes:

- `int64` types are now represented as numbers (e.g., `double` in .NET) instead of `int`
- Nested union types replaced with `object` in C# ([#19](https://github.com/pulumi/pulumi-azure-nextgen/issues/19))

New invokes:

- `getClientConfig` to retrieve current authorization context parameters ([#107](https://github.com/pulumi/pulumi-azure-nextgen/issues/107))

Fixes:

- Improved experience with enum collections ([#173](https://github.com/pulumi/pulumi-azure-nextgen/issues/173))
- Correct unknown propagation during preview ([#115](https://github.com/pulumi/pulumi-azure-nextgen/issues/115))
- Importing IotHubResource ([#176](https://github.com/pulumi/pulumi-azure-nextgen/issues/176))
- Add `subscriptionId` as an explicit SDK parameter where API declares it as a method param ([#101](https://github.com/pulumi/pulumi-azure-nextgen/issues/101)

## 0.3.1 (2020-12-18)

New Features:

- Descriptions for resources on "latest" include a reference to the actual API version

New resources:

- API Management ProductApi, ProductGroup, GroupUser and other "link" resources ([#92](https://github.com/pulumi/pulumi-azure-nextgen/issues/92))

Fixes:

- Secrets leak in state ([#185](https://github.com/pulumi/pulumi-azure-nextgen/issues/185))
- Propogate expected value for constant properties in description ([#180](https://github.com/pulumi/pulumi-azure-nextgen/issues/180))

## 0.3.0 (2020-12-11)

New features:

- Enum support ([#106](https://github.com/pulumi/pulumi-azure-nextgen/issues/106))

New resources:

- Azure KeyVault Secrets and Keys ([#54](https://github.com/pulumi/pulumi-azure-nextgen/issues/54))
- WebAppDiagnosticLogsConfiguration ([#157](https://github.com/pulumi/pulumi-azure-nextgen/issues/157) and [#150](https://github.com/pulumi/pulumi-azure-nextgen/issues/150))
- App Insights ComponentCurrentBillingFeature ([#158](https://github.com/pulumi/pulumi-azure-nextgen/issues/158))
- SQL GeoBackupPolicy ([#143](https://github.com/pulumi/pulumi-azure-nextgen/issues/143))

Fixes:

- Retrieve ID in get* invokes ([#48](https://github.com/pulumi/pulumi-azure-nextgen/issues/48))

## 0.2.8 (2020-11-23)

- Add `userAssignedIdentities` to several resources ([#136](https://github.com/pulumi/pulumi-azure-nextgen/issues/136))

## 0.2.7 (2020-11-14)

- Fix creation of UserAssignedIdentity ([#100](https://github.com/pulumi/pulumi-azure-nextgen/issues/100))
- Fix creation of Budget ([#122](https://github.com/pulumi/pulumi-azure-nextgen/issues/122))
- Default to 'public' Azure environment ([#133](https://github.com/pulumi/pulumi-azure-nextgen/issues/133))

## 0.2.6 (2020-11-13)

Bug fixes:

- Support non-public Azure environments ([#123](https://github.com/pulumi/pulumi-azure-nextgen/issues/123))
- Fix an error in Go SDK's first-class provider config ([#126](https://github.com/pulumi/pulumi-azure-nextgen/issues/126))

## 0.2.5 (2020-11-06)

Bug fixes:

- API Connection creation ([#102](https://github.com/pulumi/pulumi-azure-nextgen/issues/102))
- WebAppSwiftVirtualNetworkConnection creation ([#94](https://github.com/pulumi/pulumi-azure-nextgen/issues/94))
- Fix plugin acquisition - requires Pulumi CLI 2.13.0+ ([#70](https://github.com/pulumi/pulumi-azure-nextgen/issues/70))

New resources:

- Azure Event Hubs Namespace Network Rules ([#17](https://github.com/pulumi/pulumi-azure-nextgen/issues/17))
- WebApp SourceControl ([#79](https://github.com/pulumi/pulumi-azure-nextgen/issues/79))
- DBforPostgreSQL Configurations ([#80](https://github.com/pulumi/pulumi-azure-nextgen/issues/80))
- StorageAccount BlobServices ([#91](https://github.com/pulumi/pulumi-azure-nextgen/issues/91))

## 0.2.4 (2020-10-30)

- Fix Cannot create Azure SQL firewall rules ([#87](https://github.com/pulumi/pulumi-azure-nextgen/issues/87))
- Detect deleted resources on refresh ([#66](https://github.com/pulumi/pulumi-azure-nextgen/issues/66))

## 0.2.3 (2020-10-16)

- Fix refreshing resource state ([#60](https://github.com/pulumi/pulumi-azure-nextgen/issues/60))
- Update of a parent resource doesn't cause replacement of child resources using its `name` property anymore
  ([#64](https://github.com/pulumi/pulumi-azure-nextgen/issues/64))
- Fix choice-array type handling, e.g., creation of CDN endpoint with a delivery rules
  ([#68](https://github.com/pulumi/pulumi-azure-nextgen/issues/68))

## 0.2.2 (2020-10-09)

- Fix reading/importing resources with scope URL parameters (e.g., [#51](https://github.com/pulumi/pulumi-azure-nextgen/issues/51))
- Fix result types for invokes that list WebApp settings and similar ([#41](https://github.com/pulumi/pulumi-azure-nextgen/issues/41))
- Fix serialization of free-form dictionary arguments (e.g., [#58](https://github.com/pulumi/pulumi-azure-nextgen/issues/58))

## 0.2.1 (2020-10-03)

Bug fixes:

- Fix the packaging of the Go SDK

## 0.2.0 (2020-10-02)

Breaking changes:

- Resolved an issue when some output properties were erroneously presented as input properties too
(notable example: `tier` in `sku` of a `storage.StorageAccount`)

New features:

- Added `appplatform.getResourceUploadUrl` and other `get*` POST-based invokes ([#11](https://github.com/pulumi/pulumi-azure-nextgen/issues/11))
- Resource import is now supported ([#7](https://github.com/pulumi/pulumi-azure-nextgen/issues/7))

Bug fixes:

- Fixed handling of `@odata.type` fields ([#18](https://github.com/pulumi/pulumi-azure-nextgen/issues/18))
- Handle unspecified types in Open API specs more gracefully: accept an "any" type (e.g., `object` in .NET) for inputs and ignore
  them for outputs ([#14](https://github.com/pulumi/pulumi-azure-nextgen/issues/14))
- Fixed the serialization of `CompositeIndexes` property of `SqlResourceSqlContainer`
  ([#28](https://github.com/pulumi/pulumi-azure-nextgen/issues/28))
- Data Factory enum properties (e.g. `CompressionCodec`) changed to strings
  ([#29](https://github.com/pulumi/pulumi-azure-nextgen/issues/29))
- API types "object" are mapped to "any" type in Pulumi.
  Fixed [#30](https://github.com/pulumi/pulumi-azure-nextgen/issues/30) and [#32](https://github.com/pulumi/pulumi-azure-nextgen/issues/32)
- Fix serialization of multi-flattened properties [#31](https://github.com/pulumi/pulumi-azure-nextgen/issues/31)

## 0.1.0 (2020-09-21)

The first beta release of the next generation Azure Provider is out!
