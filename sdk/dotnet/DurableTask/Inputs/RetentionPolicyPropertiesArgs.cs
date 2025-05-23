// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DurableTask.Inputs
{

    /// <summary>
    /// The retention policy settings for the resource
    /// </summary>
    public sealed class RetentionPolicyPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("retentionPolicies")]
        private InputList<Inputs.RetentionPolicyDetailsArgs>? _retentionPolicies;

        /// <summary>
        /// The orchestration retention policies
        /// </summary>
        public InputList<Inputs.RetentionPolicyDetailsArgs> RetentionPolicies
        {
            get => _retentionPolicies ?? (_retentionPolicies = new InputList<Inputs.RetentionPolicyDetailsArgs>());
            set => _retentionPolicies = value;
        }

        public RetentionPolicyPropertiesArgs()
        {
        }
        public static new RetentionPolicyPropertiesArgs Empty => new RetentionPolicyPropertiesArgs();
    }
}
