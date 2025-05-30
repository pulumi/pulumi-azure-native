// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.Inputs
{

    public sealed class FixedScaleSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If omitted, the default value is Requeue.
        /// </summary>
        [Input("nodeDeallocationOption")]
        public Input<Pulumi.AzureNative.Batch.ComputeNodeDeallocationOption>? NodeDeallocationOption { get; set; }

        /// <summary>
        /// The default value is 15 minutes. Timeout values use ISO 8601 format. For example, use PT10M for 10 minutes. The minimum value is 5 minutes. If you specify a value less than 5 minutes, the Batch service rejects the request with an error; if you are calling the REST API directly, the HTTP status code is 400 (Bad Request).
        /// </summary>
        [Input("resizeTimeout")]
        public Input<string>? ResizeTimeout { get; set; }

        /// <summary>
        /// At least one of targetDedicatedNodes, targetLowPriorityNodes must be set.
        /// </summary>
        [Input("targetDedicatedNodes")]
        public Input<int>? TargetDedicatedNodes { get; set; }

        /// <summary>
        /// At least one of targetDedicatedNodes, targetLowPriorityNodes must be set.
        /// </summary>
        [Input("targetLowPriorityNodes")]
        public Input<int>? TargetLowPriorityNodes { get; set; }

        public FixedScaleSettingsArgs()
        {
            ResizeTimeout = "PT15M";
        }
        public static new FixedScaleSettingsArgs Empty => new FixedScaleSettingsArgs();
    }
}
