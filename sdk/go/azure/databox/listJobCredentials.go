


package databox

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListJobCredentials(ctx *pulumi.Context, args *ListJobCredentialsArgs, opts ...pulumi.InvokeOption) (*ListJobCredentialsResult, error) {
	var rv ListJobCredentialsResult
	err := ctx.Invoke("azure-native:databox:listJobCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListJobCredentialsArgs struct {
	JobName           string `pulumi:"jobName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListJobCredentialsResult struct {
	NextLink *string                          `pulumi:"nextLink"`
	Value    []UnencryptedCredentialsResponse `pulumi:"value"`
}
