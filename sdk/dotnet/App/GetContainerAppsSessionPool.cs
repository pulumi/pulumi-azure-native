// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App
{
    public static class GetContainerAppsSessionPool
    {
        /// <summary>
        /// Container App session pool.
        /// 
        /// Uses Azure REST API version 2024-10-02-preview.
        /// 
        /// Other available API versions: 2024-02-02-preview, 2024-08-02-preview, 2025-01-01, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetContainerAppsSessionPoolResult> InvokeAsync(GetContainerAppsSessionPoolArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContainerAppsSessionPoolResult>("azure-native:app:getContainerAppsSessionPool", args ?? new GetContainerAppsSessionPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Container App session pool.
        /// 
        /// Uses Azure REST API version 2024-10-02-preview.
        /// 
        /// Other available API versions: 2024-02-02-preview, 2024-08-02-preview, 2025-01-01, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetContainerAppsSessionPoolResult> Invoke(GetContainerAppsSessionPoolInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContainerAppsSessionPoolResult>("azure-native:app:getContainerAppsSessionPool", args ?? new GetContainerAppsSessionPoolInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Container App session pool.
        /// 
        /// Uses Azure REST API version 2024-10-02-preview.
        /// 
        /// Other available API versions: 2024-02-02-preview, 2024-08-02-preview, 2025-01-01, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetContainerAppsSessionPoolResult> Invoke(GetContainerAppsSessionPoolInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetContainerAppsSessionPoolResult>("azure-native:app:getContainerAppsSessionPool", args ?? new GetContainerAppsSessionPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContainerAppsSessionPoolArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the session pool.
        /// </summary>
        [Input("sessionPoolName", required: true)]
        public string SessionPoolName { get; set; } = null!;

        public GetContainerAppsSessionPoolArgs()
        {
        }
        public static new GetContainerAppsSessionPoolArgs Empty => new GetContainerAppsSessionPoolArgs();
    }

    public sealed class GetContainerAppsSessionPoolInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the session pool.
        /// </summary>
        [Input("sessionPoolName", required: true)]
        public Input<string> SessionPoolName { get; set; } = null!;

        public GetContainerAppsSessionPoolInvokeArgs()
        {
        }
        public static new GetContainerAppsSessionPoolInvokeArgs Empty => new GetContainerAppsSessionPoolInvokeArgs();
    }


    [OutputType]
    public sealed class GetContainerAppsSessionPoolResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The container type of the sessions.
        /// </summary>
        public readonly string? ContainerType;
        /// <summary>
        /// The custom container configuration if the containerType is CustomContainer.
        /// </summary>
        public readonly Outputs.CustomContainerTemplateResponse? CustomContainerTemplate;
        /// <summary>
        /// The pool configuration if the poolManagementType is dynamic.
        /// </summary>
        public readonly Outputs.DynamicPoolConfigurationResponse? DynamicPoolConfiguration;
        /// <summary>
        /// Resource ID of the session pool's environment.
        /// </summary>
        public readonly string? EnvironmentId;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Managed identities needed by a session pool to interact with other Azure services to not maintain any secrets or credentials in code.
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse? Identity;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Optional settings for a Managed Identity that is assigned to the Session pool.
        /// </summary>
        public readonly ImmutableArray<Outputs.ManagedIdentitySettingResponse> ManagedIdentitySettings;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The number of nodes the session pool is using.
        /// </summary>
        public readonly int NodeCount;
        /// <summary>
        /// The endpoint to manage the pool.
        /// </summary>
        public readonly string PoolManagementEndpoint;
        /// <summary>
        /// The pool management type of the session pool.
        /// </summary>
        public readonly string? PoolManagementType;
        /// <summary>
        /// Provisioning state of the session pool.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The scale configuration of the session pool.
        /// </summary>
        public readonly Outputs.ScaleConfigurationResponse? ScaleConfiguration;
        /// <summary>
        /// The secrets of the session pool.
        /// </summary>
        public readonly ImmutableArray<Outputs.SessionPoolSecretResponse> Secrets;
        /// <summary>
        /// The network configuration of the sessions in the session pool.
        /// </summary>
        public readonly Outputs.SessionNetworkConfigurationResponse? SessionNetworkConfiguration;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetContainerAppsSessionPoolResult(
            string azureApiVersion,

            string? containerType,

            Outputs.CustomContainerTemplateResponse? customContainerTemplate,

            Outputs.DynamicPoolConfigurationResponse? dynamicPoolConfiguration,

            string? environmentId,

            string id,

            Outputs.ManagedServiceIdentityResponse? identity,

            string location,

            ImmutableArray<Outputs.ManagedIdentitySettingResponse> managedIdentitySettings,

            string name,

            int nodeCount,

            string poolManagementEndpoint,

            string? poolManagementType,

            string provisioningState,

            Outputs.ScaleConfigurationResponse? scaleConfiguration,

            ImmutableArray<Outputs.SessionPoolSecretResponse> secrets,

            Outputs.SessionNetworkConfigurationResponse? sessionNetworkConfiguration,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            ContainerType = containerType;
            CustomContainerTemplate = customContainerTemplate;
            DynamicPoolConfiguration = dynamicPoolConfiguration;
            EnvironmentId = environmentId;
            Id = id;
            Identity = identity;
            Location = location;
            ManagedIdentitySettings = managedIdentitySettings;
            Name = name;
            NodeCount = nodeCount;
            PoolManagementEndpoint = poolManagementEndpoint;
            PoolManagementType = poolManagementType;
            ProvisioningState = provisioningState;
            ScaleConfiguration = scaleConfiguration;
            Secrets = secrets;
            SessionNetworkConfiguration = sessionNetworkConfiguration;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
