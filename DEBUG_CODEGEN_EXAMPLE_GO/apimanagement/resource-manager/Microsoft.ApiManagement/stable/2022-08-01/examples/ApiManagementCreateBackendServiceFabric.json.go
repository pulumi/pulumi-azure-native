package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := apimanagement.NewBackend(ctx, "backend", &apimanagement.BackendArgs{
BackendId: pulumi.String("sfbackend"),
Description: pulumi.String("Service Fabric Test App 1"),
Properties: apimanagement.BackendPropertiesResponse{
ServiceFabricCluster: interface{}{
ClientCertificateId: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/cert1"),
ManagementEndpoints: pulumi.StringArray{
pulumi.String("https://somecluster.com"),
},
MaxPartitionResolutionRetries: pulumi.Int(5),
ServerX509Names: apimanagement.X509CertificateNameArray{
&apimanagement.X509CertificateNameArgs{
IssuerCertificateThumbprint: pulumi.String("IssuerCertificateThumbprint1"),
Name: pulumi.String("ServerCommonName1"),
},
},
},
},
Protocol: pulumi.String("http"),
ResourceGroupName: pulumi.String("rg1"),
ServiceName: pulumi.String("apimService1"),
Url: pulumi.String("fabric:/mytestapp/mytestservice"),
})
if err != nil {
return err
}
return nil
})
}
