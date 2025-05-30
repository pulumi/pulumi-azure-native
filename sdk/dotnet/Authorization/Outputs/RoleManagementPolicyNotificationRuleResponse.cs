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
    /// The role management policy notification rule.
    /// </summary>
    [OutputType]
    public sealed class RoleManagementPolicyNotificationRuleResponse
    {
        /// <summary>
        /// The id of the rule.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Determines if the notification will be sent to the recipient type specified in the policy rule.
        /// </summary>
        public readonly bool? IsDefaultRecipientsEnabled;
        /// <summary>
        /// The notification level.
        /// </summary>
        public readonly string? NotificationLevel;
        /// <summary>
        /// The list of notification recipients.
        /// </summary>
        public readonly ImmutableArray<string> NotificationRecipients;
        /// <summary>
        /// The type of notification.
        /// </summary>
        public readonly string? NotificationType;
        /// <summary>
        /// The recipient type.
        /// </summary>
        public readonly string? RecipientType;
        /// <summary>
        /// The type of rule
        /// Expected value is 'RoleManagementPolicyNotificationRule'.
        /// </summary>
        public readonly string RuleType;
        /// <summary>
        /// The target of the current rule.
        /// </summary>
        public readonly Outputs.RoleManagementPolicyRuleTargetResponse? Target;

        [OutputConstructor]
        private RoleManagementPolicyNotificationRuleResponse(
            string? id,

            bool? isDefaultRecipientsEnabled,

            string? notificationLevel,

            ImmutableArray<string> notificationRecipients,

            string? notificationType,

            string? recipientType,

            string ruleType,

            Outputs.RoleManagementPolicyRuleTargetResponse? target)
        {
            Id = id;
            IsDefaultRecipientsEnabled = isDefaultRecipientsEnabled;
            NotificationLevel = notificationLevel;
            NotificationRecipients = notificationRecipients;
            NotificationType = notificationType;
            RecipientType = recipientType;
            RuleType = ruleType;
            Target = target;
        }
    }
}
