// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    /// <summary>
    /// Instructions for labeling job
    /// </summary>
    [OutputType]
    public sealed class LabelingJobInstructionsResponse
    {
        /// <summary>
        /// The link to a page with detailed labeling instructions for labelers.
        /// </summary>
        public readonly string? Uri;

        [OutputConstructor]
        private LabelingJobInstructionsResponse(string? uri)
        {
            Uri = uri;
        }
    }
}
