// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// Deployment container liveness/readiness probe configuration.
    /// </summary>
    public sealed class ProbeSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of failures to allow before returning an unhealthy status.
        /// </summary>
        [Input("failureThreshold")]
        public Input<int>? FailureThreshold { get; set; }

        /// <summary>
        /// The delay before the first probe in ISO 8601 format.
        /// </summary>
        [Input("initialDelay")]
        public Input<string>? InitialDelay { get; set; }

        /// <summary>
        /// The length of time between probes in ISO 8601 format.
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        /// <summary>
        /// The number of successful probes before returning a healthy status.
        /// </summary>
        [Input("successThreshold")]
        public Input<int>? SuccessThreshold { get; set; }

        /// <summary>
        /// The probe timeout in ISO 8601 format.
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        public ProbeSettingsArgs()
        {
            FailureThreshold = 30;
            Period = "PT10S";
            SuccessThreshold = 1;
            Timeout = "PT2S";
        }
        public static new ProbeSettingsArgs Empty => new ProbeSettingsArgs();
    }
}
