


package v20170101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebService(ctx *pulumi.Context, args *LookupWebServiceArgs, opts ...pulumi.InvokeOption) (*LookupWebServiceResult, error) {
	var rv LookupWebServiceResult
	err := ctx.Invoke("azure-native:machinelearning/v20170101:getWebService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebServiceArgs struct {
	Region            *string `pulumi:"region"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	WebServiceName    string  `pulumi:"webServiceName"`
}


type LookupWebServiceResult struct {
	Id         string                               `pulumi:"id"`
	Location   string                               `pulumi:"location"`
	Name       string                               `pulumi:"name"`
	Properties WebServicePropertiesForGraphResponse `pulumi:"properties"`
	Tags       map[string]string                    `pulumi:"tags"`
	Type       string                               `pulumi:"type"`
}


func (val *LookupWebServiceResult) Defaults() *LookupWebServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
