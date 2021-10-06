


package streamanalytics

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInput(ctx *pulumi.Context, args *LookupInputArgs, opts ...pulumi.InvokeOption) (*LookupInputResult, error) {
	var rv LookupInputResult
	err := ctx.Invoke("azure-native:streamanalytics:getInput", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInputArgs struct {
	InputName         string `pulumi:"inputName"`
	JobName           string `pulumi:"jobName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupInputResult struct {
	Id         string      `pulumi:"id"`
	Name       *string     `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}
