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
    /// The ARM ID for a Gallery Image.
    /// </summary>
    [OutputType]
    public sealed class ImageArmReferenceResponse
    {
        /// <summary>
        /// The ARM ID for an image resource used by the virtual machine instance.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private ImageArmReferenceResponse(string? id)
        {
            Id = id;
        }
    }
}
