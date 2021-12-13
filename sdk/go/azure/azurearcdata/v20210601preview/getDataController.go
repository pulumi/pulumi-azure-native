


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataController(ctx *pulumi.Context, args *LookupDataControllerArgs, opts ...pulumi.InvokeOption) (*LookupDataControllerResult, error) {
	var rv LookupDataControllerResult
	err := ctx.Invoke("azure-native:azurearcdata/v20210601preview:getDataController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataControllerArgs struct {
	DataControllerName string `pulumi:"dataControllerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupDataControllerResult struct {
	ExtendedLocation *ExtendedLocationResponse        `pulumi:"extendedLocation"`
	Id               string                           `pulumi:"id"`
	Location         string                           `pulumi:"location"`
	Name             string                           `pulumi:"name"`
	Properties       DataControllerPropertiesResponse `pulumi:"properties"`
	SystemData       SystemDataResponse               `pulumi:"systemData"`
	Tags             map[string]string                `pulumi:"tags"`
	Type             string                           `pulumi:"type"`
}
