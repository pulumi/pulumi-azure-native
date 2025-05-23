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
    /// Source resources for cost allocation
    /// </summary>
    [OutputType]
    public sealed class SourceCostAllocationResourceResponse
    {
        /// <summary>
        /// If resource type is dimension, this must be either ResourceGroupName or SubscriptionId. If resource type is tag, this must be a valid Azure tag
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Type of resources contained in this cost allocation rule
        /// </summary>
        public readonly string ResourceType;
        /// <summary>
        /// Source Resources for cost allocation. This list cannot contain more than 25 values.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private SourceCostAllocationResourceResponse(
            string name,

            string resourceType,

            ImmutableArray<string> values)
        {
            Name = name;
            ResourceType = resourceType;
            Values = values;
        }
    }
}
