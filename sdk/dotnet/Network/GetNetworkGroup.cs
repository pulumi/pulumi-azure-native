// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetNetworkGroup
    {
        /// <summary>
        /// Gets the specified network group.
        /// 
        /// Uses Azure REST API version 2023-02-01.
        /// 
        /// Other available API versions: 2021-02-01-preview, 2021-05-01-preview, 2022-04-01-preview, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
        /// </summary>
        public static Task<GetNetworkGroupResult> InvokeAsync(GetNetworkGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkGroupResult>("azure-native:network:getNetworkGroup", args ?? new GetNetworkGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified network group.
        /// 
        /// Uses Azure REST API version 2023-02-01.
        /// 
        /// Other available API versions: 2021-02-01-preview, 2021-05-01-preview, 2022-04-01-preview, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
        /// </summary>
        public static Output<GetNetworkGroupResult> Invoke(GetNetworkGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkGroupResult>("azure-native:network:getNetworkGroup", args ?? new GetNetworkGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified network group.
        /// 
        /// Uses Azure REST API version 2023-02-01.
        /// 
        /// Other available API versions: 2021-02-01-preview, 2021-05-01-preview, 2022-04-01-preview, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
        /// </summary>
        public static Output<GetNetworkGroupResult> Invoke(GetNetworkGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkGroupResult>("azure-native:network:getNetworkGroup", args ?? new GetNetworkGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the network group.
        /// </summary>
        [Input("networkGroupName", required: true)]
        public string NetworkGroupName { get; set; } = null!;

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

        public GetNetworkGroupArgs()
        {
        }
        public static new GetNetworkGroupArgs Empty => new GetNetworkGroupArgs();
    }

    public sealed class GetNetworkGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the network group.
        /// </summary>
        [Input("networkGroupName", required: true)]
        public Input<string> NetworkGroupName { get; set; } = null!;

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

        public GetNetworkGroupInvokeArgs()
        {
        }
        public static new GetNetworkGroupInvokeArgs Empty => new GetNetworkGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkGroupResult
    {
        /// <summary>
        /// A description of the network group.
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
        /// The provisioning state of the scope assignment resource.
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
        private GetNetworkGroupResult(
            string? description,

            string etag,

            string id,

            string name,

            string provisioningState,

            string resourceGuid,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            Description = description;
            Etag = etag;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            ResourceGuid = resourceGuid;
            SystemData = systemData;
            Type = type;
        }
    }
}
