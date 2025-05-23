// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridCloud.Outputs
{

    /// <summary>
    /// Resource reference properties.
    /// </summary>
    [OutputType]
    public sealed class ResourceReferenceResponse
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private ResourceReferenceResponse(string? id)
        {
            Id = id;
        }
    }
}
