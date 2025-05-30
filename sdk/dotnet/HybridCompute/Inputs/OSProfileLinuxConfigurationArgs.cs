// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridCompute.Inputs
{

    /// <summary>
    /// Specifies the linux configuration for update management.
    /// </summary>
    public sealed class OSProfileLinuxConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the assessment mode.
        /// </summary>
        [Input("assessmentMode")]
        public InputUnion<string, Pulumi.AzureNative.HybridCompute.AssessmentModeTypes>? AssessmentMode { get; set; }

        /// <summary>
        /// Captures the hotpatch capability enrollment intent of the customers, which enables customers to patch their Windows machines without requiring a reboot.
        /// </summary>
        [Input("enableHotpatching")]
        public Input<bool>? EnableHotpatching { get; set; }

        /// <summary>
        /// Specifies the patch mode.
        /// </summary>
        [Input("patchMode")]
        public InputUnion<string, Pulumi.AzureNative.HybridCompute.PatchModeTypes>? PatchMode { get; set; }

        public OSProfileLinuxConfigurationArgs()
        {
        }
        public static new OSProfileLinuxConfigurationArgs Empty => new OSProfileLinuxConfigurationArgs();
    }
}
