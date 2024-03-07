package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewEmailTemplate(ctx, "emailTemplate", &apimanagement.EmailTemplateArgs{
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Subject:           pulumi.String("Your request for $IssueName was successfully received."),
			TemplateName:      pulumi.String("newIssueNotificationMessage"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
