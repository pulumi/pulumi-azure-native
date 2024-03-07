package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := appplatform.NewDevToolPortal(ctx, "devToolPortal", &appplatform.DevToolPortalArgs{
DevToolPortalName: pulumi.String("default"),
Properties: appplatform.DevToolPortalPropertiesResponse{
Features: interface{}{
ApplicationAccelerator: &appplatform.DevToolPortalFeatureDetailArgs{
State: pulumi.String("Enabled"),
},
ApplicationLiveView: &appplatform.DevToolPortalFeatureDetailArgs{
State: pulumi.String("Enabled"),
},
},
Public: pulumi.Bool(true),
SsoProperties: &appplatform.DevToolPortalSsoPropertiesArgs{
ClientId: pulumi.String("00000000-0000-0000-0000-000000000000"),
ClientSecret: pulumi.String("xxxxx"),
MetadataUrl: pulumi.String("https://login.microsoftonline.com/00000000-0000-0000-0000-000000000000/v2.0/.well-known/openid-configuration"),
Scopes: pulumi.StringArray{
pulumi.String("openid"),
},
},
},
ResourceGroupName: pulumi.String("myResourceGroup"),
ServiceName: pulumi.String("myservice"),
})
if err != nil {
return err
}
return nil
})
}
