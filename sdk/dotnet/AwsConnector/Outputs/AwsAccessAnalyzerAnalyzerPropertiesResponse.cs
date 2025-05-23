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
    /// Definition of awsAccessAnalyzerAnalyzer
    /// </summary>
    [OutputType]
    public sealed class AwsAccessAnalyzerAnalyzerPropertiesResponse
    {
        /// <summary>
        /// The configuration for the analyzer
        /// </summary>
        public readonly Outputs.UnusedAccessConfigurationResponse? AnalyzerConfiguration;
        /// <summary>
        /// Analyzer name
        /// </summary>
        public readonly string? AnalyzerName;
        /// <summary>
        /// Property archiveRules
        /// </summary>
        public readonly ImmutableArray<Outputs.ArchiveRuleResponse> ArchiveRules;
        /// <summary>
        /// Amazon Resource Name (ARN) of the analyzer
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.TagResponse> Tags;
        /// <summary>
        /// The type of the analyzer, must be one of ACCOUNT, ORGANIZATION, ACCOUNT_UNUSED_ACCESS or ORGANIZATION_UNUSED_ACCESS
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private AwsAccessAnalyzerAnalyzerPropertiesResponse(
            Outputs.UnusedAccessConfigurationResponse? analyzerConfiguration,

            string? analyzerName,

            ImmutableArray<Outputs.ArchiveRuleResponse> archiveRules,

            string? arn,

            ImmutableArray<Outputs.TagResponse> tags,

            string? type)
        {
            AnalyzerConfiguration = analyzerConfiguration;
            AnalyzerName = analyzerName;
            ArchiveRules = archiveRules;
            Arn = arn;
            Tags = tags;
            Type = type;
        }
    }
}
