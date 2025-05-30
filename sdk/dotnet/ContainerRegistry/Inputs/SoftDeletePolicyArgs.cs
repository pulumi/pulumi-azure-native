// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.Inputs
{

    /// <summary>
    /// The soft delete policy for a container registry
    /// </summary>
    public sealed class SoftDeletePolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of days after which a soft-deleted item is permanently deleted.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        /// <summary>
        /// The value that indicates whether the policy is enabled or not.
        /// </summary>
        [Input("status")]
        public InputUnion<string, Pulumi.AzureNative.ContainerRegistry.PolicyStatus>? Status { get; set; }

        public SoftDeletePolicyArgs()
        {
            RetentionDays = 7;
            Status = "disabled";
        }
        public static new SoftDeletePolicyArgs Empty => new SoftDeletePolicyArgs();
    }
}
