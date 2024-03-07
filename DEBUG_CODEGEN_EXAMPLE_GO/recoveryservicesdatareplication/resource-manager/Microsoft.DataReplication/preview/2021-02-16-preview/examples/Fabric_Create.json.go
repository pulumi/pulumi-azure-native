package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datareplication/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := datareplication.NewFabric(ctx, "fabric", &datareplication.FabricArgs{
FabricName: pulumi.String("wPR"),
Location: pulumi.String("tqygutlpob"),
Properties: &datareplication.FabricModelPropertiesArgs{
CustomProperties: interface{}{
InstanceType: pulumi.String("FabricModelCustomProperties"),
},
},
ResourceGroupName: pulumi.String("rgrecoveryservicesdatareplication"),
Tags: pulumi.StringMap{
"key3917": pulumi.String("vgralu"),
},
})
if err != nil {
return err
}
return nil
})
}
