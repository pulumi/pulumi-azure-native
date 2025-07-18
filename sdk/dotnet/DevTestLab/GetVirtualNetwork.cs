// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab
{
    public static class GetVirtualNetwork
    {
        /// <summary>
        /// Get virtual network.
        /// 
        /// Uses Azure REST API version 2018-09-15.
        /// </summary>
        public static Task<GetVirtualNetworkResult> InvokeAsync(GetVirtualNetworkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualNetworkResult>("azure-native:devtestlab:getVirtualNetwork", args ?? new GetVirtualNetworkArgs(), options.WithDefaults());

        /// <summary>
        /// Get virtual network.
        /// 
        /// Uses Azure REST API version 2018-09-15.
        /// </summary>
        public static Output<GetVirtualNetworkResult> Invoke(GetVirtualNetworkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualNetworkResult>("azure-native:devtestlab:getVirtualNetwork", args ?? new GetVirtualNetworkInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get virtual network.
        /// 
        /// Uses Azure REST API version 2018-09-15.
        /// </summary>
        public static Output<GetVirtualNetworkResult> Invoke(GetVirtualNetworkInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualNetworkResult>("azure-native:devtestlab:getVirtualNetwork", args ?? new GetVirtualNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualNetworkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the $expand query. Example: 'properties($expand=externalSubnets)'
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the lab.
        /// </summary>
        [Input("labName", required: true)]
        public string LabName { get; set; } = null!;

        /// <summary>
        /// The name of the VirtualNetwork
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetVirtualNetworkArgs()
        {
        }
        public static new GetVirtualNetworkArgs Empty => new GetVirtualNetworkArgs();
    }

    public sealed class GetVirtualNetworkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the $expand query. Example: 'properties($expand=externalSubnets)'
        /// </summary>
        [Input("expand")]
        public Input<string>? Expand { get; set; }

        /// <summary>
        /// The name of the lab.
        /// </summary>
        [Input("labName", required: true)]
        public Input<string> LabName { get; set; } = null!;

        /// <summary>
        /// The name of the VirtualNetwork
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetVirtualNetworkInvokeArgs()
        {
        }
        public static new GetVirtualNetworkInvokeArgs Empty => new GetVirtualNetworkInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualNetworkResult
    {
        /// <summary>
        /// The allowed subnets of the virtual network.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubnetResponse> AllowedSubnets;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The creation date of the virtual network.
        /// </summary>
        public readonly string CreatedDate;
        /// <summary>
        /// The description of the virtual network.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The Microsoft.Network resource identifier of the virtual network.
        /// </summary>
        public readonly string? ExternalProviderResourceId;
        /// <summary>
        /// The external subnet properties.
        /// </summary>
        public readonly ImmutableArray<Outputs.ExternalSubnetResponse> ExternalSubnets;
        /// <summary>
        /// The identifier of the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The location of the resource.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning status of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The subnet overrides of the virtual network.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubnetOverrideResponse> SubnetOverrides;
        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The unique immutable identifier of a resource (Guid).
        /// </summary>
        public readonly string UniqueIdentifier;

        [OutputConstructor]
        private GetVirtualNetworkResult(
            ImmutableArray<Outputs.SubnetResponse> allowedSubnets,

            string azureApiVersion,

            string createdDate,

            string? description,

            string? externalProviderResourceId,

            ImmutableArray<Outputs.ExternalSubnetResponse> externalSubnets,

            string id,

            string? location,

            string name,

            string provisioningState,

            ImmutableArray<Outputs.SubnetOverrideResponse> subnetOverrides,

            ImmutableDictionary<string, string>? tags,

            string type,

            string uniqueIdentifier)
        {
            AllowedSubnets = allowedSubnets;
            AzureApiVersion = azureApiVersion;
            CreatedDate = createdDate;
            Description = description;
            ExternalProviderResourceId = externalProviderResourceId;
            ExternalSubnets = externalSubnets;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            SubnetOverrides = subnetOverrides;
            Tags = tags;
            Type = type;
            UniqueIdentifier = uniqueIdentifier;
        }
    }
}
