package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewWorkspacePolicy(ctx, "workspacePolicy", &apimanagement.WorkspacePolicyArgs{
			Format:            pulumi.String("rawxml"),
			PolicyId:          pulumi.String("policy"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Value: pulumi.String(`<policies>
     <inbound>
     <base />
  <set-header name="newvalue" exists-action="override">
   <value>"@(context.Request.Headers.FirstOrDefault(h => h.Ke=="Via"))" </value>
    </set-header>
  </inbound>
      </policies>`),
			WorkspaceId: pulumi.String("wks1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
