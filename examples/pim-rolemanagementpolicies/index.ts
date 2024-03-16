import * as pulumi from "@pulumi/pulumi";
import * as pim from "@pulumi/azure-native/authorization";

const clientConfig = pulumi.output(pim.getClientConfig());

const policy = new pim.RoleManagementPolicy("policy", {
    roleManagementPolicyName: new pulumi.Config().requireSecret("policy"),
    scope: pulumi.interpolate`subscriptions/${clientConfig.subscriptionId}`,
    rules: [
        {
            "id": "Expiration_Admin_Eligibility",
            "isExpirationRequired": true,
            "maximumDuration": "P365D",
            "ruleType": "RoleManagementPolicyExpirationRule",
        },
    ]
});
