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
    /// Settings for Dev Tool Portal
    /// </summary>
    [OutputType]
    public sealed class DevToolPortalFeatureSettingsResponse
    {
        /// <summary>
        /// Detail of Accelerator plugin
        /// </summary>
        public readonly Outputs.DevToolPortalFeatureDetailResponse? ApplicationAccelerator;
        /// <summary>
        /// Detail of App Live View plugin
        /// </summary>
        public readonly Outputs.DevToolPortalFeatureDetailResponse? ApplicationLiveView;

        [OutputConstructor]
        private DevToolPortalFeatureSettingsResponse(
            Outputs.DevToolPortalFeatureDetailResponse? applicationAccelerator,

            Outputs.DevToolPortalFeatureDetailResponse? applicationLiveView)
        {
            ApplicationAccelerator = applicationAccelerator;
            ApplicationLiveView = applicationLiveView;
        }
    }
}
