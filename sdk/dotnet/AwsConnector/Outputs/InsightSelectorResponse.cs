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
    /// Definition of InsightSelector
    /// </summary>
    [OutputType]
    public sealed class InsightSelectorResponse
    {
        /// <summary>
        /// The type of insight to log on a trail.
        /// </summary>
        public readonly string? InsightType;

        [OutputConstructor]
        private InsightSelectorResponse(string? insightType)
        {
            InsightType = insightType;
        }
    }
}
