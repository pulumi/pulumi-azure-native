package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := app.NewContainerAppsAuthConfig(ctx, "containerAppsAuthConfig", &app.ContainerAppsAuthConfigArgs{
AuthConfigName: pulumi.String("current"),
ContainerAppName: pulumi.String("testcanadacentral"),
GlobalValidation: &app.GlobalValidationArgs{
UnauthenticatedClientAction: app.UnauthenticatedClientActionV2AllowAnonymous,
},
IdentityProviders: app.IdentityProvidersResponse{
Facebook: interface{}{
Registration: &app.AppRegistrationArgs{
AppId: pulumi.String("123"),
AppSecretSettingName: pulumi.String("facebook-secret"),
},
},
},
Platform: &app.AuthPlatformArgs{
Enabled: pulumi.Bool(true),
},
ResourceGroupName: pulumi.String("workerapps-rg-xj"),
})
if err != nil {
return err
}
return nil
})
}
