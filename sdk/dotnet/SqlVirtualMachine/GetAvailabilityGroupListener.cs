// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SqlVirtualMachine
{
    public static class GetAvailabilityGroupListener
    {
        /// <summary>
        /// Gets an availability group listener.
        /// 
        /// Uses Azure REST API version 2023-10-01.
        /// 
        /// Other available API versions: 2022-02-01, 2022-07-01-preview, 2022-08-01-preview, 2023-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sqlvirtualmachine [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetAvailabilityGroupListenerResult> InvokeAsync(GetAvailabilityGroupListenerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAvailabilityGroupListenerResult>("azure-native:sqlvirtualmachine:getAvailabilityGroupListener", args ?? new GetAvailabilityGroupListenerArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an availability group listener.
        /// 
        /// Uses Azure REST API version 2023-10-01.
        /// 
        /// Other available API versions: 2022-02-01, 2022-07-01-preview, 2022-08-01-preview, 2023-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sqlvirtualmachine [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAvailabilityGroupListenerResult> Invoke(GetAvailabilityGroupListenerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAvailabilityGroupListenerResult>("azure-native:sqlvirtualmachine:getAvailabilityGroupListener", args ?? new GetAvailabilityGroupListenerInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an availability group listener.
        /// 
        /// Uses Azure REST API version 2023-10-01.
        /// 
        /// Other available API versions: 2022-02-01, 2022-07-01-preview, 2022-08-01-preview, 2023-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sqlvirtualmachine [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAvailabilityGroupListenerResult> Invoke(GetAvailabilityGroupListenerInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAvailabilityGroupListenerResult>("azure-native:sqlvirtualmachine:getAvailabilityGroupListener", args ?? new GetAvailabilityGroupListenerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAvailabilityGroupListenerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the availability group listener.
        /// </summary>
        [Input("availabilityGroupListenerName", required: true)]
        public string AvailabilityGroupListenerName { get; set; } = null!;

        /// <summary>
        /// The child resources to include in the response.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the SQL virtual machine group.
        /// </summary>
        [Input("sqlVirtualMachineGroupName", required: true)]
        public string SqlVirtualMachineGroupName { get; set; } = null!;

        public GetAvailabilityGroupListenerArgs()
        {
        }
        public static new GetAvailabilityGroupListenerArgs Empty => new GetAvailabilityGroupListenerArgs();
    }

    public sealed class GetAvailabilityGroupListenerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the availability group listener.
        /// </summary>
        [Input("availabilityGroupListenerName", required: true)]
        public Input<string> AvailabilityGroupListenerName { get; set; } = null!;

        /// <summary>
        /// The child resources to include in the response.
        /// </summary>
        [Input("expand")]
        public Input<string>? Expand { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the SQL virtual machine group.
        /// </summary>
        [Input("sqlVirtualMachineGroupName", required: true)]
        public Input<string> SqlVirtualMachineGroupName { get; set; } = null!;

        public GetAvailabilityGroupListenerInvokeArgs()
        {
        }
        public static new GetAvailabilityGroupListenerInvokeArgs Empty => new GetAvailabilityGroupListenerInvokeArgs();
    }


    [OutputType]
    public sealed class GetAvailabilityGroupListenerResult
    {
        /// <summary>
        /// Availability Group configuration.
        /// </summary>
        public readonly Outputs.AgConfigurationResponse? AvailabilityGroupConfiguration;
        /// <summary>
        /// Name of the availability group.
        /// </summary>
        public readonly string? AvailabilityGroupName;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Create a default availability group if it does not exist.
        /// </summary>
        public readonly bool? CreateDefaultAvailabilityGroupIfNotExist;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of load balancer configurations for an availability group listener.
        /// </summary>
        public readonly ImmutableArray<Outputs.LoadBalancerConfigurationResponse> LoadBalancerConfigurations;
        /// <summary>
        /// List of multi subnet IP configurations for an AG listener.
        /// </summary>
        public readonly ImmutableArray<Outputs.MultiSubnetIpConfigurationResponse> MultiSubnetIpConfigurations;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Listener port.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Provisioning state to track the async operation status.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAvailabilityGroupListenerResult(
            Outputs.AgConfigurationResponse? availabilityGroupConfiguration,

            string? availabilityGroupName,

            string azureApiVersion,

            bool? createDefaultAvailabilityGroupIfNotExist,

            string id,

            ImmutableArray<Outputs.LoadBalancerConfigurationResponse> loadBalancerConfigurations,

            ImmutableArray<Outputs.MultiSubnetIpConfigurationResponse> multiSubnetIpConfigurations,

            string name,

            int? port,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AvailabilityGroupConfiguration = availabilityGroupConfiguration;
            AvailabilityGroupName = availabilityGroupName;
            AzureApiVersion = azureApiVersion;
            CreateDefaultAvailabilityGroupIfNotExist = createDefaultAvailabilityGroupIfNotExist;
            Id = id;
            LoadBalancerConfigurations = loadBalancerConfigurations;
            MultiSubnetIpConfigurations = multiSubnetIpConfigurations;
            Name = name;
            Port = port;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Type = type;
        }
    }
}
