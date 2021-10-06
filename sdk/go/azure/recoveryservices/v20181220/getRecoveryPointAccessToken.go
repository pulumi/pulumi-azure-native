


package v20181220

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRecoveryPointAccessToken(ctx *pulumi.Context, args *GetRecoveryPointAccessTokenArgs, opts ...pulumi.InvokeOption) (*GetRecoveryPointAccessTokenResult, error) {
	var rv GetRecoveryPointAccessTokenResult
	err := ctx.Invoke("azure-native:recoveryservices/v20181220:getRecoveryPointAccessToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetRecoveryPointAccessTokenArgs struct {
	ContainerName     string            `pulumi:"containerName"`
	ETag              *string           `pulumi:"eTag"`
	FabricName        string            `pulumi:"fabricName"`
	Location          *string           `pulumi:"location"`
	Properties        *AADProperties    `pulumi:"properties"`
	ProtectedItemName string            `pulumi:"protectedItemName"`
	RecoveryPointId   string            `pulumi:"recoveryPointId"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	VaultName         string            `pulumi:"vaultName"`
}

type GetRecoveryPointAccessTokenResult struct {
	ETag       *string                        `pulumi:"eTag"`
	Id         string                         `pulumi:"id"`
	Location   *string                        `pulumi:"location"`
	Name       string                         `pulumi:"name"`
	Properties WorkloadCrrAccessTokenResponse `pulumi:"properties"`
	Tags       map[string]string              `pulumi:"tags"`
	Type       string                         `pulumi:"type"`
}
