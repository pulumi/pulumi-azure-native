// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CognitiveServices.Inputs
{

    /// <summary>
    /// Properties of Cognitive Services account commitment plan.
    /// </summary>
    public sealed class CommitmentPlanPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AutoRenew commitment plan.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// Commitment plan guid.
        /// </summary>
        [Input("commitmentPlanGuid")]
        public Input<string>? CommitmentPlanGuid { get; set; }

        /// <summary>
        /// Cognitive Services account commitment period.
        /// </summary>
        [Input("current")]
        public Input<Inputs.CommitmentPeriodArgs>? Current { get; set; }

        /// <summary>
        /// Account hosting model.
        /// </summary>
        [Input("hostingModel")]
        public InputUnion<string, Pulumi.AzureNative.CognitiveServices.HostingModel>? HostingModel { get; set; }

        /// <summary>
        /// Cognitive Services account commitment period.
        /// </summary>
        [Input("next")]
        public Input<Inputs.CommitmentPeriodArgs>? Next { get; set; }

        /// <summary>
        /// Commitment plan type.
        /// </summary>
        [Input("planType")]
        public Input<string>? PlanType { get; set; }

        public CommitmentPlanPropertiesArgs()
        {
        }
        public static new CommitmentPlanPropertiesArgs Empty => new CommitmentPlanPropertiesArgs();
    }
}
