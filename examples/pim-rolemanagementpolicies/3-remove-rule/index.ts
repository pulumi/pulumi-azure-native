import * as pulumi from "@pulumi/pulumi";
import * as pim from "@pulumi/azure-native/authorization";

const clientConfig = pulumi.output(pim.getClientConfig());

const policy = new pim.RoleManagementPolicy("policy", {
    roleManagementPolicyName: new pulumi.Config().requireSecret("policy"),
    scope: pulumi.interpolate`subscriptions/${clientConfig.subscriptionId}`,
    rules: [
        // rule removed, but it will still exist in Azure

        // we add another one since an empty list is not allowed
        {
            "id": "Notification_Admin_Admin_Eligibility",
            "isDefaultRecipientsEnabled": true,
            "notificationLevel": "All",
            "notificationType": "Email",
            "recipientType": "Admin",
            "ruleType": "RoleManagementPolicyNotificationRule",
        },
    ]
});
