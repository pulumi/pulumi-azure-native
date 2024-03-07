package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridcontainerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := hybridcontainerservice.NewVirtualNetworkRetrieve(ctx, "virtualNetworkRetrieve", &hybridcontainerservice.VirtualNetworkRetrieveArgs{
ExtendedLocation: &hybridcontainerservice.VirtualNetworksExtendedLocationArgs{
Name: pulumi.String("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation"),
Type: pulumi.String("CustomLocation"),
},
Location: pulumi.String("westus"),
Properties: hybridcontainerservice.VirtualNetworksPropertiesResponse{
InfraVnetProfile: interface{}{
Hci: &hybridcontainerservice.VirtualNetworksPropertiesHciArgs{
MocGroup: pulumi.String("target-group"),
MocLocation: pulumi.String("MocLocation"),
MocVnetName: pulumi.String("test-vnet"),
},
},
VipPool: hybridcontainerservice.VirtualNetworksPropertiesVipPoolArray{
&hybridcontainerservice.VirtualNetworksPropertiesVipPoolArgs{
EndIP: pulumi.String("192.168.0.50"),
StartIP: pulumi.String("192.168.0.10"),
},
},
VmipPool: hybridcontainerservice.VirtualNetworksPropertiesVmipPoolArray{
&hybridcontainerservice.VirtualNetworksPropertiesVmipPoolArgs{
EndIP: pulumi.String("192.168.0.130"),
StartIP: pulumi.String("192.168.0.110"),
},
},
},
ResourceGroupName: pulumi.String("test-arcappliance-resgrp"),
VirtualNetworksName: pulumi.String("test-vnet-static"),
})
if err != nil {
return err
}
return nil
})
}
