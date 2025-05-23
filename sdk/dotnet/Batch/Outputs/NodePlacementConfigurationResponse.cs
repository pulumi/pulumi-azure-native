// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.Outputs
{

    /// <summary>
    /// Allocation configuration used by Batch Service to provision the nodes.
    /// </summary>
    [OutputType]
    public sealed class NodePlacementConfigurationResponse
    {
        /// <summary>
        /// Allocation policy used by Batch Service to provision the nodes. If not specified, Batch will use the regional policy.
        /// </summary>
        public readonly string? Policy;

        [OutputConstructor]
        private NodePlacementConfigurationResponse(string? policy)
        {
            Policy = policy;
        }
    }
}
