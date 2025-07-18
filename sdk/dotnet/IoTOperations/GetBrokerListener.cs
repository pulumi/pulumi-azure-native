// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperations
{
    public static class GetBrokerListener
    {
        /// <summary>
        /// Get a BrokerListenerResource
        /// 
        /// Uses Azure REST API version 2024-11-01.
        /// 
        /// Other available API versions: 2024-07-01-preview, 2024-08-15-preview, 2024-09-15-preview, 2025-04-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native iotoperations [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetBrokerListenerResult> InvokeAsync(GetBrokerListenerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBrokerListenerResult>("azure-native:iotoperations:getBrokerListener", args ?? new GetBrokerListenerArgs(), options.WithDefaults());

        /// <summary>
        /// Get a BrokerListenerResource
        /// 
        /// Uses Azure REST API version 2024-11-01.
        /// 
        /// Other available API versions: 2024-07-01-preview, 2024-08-15-preview, 2024-09-15-preview, 2025-04-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native iotoperations [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetBrokerListenerResult> Invoke(GetBrokerListenerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBrokerListenerResult>("azure-native:iotoperations:getBrokerListener", args ?? new GetBrokerListenerInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a BrokerListenerResource
        /// 
        /// Uses Azure REST API version 2024-11-01.
        /// 
        /// Other available API versions: 2024-07-01-preview, 2024-08-15-preview, 2024-09-15-preview, 2025-04-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native iotoperations [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetBrokerListenerResult> Invoke(GetBrokerListenerInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetBrokerListenerResult>("azure-native:iotoperations:getBrokerListener", args ?? new GetBrokerListenerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBrokerListenerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of broker.
        /// </summary>
        [Input("brokerName", required: true)]
        public string BrokerName { get; set; } = null!;

        /// <summary>
        /// Name of instance.
        /// </summary>
        [Input("instanceName", required: true)]
        public string InstanceName { get; set; } = null!;

        /// <summary>
        /// Name of Instance broker listener resource
        /// </summary>
        [Input("listenerName", required: true)]
        public string ListenerName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetBrokerListenerArgs()
        {
        }
        public static new GetBrokerListenerArgs Empty => new GetBrokerListenerArgs();
    }

    public sealed class GetBrokerListenerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of broker.
        /// </summary>
        [Input("brokerName", required: true)]
        public Input<string> BrokerName { get; set; } = null!;

        /// <summary>
        /// Name of instance.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// Name of Instance broker listener resource
        /// </summary>
        [Input("listenerName", required: true)]
        public Input<string> ListenerName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetBrokerListenerInvokeArgs()
        {
        }
        public static new GetBrokerListenerInvokeArgs Empty => new GetBrokerListenerInvokeArgs();
    }


    [OutputType]
    public sealed class GetBrokerListenerResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Edge location of the resource.
        /// </summary>
        public readonly Outputs.ExtendedLocationResponse ExtendedLocation;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource-specific properties for this resource.
        /// </summary>
        public readonly Outputs.BrokerListenerPropertiesResponse Properties;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetBrokerListenerResult(
            string azureApiVersion,

            Outputs.ExtendedLocationResponse extendedLocation,

            string id,

            string name,

            Outputs.BrokerListenerPropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            ExtendedLocation = extendedLocation;
            Id = id;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Type = type;
        }
    }
}
