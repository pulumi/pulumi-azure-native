package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := securityinsights.NewCustomizableConnectorDefinition(ctx, "customizableConnectorDefinition", &securityinsights.CustomizableConnectorDefinitionArgs{
ConnectorUiConfig: securityinsights.CustomizableConnectorUiConfigResponse{
Availability: &securityinsights.ConnectorDefinitionsAvailabilityArgs{
IsPreview: pulumi.Bool(false),
Status: pulumi.Int(1),
},
ConnectivityCriteria: securityinsights.ConnectivityCriterionArray{
&securityinsights.ConnectivityCriterionArgs{
Type: pulumi.String("IsConnectedQuery"),
Value: pulumi.StringArray{
pulumi.String("GitHubAuditLogPolling_CL \n | summarize LastLogReceived = max(TimeGenerated)\n | project IsConnected = LastLogReceived > ago(30d)"),
},
},
},
DataTypes: securityinsights.ConnectorDataTypeArray{
&securityinsights.ConnectorDataTypeArgs{
LastDataReceivedQuery: pulumi.String("GitHubAuditLogPolling_CL \n            | summarize Time = max(TimeGenerated)\n            | where isnotempty(Time)"),
Name: pulumi.String("GitHubAuditLogPolling_CL"),
},
},
DescriptionMarkdown: pulumi.String("The GitHub audit log connector provides the capability to ingest GitHub logs into Azure Sentinel. By connecting GitHub audit logs into Azure Sentinel, you can view this data in workbooks, use it to create custom alerts, and improve your investigation process."),
GraphQueries: securityinsights.GraphQueryArray{
&securityinsights.GraphQueryArgs{
BaseQuery: pulumi.String("GitHubAuditLogPolling_CL"),
Legend: pulumi.String("GitHub audit log events"),
MetricName: pulumi.String("Total events received"),
},
},
InstructionSteps: []securityinsights.InstructionStepArgs{
{
Description: pulumi.String("Enable GitHub audit Logs. \n Follow [this](https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token) to create or find your personal key"),
Instructions: securityinsights.InstructionStepDetailsArray{
{
Parameters: {
ClientIdLabel: "Client ID",
ClientSecretLabel: "Client Secret",
ConnectButtonLabel: "Connect",
DisconnectButtonLabel: "Disconnect",
},
Type: pulumi.String("OAuthForm"),
},
},
Title: pulumi.String("Connect GitHub Enterprise Audit Log to Azure Sentinel"),
},
},
Permissions: interface{}{
Customs: securityinsights.CustomPermissionDetailsArray{
&securityinsights.CustomPermissionDetailsArgs{
Description: pulumi.String("You need access to GitHub personal token, the key should have 'admin:org' scope"),
Name: pulumi.String("GitHub API personal token Key"),
},
},
ResourceProvider: securityinsights.ConnectorDefinitionsResourceProviderArray{
interface{}{
PermissionsDisplayText: pulumi.String("read and write permissions are required."),
Provider: pulumi.String("Microsoft.OperationalInsights/workspaces"),
ProviderDisplayName: pulumi.String("Workspace"),
RequiredPermissions: &securityinsights.ResourceProviderRequiredPermissionsArgs{
Action: pulumi.Bool(false),
Delete: pulumi.Bool(false),
Read: pulumi.Bool(false),
Write: pulumi.Bool(true),
},
Scope: pulumi.String("Workspace"),
},
},
},
Publisher: pulumi.String("GitHub"),
SampleQueries: securityinsights.SampleQueryArray{
&securityinsights.SampleQueryArgs{
Description: pulumi.String("All logs"),
Query: pulumi.String("GitHubAuditLogPolling_CL \n | take 10"),
},
},
Title: pulumi.String("GitHub Enterprise Audit Log"),
},
DataConnectorDefinitionName: pulumi.String("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
Kind: pulumi.String("Customizable"),
ResourceGroupName: pulumi.String("myRg"),
WorkspaceName: pulumi.String("myWorkspace"),
})
if err != nil {
return err
}
return nil
})
}
