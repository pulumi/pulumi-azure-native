# CHANGELOG

## HEAD (Unreleased)

- Memory usage for most programs is greatly reduced [#1689](https://github.com/pulumi/pulumi-azure-native/pull/1689)

## 1.62.0 (2022-04-04)

- Deprecate older explicit API versions [#1547](https://github.com/pulumi/pulumi-azure-native/issues/1547)

### Breaking Changes published by Microsoft

- Type "azure-native:datafactory:SqlServerStoredProcedureActivity" input "storedProcedureParameters" type changed

## 1.61.0 (2022-03-24)

- Update pulumi codegen dependency to fix secret property handling in Go SDK

### Breaking Changes published by Microsoft

- `alertsmanagement` 2020-08-04 preview removed resource in spec
- `app` is in preview and is still in active development
- `dashboard` 2021-09-01 preview changed Graphana resource in spec

### New Resources

- `azure-native:authorization:PrivateLinkAssociation`
- `azure-native:azurearcdata:ActiveDirectoryConnector`
- `azure-native:containerservice:ManagedClusterSnapshot`
- `azure-native:healthcareapis:WorkspacePrivateEndpointConnection`
- `azure-native:media:Track`
- `azure-native:recommendationsservice:Account`
- `azure-native:recommendationsservice:Modeling`
- `azure-native:recommendationsservice:ServiceEndpoint`
- `azure-native:signalrservice:SignalRCustomCertificate`
- `azure-native:signalrservice:SignalRCustomDomain`

## 1.60.0 (2022-02-25)

- Fix property names starting with numbers [#1528](https://github.com/pulumi/pulumi-azure-native/issues/1528)

### New resources

- `azure-native:dashboard:Grafana`
- `azure-native:extendedlocation:ResourceSyncRule`
- `azure-native:mobilenetwork:AttachedDataNetwork`
- `azure-native:mobilenetwork:DataNetwork`
- `azure-native:mobilenetwork:MobileNetwork`
- `azure-native:mobilenetwork:PacketCoreControlPlane`
- `azure-native:mobilenetwork:PacketCoreDataPlane`
- `azure-native:mobilenetwork:Service`
- `azure-native:mobilenetwork:Sim`
- `azure-native:mobilenetwork:SimPolicy`
- `azure-native:mobilenetwork:Site`
- `azure-native:mobilenetwork:Slice`

## 1.59.0 (2022-02-22)

- Fix Diff panic when value is a map [#1515](https://github.com/pulumi/pulumi-azure-native/issues/1515)
- Fix ignores of schema and metadata go files [#1526](https://github.com/pulumi/pulumi-azure-native/issues/1513)

## 1.58.0 (2022-02-14)

- Fix go module size issue introduced in v1.57.0 (https://github.com/pulumi/pulumi-azure-native/issues/1502)

### New resources:

- `azure-native:azureactivedirectory/v20210401:GuestUsage`
- `azure-native:containerservice/v20220101:Snapshot`
- `azure-native:azureactivedirectory/v20210401:B2CTenant`
- `azure-native:containerservice/v20220101:MaintenanceConfiguration`
- `azure-native:containerservice/v20220101:PrivateEndpointConnection`
- `azure-native:containerservice/v20220101:ManagedCluster`
- `azure-native:containerservice/v20220101:AgentPool`

## 1.57.0 (2022-02-11)

- Multiple updates [#1496](https://github.com/pulumi/pulumi-azure-native/pull/1496):
  - pulumi/pulumi dependency bumped to v3.24.1
  - Updated resource specs. Includes breaking change (see below)
  - Workaround codegen bugs affecting new resource specs
  - Fix tests

### Breaking Changes published by Microsoft

- Deletion of `v2019-05-01-preview` API for `appplatform`. See [source](https://github.com/Azure/azure-rest-api-specs/pull/17506)

### New resources

- `azure-native:netapp:Subvolume`
- `azure-native:netapp:VolumeGroup`
- `azure-native:network:DnsForwardingRuleset`
- `azure-native:network:DnsResolver`
- `azure-native:network:ForwardingRule`
- `azure-native:network:InboundEndpoint`
- `azure-native:network:OutboundEndpoint`

### New functions

- `azure-native:hybridnetwork:listVendorSkusCredential`
- `azure-native:netapp:getSubvolume`
- `azure-native:netapp:getSubvolumeMetadata`
- `azure-native:netapp:getVolumeGroup`
- `azure-native:network:getDnsForwardingRuleset`
- `azure-native:network:getDnsResolver`
- `azure-native:network:getForwardingRule`
- `azure-native:network:getInboundEndpoint`
- `azure-native:network:getOutboundEndpoint`
- `azure-native:network:listDnsForwardingRulesetByVirtualNetwork`
- `azure-native:network:listDnsResolverByVirtualNetwork`

---

## 1.56.0 (2022-02-03)

### New resources

- `digitaltwins:TimeSeriesDatabaseConnection`

## 1.55.0 (2022-02-02)

- Upgrade internal pulumi/pulumi dependencies to v3.23.2

## 1.54.0 (2022-01-25)

### Breaking Changes published by Microsoft

- `security.Scanner` resource and function removed
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/17232))
- `cdn.Route` and `cdn.AFDOrigin` property types changed
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/17278))
- `botservice:BotProperties` `isDeveloperAppInsightsApiKeySet` and `cmekEncryptionStatus` properties removed
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/17112))

## 1.53.0 (2022-01-10)

### New resources

- `operationalinsights.Table`
- `security.Scanner`
- `storage.LocalUser`

### New functions

- `operationalinsights.getTable`
- `security.getScanner`
- `storage.getLocalUser`
- `storage.listLocalUserKeys`
- `web.listContainerAppSecrets`

### Enhancements

- Input/output based functions API in Go SDK

## 1.52.0 (2021-12-29)

### New resources

- `appplatform.ApiPortal`
- `appplatform.ApiPortalCustomDomain`
- `appplatform.BuildServiceBuilder`
- `appplatform.BuildpackBinding`
- `appplatform.ConfigurationService`
- `appplatform.Gateway`
- `appplatform.GatewayCustomDomain`
- `appplatform.GatewayRouteConfig`
- `appplatform.ServiceRegistry`
- `openenergyplatform.EnergyService`

### New functions

- `appplatform.getApiPortal`
- `appplatform.getApiPortalCustomDomain`
- `appplatform.getBuildServiceBuildResultLog`
- `appplatform.getBuildServiceBuilder`
- `appplatform.getBuildServiceResourceUploadUrl`
- `appplatform.getBuildpackBinding`
- `appplatform.getConfigurationService`
- `appplatform.getGateway`
- `appplatform.getGatewayCustomDomain`
- `appplatform.getGatewayRouteConfig`
- `appplatform.getServiceRegistry`
- `openenergyplatform.getEnergyService`

### Bug fixes

- Fix deletion of `sql.BackupShortTermRetentionPolicy` and `sql.LongTermRetentionPolicy`
  [#1345](https://github.com/pulumi/pulumi-azure-native/issues/1345)

## 1.51.0 (2021-12-20)

### New resources

- `datamigration.SqlMigrationService`
- `sql.IPv6FirewallRule`

### New functions

- `datamigration.getSqlMigrationService`
- `datamigration.listSqlMigrationServiceAuthKeys`
- `datamigration.listSqlMigrationServiceMonitoringData`
- `sql.getIPv6FirewallRule`

### Enhancements

- Update to `pulumi/pulumi` 3.20.0
- Go SDK improvements:
  - Removed unused input types, see https://github.com/pulumi/pulumi/pull/7943
  - All comments for resources, types, and properties restored
  - SDK size reduced by 30%

## 1.50.0 (2021-12-15)

### New functions

- `elastic.listUpgradableVersionDetails`

### Enhancements

- Remove `etag` from resource inputs
  [#1337](https://github.com/pulumi/pulumi-azure-native/issues/1337)
- Update to `pulumi/pulumi` 3.19.0

### Breaking Changes published by Microsoft

- `provisioningState` property removed from a new types in `botservice`: `MsTeamsChannel`, `DirectLineChannel`, `WebChatChannel`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/16882))
- `alertsmanagement.PrometheusRuleGroup` resource removed
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/17060))

## 1.49.0 (2021-12-06)

### Bug fixes

- Fix `web.WebAppSwiftVirtualNetworkConnectionSlot` creation
  [#866](https://github.com/pulumi/pulumi-azure-native/issues/866)

### Breaking Changes published by Microsoft

- `documentdb.GraphResourceGetPropertiesResponseResource` removed properties `rid`, `ts`, `etag`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/16841))

## 1.48.0 (2021-11-29)

### New resources

- `loadtestservice.LoadTest`

### New functions

- `loadtestservice.getLoadTest`

### Bug fixes

- Fix `resources.TagAtScope` creation
  [#1253](https://github.com/pulumi/pulumi-azure-native/issues/1253)
- Fix deletion of `sql.ExtendedServerBlobAuditingPolicy`
  [#1202](https://github.com/pulumi/pulumi-azure-native/issues/1202)

## 1.47.0 (2021-11-19)

### CRITICAL Bug fix

- Avoid provider credentials leaking into state.
  [#1309](https://github.com/pulumi/pulumi-azure-native/issues/1309)

  **PLEASE READ**

  If you set credentials through environment variables (e.g. `ARM_CLIENT_SECRET`) AND
  use the SDK to create a provider where these values are not explicitly set, (e.g. `new provider.Provider("...");`)
  prior versions of the `azure-native` provider may have included the credentials in the state in clear text.
  All users are recommended to upgrade their provider version and run a `pulumi up`. It is highly recommended to
  rotate the affected credentials after all relevant stacks have been updated.

  You can check if your state file contains credentials by running `pulumi stack export | grep -A 3 "clientSecret\|clientCertificatePassword\|clientId"`
  and checking if any unencrypted values are produced. After the update these values will either not be present
  or be stored as encrypted secrets using your stack's preferred encryption provider.

  Note that the Pulumi state backend also encrypts the state as a whole and other state backends
  support a similar mechanism which should significantly limit exposure of the credentials.
  Nonetheless, We sincerely regret the inconvenience this causes.

### New resources

- `alertsmanagement.PrometheusRuleGroup`
- `orbital.Contact`
- `orbital.ContactProfile`
- `orbital.Spacecraft`

### New functions

- `alertsmanagement.getPrometheusRuleGroup`
- `orbital.getContact`
- `orbital.getContactProfile`
- `orbital.getSpacecraft`
- `orbital.listSpacecraftAvailableContacts`

### Breaking Changes published by Microsoft

- `botservice.BotProperties` replaced `isIsolated` property with `publicNetworkAccess`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/16760))

## 1.46.0 (2021-11-15)

### Bug fixes

- Make Deletion of DNS zone more robust
  [#969](https://github.com/pulumi/pulumi-azure-native/issues/969)
- Fix creation of `sql.ServerAzureADOnlyAuthentication` resource
  [#1042](https://github.com/pulumi/pulumi-azure-native/issues/1042)

### Enhancements

- Remove aliases from `azure-nextgen` resources. If you need to update from `azure-nextgen`
  versions, please update to 1.45.0 or earlier first, run `pulumi up` successfully, then
  update to later versions.
  [#1284](https://github.com/pulumi/pulumi-azure-native/issues/1284)
- Fixed replacement of ServiceBus Topics & Queues
  [#940](https://github.com/pulumi/pulumi-azure-native/issues/940)

### New resources

- `apimanagement.Schema`
- `documentdb.MongoDBResourceMongoRoleDefinition`
- `documentdb.MongoDBResourceMongoUserDefinition`

### New functions

- `apimanagement.getSchema`
- `documentdb.getMongoDBResourceMongoRoleDefinition`
- `documentdb.getMongoDBResourceMongoUserDefinition`

### Updated API versions for top-level resources

- `edgeorder` resources updated from `2020-12-01-preview` to `2020-12-01`

### Breaking Changes published by Microsoft

- `iotsecurity.LocationSite` and `iotsecurity.SiteSensor` were removed
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/16588))
- `botservice.WebChatSite` property `enablePreview` renamed to `isWebchatPreviewEnabled`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/16377))
- `datafactory.Flowlet` property `additionalProperties` removed
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/16652))

## 1.45.0 (2021-11-05)

### New resources

- `cognitiveservices.CommitmentPlan`
- `cognitiveservices.Deployment`

### New functions

- `cognitiveservices.getCommitmentPlan`
- `cognitiveservices.getDeployment`

### Breaking Changes published by Microsoft

- `security:CustomAssessmentAutomation` properties `implementationEffort`
  and `userImpact` removed
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/16550))

## 1.44.0 (2021-11-02)

### New resources

- `web.ContainerApp`

### New functions

- `maps.listAccountSas`
- `web.getContainerApp`

### Breaking Changes published by Microsoft

- `compute.RestorePoint` property `provisioningDetails` removed
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/16099))
- `authorization.PolicyPricing` resource removed
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/16537))
- `datafactory` changed the `type` property of compression from `string`
  to `object` and removed several derived types
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/16563))

## 1.43.0 (2021-10-27)

### New resources

- `security.SecurityConnector`

### New functions

- `security.getSecurityConnector`

### Enhancements

- Update to `pulumi/pulumi` 3.16.0

## 1.42.0 (2021-10-25)

### New resources

- `avs.PlacementPolicy`
- `chaos.Capability`
- `chaos.Experiment`
- `chaos.Target`
- `deviceupdate.PrivateEndpointConnectionProxy`
- `network.RoutingIntent`
- `resourceconnector.Appliance`

### New functions

- `avs.getPlacementPolicy`
- `chaos.getCapability`
- `chaos.getExperiment`
- `chaos.getTarget`
- `deviceupdate.getPrivateEndpointConnectionProxy`
- `network.getRoutingIntent`
- `network.listFirewallPolicyIdpsSignature`
- `network.listFirewallPolicyIdpsSignaturesFilterValue`
- `resourceconnector.getAppliance`
- `resourceconnector.listApplianceClusterUserCredential`

## 1.41.0 (2021-10-20)

### New resources

- `appplatform.Storage`
- `videoindexer.Account`

### New functions

- `appplatform.getStorage`
- `videoindexer.getAccount`

## 1.40.0 (2021-10-18)

### New resources

- `eventhub.SchemaRegistry`

### New functions

- `eventhub.getSchemaRegistry`

### Bug Fixes

- Fix validation of untyped arrays
  [#1224](https://github.com/pulumi/pulumi-azure-native/issues/1224).

### Breaking Changes published by Microsoft

- `securityinsights.SourceControl` property `sourceControlId` renamed to `id`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/16280))

## 1.39.0 (2021-10-15)

### New resources

- `hybridconnectivity.Endpoint`

### New functions

- `hybridconnectivity.getEndpoint`
- `hybridconnectivity.listEndpointCredentials`

## 1.38.0 (2021-10-14)

### New resources

- `kubernetesconfiguration.FluxConfiguration`

### New functions

- `kubernetesconfiguration.getFluxConfiguration`

### Enhancements

- Add replacement annotations for `containerservice.ManagedCluster`
  [#959](https://github.com/pulumi/pulumi-azure-native/issues/959).

## 1.37.0 (2021-10-12)

### Enhancements

- Update to `pulumi/pulumi` 3.14.0 with input/output based functions API in Node.js

### Breaking Changes published by Microsoft

- Identity definition `deviceupdate.Account` was adjusted to use the common type
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/16197))
- Several fixes in the `logic` module to match the API behavior
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/16181))
- `vmwarecloudsimple.DedicatedCloudNode.created` type is set to `string`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/16090))
- `hybridnetwork.Device` has `azureStackEdge` removed
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/16316))
- `edgeorder.ShippingDetails` type split into `ForwardShippingDetails` and `ReverseShippingDetails`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15962))

## 1.36.0 (2021-10-07)

### New resources

- `sql.DistributedAvailabilityGroup`
- `sql.ServerTrustCertificate`

### New functions

- `sql.getDistributedAvailabilityGroup`
- `sql.getServerTrustCertificate`

## 1.35.0 (2021-10-06)

### Enhancements

- Update to `pulumi/pulumi` 3.13.2 with input/output based functions API in Python

### New resources

- `authorization.PolicyPricing`
- `connectedvmwarevsphere.Cluster`
- `connectedvmwarevsphere.Datastore`
- `connectedvmwarevsphere.GuestAgent`
- `connectedvmwarevsphere.Host`
- `connectedvmwarevsphere.HybridIdentityMetadatum`
- `connectedvmwarevsphere.InventoryItem`
- `connectedvmwarevsphere.MachineExtension`
- `connectedvmwarevsphere.ResourcePool`
- `connectedvmwarevsphere.VCenter`
- `connectedvmwarevsphere.VirtualMachine`
- `connectedvmwarevsphere.VirtualMachineTemplate`
- `connectedvmwarevsphere.VirtualNetwork`
- `videoanalyzer.LivePipeline`
- `videoanalyzer.PipelineJob`
- `videoanalyzer.PipelineTopology`
- `videoanalyzer.PrivateEndpointConnection`

### New functions

- `authorization.getPolicyPricing`
- `connectedvmwarevsphere.getCluster`
- `connectedvmwarevsphere.getDatastore`
- `connectedvmwarevsphere.getGuestAgent`
- `connectedvmwarevsphere.getHost`
- `connectedvmwarevsphere.getHybridIdentityMetadatum`
- `connectedvmwarevsphere.getInventoryItem`
- `connectedvmwarevsphere.getMachineExtension`
- `connectedvmwarevsphere.getResourcePool`
- `connectedvmwarevsphere.getVCenter`
- `connectedvmwarevsphere.getVirtualMachine`
- `connectedvmwarevsphere.getVirtualMachineTemplate`
- `connectedvmwarevsphere.getVirtualNetwork`
- `videoanalyzer.getLivePipeline`
- `videoanalyzer.getPipelineJob`
- `videoanalyzer.getPipelineTopology`
- `videoanalyzer.getPrivateEndpointConnection`
- `videoanalyzer.listVideoContentToken`

## 1.34.0 (2021-09-30)

### New resources

- `security.Assignment`
- `security.Standard`

### New functions

- `security.getAssignment`
- `security.getStandard`

## 1.33.0 (2021-09-28)

### New resources

- `servicelinker.Linker`
- `sql.EncryptionProtector`

### New functions

- `servicelinker.getLinker`
- `servicelinker.listLinkerConfigurations`
- `sql.getEncryptionProtector`

## 1.32.0 (2021-09-27)

### New resources

- `timeseriesinsights.PrivateEndpointConnection`
- `webpubsub.WebPubSubHub`

### New functions

- `timeseriesinsights.getPrivateEndpointConnection`
- `webpubsub.getWebPubSubHub`

## 1.31.0 (2021-09-23)

### New resources

- `apimanagement.PrivateEndpointConnectionByName`
- `deviceupdate.PrivateEndpointConnection`
- `offazure.PrivateEndpointConnection`
- `security.CustomAssessmentAutomation`
- `security.CustomEntityStoreAssignment`

### New functions

- `apimanagement.getPrivateEndpointConnectionByName`
- `deviceupdate.getPrivateEndpointConnection`
- `offazure.getPrivateEndpointConnection`
- `security.getCustomAssessmentAutomation`
- `security.getCustomEntityStoreAssignment`

### Breaking Changes published by Microsoft

- The response shape is undefine for the `offsite.MasterSite` resource,
  so we had to remove it from our SDKs until this problem is resolved
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/14776))

## 1.30.0 (2021-09-22)

### Bug Fixes

- Fix the destroy operation for `web.WebAppAuthSettings`
  [#947](https://github.com/pulumi/pulumi-azure-native/issues/947).

## 1.29.0 (2021-09-20)

### New resources

- `hdinsight.PrivateEndpointConnection`
- `labservices:LabPlan`
- `labservices:Schedule`

### New functions

- `hdinsight.getPrivateEndpointConnection`
- `labservices.getLabPlan`
- `labservices.getSchedule`

### Breaking Changes published by Microsoft

- The shape of the `documentdb.Service` resource adjusted to match API expectations
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15971))
- `redis.Redis` now has configuration properties specified correctly
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15674))
- `extendedlocation.customLocation` casing corrected to `extendedlocation.CustomLocation`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15676))

### Bug Fixes

- Fix `authorization.getClientConfig` when using Managed Identities
  [#1008](https://github.com/pulumi/pulumi-azure-native/issues/1008).

## 1.28.0 (2021-09-13)

### New functions

- `kubernetes.listConnectedClusterUserCredential`

### Breaking Changes published by Microsoft

- SKU property removed from `purview.Account`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15965))

## 1.27.0 (2021-09-10)

### New resources

- `containerservice.Snapshot`
- `documentdb.CassandraResourceCassandraView`
- `documentdb.GraphResourceGraph`

### New functions

- `containerservice.getSnapshot`
- `documentdb.getCassandraResourceCassandraView`
- `documentdb.getGraphResourceGraph`

### Dependencies

- Updated `pulumi/pulumi` dependencies to 3.10.3
- The type of discriminated union inputs in the .NET SDK has changed from `Input<object>` to `object`
- Support inputty Go enums

## 1.26.0 (2021-09-09)

### New resources

- `aadiam.PrivateEndpointConnection`
- `kusto.ManagedPrivateEndpoint`
- `kusto.PrivateEndpointConnection`

### New functions

- `aadiam:getPrivateEndpointConnection`
- `kusto.getManagedPrivateEndpoint`
- `kusto.getPrivateEndpointConnection`
- `redhatopenshift.listOpenShiftClusterAdminCredentials`

### Breaking Changes published by Microsoft

- Kusto-related resources and functions in the `synapse` are renamed
  (e.g. `synapse.AttachedDatabaseConfiguration` to `synapse.KustoPoolAttachedDatabaseConfiguration`)
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15816))
- `v20210513preview` version removed from `desktopvirtualization`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15811))
- Hub and NetworkGroup resources restructured in `network`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15840))
- App Service Auth shape restructured for `web` resources
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15283))

### Updated API versions for top-level resources

- `authorization:ManagementLock*` resources and functions updated from `2016-09-01` to `2017-04-01`
- `aadiam.privateLinkForAzureAd" changed from `2020-03-01-preview`to`2020-03-01`

---

## 1.25.0 (2021-09-01)

### New resources

- `security.AdvancedThreatProtection`

### New functions

- `security.getAdvancedThreatProtection`

### Enhancements

- Add missing API versions of Service Fabric
  [#922](https://github.com/pulumi/pulumi-azure-native/issues/922).

## 1.24.0 (2021-08-31)

### New resources

- `automation.HybridRunbookWorker`
- `automation.HybridRunbookWorkerGroup`

### New functions

- `automation.getHybridRunbookWorker`
- `automation.getHybridRunbookWorkerGroup`

### Updated API versions for top-level resources

- `logz` resources and functions updated from `2020-10-01-preview` to `2020-10-01`

### Breaking Changes published by Microsoft

- Enum `securityinsights:CaseSeverity` renamed to `securityinsights:IncidentSeverityEnum`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15761))

## 1.23.1 (2021-08-26)

Rollback code generation changes below to work around
[#1092](https://github.com/pulumi/pulumi-azure-native/issues/1092).
This removes the following changes of 1.23.0:

- The type of discriminated union inputs in the .NET SDK has changed from `Input<object>` to `object`
- Support inputty Go enums

## 1.23.0 (2021-08-20)

### New resources

- `dataprotection.ResourceGuard`
- `powerbi.PowerBIResource`
- `powerbi.PrivateEndpointConnection`
- `sql.DatabaseAdvisor`
- `sql.ServerAdvisor`

### New functions

- `dataprotection.getResourceGuard`
- `powerbi.getPrivateEndpointConnection`
- `sql.getDatabaseAdvisor`
- `sql.getServerAdvisor`

### Bug Fixes

- Fix `sql.ServerVulnerabilityAssessment` existence check
  [#1050](https://github.com/pulumi/pulumi-azure-native/issues/1050).

### Dependencies

- Updated `pulumi/pulumi` dependencies to 3.10.3
- The type of discriminated union inputs in the .NET SDK has changed from `Input<object>` to `object`
- Support inputty Go enums

### Breaking Changes published by Microsoft

- `network.listEffectiveSecurityUserRuleBySubnet` was removed
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15668))

## 1.22.0 (2021-08-17)

### Bug Fixes

- Remove validation of properties with type `any`
  [#1058](https://github.com/pulumi/pulumi-azure-native/issues/1058).

- Better support for initialization failures by checkpointing partially created resources to state
  [#938](https://github.com/pulumi/pulumi-azure-native/issues/938)

### Breaking Changes published by Microsoft

- All hierchicalQueue-related fields removed from `datalakeanalytics.Account`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15591))

- Fixes in SKU schema of providerhub resources
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15608))

## 1.21.0 (2021-08-12)

### New functions

- `network.listActiveConnectivityConfiguration`
- `network.listActiveSecurityAdminRule`
- `network.listActiveSecurityUserRule`
- `network.listEffectiveConnectivityConfiguration`
- `network.listEffectiveSecurityUserRuleBySubnet`
- `network.listNetworkManagerEffectiveSecurityAdminRule`

### Breaking Changes published by Microsoft

- `network.listEffectiveConfiguration`, `network.listActiveConfiguration`,
  `network.listEffectiveConfigurationBySubnet` were removed in favor of the functions listed
  in "New functions" above
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15282))

## 1.20.0 (2021-08-06)

### New resources

- `iotsecurity.LocationSite`
- `iotsecurity.SiteSensor`

### New functions

- `iotsecurity.getLocationSite`
- `iotsecurity.getSiteSensor`

### Enhancements

- Support initialization failures by checkpointing partially created resources to state
  [#938](https://github.com/pulumi/pulumi-azure-native/issues/938)

### Breaking Changes published by Microsoft

- `datafactory` adjusted the type of `additionalColumns` and `compressionType` properties
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15264))
- Version `2020-08-06-preview` of the `security` module is deprecated, all resource removed
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15434))

## 1.19.0 (2021-07-22)

### Enhancements

- Updates Authorization::RoleAssignment to force replace on principalId or scope change
  [#771](https://github.com/pulumi/pulumi-azure-native/issues/771)

### New resources

- `fluidrelay.FluidRelayServer`
- `healthcareapis.DicomService`
- `healthcareapis.FhirService`
- `healthcareapis.IotConnector`
- `healthcareapis.IotConnectorFhirDestination`
- `healthcareapis.Workspace`

### New functions

- `fluidrelay.getFluidRelayServer`
- `fluidrelay.getFluidRelayServerKeys`
- `healthcareapis.getDicomService`
- `healthcareapis.getFhirService`
- `healthcareapis.getIotConnector`
- `healthcareapis.getIotConnectorFhirDestination`
- `healthcareapis.getWorkspace`

### Breaking Changes published by Microsoft

- Several properties in `hanaonazure` are marked as read-only
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15205))
- `maxUnusedVersionsToKeep` in `servicefabric.ApplicationTypeVersionsCleanupPolicy` changed the type from
  integer to float (expects int64)
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/15128))

## 1.18.0 (2021-07-19)

### New resources

- `authorization.ResourceManagementPrivateLink`
- `securityinsights.Anomalies`

### New functions

- `authorization.getResourceManagementPrivateLink`
- `securityinsights.getAnomalies`

## 1.17.0 (2021-07-16)

### New resources

- `compute.CapacityReservation`
- `compute.CapacityReservationGroup`
- `network.NetworkSecurityPerimeter`

### New functions

- `compute.getCapacityReservation`
- `compute.getCapacityReservationGroup`
- `dbformysql.getGetPrivateDnsZoneSuffixExecute`
- `network.getNetworkSecurityPerimeter`

### Breaking Changes published by Microsoft

- `maxCapacity` is removed from `operationalinsights.outputs.CapacityReservationPropertiesResponse`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/14953))
- `managedCredential` renamed to `credential` in `datafactory.IntegrationRuntimeSsisProperties`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/14589))

## 1.16.0 (2021-07-06)

### New Resources

- `synapse.AttachedDatabaseConfiguration`

### New Functions

- `synapse.getAttachedDatabaseConfiguration`
- `synapse.listKustoPoolFollowerDatabases`
- `synapse.listKustoPoolLanguageExtensions`

### Breaking Changes published by Microsoft

- A number of breaking changes in `edgeorder` module (preview API version)
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/14935))

### Updated API versions for top-level resources

- `elastic` resources and functions updated from `2020-07-01-preview` to `2020-07-01`

## 1.15.0 (2021-07-01)

### New Resources

- `features.SubscriptionFeatureRegistration`
- `peering.ConnectionMonitorTest`

### New Functions

- `features.getSubscriptionFeatureRegistration`
- `peering.getConnectionMonitorTest`

### Bug Fixes

- Fix import of resources with unorthodox parameter capitalization, e.g. `network.VirtualWan`
  [#942](https://github.com/pulumi/pulumi-azure-native/issues/942).

### Breaking Changes published by Microsoft

- `iotspaces` resource provider was removed
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/13993))

## 1.14.0 (2021-06-23)

### New Resources

- `azurestackhci.ArcSetting`
- `azurestackhci.Extension`

### New Functions

- `azurestackhci.getArcSetting`
- `azurestackhci.getExtension`

### Breaking Changes published by Microsoft

- `datafactory` changed the type for `avroCompressionCodec` and `orcCompressionCodec`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/14770))
- Type `compute.PublicIPAddressSku` changed the properties `publicIPAddressSkuName` to `name` and `tier`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/14782))

## 1.13.0 (2021-06-17)

### Bug Fixes

- Add missing types for complete definition of dashboards
  [#858](https://github.com/pulumi/pulumi-azure-native/issues/858).

### New Resources

- `desktopvirtualization.SessionHostConfiguration`

### New Functions

- `desktopvirtualization.getSessionHostConfiguration`
- `eventgrid.getEventSubscriptionDeliveryAttributes`
- `eventgrid.getPartnerTopicEventSubscriptionDeliveryAttributes`
- `eventgrid.getSystemTopicEventSubscriptionDeliveryAttributes`

### Updated API versions for top-level resources

- `securityinsights.IncidentComment` and `securityinsights.IncidentRelation` promoted
  from `2019-01-01-preview` to `2021-03-01-preview` to unify with other resources of
  the `securityinsights` module. New API contains no breaking changes.

### Breaking Changes published by Microsoft

- `batchai` resource provider was deprecated and removed
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/14788))

## 1.12.0 (2021-06-15)

### New Resources

- `databricks.PrivateEndpointConnection`

### New Functions

- `apimanagement.getUserSharedAccessToken`
- `databricks.getPrivateEndpointConnection`
- `insights.getDiagnosticServiceTokenReadOnly`
- `insights.getDiagnosticServiceTokenReadWrite`
- `insights.getLiveToken`
- `machinelearningservices.getOnlineEndpointToken`
- `migrate.getProjectKeys`
- `notificationhubs.getNotificationHubPnsCredentials`
- `operationalinsights.getSharedKeys`
- `recoveryservices.getRecoveryPointAccessToken`
- `storsimple.getManagerDevicePublicEncryptionKey`

### Removed Functions

- `operationalinsights.listWorkspaceKeys` - this function was deprecated on
  the Azure side, and the API doesn't work anymore. Use `operationalinsights.getSharedKeys`
  instead. See [#882](https://github.com/pulumi/pulumi-azure-native/issues/882).

## 1.11.0 (2021-06-09)

### New Resources

- `azurearcdata.DataController`
- `azurearcdata.PostgresInstance`
- `azurearcdata.SqlManagedInstance`
- `azurearcdata.SqlServerInstance`

### New Functions

- `azurearcdata.getDataController`
- `azurearcdata.getPostgresInstance`
- `azurearcdata.getSqlManagedInstance`
- `azurearcdata.getSqlServerInstance`

## 1.10.0 (2021-06-08)

### New Resources

- `marketplace:PrivateStoreCollection`
- `marketplace:PrivateStoreCollectionOffer`
- `powerplatform:Account`
- `securityinsights:ActivityCustomEntityQuery`
- `securityinsights:EntityQuery`
- `testbase:CustomerEvent`

### New Functions

- `marketplace:getPrivateStoreCollection`
- `marketplace:getPrivateStoreCollectionOffer`
- `powerplatform:getAccount`
- `securityinsights:getActivityCustomEntityQuery`
- `securityinsights:getEntityQuery`
- `testbase:getCustomerEvent`

## 1.9.0 (2021-05-31)

### New Resources

- `botservice.PrivateEndpointConnection`
- `documentdb.Service`
- `insights.DataCollectionEndpoint`
- `machinelearningservices.BatchDeployment`
- `machinelearningservices.BatchEndpoint`
- `machinelearningservices.CodeContainer`
- `machinelearningservices.CodeVersion`
- `machinelearningservices.DataContainer`
- `machinelearningservices.DataVersion`
- `machinelearningservices.EnvironmentContainer`
- `machinelearningservices.EnvironmentSpecificationVersion`
- `machinelearningservices.Job`
- `machinelearningservices.ModelContainer`
- `machinelearningservices.ModelVersion`
- `machinelearningservices.OnlineDeployment`
- `machinelearningservices.OnlineEndpoint`
- `securityinsights.Metadata`
- `securityinsights.SourceControl`
- `sql.OutboundFirewallRule`

### New Functions

- `botservice.getPrivateEndpointConnection`
- `documentdb.getService`
- `insights.getDataCollectionEndpoint`
- `machinelearningservices.getBatchDeployment`
- `machinelearningservices.getBatchEndpoint`
- `machinelearningservices.getCodeContainer`
- `machinelearningservices.getCodeVersion`
- `machinelearningservices.getDataContainer`
- `machinelearningservices.getDataVersion`
- `machinelearningservices.getEnvironmentContainer`
- `machinelearningservices.getEnvironmentSpecificationVersion`
- `machinelearningservices.getJob`
- `machinelearningservices.getModelContainer`
- `machinelearningservices.getModelVersion`
- `machinelearningservices.getOnlineDeployment`
- `machinelearningservices.getOnlineDeploymentLogs`
- `machinelearningservices.getOnlineEndpoint`
- `machinelearningservices.listBatchEndpointKeys`
- `machinelearningservices.listDatastoreSecrets`
- `machinelearningservices.listOnlineEndpointKeys`
- `securityinsights.getMetadata`
- `securityinsights.getSourceControl`
- `securityinsights.listSourceControlRepositories`
- `sql.getOutboundFirewallRule`

## 1.8.0 (2021-05-25)

### New Resources

- `web.KubeEnvironment`

### New Functions

- `logz.listMonitorUserRoles`
- `logz.listMonitorVMHosts`
- `logz.listSubAccountVMHosts`
- `web.getKubeEnvironment`

### Updated API versions for top-level resources

- Azure now indicates that the version `2021-01-01-preview` of `servicebus.PrivateEndpointConnection`
  isn't supported yet, so we switched it to the older supported version `2018-01-01-preview`

## 1.7.0 (2021-05-19)

### Improvements

- Do not count 429 (Too Many Requests) towards the limit of retry attempts, retry them indefinitely
  [#849](https://github.com/pulumi/pulumi-azure-native/issues/849)

### New Resources

- `avs.CloudLink`
- `avs.ScriptExecution`
- `avs.WorkloadNetworkPublicIP`
- `blueprint.PolicyAssignmentArtifact`
- `blueprint.RoleAssignmentArtifact`
- `blueprint.TemplateArtifact`
- `databoxedge.ArcAddon`
- `databoxedge.CloudEdgeManagementRole`
- `databoxedge.FileEventTrigger`
- `databoxedge.IoTAddon`
- `databoxedge.IoTRole`
- `databoxedge.KubernetesRole`
- `databoxedge.MECRole`
- `databoxedge.PeriodicTimerEventTrigger`
- `datashare.ADLSGen1FileDataSet`
- `datashare.ADLSGen1FolderDataSet`
- `datashare.ADLSGen2FileDataSet`
- `datashare.ADLSGen2FileDataSetMapping`
- `datashare.ADLSGen2FileSystemDataSet`
- `datashare.ADLSGen2FileSystemDataSetMapping`
- `datashare.ADLSGen2FolderDataSet`
- `datashare.ADLSGen2FolderDataSetMapping`
- `datashare.BlobContainerDataSet`
- `datashare.BlobContainerDataSetMapping`
- `datashare.BlobDataSet`
- `datashare.BlobDataSetMapping`
- `datashare.BlobFolderDataSet`
- `datashare.BlobFolderDataSetMapping`
- `datashare.KustoClusterDataSet`
- `datashare.KustoClusterDataSetMapping`
- `datashare.KustoDatabaseDataSet`
- `datashare.KustoDatabaseDataSetMapping`
- `datashare.ScheduledSynchronizationSetting`
- `datashare.ScheduledTrigger`
- `datashare.SqlDBTableDataSet`
- `datashare.SqlDBTableDataSetMapping`
- `datashare.SqlDWTableDataSet`
- `datashare.SqlDWTableDataSetMapping`
- `datashare.SynapseWorkspaceSqlPoolTableDataSet`
- `datashare.SynapseWorkspaceSqlPoolTableDataSetMapping`
- `kusto.EventGridDataConnection`
- `kusto.EventHubDataConnection`
- `kusto.IotHubDataConnection`
- `kusto.ReadOnlyFollowingDatabase`
- `kusto.ReadWriteDatabase`
- `machinelearningservices.ACIService`
- `machinelearningservices.AKSService`
- `machinelearningservices.EndpointVariant`
- `network.DefaultAdminRule`
- `network.DefaultUserRule`
- `resources.AzureCliScript`
- `resources.AzurePowerShellScript`
- `security.IngestionSetting`
- `securityinsights.AADDataConnector`
- `securityinsights.AATPDataConnector`
- `securityinsights.ASCDataConnector`
- `securityinsights.AwsCloudTrailDataConnector`
- `securityinsights.EntityAnalytics`
- `securityinsights.EyesOn`
- `securityinsights.FusionAlertRule`
- `securityinsights.MCASDataConnector`
- `securityinsights.MDATPDataConnector`
- `securityinsights.MicrosoftSecurityIncidentCreationAlertRule`
- `securityinsights.OfficeDataConnector`
- `securityinsights.ScheduledAlertRule`
- `securityinsights.TIDataConnector`
- `securityinsights.Ueba`
- `synapse.EventGridDataConnection`
- `synapse.EventHubDataConnection`
- `synapse.IotHubDataConnection`
- `synapse.ReadWriteDatabase`
- `timeseriesinsights.EventHubEventSource`
- `timeseriesinsights.Gen1Environment`
- `timeseriesinsights.Gen2Environment`
- `timeseriesinsights.IoTHubEventSource`

### New Functions

- `avs.getCloudLink`
- `avs.getScriptExecution`
- `avs.getScriptExecutionLogs`
- `avs.getWorkloadNetworkPublicIP`
- `blueprint.getPolicyAssignmentArtifact`
- `blueprint.getRoleAssignmentArtifact`
- `blueprint.getTemplateArtifact`
- `databoxedge.getArcAddon`
- `databoxedge.getCloudEdgeManagementRole`
- `databoxedge.getFileEventTrigger`
- `databoxedge.getIoTAddon`
- `databoxedge.getIoTRole`
- `databoxedge.getKubernetesRole`
- `databoxedge.getMECRole`
- `databoxedge.getPeriodicTimerEventTrigger`
- `datashare.getADLSGen1FileDataSet`
- `datashare.getADLSGen1FolderDataSet`
- `datashare.getADLSGen2FileDataSet`
- `datashare.getADLSGen2FileDataSetMapping`
- `datashare.getADLSGen2FileSystemDataSet`
- `datashare.getADLSGen2FileSystemDataSetMapping`
- `datashare.getADLSGen2FolderDataSet`
- `datashare.getADLSGen2FolderDataSetMapping`
- `datashare.getBlobContainerDataSet`
- `datashare.getBlobContainerDataSetMapping`
- `datashare.getBlobDataSet`
- `datashare.getBlobDataSetMapping`
- `datashare.getBlobFolderDataSet`
- `datashare.getBlobFolderDataSetMapping`
- `datashare.getKustoClusterDataSet`
- `datashare.getKustoClusterDataSetMapping`
- `datashare.getKustoDatabaseDataSet`
- `datashare.getKustoDatabaseDataSetMapping`
- `datashare.getScheduledSynchronizationSetting`
- `datashare.getScheduledTrigger`
- `datashare.getSqlDBTableDataSet`
- `datashare.getSqlDBTableDataSetMapping`
- `datashare.getSqlDWTableDataSet`
- `datashare.getSqlDWTableDataSetMapping`
- `datashare.getSynapseWorkspaceSqlPoolTableDataSet`
- `datashare.getSynapseWorkspaceSqlPoolTableDataSetMapping`
- `kusto.getEventGridDataConnection`
- `kusto.getEventHubDataConnection`
- `kusto.getIotHubDataConnection`
- `kusto.getReadOnlyFollowingDatabase`
- `kusto.getReadWriteDatabase`
- `machinelearningservices.getACIService`
- `machinelearningservices.getAKSService`
- `machinelearningservices.getEndpointVariant`
- `network.getDefaultAdminRule`
- `network.getDefaultUserRule`
- `resources.getAzureCliScript`
- `resources.getAzurePowerShellScript`
- `security.getIngestionSetting`
- `security.listIngestionSettingConnectionStrings`
- `security.listIngestionSettingTokens`
- `securityinsights.getAADDataConnector`
- `securityinsights.getAATPDataConnector`
- `securityinsights.getASCDataConnector`
- `securityinsights.getAwsCloudTrailDataConnector`
- `securityinsights.getEntityAnalytics`
- `securityinsights.getEyesOn`
- `securityinsights.getFusionAlertRule`
- `securityinsights.getMCASDataConnector`
- `securityinsights.getMDATPDataConnector`
- `securityinsights.getMicrosoftSecurityIncidentCreationAlertRule`
- `securityinsights.getOfficeDataConnector`
- `securityinsights.getScheduledAlertRule`
- `securityinsights.getTIDataConnector`
- `securityinsights.getUeba`
- `synapse.getEventGridDataConnection`
- `synapse.getEventHubDataConnection`
- `synapse.getIotHubDataConnection`
- `synapse.getReadWriteDatabase`
- `timeseriesinsights.getEventHubEventSource`
- `timeseriesinsights.getGen1Environment`
- `timeseriesinsights.getGen2Environment`
- `timeseriesinsights.getIoTHubEventSource`

### Deprecated resources

The following resources are now deprecated in favor of their specialized variants listed above.

- `blueprint.Artifact`
- `databoxedge.Addon`
- `databoxedge.Role`
- `databoxedge.Trigger`
- `datashare.DataSet`
- `datashare.DataSetMapping`
- `datashare.SynchronizationSetting`
- `datashare.Trigger`
- `kusto.DataConnection`
- `kusto.Database`
- `machinelearningservices.MachineLearningService`
- `network.AdminRule`
- `network.UserRule`
- `resources.DeploymentScript`
- `securityinsights.AlertRule`
- `securityinsights.DataConnector`
- `securityinsights.ProductSetting`
- `synapse.DataConnection`
- `synapse.Database`
- `timeseriesinsights.Environment`
- `timeseriesinsights.EventSource`

### Breaking Changes published by Microsoft

- `sql` resource provider renamed the type `ResourceIdentityWithUserAssignedIdentities` to `ResourceIdentity`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/14439))

## 1.6.0 (2021-05-17)

### New Features

- Add support for authentication against private and hybrid cloud environments.
  [#831](https://github.com/pulumi/pulumi-azure-native/pull/831/)

### New Resources

- `compute:RestorePoint`
- `compute:RestorePointCollection`
- `desktopvirtualization:PrivateEndpointConnectionByHostPool`
- `desktopvirtualization:PrivateEndpointConnectionByWorkspace`
- `network:VirtualNetworkGatewayNatRule`

### New Functions

- `compute:getRestorePoint`
- `compute:getRestorePointCollection`
- `desktopvirtualization:getPrivateEndpointConnectionByHostPool`
- `desktopvirtualization:getPrivateEndpointConnectionByWorkspace`
- `network:getVirtualNetworkGatewayNatRule`

### Updated API versions for top-level resources

- `compute:VirtualMachine`, `compute:VirtualMachineExtension`, `compute:VirtualMachineRunCommandByVirtualMachine`,
  `compute:VirtualMachineScaleSet`, `compute:VirtualMachineScaleSetExtension`, `compute:VirtualMachineScaleSetVM`,
  `compute:VirtualMachineScaleSetVMExtension`, `compute:VirtualMachineScaleSetVMRunCommand` are upgraded from version
  `2020-12-01` to `2021-03-01` to enable `RestorePoint` and `RestorePointCollection` resources

### Breaking Changes published by Microsoft

- `appplatform` version `2021-03-03-preview` is renamed to `2021-06-01-preview`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/14323))
- `operationalinsights.Workspace` reshaped the `features` property
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/14326))
- Several properties in `datafactory` resources were changed from enums to unstructured dictionaries
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/14160))
- `botservice` removed the `cognitiveServicesSubscriptionId` property from `DirectLineSpeechChannelProperties`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/14317))
- `devtestlab:ScheduleCreationParameter`'s `location` property is now read-only
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/14294))

## 1.5.0 (2021-05-08)

### New Resources

- `agfoodplatform.Extension`
- `agfoodplatform.FarmBeatsModel`

### New Functions

- `agfoodplatform.getExtension`
- `agfoodplatform.getFarmBeatsModel`

### Bug fixes

- Don't set auto-location of LoadBalancerBackendAddressPool as the service doesn't expect it
  [#819](https://github.com/pulumi/pulumi-azure-native/issues/819)

## 1.4.0 (2021-05-05)

### New Resources

- `securityinsights.SentinelOnboardingState`

### New Functions

- `securityinsights.getSentinelOnboardingState`

### Improvements

- Load Python modules lazily to speed up program execution
  [#738](https://github.com/pulumi/pulumi-azure-native/issues/738)

### Bug fixes

- Allow 204 as a valid response for a non-existing resource
  [#808](https://github.com/pulumi/pulumi-azure-native/issues/808)

## 1.3.0 (2021-04-30)

### New Resources

- `synapse.DatabasePrincipalAssignment`
- `synapse.KustoPoolPrincipalAssignment`

### New Functions

- `synapse.getDatabasePrincipalAssignment`
- `synapse.getKustoPoolPrincipalAssignment`

### Bug fixes

- Fix constant value discriminator resolution for multi-level discriminated unions
  [#765](https://github.com/pulumi/pulumi-azure-native/issues/765)

### Breaking Changes published by Microsoft

- `management.ManagementGroup` moved the location of the `path` property
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/13897))
- Changes in `network` resources that are still in private preview:
  `AdminRule`, `UserRule`, `SecurityAdminConfigurations`, `ConnectivityConfiguration`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/13827))

## 1.2.0 (2021-04-28)

### New Resources

- `dbforpostgresql.ServerSecurityAlertPolicy`
  [#788](https://github.com/pulumi/pulumi-azure-native/issues/788)

New module for Video Analyzer management:

- `videoanalyzer:AccessPolicy`
- `videoanalyzer:EdgeModule`
- `videoanalyzer:Video`
- `videoanalyzer:VideoAnalyzer`

New module for Test Base management:

- `testbase.TestBaseAccount`
- `testbase.Package`
- `testbase.FavoriteProcess`

### New Functions

- `elastic.listVMHost`

New module for Video Analyzer management:

- `videoanalyzer.getAccessPolicy`
- `videoanalyzer.getEdgeModule`
- `videoanalyzer.getVideo`
- `videoanalyzer.getVideoAnalyzer`
- `videoanalyzer.listEdgeModuleProvisioningToken`
- `videoanalyzer.listVideoStreamingToken`

New module for Test Base management:

- `testbase.getTestResultDownloadURL`
- `testbase.getTestBaseAccountFileUploadUrl`
- `testbase.getTestResultVideoDownloadURL`
- `testbase.getPackageDownloadURL`
- `testbase.getTestBaseAccount`
- `testbase.getPackage`
- `testbase.getFavoriteProcess`

### Bug fixes

- Fix import of SQL/Mongo databases and collections
  [#741](https://github.com/pulumi/pulumi-azure-native/issues/741) and
  [#777](https://github.com/pulumi/pulumi-azure-native/issues/777)

### Breaking Changes published by Microsoft

- `batchai.FileServer` resource is deprecated ([source](https://github.com/Azure/azure-rest-api-specs/pull/13944))
- Extended location hierarchy for virtual network local gateway ([source](https://github.com/Azure/azure-rest-api-specs/pull/13943))
- `insights` has reconciled two separate types `MyManagedIdentity` and `MyWorkbookManagedIdentity` into the single type
  `MyWorkbookManagedIdentity`. They represent the same thing but had different name by a mistake in design.
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/13951))
- The `maxCapacityReservationLevel` output property was removed from `operationalinsights.WorkspaceSkuResponse`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/13812))
- `security.AssessmentMetadataInSubscription` fixed the property name, changing from `category` to `categories`
  ([source](https://github.com/Azure/azure-rest-api-specs/pull/13883))

## 1.1.0 (2021-04-20)

### New Resources

- `dbforpostgresql.v20200214preview.Configuration` - PostgreSQL Flexible Server Configuration (preview)
  [#711](https://github.com/pulumi/pulumi-azure-native/issues/711)

## 1.0.1 (2021-04-19)

- Fix SDK regression for .NET, Python and Typescript introduced in [#pulumi/6686](https://github.com/pulumi/pulumi/pull/6686)
  [#pulumi/6811](https://github.com/pulumi/pulumi/pull/6811)

## 1.0.0 (2021-04-19)

The native Azure provider for Pulumi is now generally available.

### Breaking Changes

- Depend on Pulumi 3.0, which includes improvements to Python resource arguments and key translation, Go SDK performance, Node SDK performance, general availability of Automation API, and more.
- Generate consistent names for object types regardless on whether or not they are
  transitively reachable from resources or functions. Many input type names have changed
  across Node.js, .NET, and Python SDKs
  [#pulumi/6686](https://github.com/pulumi/pulumi/pull/6686)

## 0.9.0 (2021-04-16)

### Breaking Changes

- All `*/latest` modules are now removed. They have been deprecated since 0.7.0. If you are migrating
  from earlier versions, first upgrade to 0.8.0, replace all resources from `latest` modules with their
  counterparts in top-level unnamed modules, apply the change with `pulumi up`, and then upgrade to 0.9.0+.
  [#712](https://github.com/pulumi/pulumi-azure-native/issues/712)
- The ambiguous resource `network.RecordSet` is split into `network.PrivateRecordSet` (private DNS)
  and `network.RecordSet` (public DNS).
  [#583](https://github.com/pulumi/pulumi-azure-native/issues/583)

### Improvements

- New resources: `sql.BackupShortTermRetentionPolicy`, `sql.DatabaseBlobAuditingPolicy`,
  `sql.ExtendedDatabaseBlobAuditingPolicy`, `sql.ExtendedServerBlobAuditingPolicy`,
  `sql.LongTermRetentionPolicy`, `sql.ServerBlobAuditingPolicy`, `sql.ServerSecurityAlertPolicy`.
  [#725](https://github.com/pulumi/pulumi-azure-native/issues/725)
- Use the `2019-06-01-preview` version for the `containerregistry.Task` resource
  [#736](https://github.com/pulumi/pulumi-azure-native/issues/736)

## 0.8.0 (2021-04-02)

### Breaking Changes

- Updated all top-level resources to the latest API versions.
- Prevent silent clashes of different types under the same name. The types are now disambiguated.
  [#641](https://github.com/pulumi/pulumi-azure-native/pull/641)
  [#673](https://github.com/pulumi/pulumi-azure-native/pull/673)
- The following resources will not appear in the top-level modules anymore:
  `apimanagement:TenantPolicy`, `consumption:BudgetByResourceGroupName`, `containerregistry:BuildStep`,
  `containerservice:ContainerService`, `costmanagement:Budget`, `costmanagement:ReportConfig`,
  `costmanagement:ReportConfigByResourceGroupName`, `datamigration:ServiceTask`, `synapse:SqlDatabase`,
  `web:CertificateCsr`, `web:SiteInstanceDeployment`, `web:SiteInstanceDeploymentSlot`.
- Remove env var support for auxiliary tenant IDs (ARM_AUXILIARY_TENANT_IDS).
  This option can still be set directly as a Provider argument or config.
  [#624](https://github.com/pulumi/pulumi-azure-native/pull/624)
- Respect x-ms-mutability annotations for write-only properties.
  [#679](https://github.com/pulumi/pulumi-azure-native/pull/679)

### Improvements

- Upgrade to Go 1.16.
  [#630](https://github.com/pulumi/pulumi-azure-native/issues/630)
- New resource: `sql.DataMaskingPolicy`.
  [#605](https://github.com/pulumi/pulumi-azure-native/issues/605)
- New resource: `storage.EncryptionSource`.
  [#637](https://github.com/pulumi/pulumi-azure-native/issues/637)
- New resource: `insights.ProactiveDetectionConfiguration`.
  [#704](https://github.com/pulumi/pulumi-azure-native/issues/704)
- `Endpoint` and `CustomDomain` resources are auto-named with a random suffix.
  [#629](https://github.com/pulumi/pulumi-azure-native/issues/629)
- Support auto-naming for UUID properties.
  [#625](https://github.com/pulumi/pulumi-azure-native/issues/625)
- Add support for arm64 plugin binaries.
  [#652](https://github.com/pulumi/pulumi-azure-native/pull/652)
- Add nested collection types in Go to fix SDK build issues
  [#707](https://github.com/pulumi/pulumi-azure-native/pull/707)

### Fixes

- Support migration from `azure-nextgen` with custom providers.
  [#617](https://github.com/pulumi/pulumi-azure-native/issues/617)
- Do not auto-populate the `location` of the `Deployment` resource.
  [#643](https://github.com/pulumi/pulumi-azure-native/issues/643)

## 0.7.1 (2021-03-02)

Fixed the aliases for smooth migration from `azure-nextgen` to `azure-native`.

## 0.7.0 (2021-02-27)

### Azure NextGen is renamed to Azure-Native

See the [announcement](https://pulumi.com/blog/full-coverage-of-azure-resources-with-azure-native/) in the official Pulumi blog.

### Top-Level Resources

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

### Auto-naming ([#5](https://github.com/pulumi/pulumi-azure-nextgen/issues/5))

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

### Breaking Changes

- `web.WebApplicationSettings` renamed to `web.WebAppApplicationSettings` ([#282](https://github.com/pulumi/pulumi-azure-nextgen/issues/282))
- `authorization.GetClientConfig`, `authorization.GetClientToken`, `Storage.Blob`, `Storage.StorageAccountStaticWebsite` were removed from
  corresponding `latest` modules and are only available in the top-level modules

### Bug Fixes

- Relax the enum check to accomodate irregularities like in [#294](https://github.com/pulumi/pulumi-azure-nextgen/issues/294)

## 0.6.1 (2021-02-11)

New features:

- Automatic location propagation ([#6](https://github.com/pulumi/pulumi-azure-nextgen/issues/6))
- New resource: `Storage.Blob` ([#13](https://github.com/pulumi/pulumi-azure-nextgen/issues/13))

## 0.6.0 (2021-02-02)

### Breaking changes in SDKs

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

### Fixes

- Fix unknown propagation in Update's preview ([#115](https://github.com/pulumi/pulumi-azure-nextgen/issues/115))
- Re-model Identity and Encryption properties of EventHubs/ServiceBus namespace ([#243](https://github.com/pulumi/pulumi-azure-nextgen/issues/243))
- Various fixes to Go SDK to address compilation issues due to name collisions. Conflicting types have been renamed for various resource providers but Go programs would not have been able to successfully compile with previous names anyway. ([#137](https://github.com/pulumi/pulumi-azure-nextgen/issues/137))
- Respect custom timeouts ([#252](https://github.com/pulumi/pulumi-azure-nextgen/issues/252))

### New resources

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

### New invoke functions

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

- Retrieve ID in get\* invokes ([#48](https://github.com/pulumi/pulumi-azure-nextgen/issues/48))

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
