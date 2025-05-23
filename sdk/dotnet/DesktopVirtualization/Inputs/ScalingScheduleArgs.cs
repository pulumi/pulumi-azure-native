// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DesktopVirtualization.Inputs
{

    /// <summary>
    /// A ScalingPlanPooledSchedule.
    /// </summary>
    public sealed class ScalingScheduleArgs : global::Pulumi.ResourceArgs
    {
        [Input("daysOfWeek")]
        private InputList<string>? _daysOfWeek;

        /// <summary>
        /// Set of days of the week on which this schedule is active.
        /// </summary>
        public InputList<string> DaysOfWeek
        {
            get => _daysOfWeek ?? (_daysOfWeek = new InputList<string>());
            set => _daysOfWeek = value;
        }

        /// <summary>
        /// Name of the ScalingPlanPooledSchedule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Load balancing algorithm for off-peak period.
        /// </summary>
        [Input("offPeakLoadBalancingAlgorithm")]
        public InputUnion<string, Pulumi.AzureNative.DesktopVirtualization.SessionHostLoadBalancingAlgorithm>? OffPeakLoadBalancingAlgorithm { get; set; }

        /// <summary>
        /// Starting time for off-peak period.
        /// </summary>
        [Input("offPeakStartTime")]
        public Input<Inputs.TimeArgs>? OffPeakStartTime { get; set; }

        /// <summary>
        /// Load balancing algorithm for peak period.
        /// </summary>
        [Input("peakLoadBalancingAlgorithm")]
        public InputUnion<string, Pulumi.AzureNative.DesktopVirtualization.SessionHostLoadBalancingAlgorithm>? PeakLoadBalancingAlgorithm { get; set; }

        /// <summary>
        /// Starting time for peak period.
        /// </summary>
        [Input("peakStartTime")]
        public Input<Inputs.TimeArgs>? PeakStartTime { get; set; }

        /// <summary>
        /// Capacity threshold for ramp down period.
        /// </summary>
        [Input("rampDownCapacityThresholdPct")]
        public Input<int>? RampDownCapacityThresholdPct { get; set; }

        /// <summary>
        /// Should users be logged off forcefully from hosts.
        /// </summary>
        [Input("rampDownForceLogoffUsers")]
        public Input<bool>? RampDownForceLogoffUsers { get; set; }

        /// <summary>
        /// Load balancing algorithm for ramp down period.
        /// </summary>
        [Input("rampDownLoadBalancingAlgorithm")]
        public InputUnion<string, Pulumi.AzureNative.DesktopVirtualization.SessionHostLoadBalancingAlgorithm>? RampDownLoadBalancingAlgorithm { get; set; }

        /// <summary>
        /// Minimum host percentage for ramp down period.
        /// </summary>
        [Input("rampDownMinimumHostsPct")]
        public Input<int>? RampDownMinimumHostsPct { get; set; }

        /// <summary>
        /// Notification message for users during ramp down period.
        /// </summary>
        [Input("rampDownNotificationMessage")]
        public Input<string>? RampDownNotificationMessage { get; set; }

        /// <summary>
        /// Starting time for ramp down period.
        /// </summary>
        [Input("rampDownStartTime")]
        public Input<Inputs.TimeArgs>? RampDownStartTime { get; set; }

        /// <summary>
        /// Specifies when to stop hosts during ramp down period.
        /// </summary>
        [Input("rampDownStopHostsWhen")]
        public InputUnion<string, Pulumi.AzureNative.DesktopVirtualization.StopHostsWhen>? RampDownStopHostsWhen { get; set; }

        /// <summary>
        /// Number of minutes to wait to stop hosts during ramp down period.
        /// </summary>
        [Input("rampDownWaitTimeMinutes")]
        public Input<int>? RampDownWaitTimeMinutes { get; set; }

        /// <summary>
        /// Capacity threshold for ramp up period.
        /// </summary>
        [Input("rampUpCapacityThresholdPct")]
        public Input<int>? RampUpCapacityThresholdPct { get; set; }

        /// <summary>
        /// Load balancing algorithm for ramp up period.
        /// </summary>
        [Input("rampUpLoadBalancingAlgorithm")]
        public InputUnion<string, Pulumi.AzureNative.DesktopVirtualization.SessionHostLoadBalancingAlgorithm>? RampUpLoadBalancingAlgorithm { get; set; }

        /// <summary>
        /// Minimum host percentage for ramp up period.
        /// </summary>
        [Input("rampUpMinimumHostsPct")]
        public Input<int>? RampUpMinimumHostsPct { get; set; }

        /// <summary>
        /// Starting time for ramp up period.
        /// </summary>
        [Input("rampUpStartTime")]
        public Input<Inputs.TimeArgs>? RampUpStartTime { get; set; }

        public ScalingScheduleArgs()
        {
        }
        public static new ScalingScheduleArgs Empty => new ScalingScheduleArgs();
    }
}
