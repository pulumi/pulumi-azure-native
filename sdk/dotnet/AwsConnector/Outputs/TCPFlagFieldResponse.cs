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
    /// Definition of TCPFlagField
    /// </summary>
    [OutputType]
    public sealed class TCPFlagFieldResponse
    {
        /// <summary>
        /// Property flags
        /// </summary>
        public readonly ImmutableArray<string> Flags;
        /// <summary>
        /// Property masks
        /// </summary>
        public readonly ImmutableArray<string> Masks;

        [OutputConstructor]
        private TCPFlagFieldResponse(
            ImmutableArray<string> flags,

            ImmutableArray<string> masks)
        {
            Flags = flags;
            Masks = masks;
        }
    }
}
