// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Inputs
{

    /// <summary>
    /// Soft delete Settings of vault
    /// </summary>
    public sealed class SoftDeleteSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("enhancedSecurityState")]
        public InputUnion<string, Pulumi.AzureNative.RecoveryServices.EnhancedSecurityState>? EnhancedSecurityState { get; set; }

        /// <summary>
        /// Soft delete retention period in days
        /// </summary>
        [Input("softDeleteRetentionPeriodInDays")]
        public Input<int>? SoftDeleteRetentionPeriodInDays { get; set; }

        [Input("softDeleteState")]
        public InputUnion<string, Pulumi.AzureNative.RecoveryServices.SoftDeleteState>? SoftDeleteState { get; set; }

        public SoftDeleteSettingsArgs()
        {
        }
        public static new SoftDeleteSettingsArgs Empty => new SoftDeleteSettingsArgs();
    }
}
