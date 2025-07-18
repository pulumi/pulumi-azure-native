// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Aad
{
    public static class GetDomainService
    {
        /// <summary>
        /// The Get Domain Service operation retrieves a json representation of the Domain Service.
        /// 
        /// Uses Azure REST API version 2022-12-01.
        /// 
        /// Other available API versions: 2025-05-01, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native aad [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetDomainServiceResult> InvokeAsync(GetDomainServiceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDomainServiceResult>("azure-native:aad:getDomainService", args ?? new GetDomainServiceArgs(), options.WithDefaults());

        /// <summary>
        /// The Get Domain Service operation retrieves a json representation of the Domain Service.
        /// 
        /// Uses Azure REST API version 2022-12-01.
        /// 
        /// Other available API versions: 2025-05-01, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native aad [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDomainServiceResult> Invoke(GetDomainServiceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainServiceResult>("azure-native:aad:getDomainService", args ?? new GetDomainServiceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The Get Domain Service operation retrieves a json representation of the Domain Service.
        /// 
        /// Uses Azure REST API version 2022-12-01.
        /// 
        /// Other available API versions: 2025-05-01, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native aad [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDomainServiceResult> Invoke(GetDomainServiceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainServiceResult>("azure-native:aad:getDomainService", args ?? new GetDomainServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainServiceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the domain service.
        /// </summary>
        [Input("domainServiceName", required: true)]
        public string DomainServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDomainServiceArgs()
        {
        }
        public static new GetDomainServiceArgs Empty => new GetDomainServiceArgs();
    }

    public sealed class GetDomainServiceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the domain service.
        /// </summary>
        [Input("domainServiceName", required: true)]
        public Input<string> DomainServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetDomainServiceInvokeArgs()
        {
        }
        public static new GetDomainServiceInvokeArgs Empty => new GetDomainServiceInvokeArgs();
    }


    [OutputType]
    public sealed class GetDomainServiceResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Configuration diagnostics data containing latest execution from client.
        /// </summary>
        public readonly Outputs.ConfigDiagnosticsResponse? ConfigDiagnostics;
        /// <summary>
        /// Deployment Id
        /// </summary>
        public readonly string DeploymentId;
        /// <summary>
        /// Domain Configuration Type
        /// </summary>
        public readonly string? DomainConfigurationType;
        /// <summary>
        /// The name of the Azure domain that the user would like to deploy Domain Services to.
        /// </summary>
        public readonly string? DomainName;
        /// <summary>
        /// DomainSecurity Settings
        /// </summary>
        public readonly Outputs.DomainSecuritySettingsResponse? DomainSecuritySettings;
        /// <summary>
        /// Resource etag
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Enabled or Disabled flag to turn on Group-based filtered sync
        /// </summary>
        public readonly string? FilteredSync;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Secure LDAP Settings
        /// </summary>
        public readonly Outputs.LdapsSettingsResponse? LdapsSettings;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Migration Properties
        /// </summary>
        public readonly Outputs.MigrationPropertiesResponse MigrationProperties;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Notification Settings
        /// </summary>
        public readonly Outputs.NotificationSettingsResponse? NotificationSettings;
        /// <summary>
        /// the current deployment or provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// List of ReplicaSets
        /// </summary>
        public readonly ImmutableArray<Outputs.ReplicaSetResponse> ReplicaSets;
        /// <summary>
        /// Resource Forest Settings
        /// </summary>
        public readonly Outputs.ResourceForestSettingsResponse? ResourceForestSettings;
        /// <summary>
        /// Sku Type
        /// </summary>
        public readonly string? Sku;
        /// <summary>
        /// The unique sync application id of the Azure AD Domain Services deployment.
        /// </summary>
        public readonly string SyncApplicationId;
        /// <summary>
        /// SyncOwner ReplicaSet Id
        /// </summary>
        public readonly string SyncOwner;
        /// <summary>
        /// All or CloudOnly, All users in AAD are synced to AAD DS domain or only users actively syncing in the cloud
        /// </summary>
        public readonly string? SyncScope;
        /// <summary>
        /// The system meta data relating to this resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Azure Active Directory Tenant Id
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Data Model Version
        /// </summary>
        public readonly int Version;

        [OutputConstructor]
        private GetDomainServiceResult(
            string azureApiVersion,

            Outputs.ConfigDiagnosticsResponse? configDiagnostics,

            string deploymentId,

            string? domainConfigurationType,

            string? domainName,

            Outputs.DomainSecuritySettingsResponse? domainSecuritySettings,

            string? etag,

            string? filteredSync,

            string id,

            Outputs.LdapsSettingsResponse? ldapsSettings,

            string? location,

            Outputs.MigrationPropertiesResponse migrationProperties,

            string name,

            Outputs.NotificationSettingsResponse? notificationSettings,

            string provisioningState,

            ImmutableArray<Outputs.ReplicaSetResponse> replicaSets,

            Outputs.ResourceForestSettingsResponse? resourceForestSettings,

            string? sku,

            string syncApplicationId,

            string syncOwner,

            string? syncScope,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string tenantId,

            string type,

            int version)
        {
            AzureApiVersion = azureApiVersion;
            ConfigDiagnostics = configDiagnostics;
            DeploymentId = deploymentId;
            DomainConfigurationType = domainConfigurationType;
            DomainName = domainName;
            DomainSecuritySettings = domainSecuritySettings;
            Etag = etag;
            FilteredSync = filteredSync;
            Id = id;
            LdapsSettings = ldapsSettings;
            Location = location;
            MigrationProperties = migrationProperties;
            Name = name;
            NotificationSettings = notificationSettings;
            ProvisioningState = provisioningState;
            ReplicaSets = replicaSets;
            ResourceForestSettings = resourceForestSettings;
            Sku = sku;
            SyncApplicationId = syncApplicationId;
            SyncOwner = syncOwner;
            SyncScope = syncScope;
            SystemData = systemData;
            Tags = tags;
            TenantId = tenantId;
            Type = type;
            Version = version;
        }
    }
}
