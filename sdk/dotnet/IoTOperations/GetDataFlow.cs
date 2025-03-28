// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperations
{
    public static class GetDataFlow
    {
        /// <summary>
        /// Get a DataFlowResource
        /// 
        /// Uses Azure REST API version 2024-07-01-preview.
        /// </summary>
        public static Task<GetDataFlowResult> InvokeAsync(GetDataFlowArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDataFlowResult>("azure-native:iotoperations:getDataFlow", args ?? new GetDataFlowArgs(), options.WithDefaults());

        /// <summary>
        /// Get a DataFlowResource
        /// 
        /// Uses Azure REST API version 2024-07-01-preview.
        /// </summary>
        public static Output<GetDataFlowResult> Invoke(GetDataFlowInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataFlowResult>("azure-native:iotoperations:getDataFlow", args ?? new GetDataFlowInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a DataFlowResource
        /// 
        /// Uses Azure REST API version 2024-07-01-preview.
        /// </summary>
        public static Output<GetDataFlowResult> Invoke(GetDataFlowInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataFlowResult>("azure-native:iotoperations:getDataFlow", args ?? new GetDataFlowInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataFlowArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of Instance dataflowProfile dataflow resource
        /// </summary>
        [Input("dataflowName", required: true)]
        public string DataflowName { get; set; } = null!;

        /// <summary>
        /// Name of Instance dataflowProfile resource
        /// </summary>
        [Input("dataflowProfileName", required: true)]
        public string DataflowProfileName { get; set; } = null!;

        /// <summary>
        /// Name of instance.
        /// </summary>
        [Input("instanceName", required: true)]
        public string InstanceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDataFlowArgs()
        {
        }
        public static new GetDataFlowArgs Empty => new GetDataFlowArgs();
    }

    public sealed class GetDataFlowInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of Instance dataflowProfile dataflow resource
        /// </summary>
        [Input("dataflowName", required: true)]
        public Input<string> DataflowName { get; set; } = null!;

        /// <summary>
        /// Name of Instance dataflowProfile resource
        /// </summary>
        [Input("dataflowProfileName", required: true)]
        public Input<string> DataflowProfileName { get; set; } = null!;

        /// <summary>
        /// Name of instance.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetDataFlowInvokeArgs()
        {
        }
        public static new GetDataFlowInvokeArgs Empty => new GetDataFlowInvokeArgs();
    }


    [OutputType]
    public sealed class GetDataFlowResult
    {
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
        public readonly Outputs.DataFlowPropertiesResponse Properties;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDataFlowResult(
            Outputs.ExtendedLocationResponse extendedLocation,

            string id,

            string name,

            Outputs.DataFlowPropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            ExtendedLocation = extendedLocation;
            Id = id;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Type = type;
        }
    }
}
