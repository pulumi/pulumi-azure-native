// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerInstance.Outputs
{

    /// <summary>
    /// Container group diagnostic information.
    /// </summary>
    [OutputType]
    public sealed class ContainerGroupDiagnosticsResponse
    {
        /// <summary>
        /// Container group log analytics information.
        /// </summary>
        public readonly Outputs.LogAnalyticsResponse? LogAnalytics;

        [OutputConstructor]
        private ContainerGroupDiagnosticsResponse(Outputs.LogAnalyticsResponse? logAnalytics)
        {
            LogAnalytics = logAnalytics;
        }
    }
}
