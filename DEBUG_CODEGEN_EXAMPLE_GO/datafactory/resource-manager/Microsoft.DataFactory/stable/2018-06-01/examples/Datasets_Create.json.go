package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datafactory/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := datafactory.NewDataset(ctx, "dataset", &datafactory.DatasetArgs{
DatasetName: pulumi.String("exampleDataset"),
FactoryName: pulumi.String("exampleFactoryName"),
Properties: datafactory.AzureBlobDataset{
FileName: map[string]interface{}{
"type": "Expression",
"value": "@dataset().MyFileName",
},
FolderPath: map[string]interface{}{
"type": "Expression",
"value": "@dataset().MyFolderPath",
},
Format: datafactory.TextFormat{
Type: "TextFormat",
},
LinkedServiceName: datafactory.LinkedServiceReference{
ReferenceName: "exampleLinkedService",
Type: "LinkedServiceReference",
},
Parameters: interface{}{
MyFileName: datafactory.ParameterSpecification{
Type: "String",
},
MyFolderPath: datafactory.ParameterSpecification{
Type: "String",
},
},
Type: "AzureBlob",
},
ResourceGroupName: pulumi.String("exampleResourceGroup"),
})
if err != nil {
return err
}
return nil
})
}
