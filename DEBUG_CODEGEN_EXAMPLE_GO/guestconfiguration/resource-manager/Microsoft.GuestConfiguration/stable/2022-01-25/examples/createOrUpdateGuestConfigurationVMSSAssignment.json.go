package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/guestconfiguration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := guestconfiguration.NewGuestConfigurationAssignmentsVMSS(ctx, "guestConfigurationAssignmentsVMSS", &guestconfiguration.GuestConfigurationAssignmentsVMSSArgs{
Location: pulumi.String("westcentralus"),
Name: pulumi.String("NotInstalledApplicationForWindows"),
Properties: guestconfiguration.GuestConfigurationAssignmentPropertiesResponse{
Context: pulumi.String("Azure policy"),
GuestConfiguration: interface{}{
AssignmentType: pulumi.String("ApplyAndAutoCorrect"),
ConfigurationParameter: guestconfiguration.ConfigurationParameterArray{
&guestconfiguration.ConfigurationParameterArgs{
Name: pulumi.String("[InstalledApplication]NotInstalledApplicationResource1;Name"),
Value: pulumi.String("NotePad,sql"),
},
},
ContentHash: pulumi.String("123contenthash"),
ContentUri: pulumi.String("https://thisisfake/pacakge"),
Name: pulumi.String("NotInstalledApplicationForWindows"),
Version: pulumi.String("1.0.0.3"),
},
},
ResourceGroupName: pulumi.String("myResourceGroupName"),
VmssName: pulumi.String("myVMSSName"),
})
if err != nil {
return err
}
return nil
})
}
