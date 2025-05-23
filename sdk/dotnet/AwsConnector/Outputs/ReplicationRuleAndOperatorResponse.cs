// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of ReplicationRuleAndOperator
    /// </summary>
    [OutputType]
    public sealed class ReplicationRuleAndOperatorResponse
    {
        /// <summary>
        /// An object key name prefix that identifies the subset of objects to which the rule applies.
        /// </summary>
        public readonly string? Prefix;
        /// <summary>
        /// An array of tags containing key and value pairs.
        /// </summary>
        public readonly ImmutableArray<Outputs.TagFilterResponse> TagFilters;

        [OutputConstructor]
        private ReplicationRuleAndOperatorResponse(
            string? prefix,

            ImmutableArray<Outputs.TagFilterResponse> tagFilters)
        {
            Prefix = prefix;
            TagFilters = tagFilters;
        }
    }
}
