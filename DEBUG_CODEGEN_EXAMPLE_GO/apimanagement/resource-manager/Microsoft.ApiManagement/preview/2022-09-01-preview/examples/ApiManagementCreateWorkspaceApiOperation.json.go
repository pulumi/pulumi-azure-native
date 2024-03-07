package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewWorkspaceApiOperation(ctx, "workspaceApiOperation", &apimanagement.WorkspaceApiOperationArgs{
			ApiId:       pulumi.String("PetStoreTemplate2"),
			Description: pulumi.String("This can only be done by the logged in user."),
			DisplayName: pulumi.String("createUser2"),
			Method:      pulumi.String("POST"),
			OperationId: pulumi.String("newoperations"),
			Request: &apimanagement.RequestContractArgs{
				Description:     pulumi.String("Created user object"),
				Headers:         apimanagement.ParameterContractArray{},
				QueryParameters: apimanagement.ParameterContractArray{},
				Representations: apimanagement.RepresentationContractArray{
					&apimanagement.RepresentationContractArgs{
						ContentType: pulumi.String("application/json"),
						SchemaId:    pulumi.String("592f6c1d0af5840ca8897f0c"),
						TypeName:    pulumi.String("User"),
					},
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			Responses: apimanagement.ResponseContractArray{
				&apimanagement.ResponseContractArgs{
					Description: pulumi.String("successful operation"),
					Headers:     apimanagement.ParameterContractArray{},
					Representations: apimanagement.RepresentationContractArray{
						&apimanagement.RepresentationContractArgs{
							ContentType: pulumi.String("application/xml"),
						},
						&apimanagement.RepresentationContractArgs{
							ContentType: pulumi.String("application/json"),
						},
					},
					StatusCode: pulumi.Int(200),
				},
			},
			ServiceName:        pulumi.String("apimService1"),
			TemplateParameters: apimanagement.ParameterContractArray{},
			UrlTemplate:        pulumi.String("/user1"),
			WorkspaceId:        pulumi.String("wks1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
