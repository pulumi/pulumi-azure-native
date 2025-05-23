// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageCache.Outputs
{

    /// <summary>
    /// Properties pertaining to the ClfsTarget
    /// </summary>
    [OutputType]
    public sealed class ClfsTargetResponse
    {
        /// <summary>
        /// Resource ID of storage container.
        /// </summary>
        public readonly string? Target;

        [OutputConstructor]
        private ClfsTargetResponse(string? target)
        {
            Target = target;
        }
    }
}
