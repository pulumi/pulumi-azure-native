// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Authorization.Outputs
{

    /// <summary>
    /// The approval stage.
    /// </summary>
    [OutputType]
    public sealed class ApprovalStageResponse
    {
        /// <summary>
        /// The time in days when approval request would be timed out
        /// </summary>
        public readonly int? ApprovalStageTimeOutInDays;
        /// <summary>
        /// The escalation approver of the request.
        /// </summary>
        public readonly ImmutableArray<Outputs.UserSetResponse> EscalationApprovers;
        /// <summary>
        /// The time in minutes when the approval request would be escalated if the primary approver does not approve
        /// </summary>
        public readonly int? EscalationTimeInMinutes;
        /// <summary>
        /// Determines whether approver need to provide justification for his decision.
        /// </summary>
        public readonly bool? IsApproverJustificationRequired;
        /// <summary>
        /// The value determine whether escalation feature is enabled.
        /// </summary>
        public readonly bool? IsEscalationEnabled;
        /// <summary>
        /// The primary approver of the request.
        /// </summary>
        public readonly ImmutableArray<Outputs.UserSetResponse> PrimaryApprovers;

        [OutputConstructor]
        private ApprovalStageResponse(
            int? approvalStageTimeOutInDays,

            ImmutableArray<Outputs.UserSetResponse> escalationApprovers,

            int? escalationTimeInMinutes,

            bool? isApproverJustificationRequired,

            bool? isEscalationEnabled,

            ImmutableArray<Outputs.UserSetResponse> primaryApprovers)
        {
            ApprovalStageTimeOutInDays = approvalStageTimeOutInDays;
            EscalationApprovers = escalationApprovers;
            EscalationTimeInMinutes = escalationTimeInMinutes;
            IsApproverJustificationRequired = isApproverJustificationRequired;
            IsEscalationEnabled = isEscalationEnabled;
            PrimaryApprovers = primaryApprovers;
        }
    }
}
