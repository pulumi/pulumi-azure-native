package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewAuthorizationAccessPolicy(ctx, "authorizationAccessPolicy", &apimanagement.AuthorizationAccessPolicyArgs{
			AuthorizationAccessPolicyId: pulumi.String("fe0bed83-631f-4149-bd0b-0464b1bc7cab"),
			AuthorizationId:             pulumi.String("authz1"),
			AuthorizationProviderId:     pulumi.String("aadwithauthcode"),
			ObjectId:                    pulumi.String("fe0bed83-631f-4149-bd0b-0464b1bc7cab"),
			ResourceGroupName:           pulumi.String("rg1"),
			ServiceName:                 pulumi.String("apimService1"),
			TenantId:                    pulumi.String("13932a0d-5c63-4d37-901d-1df9c97722ff"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
