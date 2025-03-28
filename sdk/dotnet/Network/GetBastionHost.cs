// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetBastionHost
    {
        /// <summary>
        /// Gets the specified Bastion Host.
        /// 
        /// Uses Azure REST API version 2023-02-01.
        /// 
        /// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
        /// </summary>
        public static Task<GetBastionHostResult> InvokeAsync(GetBastionHostArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBastionHostResult>("azure-native:network:getBastionHost", args ?? new GetBastionHostArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified Bastion Host.
        /// 
        /// Uses Azure REST API version 2023-02-01.
        /// 
        /// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
        /// </summary>
        public static Output<GetBastionHostResult> Invoke(GetBastionHostInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBastionHostResult>("azure-native:network:getBastionHost", args ?? new GetBastionHostInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified Bastion Host.
        /// 
        /// Uses Azure REST API version 2023-02-01.
        /// 
        /// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
        /// </summary>
        public static Output<GetBastionHostResult> Invoke(GetBastionHostInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetBastionHostResult>("azure-native:network:getBastionHost", args ?? new GetBastionHostInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBastionHostArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Bastion Host.
        /// </summary>
        [Input("bastionHostName", required: true)]
        public string BastionHostName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetBastionHostArgs()
        {
        }
        public static new GetBastionHostArgs Empty => new GetBastionHostArgs();
    }

    public sealed class GetBastionHostInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Bastion Host.
        /// </summary>
        [Input("bastionHostName", required: true)]
        public Input<string> BastionHostName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetBastionHostInvokeArgs()
        {
        }
        public static new GetBastionHostInvokeArgs Empty => new GetBastionHostInvokeArgs();
    }


    [OutputType]
    public sealed class GetBastionHostResult
    {
        /// <summary>
        /// Enable/Disable Copy/Paste feature of the Bastion Host resource.
        /// </summary>
        public readonly bool? DisableCopyPaste;
        /// <summary>
        /// FQDN for the endpoint on which bastion host is accessible.
        /// </summary>
        public readonly string? DnsName;
        /// <summary>
        /// Enable/Disable File Copy feature of the Bastion Host resource.
        /// </summary>
        public readonly bool? EnableFileCopy;
        /// <summary>
        /// Enable/Disable IP Connect feature of the Bastion Host resource.
        /// </summary>
        public readonly bool? EnableIpConnect;
        /// <summary>
        /// Enable/Disable Kerberos feature of the Bastion Host resource.
        /// </summary>
        public readonly bool? EnableKerberos;
        /// <summary>
        /// Enable/Disable Shareable Link of the Bastion Host resource.
        /// </summary>
        public readonly bool? EnableShareableLink;
        /// <summary>
        /// Enable/Disable Tunneling feature of the Bastion Host resource.
        /// </summary>
        public readonly bool? EnableTunneling;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// IP configuration of the Bastion Host resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.BastionHostIPConfigurationResponse> IpConfigurations;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the bastion host resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The scale units for the Bastion Host resource.
        /// </summary>
        public readonly int? ScaleUnits;
        /// <summary>
        /// The sku of this Bastion Host.
        /// </summary>
        public readonly Outputs.SkuResponse? Sku;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetBastionHostResult(
            bool? disableCopyPaste,

            string? dnsName,

            bool? enableFileCopy,

            bool? enableIpConnect,

            bool? enableKerberos,

            bool? enableShareableLink,

            bool? enableTunneling,

            string etag,

            string? id,

            ImmutableArray<Outputs.BastionHostIPConfigurationResponse> ipConfigurations,

            string? location,

            string name,

            string provisioningState,

            int? scaleUnits,

            Outputs.SkuResponse? sku,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            DisableCopyPaste = disableCopyPaste;
            DnsName = dnsName;
            EnableFileCopy = enableFileCopy;
            EnableIpConnect = enableIpConnect;
            EnableKerberos = enableKerberos;
            EnableShareableLink = enableShareableLink;
            EnableTunneling = enableTunneling;
            Etag = etag;
            Id = id;
            IpConfigurations = ipConfigurations;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            ScaleUnits = scaleUnits;
            Sku = sku;
            Tags = tags;
            Type = type;
        }
    }
}
