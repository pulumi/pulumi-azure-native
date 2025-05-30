// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Outputs
{

    /// <summary>
    /// Describes a virtual machines scale set IP Configuration's PublicIPAddress configuration
    /// </summary>
    [OutputType]
    public sealed class VirtualMachineScaleSetPublicIPAddressConfigurationResponse
    {
        /// <summary>
        /// Specify what happens to the public IP when the VM is deleted
        /// </summary>
        public readonly string? DeleteOption;
        /// <summary>
        /// The dns settings to be applied on the publicIP addresses .
        /// </summary>
        public readonly Outputs.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse? DnsSettings;
        /// <summary>
        /// The idle timeout of the public IP address.
        /// </summary>
        public readonly int? IdleTimeoutInMinutes;
        /// <summary>
        /// The list of IP tags associated with the public IP address.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualMachineScaleSetIpTagResponse> IpTags;
        /// <summary>
        /// The publicIP address configuration name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Available from Api-Version 2019-07-01 onwards, it represents whether the specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4. Possible values are: 'IPv4' and 'IPv6'.
        /// </summary>
        public readonly string? PublicIPAddressVersion;
        /// <summary>
        /// The PublicIPPrefix from which to allocate publicIP addresses.
        /// </summary>
        public readonly Outputs.SubResourceResponse? PublicIPPrefix;
        /// <summary>
        /// Describes the public IP Sku. It can only be set with OrchestrationMode as Flexible.
        /// </summary>
        public readonly Outputs.PublicIPAddressSkuResponse? Sku;

        [OutputConstructor]
        private VirtualMachineScaleSetPublicIPAddressConfigurationResponse(
            string? deleteOption,

            Outputs.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse? dnsSettings,

            int? idleTimeoutInMinutes,

            ImmutableArray<Outputs.VirtualMachineScaleSetIpTagResponse> ipTags,

            string name,

            string? publicIPAddressVersion,

            Outputs.SubResourceResponse? publicIPPrefix,

            Outputs.PublicIPAddressSkuResponse? sku)
        {
            DeleteOption = deleteOption;
            DnsSettings = dnsSettings;
            IdleTimeoutInMinutes = idleTimeoutInMinutes;
            IpTags = ipTags;
            Name = name;
            PublicIPAddressVersion = publicIPAddressVersion;
            PublicIPPrefix = publicIPPrefix;
            Sku = sku;
        }
    }
}
