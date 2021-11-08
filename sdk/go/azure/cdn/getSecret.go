


package cdn

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecret(ctx *pulumi.Context, args *LookupSecretArgs, opts ...pulumi.InvokeOption) (*LookupSecretResult, error) {
	var rv LookupSecretResult
	err := ctx.Invoke("azure-native:cdn:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecretArgs struct {
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SecretName        string `pulumi:"secretName"`
}


type LookupSecretResult struct {
	DeploymentStatus  string             `pulumi:"deploymentStatus"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	Parameters        interface{}        `pulumi:"parameters"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}
