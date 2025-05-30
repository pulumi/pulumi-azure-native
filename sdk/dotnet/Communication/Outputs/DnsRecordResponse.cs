// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Communication.Outputs
{

    /// <summary>
    /// A class that represents a VerificationStatus record.
    /// </summary>
    [OutputType]
    public sealed class DnsRecordResponse
    {
        /// <summary>
        /// Name of the DNS record.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Represents an expiry time in seconds to represent how long this entry can be cached by the resolver, default = 3600sec.
        /// </summary>
        public readonly int Ttl;
        /// <summary>
        /// Type of the DNS record. Example: TXT
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Value of the DNS record.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private DnsRecordResponse(
            string name,

            int ttl,

            string type,

            string value)
        {
            Name = name;
            Ttl = ttl;
            Type = type;
            Value = value;
        }
    }
}
