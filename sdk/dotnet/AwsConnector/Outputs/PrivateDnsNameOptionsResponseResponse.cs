// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of PrivateDnsNameOptionsResponse
    /// </summary>
    [OutputType]
    public sealed class PrivateDnsNameOptionsResponseResponse
    {
        /// <summary>
        /// &lt;p&gt;Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records.&lt;/p&gt;
        /// </summary>
        public readonly bool? EnableResourceNameDnsAAAARecord;
        /// <summary>
        /// &lt;p&gt;Indicates whether to respond to DNS queries for instance hostnames with DNS A records.&lt;/p&gt;
        /// </summary>
        public readonly bool? EnableResourceNameDnsARecord;
        /// <summary>
        /// &lt;p&gt;The type of hostname to assign to an instance.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.HostnameTypeEnumValueResponse? HostnameType;

        [OutputConstructor]
        private PrivateDnsNameOptionsResponseResponse(
            bool? enableResourceNameDnsAAAARecord,

            bool? enableResourceNameDnsARecord,

            Outputs.HostnameTypeEnumValueResponse? hostnameType)
        {
            EnableResourceNameDnsAAAARecord = enableResourceNameDnsAAAARecord;
            EnableResourceNameDnsARecord = enableResourceNameDnsARecord;
            HostnameType = hostnameType;
        }
    }
}
