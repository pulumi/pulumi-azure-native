package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := devtestlab.NewFormula(ctx, "formula", &devtestlab.FormulaArgs{
Description: pulumi.String("Formula using a Linux base"),
FormulaContent: devtestlab.LabVirtualMachineCreationParameterResponse{
AllowClaim: pulumi.Bool(false),
Artifacts: devtestlab.ArtifactInstallPropertiesArray{
interface{}{
ArtifactId: pulumi.String("/artifactsources/{artifactSourceName}/artifacts/linux-install-nodejs"),
Parameters: devtestlab.ArtifactParameterPropertiesArray{
},
},
},
DisallowPublicIpAddress: pulumi.Bool(true),
GalleryImageReference: &devtestlab.GalleryImageReferenceArgs{
Offer: pulumi.String("0001-com-ubuntu-server-groovy"),
OsType: pulumi.String("Linux"),
Publisher: pulumi.String("canonical"),
Sku: pulumi.String("20_10"),
Version: pulumi.String("latest"),
},
IsAuthenticationWithSshKey: pulumi.Bool(false),
LabSubnetName: pulumi.String("Dtl{labName}Subnet"),
LabVirtualNetworkId: pulumi.String("/virtualnetworks/dtl{labName}"),
Location: pulumi.String("{location}"),
NetworkInterface: interface{}{
SharedPublicIpAddressConfiguration: interface{}{
InboundNatRules: devtestlab.InboundNatRuleArray{
&devtestlab.InboundNatRuleArgs{
BackendPort: pulumi.Int(22),
TransportProtocol: pulumi.String("Tcp"),
},
},
},
},
Notes: pulumi.String("Ubuntu Server 20.10"),
Size: pulumi.String("Standard_B1ms"),
StorageType: pulumi.String("Standard"),
UserName: pulumi.String("user"),
},
LabName: pulumi.String("{labName}"),
Location: pulumi.String("{location}"),
Name: pulumi.String("{formulaName}"),
ResourceGroupName: pulumi.String("resourceGroupName"),
})
if err != nil {
return err
}
return nil
})
}
