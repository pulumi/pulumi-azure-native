// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridCompute.V20220510Preview
{
    public static class GetMachine
    {
        /// <summary>
        /// Retrieves information about the model view or the instance view of a hybrid machine.
        /// </summary>
        public static Task<GetMachineResult> InvokeAsync(GetMachineArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMachineResult>("azure-native:hybridcompute/v20220510preview:getMachine", args ?? new GetMachineArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves information about the model view or the instance view of a hybrid machine.
        /// </summary>
        public static Output<GetMachineResult> Invoke(GetMachineInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMachineResult>("azure-native:hybridcompute/v20220510preview:getMachine", args ?? new GetMachineInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves information about the model view or the instance view of a hybrid machine.
        /// </summary>
        public static Output<GetMachineResult> Invoke(GetMachineInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMachineResult>("azure-native:hybridcompute/v20220510preview:getMachine", args ?? new GetMachineInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMachineArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The expand expression to apply on the operation.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the hybrid machine.
        /// </summary>
        [Input("machineName", required: true)]
        public string MachineName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetMachineArgs()
        {
        }
        public static new GetMachineArgs Empty => new GetMachineArgs();
    }

    public sealed class GetMachineInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The expand expression to apply on the operation.
        /// </summary>
        [Input("expand")]
        public Input<string>? Expand { get; set; }

        /// <summary>
        /// The name of the hybrid machine.
        /// </summary>
        [Input("machineName", required: true)]
        public Input<string> MachineName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetMachineInvokeArgs()
        {
        }
        public static new GetMachineInvokeArgs Empty => new GetMachineInvokeArgs();
    }


    [OutputType]
    public sealed class GetMachineResult
    {
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Identity for the resource.
        /// </summary>
        public readonly Outputs.IdentityResponse? Identity;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Hybrid Compute Machine properties
        /// </summary>
        public readonly Outputs.MachinePropertiesResponse Properties;
        /// <summary>
        /// The list of extensions affiliated to the machine
        /// </summary>
        public readonly ImmutableArray<Outputs.MachineExtensionResponse> Resources;
        /// <summary>
        /// The system meta data relating to this resource.
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
        private GetMachineResult(
            string id,

            Outputs.IdentityResponse? identity,

            string location,

            string name,

            Outputs.MachinePropertiesResponse properties,

            ImmutableArray<Outputs.MachineExtensionResponse> resources,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Id = id;
            Identity = identity;
            Location = location;
            Name = name;
            Properties = properties;
            Resources = resources;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
