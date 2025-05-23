// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EdgeOrder.Outputs
{

    /// <summary>
    /// Resource stage details.
    /// </summary>
    [OutputType]
    public sealed class StageDetailsResponse
    {
        /// <summary>
        /// Display name of the resource stage.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Stage name.
        /// </summary>
        public readonly string StageName;
        /// <summary>
        /// Stage status.
        /// </summary>
        public readonly string StageStatus;
        /// <summary>
        /// Stage start time.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private StageDetailsResponse(
            string displayName,

            string stageName,

            string stageStatus,

            string startTime)
        {
            DisplayName = displayName;
            StageName = stageName;
            StageStatus = stageStatus;
            StartTime = startTime;
        }
    }
}
