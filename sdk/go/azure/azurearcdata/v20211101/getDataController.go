


package v20211101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataController(ctx *pulumi.Context, args *LookupDataControllerArgs, opts ...pulumi.InvokeOption) (*LookupDataControllerResult, error) {
	var rv LookupDataControllerResult
	err := ctx.Invoke("azure-native:azurearcdata/v20211101:getDataController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
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


func (val *LookupDataControllerResult) Defaults() *LookupDataControllerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
