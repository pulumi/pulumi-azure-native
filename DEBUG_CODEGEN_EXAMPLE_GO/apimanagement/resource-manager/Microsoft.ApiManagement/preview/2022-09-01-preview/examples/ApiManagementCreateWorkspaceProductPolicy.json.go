package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewWorkspaceProductPolicy(ctx, "workspaceProductPolicy", &apimanagement.WorkspaceProductPolicyArgs{
			Format:            pulumi.String("xml"),
			PolicyId:          pulumi.String("policy"),
			ProductId:         pulumi.String("5702e97e5157a50f48dce801"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Value: pulumi.String(`<policies>
  <inbound>
    <rate-limit calls="{{call-count}}" renewal-period="15"></rate-limit>
    <log-to-eventhub logger-id="16">
                      @( string.Join(",", DateTime.UtcNow, context.Deployment.ServiceName, context.RequestId, context.Request.IpAddress, context.Operation.Name) ) 
                  </log-to-eventhub>
    <quota-by-key calls="40" counter-key="cc" renewal-period="3600" increment-count="@(context.Request.Method == &quot;POST&quot; ? 1:2)" />
    <base />
  </inbound>
  <backend>
    <base />
  </backend>
  <outbound>
    <base />
  </outbound>
</policies>`),
			WorkspaceId: pulumi.String("wks1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
