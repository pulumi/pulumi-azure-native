// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform.Outputs
{

    /// <summary>
    /// Application Insights agent versions properties payload
    /// </summary>
    [OutputType]
    public sealed class ApplicationInsightsAgentVersionsResponse
    {
        /// <summary>
        /// Indicates the version of application insight java agent
        /// </summary>
        public readonly string Java;

        [OutputConstructor]
        private ApplicationInsightsAgentVersionsResponse(string java)
        {
            Java = java;
        }
    }
}
