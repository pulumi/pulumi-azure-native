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
    /// Definition of RulesSourceList
    /// </summary>
    [OutputType]
    public sealed class RulesSourceListResponse
    {
        /// <summary>
        /// Property generatedRulesType
        /// </summary>
        public readonly string? GeneratedRulesType;
        /// <summary>
        /// Property targetTypes
        /// </summary>
        public readonly ImmutableArray<string> TargetTypes;
        /// <summary>
        /// Property targets
        /// </summary>
        public readonly ImmutableArray<string> Targets;

        [OutputConstructor]
        private RulesSourceListResponse(
            string? generatedRulesType,

            ImmutableArray<string> targetTypes,

            ImmutableArray<string> targets)
        {
            GeneratedRulesType = generatedRulesType;
            TargetTypes = targetTypes;
            Targets = targets;
        }
    }
}
