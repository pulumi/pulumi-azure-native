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
    /// Definition of ReplicationConfiguration
    /// </summary>
    public sealed class ReplicationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinations")]
        private InputList<Inputs.ReplicationDestinationArgs>? _destinations;

        /// <summary>
        /// An array of destination objects. Only one destination object is supported.
        /// </summary>
        public InputList<Inputs.ReplicationDestinationArgs> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<Inputs.ReplicationDestinationArgs>());
            set => _destinations = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAMlong (IAM) role that Amazon S3 assumes when replicating objects. For more information, see [How to Set Up Replication](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-how-setup.html) in the *Amazon S3 User Guide*.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("rules")]
        private InputList<Inputs.ReplicationRuleArgs>? _rules;

        /// <summary>
        /// A container for one or more replication rules. A replication configuration must have at least one rule and can contain a maximum of 1,000 rules.
        /// </summary>
        public InputList<Inputs.ReplicationRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.ReplicationRuleArgs>());
            set => _rules = value;
        }

        public ReplicationConfigurationArgs()
        {
        }
        public static new ReplicationConfigurationArgs Empty => new ReplicationConfigurationArgs();
    }
}
