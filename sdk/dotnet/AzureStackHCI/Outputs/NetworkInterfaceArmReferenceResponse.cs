// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Outputs
{

    /// <summary>
    /// The ARM ID for a Network Interface.
    /// </summary>
    [OutputType]
    public sealed class NetworkInterfaceArmReferenceResponse
    {
        /// <summary>
        /// The ARM ID for a Network Interface.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private NetworkInterfaceArmReferenceResponse(string? id)
        {
            Id = id;
        }
    }
}
