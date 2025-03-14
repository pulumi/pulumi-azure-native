// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20240101Preview.Inputs
{

    /// <summary>
    /// Used for sweep over component
    /// </summary>
    public sealed class ComponentConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Pipeline settings, for things like ContinueRunOnStepFailure etc.
        /// </summary>
        [Input("pipelineSettings")]
        public Input<object>? PipelineSettings { get; set; }

        public ComponentConfigurationArgs()
        {
        }
        public static new ComponentConfigurationArgs Empty => new ComponentConfigurationArgs();
    }
}
