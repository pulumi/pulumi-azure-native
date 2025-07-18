// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage.Inputs
{

    /// <summary>
    /// The trigger parameters update for the storage task assignment execution
    /// </summary>
    public sealed class TriggerParametersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When to end task execution. This is a required field when ExecutionTrigger.properties.type is 'OnSchedule'; this property should not be present when ExecutionTrigger.properties.type is 'RunOnce'
        /// </summary>
        [Input("endBy")]
        public Input<string>? EndBy { get; set; }

        /// <summary>
        /// Run interval of task execution. This is a required field when ExecutionTrigger.properties.type is 'OnSchedule'; this property should not be present when ExecutionTrigger.properties.type is 'RunOnce'
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// Run interval unit of task execution. This is a required field when ExecutionTrigger.properties.type is 'OnSchedule'; this property should not be present when ExecutionTrigger.properties.type is 'RunOnce'
        /// </summary>
        [Input("intervalUnit")]
        public Input<Pulumi.AzureNative.Storage.IntervalUnit>? IntervalUnit { get; set; }

        /// <summary>
        /// When to start task execution. This is a required field when ExecutionTrigger.properties.type is 'OnSchedule'; this property should not be present when ExecutionTrigger.properties.type is 'RunOnce'
        /// </summary>
        [Input("startFrom")]
        public Input<string>? StartFrom { get; set; }

        /// <summary>
        /// When to start task execution. This is a required field when ExecutionTrigger.properties.type is 'RunOnce'; this property should not be present when ExecutionTrigger.properties.type is 'OnSchedule'
        /// </summary>
        [Input("startOn")]
        public Input<string>? StartOn { get; set; }

        public TriggerParametersArgs()
        {
        }
        public static new TriggerParametersArgs Empty => new TriggerParametersArgs();
    }
}
