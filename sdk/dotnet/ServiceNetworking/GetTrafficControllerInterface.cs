// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceNetworking
{
    public static class GetTrafficControllerInterface
    {
        /// <summary>
        /// Get a TrafficController
        /// 
        /// Uses Azure REST API version 2025-01-01.
        /// 
        /// Other available API versions: 2023-05-01-preview, 2023-11-01, 2024-05-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native servicenetworking [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetTrafficControllerInterfaceResult> InvokeAsync(GetTrafficControllerInterfaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTrafficControllerInterfaceResult>("azure-native:servicenetworking:getTrafficControllerInterface", args ?? new GetTrafficControllerInterfaceArgs(), options.WithDefaults());

        /// <summary>
        /// Get a TrafficController
        /// 
        /// Uses Azure REST API version 2025-01-01.
        /// 
        /// Other available API versions: 2023-05-01-preview, 2023-11-01, 2024-05-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native servicenetworking [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTrafficControllerInterfaceResult> Invoke(GetTrafficControllerInterfaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTrafficControllerInterfaceResult>("azure-native:servicenetworking:getTrafficControllerInterface", args ?? new GetTrafficControllerInterfaceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a TrafficController
        /// 
        /// Uses Azure REST API version 2025-01-01.
        /// 
        /// Other available API versions: 2023-05-01-preview, 2023-11-01, 2024-05-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native servicenetworking [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTrafficControllerInterfaceResult> Invoke(GetTrafficControllerInterfaceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTrafficControllerInterfaceResult>("azure-native:servicenetworking:getTrafficControllerInterface", args ?? new GetTrafficControllerInterfaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTrafficControllerInterfaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// traffic controller name for path
        /// </summary>
        [Input("trafficControllerName", required: true)]
        public string TrafficControllerName { get; set; } = null!;

        public GetTrafficControllerInterfaceArgs()
        {
        }
        public static new GetTrafficControllerInterfaceArgs Empty => new GetTrafficControllerInterfaceArgs();
    }

    public sealed class GetTrafficControllerInterfaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// traffic controller name for path
        /// </summary>
        [Input("trafficControllerName", required: true)]
        public Input<string> TrafficControllerName { get; set; } = null!;

        public GetTrafficControllerInterfaceInvokeArgs()
        {
        }
        public static new GetTrafficControllerInterfaceInvokeArgs Empty => new GetTrafficControllerInterfaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetTrafficControllerInterfaceResult
    {
        /// <summary>
        /// Associations References List
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceIdResponse> Associations;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Configuration Endpoints.
        /// </summary>
        public readonly ImmutableArray<string> ConfigurationEndpoints;
        /// <summary>
        /// Frontends References List
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceIdResponse> Frontends;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The status of the last operation.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Security Policies References List
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceIdResponse> SecurityPolicies;
        /// <summary>
        /// Security Policy Configuration
        /// </summary>
        public readonly Outputs.SecurityPolicyConfigurationsResponse? SecurityPolicyConfigurations;
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
        private GetTrafficControllerInterfaceResult(
            ImmutableArray<Outputs.ResourceIdResponse> associations,

            string azureApiVersion,

            ImmutableArray<string> configurationEndpoints,

            ImmutableArray<Outputs.ResourceIdResponse> frontends,

            string id,

            string location,

            string name,

            string provisioningState,

            ImmutableArray<Outputs.ResourceIdResponse> securityPolicies,

            Outputs.SecurityPolicyConfigurationsResponse? securityPolicyConfigurations,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Associations = associations;
            AzureApiVersion = azureApiVersion;
            ConfigurationEndpoints = configurationEndpoints;
            Frontends = frontends;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            SecurityPolicies = securityPolicies;
            SecurityPolicyConfigurations = securityPolicyConfigurations;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
