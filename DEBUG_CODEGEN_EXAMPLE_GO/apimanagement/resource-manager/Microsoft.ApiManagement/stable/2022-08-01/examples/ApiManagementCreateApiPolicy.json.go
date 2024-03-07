package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApiPolicy(ctx, "apiPolicy", &apimanagement.ApiPolicyArgs{
			ApiId:             pulumi.String("5600b57e7e8880006a040001"),
			Format:            pulumi.String("xml"),
			PolicyId:          pulumi.String("policy"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Value:             pulumi.String("<policies> <inbound /> <backend>    <forward-request />  </backend>  <outbound /></policies>"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
