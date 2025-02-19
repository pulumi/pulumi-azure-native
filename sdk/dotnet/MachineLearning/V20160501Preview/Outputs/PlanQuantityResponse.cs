// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearning.V20160501Preview.Outputs
{

    /// <summary>
    /// Represents the quantity a commitment plan provides of a metered resource.
    /// </summary>
    [OutputType]
    public sealed class PlanQuantityResponse
    {
        /// <summary>
        /// The quantity added to the commitment plan at an interval specified by its allowance frequency.
        /// </summary>
        public readonly double Allowance;
        /// <summary>
        /// The quantity available to the plan the last time usage was calculated.
        /// </summary>
        public readonly double Amount;
        /// <summary>
        /// The Azure meter for usage against included quantities.
        /// </summary>
        public readonly string IncludedQuantityMeter;
        /// <summary>
        /// The Azure meter for usage which exceeds included quantities.
        /// </summary>
        public readonly string OverageMeter;

        [OutputConstructor]
        private PlanQuantityResponse(
            double allowance,

            double amount,

            string includedQuantityMeter,

            string overageMeter)
        {
            Allowance = allowance;
            Amount = amount;
            IncludedQuantityMeter = includedQuantityMeter;
            OverageMeter = overageMeter;
        }
    }
}
