// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CloudHealth.Inputs
{

    /// <summary>
    /// Properties for dependent entities, i.e. child entities
    /// </summary>
    public sealed class DependenciesSignalGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Aggregation type for child dependencies.
        /// </summary>
        [Input("aggregationType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.CloudHealth.DependenciesAggregationType> AggregationType { get; set; } = null!;

        /// <summary>
        /// Degraded threshold for aggregating the propagated health state of child dependencies. Can be either an absolute number that is greater than 0, or a percentage between 1-100%. The entity will be considered degraded when the number of not healthy child dependents (unhealthy, degraded, unknown) is equal to or above the threshold value. Must only be set when AggregationType is 'Thresholds'.
        /// </summary>
        [Input("degradedThreshold")]
        public Input<string>? DegradedThreshold { get; set; }

        /// <summary>
        /// Unhealthy threshold for aggregating the propagated health state of child dependencies. Can be either an absolute number that is greater than 0, or a percentage between 1-100%. The entity will be considered unhealthy when the number of not healthy child dependents (unhealthy, degraded, unknown) is equal to or above the threshold value. Must only be set when AggregationType is 'Thresholds'.
        /// </summary>
        [Input("unhealthyThreshold")]
        public Input<string>? UnhealthyThreshold { get; set; }

        public DependenciesSignalGroupArgs()
        {
            AggregationType = "WorstOf";
        }
        public static new DependenciesSignalGroupArgs Empty => new DependenciesSignalGroupArgs();
    }
}
