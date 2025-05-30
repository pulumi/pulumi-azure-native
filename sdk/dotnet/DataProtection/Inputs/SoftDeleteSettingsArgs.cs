// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataProtection.Inputs
{

    /// <summary>
    /// Soft delete related settings
    /// </summary>
    public sealed class SoftDeleteSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Soft delete retention duration
        /// </summary>
        [Input("retentionDurationInDays")]
        public Input<double>? RetentionDurationInDays { get; set; }

        /// <summary>
        /// State of soft delete
        /// </summary>
        [Input("state")]
        public InputUnion<string, Pulumi.AzureNative.DataProtection.SoftDeleteState>? State { get; set; }

        public SoftDeleteSettingsArgs()
        {
        }
        public static new SoftDeleteSettingsArgs Empty => new SoftDeleteSettingsArgs();
    }
}
