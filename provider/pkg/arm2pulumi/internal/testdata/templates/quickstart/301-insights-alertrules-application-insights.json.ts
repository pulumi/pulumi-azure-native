import * as pulumi from "@pulumi/pulumi";
import * as azure_native from "@pulumi/azure-native";

const config = new pulumi.Config();
const applicationInsightsNameParam = config.get("applicationInsightsNameParam") || "myApplicationInsights";
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const resourceGroupVar = azure_native.resources.getResourceGroup({
    resourceGroupName: resourceGroupNameParam,
});
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);
const workspaceNameParam = config.get("workspaceNameParam") || "myWorkspace";
const workspaceResource = new azure_native.operationalinsights.v20211201preview.Workspace("workspaceResource", {
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,
    sku: {
        name: "Free",
    },
    workspaceName: workspaceNameParam,
});
const componentResource = new azure_native.insights.v20200202preview.Component("componentResource", {
    applicationType: "web",
    kind: "web",
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,
    resourceName: applicationInsightsNameParam,
    workspaceResourceId: workspaceResource.id,
}, {
    dependsOn: [workspaceResource],
});
const emailActionGroup = new azure_native.insights.v20230101.ActionGroup("emailActionGroup", {
    actionGroupName: "emailActionGroup",
    emailReceivers: [{
        emailAddress: "example@test.com",
        name: "Example",
        useCommonAlertSchema: true,
    }],
    enabled: true,
    groupShortName: "string",
    location: "global",
    resourceGroupName: resourceGroupNameParam,
});
const responseAlertNameVar = "[concat('ResponseTime-', toLower(parameters('applicationInsightsName')))]";
const responseTimeParam = config.getNumber("responseTimeParam") || 3;
const metricAlertResource = new azure_native.insights.v20180301.MetricAlert("metricAlertResource", {
    actions: [{
        actionGroupId: emailActionGroup.id,
    }],
    criteria: {
        allOf: [{
            criterionType: "StaticThresholdCriterion",
            metricName: "requests/duration",
            name: "1st criterion",
            operator: "GreaterThan",
            threshold: responseTimeParam,
            timeAggregation: "Average",
        }],
        odataType: "Microsoft.Azure.Monitor.SingleResourceMultipleMetricCriteria",
    },
    description: "response time alert",
    enabled: true,
    evaluationFrequency: "PT1M",
    location: "global",
    resourceGroupName: resourceGroupNameParam,
    ruleName: responseAlertNameVar,
    scopes: [componentResource.id],
    severity: 0,
    windowSize: "PT5M",
}, {
    dependsOn: [
        componentResource,
        emailActionGroup,
    ],
});
