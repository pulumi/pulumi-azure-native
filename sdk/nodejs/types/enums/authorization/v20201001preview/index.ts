// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ApprovalMode = {
    SingleStage: "SingleStage",
    Serial: "Serial",
    Parallel: "Parallel",
    NoApproval: "NoApproval",
} as const;

/**
 * The type of rule
 */
export type ApprovalMode = (typeof ApprovalMode)[keyof typeof ApprovalMode];

export const EnablementRules = {
    MultiFactorAuthentication: "MultiFactorAuthentication",
    Justification: "Justification",
    Ticketing: "Ticketing",
} as const;

/**
 * The type of enablement rule
 */
export type EnablementRules = (typeof EnablementRules)[keyof typeof EnablementRules];

export const NotificationDeliveryMechanism = {
    Email: "Email",
} as const;

/**
 * The type of notification.
 */
export type NotificationDeliveryMechanism = (typeof NotificationDeliveryMechanism)[keyof typeof NotificationDeliveryMechanism];

export const NotificationLevel = {
    None: "None",
    Critical: "Critical",
    All: "All",
} as const;

/**
 * The notification level.
 */
export type NotificationLevel = (typeof NotificationLevel)[keyof typeof NotificationLevel];

export const RecipientType = {
    Requestor: "Requestor",
    Approver: "Approver",
    Admin: "Admin",
} as const;

/**
 * The recipient type.
 */
export type RecipientType = (typeof RecipientType)[keyof typeof RecipientType];

export const RoleManagementPolicyRuleType = {
    RoleManagementPolicyApprovalRule: "RoleManagementPolicyApprovalRule",
    RoleManagementPolicyAuthenticationContextRule: "RoleManagementPolicyAuthenticationContextRule",
    RoleManagementPolicyEnablementRule: "RoleManagementPolicyEnablementRule",
    RoleManagementPolicyExpirationRule: "RoleManagementPolicyExpirationRule",
    RoleManagementPolicyNotificationRule: "RoleManagementPolicyNotificationRule",
} as const;

/**
 * The type of rule
 */
export type RoleManagementPolicyRuleType = (typeof RoleManagementPolicyRuleType)[keyof typeof RoleManagementPolicyRuleType];

export const UserType = {
    User: "User",
    Group: "Group",
} as const;

/**
 * The type of user.
 */
export type UserType = (typeof UserType)[keyof typeof UserType];
