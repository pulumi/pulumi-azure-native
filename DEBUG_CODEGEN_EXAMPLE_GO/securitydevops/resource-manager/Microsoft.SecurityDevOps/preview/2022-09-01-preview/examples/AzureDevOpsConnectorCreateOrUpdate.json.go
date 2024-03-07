package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securitydevops/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := securitydevops.NewAzureDevOpsConnector(ctx, "azureDevOpsConnector", &securitydevops.AzureDevOpsConnectorArgs{
AzureDevOpsConnectorName: pulumi.String("testconnector"),
Location: pulumi.String("West US"),
Properties: securitydevops.AzureDevOpsConnectorPropertiesResponse{
Authorization: &securitydevops.AuthorizationInfoArgs{
Code: pulumi.String("00000000000000000000"),
},
Orgs: securitydevops.AzureDevOpsOrgMetadataArray{
interface{}{
Name: pulumi.String("testOrg"),
Projects: securitydevops.AzureDevOpsProjectMetadataArray{
&securitydevops.AzureDevOpsProjectMetadataArgs{
Name: pulumi.String("testProject"),
Repos: pulumi.StringArray{
pulumi.String("testRepo"),
},
},
},
},
},
},
ResourceGroupName: pulumi.String("westusrg"),
})
if err != nil {
return err
}
return nil
})
}
