// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CostManagement.Outputs
{

    /// <summary>
    /// Target resources and allocation
    /// </summary>
    [OutputType]
    public sealed class CostAllocationProportionResponse
    {
        /// <summary>
        /// Target resource for cost allocation
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Percentage of source cost to allocate to this resource. This value can be specified to two decimal places and the total percentage of all resources in this rule must sum to 100.00.
        /// </summary>
        public readonly double Percentage;

        [OutputConstructor]
        private CostAllocationProportionResponse(
            string name,

            double percentage)
        {
            Name = name;
            Percentage = percentage;
        }
    }
}
