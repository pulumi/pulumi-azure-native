using Pulumi;
using Pulumi.AzureNative.Authorization;
using Pulumi.AzureNative.Authorization.Inputs;
using System.Collections.Generic;

return await Pulumi.Deployment.RunAsync(async () =>
{
    var clientConfig = await GetClientConfig.InvokeAsync();

    var policy = new RoleManagementPolicy("policy", new RoleManagementPolicyArgs
    {
        RoleManagementPolicyName = "3faafb81-7f6f-4c66-b936-fb41ef4e4734",
        Scope =  Output.Format($"subscriptions/{clientConfig.SubscriptionId}"),
        Rules = new InputList<ResourceArgs>
        {
            new RoleManagementPolicyExpirationRuleArgs
            {
                Id = "Expiration_Admin_Eligibility",
                IsExpirationRequired = true,
                MaximumDuration = "P90D",
                RuleType = "RoleManagementPolicyExpirationRule",
            },
            new RoleManagementPolicyNotificationRuleArgs
            {
                Id = "Notification_Admin_Admin_Eligibility",
                IsDefaultRecipientsEnabled = true,
                NotificationLevel = "All",
                NotificationType = "Email",
                RecipientType = "Admin",
                RuleType = "RoleManagementPolicyNotificationRule",
            },
        },
    });

    // Export the primary key of the Storage Account
    return new Dictionary<string, object?>
    {
        {"policy", policy.Name}
    };
});