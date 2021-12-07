


package v20211201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLoadTest(ctx *pulumi.Context, args *LookupLoadTestArgs, opts ...pulumi.InvokeOption) (*LookupLoadTestResult, error) {
	var rv LookupLoadTestResult
	err := ctx.Invoke("azure-native:loadtestservice/v20211201preview:getLoadTest", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoadTestArgs struct {
	LoadTestName      string `pulumi:"loadTestName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLoadTestResult struct {
	DataPlaneURI      string                                 `pulumi:"dataPlaneURI"`
	Description       *string                                `pulumi:"description"`
	Id                string                                 `pulumi:"id"`
	Identity          *SystemAssignedServiceIdentityResponse `pulumi:"identity"`
	Location          string                                 `pulumi:"location"`
	Name              string                                 `pulumi:"name"`
	ProvisioningState string                                 `pulumi:"provisioningState"`
	SystemData        SystemDataResponse                     `pulumi:"systemData"`
	Tags              map[string]string                      `pulumi:"tags"`
	Type              string                                 `pulumi:"type"`
}
