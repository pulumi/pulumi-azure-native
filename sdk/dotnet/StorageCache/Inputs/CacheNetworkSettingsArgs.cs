// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageCache.Inputs
{

    /// <summary>
    /// Cache network settings.
    /// </summary>
    public sealed class CacheNetworkSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS search domain
        /// </summary>
        [Input("dnsSearchDomain")]
        public Input<string>? DnsSearchDomain { get; set; }

        [Input("dnsServers")]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// DNS servers for the cache to use.  It will be set from the network configuration if no value is provided.
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        /// <summary>
        /// The IPv4 maximum transmission unit configured for the subnet.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// NTP server IP Address or FQDN for the cache to use. The default is time.windows.com.
        /// </summary>
        [Input("ntpServer")]
        public Input<string>? NtpServer { get; set; }

        public CacheNetworkSettingsArgs()
        {
            Mtu = 1500;
            NtpServer = "time.windows.com";
        }
        public static new CacheNetworkSettingsArgs Empty => new CacheNetworkSettingsArgs();
    }
}
