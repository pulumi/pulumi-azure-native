// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Inputs
{

    /// <summary>
    /// Windows gMSA Profile in the managed cluster.
    /// </summary>
    public sealed class WindowsGmsaProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the DNS server for Windows gMSA. &lt;br&gt;&lt;br&gt; Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster.
        /// </summary>
        [Input("dnsServer")]
        public Input<string>? DnsServer { get; set; }

        /// <summary>
        /// Specifies whether to enable Windows gMSA in the managed cluster.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Specifies the root domain name for Windows gMSA. &lt;br&gt;&lt;br&gt; Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster.
        /// </summary>
        [Input("rootDomainName")]
        public Input<string>? RootDomainName { get; set; }

        public WindowsGmsaProfileArgs()
        {
        }
        public static new WindowsGmsaProfileArgs Empty => new WindowsGmsaProfileArgs();
    }
}
