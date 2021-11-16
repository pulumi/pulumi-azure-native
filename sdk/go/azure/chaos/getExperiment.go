


package chaos

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExperiment(ctx *pulumi.Context, args *LookupExperimentArgs, opts ...pulumi.InvokeOption) (*LookupExperimentResult, error) {
	var rv LookupExperimentResult
	err := ctx.Invoke("azure-native:chaos:getExperiment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExperimentArgs struct {
	ExperimentName    string `pulumi:"experimentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupExperimentResult struct {
	Id         string                       `pulumi:"id"`
	Identity   *ResourceIdentityResponse    `pulumi:"identity"`
	Location   string                       `pulumi:"location"`
	Name       string                       `pulumi:"name"`
	Properties ExperimentPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse           `pulumi:"systemData"`
	Tags       map[string]string            `pulumi:"tags"`
	Type       string                       `pulumi:"type"`
}
