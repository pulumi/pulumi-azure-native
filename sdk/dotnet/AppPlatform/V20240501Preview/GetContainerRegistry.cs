// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform.V20240501Preview
{
    public static class GetContainerRegistry
    {
        /// <summary>
        /// Get the container registries resource.
        /// </summary>
        public static Task<GetContainerRegistryResult> InvokeAsync(GetContainerRegistryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContainerRegistryResult>("azure-native:appplatform/v20240501preview:getContainerRegistry", args ?? new GetContainerRegistryArgs(), options.WithDefaults());

        /// <summary>
        /// Get the container registries resource.
        /// </summary>
        public static Output<GetContainerRegistryResult> Invoke(GetContainerRegistryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContainerRegistryResult>("azure-native:appplatform/v20240501preview:getContainerRegistry", args ?? new GetContainerRegistryInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the container registries resource.
        /// </summary>
        public static Output<GetContainerRegistryResult> Invoke(GetContainerRegistryInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetContainerRegistryResult>("azure-native:appplatform/v20240501preview:getContainerRegistry", args ?? new GetContainerRegistryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContainerRegistryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("containerRegistryName", required: true)]
        public string ContainerRegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Service resource.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetContainerRegistryArgs()
        {
        }
        public static new GetContainerRegistryArgs Empty => new GetContainerRegistryArgs();
    }

    public sealed class GetContainerRegistryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("containerRegistryName", required: true)]
        public Input<string> ContainerRegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Service resource.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetContainerRegistryInvokeArgs()
        {
        }
        public static new GetContainerRegistryInvokeArgs Empty => new GetContainerRegistryInvokeArgs();
    }


    [OutputType]
    public sealed class GetContainerRegistryResult
    {
        /// <summary>
        /// Fully qualified resource Id for the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the container registry resource payload.
        /// </summary>
        public readonly Outputs.ContainerRegistryPropertiesResponse Properties;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetContainerRegistryResult(
            string id,

            string name,

            Outputs.ContainerRegistryPropertiesResponse properties,

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
