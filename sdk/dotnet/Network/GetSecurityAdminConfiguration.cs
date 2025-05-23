// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetSecurityAdminConfiguration
    {
        /// <summary>
        /// Retrieves a network manager security admin configuration.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2021-02-01-preview, 2022-01-01, 2022-02-01-preview, 2022-04-01-preview, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-01-01-preview, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetSecurityAdminConfigurationResult> InvokeAsync(GetSecurityAdminConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecurityAdminConfigurationResult>("azure-native:network:getSecurityAdminConfiguration", args ?? new GetSecurityAdminConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a network manager security admin configuration.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2021-02-01-preview, 2022-01-01, 2022-02-01-preview, 2022-04-01-preview, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-01-01-preview, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSecurityAdminConfigurationResult> Invoke(GetSecurityAdminConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityAdminConfigurationResult>("azure-native:network:getSecurityAdminConfiguration", args ?? new GetSecurityAdminConfigurationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a network manager security admin configuration.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2021-02-01-preview, 2022-01-01, 2022-02-01-preview, 2022-04-01-preview, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-01-01-preview, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSecurityAdminConfigurationResult> Invoke(GetSecurityAdminConfigurationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityAdminConfigurationResult>("azure-native:network:getSecurityAdminConfiguration", args ?? new GetSecurityAdminConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityAdminConfigurationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the network manager Security Configuration.
        /// </summary>
        [Input("configurationName", required: true)]
        public string ConfigurationName { get; set; } = null!;

        /// <summary>
        /// The name of the network manager.
        /// </summary>
        [Input("networkManagerName", required: true)]
        public string NetworkManagerName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetSecurityAdminConfigurationArgs()
        {
        }
        public static new GetSecurityAdminConfigurationArgs Empty => new GetSecurityAdminConfigurationArgs();
    }

    public sealed class GetSecurityAdminConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the network manager Security Configuration.
        /// </summary>
        [Input("configurationName", required: true)]
        public Input<string> ConfigurationName { get; set; } = null!;

        /// <summary>
        /// The name of the network manager.
        /// </summary>
        [Input("networkManagerName", required: true)]
        public Input<string> NetworkManagerName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetSecurityAdminConfigurationInvokeArgs()
        {
        }
        public static new GetSecurityAdminConfigurationInvokeArgs Empty => new GetSecurityAdminConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecurityAdminConfigurationResult
    {
        /// <summary>
        /// Enum list of network intent policy based services.
        /// </summary>
        public readonly ImmutableArray<string> ApplyOnNetworkIntentPolicyBasedServices;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// A description of the security configuration.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Determine update behavior for changes to network groups referenced within the rules in this configuration.
        /// </summary>
        public readonly string? NetworkGroupAddressSpaceAggregationOption;
        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Unique identifier for this resource.
        /// </summary>
        public readonly string ResourceGuid;
        /// <summary>
        /// The system metadata related to this resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSecurityAdminConfigurationResult(
            ImmutableArray<string> applyOnNetworkIntentPolicyBasedServices,

            string azureApiVersion,

            string? description,

            string etag,

            string id,

            string name,

            string? networkGroupAddressSpaceAggregationOption,

            string provisioningState,

            string resourceGuid,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            ApplyOnNetworkIntentPolicyBasedServices = applyOnNetworkIntentPolicyBasedServices;
            AzureApiVersion = azureApiVersion;
            Description = description;
            Etag = etag;
            Id = id;
            Name = name;
            NetworkGroupAddressSpaceAggregationOption = networkGroupAddressSpaceAggregationOption;
            ProvisioningState = provisioningState;
            ResourceGuid = resourceGuid;
            SystemData = systemData;
            Type = type;
        }
    }
}
