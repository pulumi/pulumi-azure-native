// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Authorization.Inputs
{

    /// <summary>
    /// The approval settings.
    /// </summary>
    public sealed class ApprovalSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of rule
        /// </summary>
        [Input("approvalMode")]
        public InputUnion<string, Pulumi.AzureNative.Authorization.ApprovalMode>? ApprovalMode { get; set; }

        [Input("approvalStages")]
        private InputList<Inputs.ApprovalStageArgs>? _approvalStages;

        /// <summary>
        /// The approval stages of the request.
        /// </summary>
        public InputList<Inputs.ApprovalStageArgs> ApprovalStages
        {
            get => _approvalStages ?? (_approvalStages = new InputList<Inputs.ApprovalStageArgs>());
            set => _approvalStages = value;
        }

        /// <summary>
        /// Determines whether approval is required or not.
        /// </summary>
        [Input("isApprovalRequired")]
        public Input<bool>? IsApprovalRequired { get; set; }

        /// <summary>
        /// Determines whether approval is required for assignment extension.
        /// </summary>
        [Input("isApprovalRequiredForExtension")]
        public Input<bool>? IsApprovalRequiredForExtension { get; set; }

        /// <summary>
        /// Determine whether requestor justification is required.
        /// </summary>
        [Input("isRequestorJustificationRequired")]
        public Input<bool>? IsRequestorJustificationRequired { get; set; }

        public ApprovalSettingsArgs()
        {
        }
        public static new ApprovalSettingsArgs Empty => new ApprovalSettingsArgs();
    }
}
