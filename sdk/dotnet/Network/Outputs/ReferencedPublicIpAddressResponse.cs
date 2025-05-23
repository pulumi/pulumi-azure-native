// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Outputs
{

    /// <summary>
    /// Reference to a public IP address.
    /// </summary>
    [OutputType]
    public sealed class ReferencedPublicIpAddressResponse
    {
        /// <summary>
        /// The PublicIPAddress Reference.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private ReferencedPublicIpAddressResponse(string? id)
        {
            Id = id;
        }
    }
}
