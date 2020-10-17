import * as pulumi from "@pulumi/pulumi";
import * as azure_nextgen from "@pulumi/azure-nextgen";

const config = new pulumi.Config();
const applicationInsightsNameParam = config.get("applicationInsightsNameParam") || "myApplicationInsights";
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const resourceGroupVar = azure_nextgen.resources.latest.getResourceGroup({
    resourceGroupName: resourceGroupNameParam,
});
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);
const workspaceNameParam = config.get("workspaceNameParam") || "myWorkspace";
const workspaceResource = new azure_nextgen.operationalinsights.v20200301preview.Workspace("workspaceResource", {
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,
    sku: {
        name: "Free",
    },
    workspaceName: workspaceNameParam,
});
const componentResource = new azure_nextgen.insights.v20200202preview.Component("componentResource", {
    applicationType: "web",
    kind: "web",
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,
    resourceName: applicationInsightsNameParam,
    workspaceResourceId: "[resourceId('Microsoft.OperationalInsights/workspaces',parameters('workspaceName'))]",
}, {
    dependsOn: [workspaceResource],
});
const emailActionGroup = new azure_nextgen.insights.latest.ActionGroup("emailActionGroup", {
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
const metricAlertResource = new azure_nextgen.insights.latest.MetricAlert("metricAlertResource", {
    actions: [{
        actionGroupId: "[resourceId('microsoft.insights/actionGroups','emailActionGroup')]",
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
    scopes: ["[resourceId('Microsoft.Insights/components',parameters('applicationInsightsName'))]"],
    severity: 0,
    windowSize: "PT5M",
});
