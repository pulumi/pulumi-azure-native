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
    /// Definition of Ipv4PrefixSpecification
    /// </summary>
    [OutputType]
    public sealed class Ipv4PrefixSpecificationResponse
    {
        /// <summary>
        /// Property ipv4Prefix
        /// </summary>
        public readonly string? Ipv4Prefix;

        [OutputConstructor]
        private Ipv4PrefixSpecificationResponse(string? ipv4Prefix)
        {
            Ipv4Prefix = ipv4Prefix;
        }
    }
}
