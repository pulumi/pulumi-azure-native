


package v20220615preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupActiveDirectoryConnector(ctx *pulumi.Context, args *LookupActiveDirectoryConnectorArgs, opts ...pulumi.InvokeOption) (*LookupActiveDirectoryConnectorResult, error) {
	var rv LookupActiveDirectoryConnectorResult
	err := ctx.Invoke("azure-native:azurearcdata/v20220615preview:getActiveDirectoryConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupActiveDirectoryConnectorArgs struct {
	ActiveDirectoryConnectorName string `pulumi:"activeDirectoryConnectorName"`
	DataControllerName           string `pulumi:"dataControllerName"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
}


type LookupActiveDirectoryConnectorResult struct {
	Id         string                                     `pulumi:"id"`
	Name       string                                     `pulumi:"name"`
	Properties ActiveDirectoryConnectorPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                         `pulumi:"systemData"`
	Type       string                                     `pulumi:"type"`
}


func (val *LookupActiveDirectoryConnectorResult) Defaults() *LookupActiveDirectoryConnectorResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupActiveDirectoryConnectorOutput(ctx *pulumi.Context, args LookupActiveDirectoryConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupActiveDirectoryConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupActiveDirectoryConnectorResult, error) {
			args := v.(LookupActiveDirectoryConnectorArgs)
			r, err := LookupActiveDirectoryConnector(ctx, &args, opts...)
			var s LookupActiveDirectoryConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupActiveDirectoryConnectorResultOutput)
}

type LookupActiveDirectoryConnectorOutputArgs struct {
	ActiveDirectoryConnectorName pulumi.StringInput `pulumi:"activeDirectoryConnectorName"`
	DataControllerName           pulumi.StringInput `pulumi:"dataControllerName"`
	ResourceGroupName            pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupActiveDirectoryConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActiveDirectoryConnectorArgs)(nil)).Elem()
}


type LookupActiveDirectoryConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupActiveDirectoryConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActiveDirectoryConnectorResult)(nil)).Elem()
}

func (o LookupActiveDirectoryConnectorResultOutput) ToLookupActiveDirectoryConnectorResultOutput() LookupActiveDirectoryConnectorResultOutput {
	return o
}

func (o LookupActiveDirectoryConnectorResultOutput) ToLookupActiveDirectoryConnectorResultOutputWithContext(ctx context.Context) LookupActiveDirectoryConnectorResultOutput {
	return o
}

func (o LookupActiveDirectoryConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActiveDirectoryConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupActiveDirectoryConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActiveDirectoryConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupActiveDirectoryConnectorResultOutput) Properties() ActiveDirectoryConnectorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupActiveDirectoryConnectorResult) ActiveDirectoryConnectorPropertiesResponse {
		return v.Properties
	}).(ActiveDirectoryConnectorPropertiesResponseOutput)
}

func (o LookupActiveDirectoryConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupActiveDirectoryConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupActiveDirectoryConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActiveDirectoryConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupActiveDirectoryConnectorResultOutput{})
}
