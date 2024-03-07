package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := appplatform.NewContainerRegistry(ctx, "containerRegistry", &appplatform.ContainerRegistryArgs{
ContainerRegistryName: pulumi.String("my-container-registry"),
Properties: appplatform.ContainerRegistryPropertiesResponse{
Credentials: interface{}{
Password: pulumi.String("myPassword"),
Server: pulumi.String("myServer"),
Type: pulumi.String("BasicAuth"),
Username: pulumi.String("myUsername"),
},
},
ResourceGroupName: pulumi.String("myResourceGroup"),
ServiceName: pulumi.String("my-service"),
})
if err != nil {
return err
}
return nil
})
}
