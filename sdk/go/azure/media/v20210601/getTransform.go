


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTransform(ctx *pulumi.Context, args *LookupTransformArgs, opts ...pulumi.InvokeOption) (*LookupTransformResult, error) {
	var rv LookupTransformResult
	err := ctx.Invoke("azure-native:media/v20210601:getTransform", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTransformArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TransformName     string `pulumi:"transformName"`
}


type LookupTransformResult struct {
	Created      string                    `pulumi:"created"`
	Description  *string                   `pulumi:"description"`
	Id           string                    `pulumi:"id"`
	LastModified string                    `pulumi:"lastModified"`
	Name         string                    `pulumi:"name"`
	Outputs      []TransformOutputResponse `pulumi:"outputs"`
	SystemData   SystemDataResponse        `pulumi:"systemData"`
	Type         string                    `pulumi:"type"`
}
