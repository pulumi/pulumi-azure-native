package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApiIssueAttachment(ctx, "apiIssueAttachment", &apimanagement.ApiIssueAttachmentArgs{
			ApiId:             pulumi.String("57d1f7558aa04f15146d9d8a"),
			AttachmentId:      pulumi.String("57d2ef278aa04f0888cba3f3"),
			Content:           pulumi.String("IEJhc2U2NA=="),
			ContentFormat:     pulumi.String("image/jpeg"),
			IssueId:           pulumi.String("57d2ef278aa04f0ad01d6cdc"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Title:             pulumi.String("Issue attachment."),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
