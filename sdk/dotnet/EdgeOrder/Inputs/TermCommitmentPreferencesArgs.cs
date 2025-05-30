// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EdgeOrder.Inputs
{

    /// <summary>
    /// Term Commitment preference received from customer.
    /// </summary>
    public sealed class TermCommitmentPreferencesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Customer preferred Term Duration.
        /// </summary>
        [Input("preferredTermCommitmentDuration")]
        public Input<string>? PreferredTermCommitmentDuration { get; set; }

        /// <summary>
        /// Term Commitment Type
        /// </summary>
        [Input("preferredTermCommitmentType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.EdgeOrder.TermCommitmentType> PreferredTermCommitmentType { get; set; } = null!;

        public TermCommitmentPreferencesArgs()
        {
        }
        public static new TermCommitmentPreferencesArgs Empty => new TermCommitmentPreferencesArgs();
    }
}
