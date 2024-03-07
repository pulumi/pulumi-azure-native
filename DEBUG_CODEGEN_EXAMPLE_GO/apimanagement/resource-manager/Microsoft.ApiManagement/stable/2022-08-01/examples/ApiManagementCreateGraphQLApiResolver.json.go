package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewGraphQLApiResolver(ctx, "graphQLApiResolver", &apimanagement.GraphQLApiResolverArgs{
			ApiId:             pulumi.String("someAPI"),
			Description:       pulumi.String("A GraphQL Resolver example"),
			DisplayName:       pulumi.String("Query Users"),
			Path:              pulumi.String("Query/users"),
			ResolverId:        pulumi.String("newResolver"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
