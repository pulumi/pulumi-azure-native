// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20240101
{
    public static class GetNetworkProfile
    {
        /// <summary>
        /// Gets the specified network profile in a specified resource group.
        /// </summary>
        public static Task<GetNetworkProfileResult> InvokeAsync(GetNetworkProfileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkProfileResult>("azure-native:network/v20240101:getNetworkProfile", args ?? new GetNetworkProfileArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified network profile in a specified resource group.
        /// </summary>
        public static Output<GetNetworkProfileResult> Invoke(GetNetworkProfileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkProfileResult>("azure-native:network/v20240101:getNetworkProfile", args ?? new GetNetworkProfileInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified network profile in a specified resource group.
        /// </summary>
        public static Output<GetNetworkProfileResult> Invoke(GetNetworkProfileInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkProfileResult>("azure-native:network/v20240101:getNetworkProfile", args ?? new GetNetworkProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkProfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Expands referenced resources.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the public IP prefix.
        /// </summary>
        [Input("networkProfileName", required: true)]
        public string NetworkProfileName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNetworkProfileArgs()
        {
        }
        public static new GetNetworkProfileArgs Empty => new GetNetworkProfileArgs();
    }

    public sealed class GetNetworkProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Expands referenced resources.
        /// </summary>
        [Input("expand")]
        public Input<string>? Expand { get; set; }

        /// <summary>
        /// The name of the public IP prefix.
        /// </summary>
        [Input("networkProfileName", required: true)]
        public Input<string> NetworkProfileName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetNetworkProfileInvokeArgs()
        {
        }
        public static new GetNetworkProfileInvokeArgs Empty => new GetNetworkProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkProfileResult
    {
        /// <summary>
        /// List of chid container network interface configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerNetworkInterfaceConfigurationResponse> ContainerNetworkInterfaceConfigurations;
        /// <summary>
        /// List of child container network interfaces.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerNetworkInterfaceResponse> ContainerNetworkInterfaces;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the network profile resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The resource GUID property of the network profile resource.
        /// </summary>
        public readonly string ResourceGuid;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetNetworkProfileResult(
            ImmutableArray<Outputs.ContainerNetworkInterfaceConfigurationResponse> containerNetworkInterfaceConfigurations,

            ImmutableArray<Outputs.ContainerNetworkInterfaceResponse> containerNetworkInterfaces,

            string etag,

            string? id,

            string? location,

            string name,

            string provisioningState,

            string resourceGuid,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            ContainerNetworkInterfaceConfigurations = containerNetworkInterfaceConfigurations;
            ContainerNetworkInterfaces = containerNetworkInterfaces;
            Etag = etag;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            ResourceGuid = resourceGuid;
            Tags = tags;
            Type = type;
        }
    }
}
