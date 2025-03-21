// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Inputs
{

    /// <summary>
    /// configuration for Microsoft Defender for Server VM scanning
    /// </summary>
    public sealed class DefenderCspmAwsOfferingConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cloud role ARN in AWS for this feature
        /// </summary>
        [Input("cloudRoleArn")]
        public Input<string>? CloudRoleArn { get; set; }

        [Input("exclusionTags")]
        private InputMap<string>? _exclusionTags;

        /// <summary>
        /// VM tags that indicates that VM should not be scanned
        /// </summary>
        public InputMap<string> ExclusionTags
        {
            get => _exclusionTags ?? (_exclusionTags = new InputMap<string>());
            set => _exclusionTags = value;
        }

        /// <summary>
        /// The scanning mode for the VM scan.
        /// </summary>
        [Input("scanningMode")]
        public InputUnion<string, Pulumi.AzureNative.Security.ScanningMode>? ScanningMode { get; set; }

        public DefenderCspmAwsOfferingConfigurationArgs()
        {
        }
        public static new DefenderCspmAwsOfferingConfigurationArgs Empty => new DefenderCspmAwsOfferingConfigurationArgs();
    }
}
