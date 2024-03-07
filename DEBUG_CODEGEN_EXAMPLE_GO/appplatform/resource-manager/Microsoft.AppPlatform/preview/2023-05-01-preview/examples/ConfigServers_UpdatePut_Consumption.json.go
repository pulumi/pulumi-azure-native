package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := appplatform.NewConfigServer(ctx, "configServer", &appplatform.ConfigServerArgs{
Properties: appplatform.ConfigServerPropertiesResponse{
ConfigServer: interface{}{
GitProperty: &appplatform.ConfigServerGitPropertyArgs{
Label: pulumi.String("master"),
SearchPaths: pulumi.StringArray{
pulumi.String("/"),
},
Uri: pulumi.String("https://github.com/fake-user/fake-repository.git"),
},
},
EnabledState: pulumi.String("Enabled"),
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
