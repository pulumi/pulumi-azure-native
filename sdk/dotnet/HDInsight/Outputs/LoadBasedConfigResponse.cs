// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Outputs
{

    /// <summary>
    /// Profile of load based Autoscale.
    /// </summary>
    [OutputType]
    public sealed class LoadBasedConfigResponse
    {
        /// <summary>
        /// This is a cool down period, this is a time period in seconds, which determines the amount of time that must elapse between a scaling activity started by a rule and the start of the next scaling activity, regardless of the rule that triggers it. The default value is 300 seconds.
        /// </summary>
        public readonly int? CooldownPeriod;
        /// <summary>
        /// User needs to set the maximum number of nodes for load based scaling, the load based scaling will use this to scale up and scale down between minimum and maximum number of nodes.
        /// </summary>
        public readonly int MaxNodes;
        /// <summary>
        /// User needs to set the minimum number of nodes for load based scaling, the load based scaling will use this to scale up and scale down between minimum and maximum number of nodes.
        /// </summary>
        public readonly int MinNodes;
        /// <summary>
        /// User can specify the poll interval, this is the time period (in seconds) after which scaling metrics are polled for triggering a scaling operation.
        /// </summary>
        public readonly int? PollInterval;
        /// <summary>
        /// The scaling rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScalingRuleResponse> ScalingRules;

        [OutputConstructor]
        private LoadBasedConfigResponse(
            int? cooldownPeriod,

            int maxNodes,

            int minNodes,

            int? pollInterval,

            ImmutableArray<Outputs.ScalingRuleResponse> scalingRules)
        {
            CooldownPeriod = cooldownPeriod;
            MaxNodes = maxNodes;
            MinNodes = minNodes;
            PollInterval = pollInterval;
            ScalingRules = scalingRules;
        }
    }
}
