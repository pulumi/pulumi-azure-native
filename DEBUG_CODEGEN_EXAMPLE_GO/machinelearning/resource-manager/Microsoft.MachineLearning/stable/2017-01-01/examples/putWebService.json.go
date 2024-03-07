package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearning/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := machinelearning.NewWebService(ctx, "webService", &machinelearning.WebServiceArgs{
Location: pulumi.String("West US"),
Properties: &machinelearning.WebServicePropertiesForGraphArgs{
Assets: interface{}{
Asset1: &machinelearning.AssetItemArgs{
LocationInfo: &machinelearning.BlobLocationArgs{
Credentials: pulumi.String(""),
Uri: pulumi.String("aml://module/moduleId-1"),
},
Name: pulumi.String("Execute R Script"),
Type: pulumi.String("Module"),
},
Asset2: &machinelearning.AssetItemArgs{
LocationInfo: &machinelearning.BlobLocationArgs{
Credentials: pulumi.String(""),
Uri: pulumi.String("aml://module/moduleId-2"),
},
Name: pulumi.String("Import Data"),
Type: pulumi.String("Module"),
},
},
CommitmentPlan: &machinelearning.CommitmentPlanTypeArgs{
Id: pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.MachineLearning/commitmentPlans/commitmentPlanName"),
},
Description: pulumi.String("Web Service Description"),
Diagnostics: &machinelearning.DiagnosticsConfigurationArgs{
Level: pulumi.String("None"),
},
ExampleRequest: &machinelearning.ExampleRequestArgs{
Inputs: pulumi.ArrayArrayMap{
"input1": pulumi.ArrayArray{
pulumi.Array{
pulumi.Any("age"),
},
pulumi.Array{
pulumi.Any("workclass"),
},
pulumi.Array{
pulumi.Any("fnlwgt"),
},
pulumi.Array{
pulumi.Any("education"),
},
pulumi.Array{
pulumi.Any("education-num"),
},
},
},
},
ExposeSampleData: pulumi.Bool(true),
Input: &machinelearning.ServiceInputOutputSpecificationArgs{
Description: pulumi.String(""),
Properties: machinelearning.TableSpecificationMap{
"input1": &machinelearning.TableSpecificationArgs{
Description: pulumi.String(""),
Properties: machinelearning.ColumnSpecificationMap{
"column_name": &machinelearning.ColumnSpecificationArgs{
Type: pulumi.String("String"),
XMsIsnullable: pulumi.Bool(false),
},
},
Title: pulumi.String(""),
Type: pulumi.String("object"),
},
},
Title: pulumi.String(""),
Type: pulumi.String("object"),
},
MachineLearningWorkspace: &machinelearning.MachineLearningWorkspaceArgs{
Id: pulumi.String("workspaceId"),
},
Output: &machinelearning.ServiceInputOutputSpecificationArgs{
Description: pulumi.String(""),
Properties: machinelearning.TableSpecificationMap{
"output1": &machinelearning.TableSpecificationArgs{
Description: pulumi.String(""),
Properties: machinelearning.ColumnSpecificationMap{
"age": &machinelearning.ColumnSpecificationArgs{
Format: pulumi.String("Int32"),
Type: pulumi.String("Integer"),
XMsIsnullable: pulumi.Bool(true),
},
"workclass": &machinelearning.ColumnSpecificationArgs{
Type: pulumi.String("String"),
XMsIsnullable: pulumi.Bool(false),
},
},
Title: pulumi.String(""),
Type: pulumi.String("object"),
},
},
Title: pulumi.String(""),
Type: pulumi.String("object"),
},
Package: &machinelearning.GraphPackageArgs{
Edges: machinelearning.GraphEdgeArray{
&machinelearning.GraphEdgeArgs{
SourceNodeId: pulumi.String("node2"),
SourcePortId: pulumi.String("Results dataset"),
TargetNodeId: pulumi.String("node1"),
TargetPortId: pulumi.String("Dataset2"),
},
&machinelearning.GraphEdgeArgs{
SourceNodeId: pulumi.String("node3"),
TargetNodeId: pulumi.String("node1"),
TargetPortId: pulumi.String("Dataset1"),
},
&machinelearning.GraphEdgeArgs{
SourceNodeId: pulumi.String("node1"),
SourcePortId: pulumi.String("Result Dataset"),
TargetNodeId: pulumi.String("node4"),
},
},
GraphParameters: nil,
Nodes: machinelearning.GraphNodeMap{
"node1": &machinelearning.GraphNodeArgs{
AssetId: pulumi.String("asset1"),
Parameters: machinelearning.WebServiceParameterMap{
"R Script": &machinelearning.WebServiceParameterArgs{
CertificateThumbprint: pulumi.String(""),
Value: pulumi.Any("The R Script"),
},
"R Version": &machinelearning.WebServiceParameterArgs{
CertificateThumbprint: pulumi.String(""),
Value: pulumi.Any("CRAN R 3.1.0"),
},
},
},
"node2": &machinelearning.GraphNodeArgs{
AssetId: pulumi.String("asset2"),
Parameters: machinelearning.WebServiceParameterMap{
"Account Key": &machinelearning.WebServiceParameterArgs{
CertificateThumbprint: pulumi.String("TheThumbprint"),
Value: pulumi.Any("Encrypted Key"),
},
"Account Name": &machinelearning.WebServiceParameterArgs{
CertificateThumbprint: pulumi.String(""),
Value: pulumi.Any("accountName"),
},
"Please Specify Authentication Type": &machinelearning.WebServiceParameterArgs{
CertificateThumbprint: pulumi.String(""),
Value: pulumi.Any("Account"),
},
"Please Specify Data Source": &machinelearning.WebServiceParameterArgs{
CertificateThumbprint: pulumi.String(""),
Value: pulumi.Any("AzureBlobStorage"),
},
},
},
"node3": &machinelearning.GraphNodeArgs{
InputId: pulumi.String("input1"),
},
"node4": &machinelearning.GraphNodeArgs{
OutputId: pulumi.String("output1"),
},
},
},
PackageType: pulumi.String("Graph"),
Parameters: nil,
PayloadsInBlobStorage: pulumi.Bool(false),
ReadOnly: pulumi.Bool(false),
RealtimeConfiguration: &machinelearning.RealtimeConfigurationArgs{
MaxConcurrentCalls: pulumi.Int(4),
},
StorageAccount: &machinelearning.StorageAccountArgs{
Key: pulumi.String("Storage_Key"),
Name: pulumi.String("Storage_Name"),
},
Title: pulumi.String("Web Service Title"),
},
ResourceGroupName: pulumi.String("OneResourceGroupName"),
Tags: pulumi.StringMap{
"tag1": pulumi.String("value1"),
"tag2": pulumi.String("value2"),
},
WebServiceName: pulumi.String("TargetWebServiceName"),
})
if err != nil {
return err
}
return nil
})
}
