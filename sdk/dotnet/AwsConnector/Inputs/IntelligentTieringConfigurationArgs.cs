// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of IntelligentTieringConfiguration
    /// </summary>
    public sealed class IntelligentTieringConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID used to identify the S3 Intelligent-Tiering configuration.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// An object key name prefix that identifies the subset of objects to which the rule applies.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// Specifies the status of the configuration.
        /// </summary>
        [Input("status")]
        public InputUnion<string, Pulumi.AzureNative.AwsConnector.IntelligentTieringConfigurationStatus>? Status { get; set; }

        [Input("tagFilters")]
        private InputList<Inputs.TagFilterArgs>? _tagFilters;

        /// <summary>
        /// A container for a key-value pair.
        /// </summary>
        public InputList<Inputs.TagFilterArgs> TagFilters
        {
            get => _tagFilters ?? (_tagFilters = new InputList<Inputs.TagFilterArgs>());
            set => _tagFilters = value;
        }

        [Input("tierings")]
        private InputList<Inputs.TieringArgs>? _tierings;

        /// <summary>
        /// Specifies a list of S3 Intelligent-Tiering storage class tiers in the configuration. At least one tier must be defined in the list. At most, you can specify two tiers in the list, one for each available AccessTier: ``ARCHIVE_ACCESS`` and ``DEEP_ARCHIVE_ACCESS``.  You only need Intelligent Tiering Configuration enabled on a bucket if you want to automatically move objects stored in the Intelligent-Tiering storage class to Archive Access or Deep Archive Access tiers.
        /// </summary>
        public InputList<Inputs.TieringArgs> Tierings
        {
            get => _tierings ?? (_tierings = new InputList<Inputs.TieringArgs>());
            set => _tierings = value;
        }

        public IntelligentTieringConfigurationArgs()
        {
        }
        public static new IntelligentTieringConfigurationArgs Empty => new IntelligentTieringConfigurationArgs();
    }
}
