// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Scom.Inputs
{

    /// <summary>
    /// The properties of domain controller to which SCOM and SQL servers join for AuthN/AuthZ.
    /// </summary>
    public sealed class DomainControllerPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address of DNS server 
        /// </summary>
        [Input("dnsServer")]
        public Input<string>? DnsServer { get; set; }

        /// <summary>
        /// Fully qualified domain name
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Organizational Unit path in which the SCOM servers will be present
        /// </summary>
        [Input("ouPath")]
        public Input<string>? OuPath { get; set; }

        public DomainControllerPropertiesArgs()
        {
            OuPath = "";
        }
        public static new DomainControllerPropertiesArgs Empty => new DomainControllerPropertiesArgs();
    }
}
