


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCodeVersion(ctx *pulumi.Context, args *LookupCodeVersionArgs, opts ...pulumi.InvokeOption) (*LookupCodeVersionResult, error) {
	var rv LookupCodeVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:getCodeVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCodeVersionArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Version           string `pulumi:"version"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupCodeVersionResult struct {
	Id         string              `pulumi:"id"`
	Name       string              `pulumi:"name"`
	Properties CodeVersionResponse `pulumi:"properties"`
	SystemData SystemDataResponse  `pulumi:"systemData"`
	Type       string              `pulumi:"type"`
}
