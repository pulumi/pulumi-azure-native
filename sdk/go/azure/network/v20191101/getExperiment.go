


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExperiment(ctx *pulumi.Context, args *LookupExperimentArgs, opts ...pulumi.InvokeOption) (*LookupExperimentResult, error) {
	var rv LookupExperimentResult
	err := ctx.Invoke("azure-native:network/v20191101:getExperiment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExperimentArgs struct {
	ExperimentName    string `pulumi:"experimentName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupExperimentResult struct {
	Description   *string                     `pulumi:"description"`
	EnabledState  *string                     `pulumi:"enabledState"`
	EndpointA     *ExperimentEndpointResponse `pulumi:"endpointA"`
	EndpointB     *ExperimentEndpointResponse `pulumi:"endpointB"`
	Id            string                      `pulumi:"id"`
	Location      *string                     `pulumi:"location"`
	Name          string                      `pulumi:"name"`
	ResourceState string                      `pulumi:"resourceState"`
	ScriptFileUri string                      `pulumi:"scriptFileUri"`
	Status        string                      `pulumi:"status"`
	Tags          map[string]string           `pulumi:"tags"`
	Type          string                      `pulumi:"type"`
}
