import * as pulumi from "@pulumi/pulumi";
import * as pim from "@pulumi/azure-native/authorization";

const clientConfig = pulumi.output(pim.getClientConfig());

const policy = new pim.RoleManagementPolicy("policy", {
    roleManagementPolicyName: "3faafb81-7f6f-4c66-b936-fb41ef4e4734",
    scope: pulumi.interpolate`subscriptions/${clientConfig.subscriptionId}`,
    rules: [
        {
            "id": "Expiration_Admin_Eligibility",
            "maximumDuration": "P90D", // new
            "isExpirationRequired": true,
            "ruleType": "RoleManagementPolicyExpirationRule",
        },
    ]
});
