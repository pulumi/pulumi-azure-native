// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric
{
    public static class GetAccessControlList
    {
        /// <summary>
        /// Implements Access Control List GET method.
        /// 
        /// Uses Azure REST API version 2023-06-15.
        /// 
        /// Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetAccessControlListResult> InvokeAsync(GetAccessControlListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccessControlListResult>("azure-native:managednetworkfabric:getAccessControlList", args ?? new GetAccessControlListArgs(), options.WithDefaults());

        /// <summary>
        /// Implements Access Control List GET method.
        /// 
        /// Uses Azure REST API version 2023-06-15.
        /// 
        /// Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAccessControlListResult> Invoke(GetAccessControlListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessControlListResult>("azure-native:managednetworkfabric:getAccessControlList", args ?? new GetAccessControlListInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Implements Access Control List GET method.
        /// 
        /// Uses Azure REST API version 2023-06-15.
        /// 
        /// Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAccessControlListResult> Invoke(GetAccessControlListInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessControlListResult>("azure-native:managednetworkfabric:getAccessControlList", args ?? new GetAccessControlListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccessControlListArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Access Control List.
        /// </summary>
        [Input("accessControlListName", required: true)]
        public string AccessControlListName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAccessControlListArgs()
        {
        }
        public static new GetAccessControlListArgs Empty => new GetAccessControlListArgs();
    }

    public sealed class GetAccessControlListInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Access Control List.
        /// </summary>
        [Input("accessControlListName", required: true)]
        public Input<string> AccessControlListName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAccessControlListInvokeArgs()
        {
        }
        public static new GetAccessControlListInvokeArgs Empty => new GetAccessControlListInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccessControlListResult
    {
        /// <summary>
        /// Access Control List file URL.
        /// </summary>
        public readonly string? AclsUrl;
        /// <summary>
        /// Administrative state of the resource.
        /// </summary>
        public readonly string AdministrativeState;
        /// <summary>
        /// Switch configuration description.
        /// </summary>
        public readonly string? Annotation;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Configuration state of the resource.
        /// </summary>
        public readonly string ConfigurationState;
        /// <summary>
        /// Input method to configure Access Control List.
        /// </summary>
        public readonly string ConfigurationType;
        /// <summary>
        /// Default action that needs to be applied when no condition is matched. Example: Permit | Deny.
        /// </summary>
        public readonly string? DefaultAction;
        /// <summary>
        /// List of dynamic match configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.CommonDynamicMatchConfigurationResponse> DynamicMatchConfigurations;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The last synced timestamp.
        /// </summary>
        public readonly string LastSyncedTime;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// List of match configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.AccessControlListMatchConfigurationResponse> MatchConfigurations;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
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
        private GetAccessControlListResult(
            string? aclsUrl,

            string administrativeState,

            string? annotation,

            string azureApiVersion,

            string configurationState,

            string configurationType,

            string? defaultAction,

            ImmutableArray<Outputs.CommonDynamicMatchConfigurationResponse> dynamicMatchConfigurations,

            string id,

            string lastSyncedTime,

            string location,

            ImmutableArray<Outputs.AccessControlListMatchConfigurationResponse> matchConfigurations,

            string name,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AclsUrl = aclsUrl;
            AdministrativeState = administrativeState;
            Annotation = annotation;
            AzureApiVersion = azureApiVersion;
            ConfigurationState = configurationState;
            ConfigurationType = configurationType;
            DefaultAction = defaultAction;
            DynamicMatchConfigurations = dynamicMatchConfigurations;
            Id = id;
            LastSyncedTime = lastSyncedTime;
            Location = location;
            MatchConfigurations = matchConfigurations;
            Name = name;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
