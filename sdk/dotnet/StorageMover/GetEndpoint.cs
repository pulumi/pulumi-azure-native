// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageMover
{
    public static class GetEndpoint
    {
        /// <summary>
        /// Gets an Endpoint resource.
        /// 
        /// Uses Azure REST API version 2023-03-01.
        /// 
        /// Other available API versions: 2023-07-01-preview, 2023-10-01, 2024-07-01.
        /// </summary>
        public static Task<GetEndpointResult> InvokeAsync(GetEndpointArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEndpointResult>("azure-native:storagemover:getEndpoint", args ?? new GetEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an Endpoint resource.
        /// 
        /// Uses Azure REST API version 2023-03-01.
        /// 
        /// Other available API versions: 2023-07-01-preview, 2023-10-01, 2024-07-01.
        /// </summary>
        public static Output<GetEndpointResult> Invoke(GetEndpointInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEndpointResult>("azure-native:storagemover:getEndpoint", args ?? new GetEndpointInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an Endpoint resource.
        /// 
        /// Uses Azure REST API version 2023-03-01.
        /// 
        /// Other available API versions: 2023-07-01-preview, 2023-10-01, 2024-07-01.
        /// </summary>
        public static Output<GetEndpointResult> Invoke(GetEndpointInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetEndpointResult>("azure-native:storagemover:getEndpoint", args ?? new GetEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEndpointArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Endpoint resource.
        /// </summary>
        [Input("endpointName", required: true)]
        public string EndpointName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Storage Mover resource.
        /// </summary>
        [Input("storageMoverName", required: true)]
        public string StorageMoverName { get; set; } = null!;

        public GetEndpointArgs()
        {
        }
        public static new GetEndpointArgs Empty => new GetEndpointArgs();
    }

    public sealed class GetEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Endpoint resource.
        /// </summary>
        [Input("endpointName", required: true)]
        public Input<string> EndpointName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Storage Mover resource.
        /// </summary>
        [Input("storageMoverName", required: true)]
        public Input<string> StorageMoverName { get; set; } = null!;

        public GetEndpointInvokeArgs()
        {
        }
        public static new GetEndpointInvokeArgs Empty => new GetEndpointInvokeArgs();
    }


    [OutputType]
    public sealed class GetEndpointResult
    {
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource specific properties for the Storage Mover resource.
        /// </summary>
        public readonly Union<Outputs.AzureStorageBlobContainerEndpointPropertiesResponse, Outputs.NfsMountEndpointPropertiesResponse> Properties;
        /// <summary>
        /// Resource system metadata.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetEndpointResult(
            string id,

            string name,

            Union<Outputs.AzureStorageBlobContainerEndpointPropertiesResponse, Outputs.NfsMountEndpointPropertiesResponse> properties,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            Id = id;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Type = type;
        }
    }
}
