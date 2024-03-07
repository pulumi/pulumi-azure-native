package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApiTagDescription(ctx, "apiTagDescription", &apimanagement.ApiTagDescriptionArgs{
			ApiId:                   pulumi.String("5931a75ae4bbd512a88c680b"),
			Description:             pulumi.String("Some description that will be displayed for operation's tag if the tag is assigned to operation of the API"),
			ExternalDocsDescription: pulumi.String("Description of the external docs resource"),
			ExternalDocsUrl:         pulumi.String("http://some.url/additionaldoc"),
			ResourceGroupName:       pulumi.String("rg1"),
			ServiceName:             pulumi.String("apimService1"),
			TagDescriptionId:        pulumi.String("tagId1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
