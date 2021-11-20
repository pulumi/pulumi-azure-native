


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDatastoreSecrets(ctx *pulumi.Context, args *ListDatastoreSecretsArgs, opts ...pulumi.InvokeOption) (*ListDatastoreSecretsResult, error) {
	var rv ListDatastoreSecretsResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:listDatastoreSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatastoreSecretsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListDatastoreSecretsResult struct {
	SecretsType string `pulumi:"secretsType"`
}
