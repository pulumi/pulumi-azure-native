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
    /// IP Pool info
    /// </summary>
    [OutputType]
    public sealed class IPPoolInfoResponse
    {
        /// <summary>
        /// Number of IP addresses available in the IP Pool
        /// </summary>
        public readonly string Available;
        /// <summary>
        /// Number of IP addresses allocated from the IP Pool
        /// </summary>
        public readonly string Used;

        [OutputConstructor]
        private IPPoolInfoResponse(
            string available,

            string used)
        {
            Available = available;
            Used = used;
        }
    }
}
