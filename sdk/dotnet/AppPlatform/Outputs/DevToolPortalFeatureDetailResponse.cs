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
    /// Detail settings for Dev Tool Portal feature
    /// </summary>
    [OutputType]
    public sealed class DevToolPortalFeatureDetailResponse
    {
        /// <summary>
        /// Route path to visit the plugin
        /// </summary>
        public readonly string Route;
        /// <summary>
        /// State of the plugin
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private DevToolPortalFeatureDetailResponse(
            string route,

            string? state)
        {
            Route = route;
            State = state;
        }
    }
}
