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
    /// The reference to a user assigned identity associated with the Batch pool which a compute node will use.
    /// </summary>
    [OutputType]
    public sealed class ComputeNodeIdentityReferenceResponse
    {
        /// <summary>
        /// The ARM resource id of the user assigned identity.
        /// </summary>
        public readonly string? ResourceId;

        [OutputConstructor]
        private ComputeNodeIdentityReferenceResponse(string? resourceId)
        {
            ResourceId = resourceId;
        }
    }
}
