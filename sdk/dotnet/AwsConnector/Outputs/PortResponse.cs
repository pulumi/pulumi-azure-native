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
    /// Definition of Port
    /// </summary>
    [OutputType]
    public sealed class PortResponse
    {
        /// <summary>
        /// Access Direction for Protocol of the Instance(inbound/outbound).
        /// </summary>
        public readonly string? AccessDirection;
        /// <summary>
        /// Access From Protocol of the Instance.
        /// </summary>
        public readonly string? AccessFrom;
        /// <summary>
        /// Access Type Protocol of the Instance.
        /// </summary>
        public readonly string? AccessType;
        /// <summary>
        /// cidr List Aliases
        /// </summary>
        public readonly ImmutableArray<string> CidrListAliases;
        /// <summary>
        /// Property cidrs
        /// </summary>
        public readonly ImmutableArray<string> Cidrs;
        /// <summary>
        /// CommonName for Protocol of the Instance.
        /// </summary>
        public readonly string? CommonName;
        /// <summary>
        /// From Port of the Instance.
        /// </summary>
        public readonly int? FromPort;
        /// <summary>
        /// IPv6 Cidrs
        /// </summary>
        public readonly ImmutableArray<string> Ipv6Cidrs;
        /// <summary>
        /// Port Protocol of the Instance.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// To Port of the Instance.
        /// </summary>
        public readonly int? ToPort;

        [OutputConstructor]
        private PortResponse(
            string? accessDirection,

            string? accessFrom,

            string? accessType,

            ImmutableArray<string> cidrListAliases,

            ImmutableArray<string> cidrs,

            string? commonName,

            int? fromPort,

            ImmutableArray<string> ipv6Cidrs,

            string? protocol,

            int? toPort)
        {
            AccessDirection = accessDirection;
            AccessFrom = accessFrom;
            AccessType = accessType;
            CidrListAliases = cidrListAliases;
            Cidrs = cidrs;
            CommonName = commonName;
            FromPort = fromPort;
            Ipv6Cidrs = ipv6Cidrs;
            Protocol = protocol;
            ToPort = toPort;
        }
    }
}
