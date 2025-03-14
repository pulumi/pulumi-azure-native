// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Confluent.V20240701.Outputs
{

    /// <summary>
    /// Stream governance configuration
    /// </summary>
    [OutputType]
    public sealed class StreamGovernanceConfigResponse
    {
        /// <summary>
        /// Stream governance configuration
        /// </summary>
        public readonly string? Package;

        [OutputConstructor]
        private StreamGovernanceConfigResponse(string? package)
        {
            Package = package;
        }
    }
}
