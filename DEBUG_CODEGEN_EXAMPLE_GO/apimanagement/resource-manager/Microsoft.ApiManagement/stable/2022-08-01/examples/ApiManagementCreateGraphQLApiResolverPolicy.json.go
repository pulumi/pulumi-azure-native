package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewGraphQLApiResolverPolicy(ctx, "graphQLApiResolverPolicy", &apimanagement.GraphQLApiResolverPolicyArgs{
			ApiId:             pulumi.String("5600b57e7e8880006a040001"),
			Format:            pulumi.String("xml"),
			PolicyId:          pulumi.String("policy"),
			ResolverId:        pulumi.String("5600b57e7e8880006a080001"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Value:             pulumi.String("<http-data-source><http-request><set-method>GET</set-method><set-backend-service base-url=\"https://some.service.com\" /><set-url>/api/users</set-url></http-request></http-data-source>"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
