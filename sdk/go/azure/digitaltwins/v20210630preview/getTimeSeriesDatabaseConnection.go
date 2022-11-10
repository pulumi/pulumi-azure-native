


package v20210630preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTimeSeriesDatabaseConnection(ctx *pulumi.Context, args *LookupTimeSeriesDatabaseConnectionArgs, opts ...pulumi.InvokeOption) (*LookupTimeSeriesDatabaseConnectionResult, error) {
	var rv LookupTimeSeriesDatabaseConnectionResult
	err := ctx.Invoke("azure-native:digitaltwins/v20210630preview:getTimeSeriesDatabaseConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupTimeSeriesDatabaseConnectionArgs struct {
	ResourceGroupName                string `pulumi:"resourceGroupName"`
	ResourceName                     string `pulumi:"resourceName"`
	TimeSeriesDatabaseConnectionName string `pulumi:"timeSeriesDatabaseConnectionName"`
}


type LookupTimeSeriesDatabaseConnectionResult struct {
	Id         string                                        `pulumi:"id"`
	Name       string                                        `pulumi:"name"`
	Properties AzureDataExplorerConnectionPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                            `pulumi:"systemData"`
	Type       string                                        `pulumi:"type"`
}


func (val *LookupTimeSeriesDatabaseConnectionResult) Defaults() *LookupTimeSeriesDatabaseConnectionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupTimeSeriesDatabaseConnectionOutput(ctx *pulumi.Context, args LookupTimeSeriesDatabaseConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupTimeSeriesDatabaseConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTimeSeriesDatabaseConnectionResult, error) {
			args := v.(LookupTimeSeriesDatabaseConnectionArgs)
			r, err := LookupTimeSeriesDatabaseConnection(ctx, &args, opts...)
			var s LookupTimeSeriesDatabaseConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTimeSeriesDatabaseConnectionResultOutput)
}

type LookupTimeSeriesDatabaseConnectionOutputArgs struct {
	ResourceGroupName                pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                     pulumi.StringInput `pulumi:"resourceName"`
	TimeSeriesDatabaseConnectionName pulumi.StringInput `pulumi:"timeSeriesDatabaseConnectionName"`
}

func (LookupTimeSeriesDatabaseConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTimeSeriesDatabaseConnectionArgs)(nil)).Elem()
}


type LookupTimeSeriesDatabaseConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupTimeSeriesDatabaseConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTimeSeriesDatabaseConnectionResult)(nil)).Elem()
}

func (o LookupTimeSeriesDatabaseConnectionResultOutput) ToLookupTimeSeriesDatabaseConnectionResultOutput() LookupTimeSeriesDatabaseConnectionResultOutput {
	return o
}

func (o LookupTimeSeriesDatabaseConnectionResultOutput) ToLookupTimeSeriesDatabaseConnectionResultOutputWithContext(ctx context.Context) LookupTimeSeriesDatabaseConnectionResultOutput {
	return o
}

func (o LookupTimeSeriesDatabaseConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeSeriesDatabaseConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTimeSeriesDatabaseConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeSeriesDatabaseConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTimeSeriesDatabaseConnectionResultOutput) Properties() AzureDataExplorerConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupTimeSeriesDatabaseConnectionResult) AzureDataExplorerConnectionPropertiesResponse {
		return v.Properties
	}).(AzureDataExplorerConnectionPropertiesResponseOutput)
}

func (o LookupTimeSeriesDatabaseConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTimeSeriesDatabaseConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTimeSeriesDatabaseConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeSeriesDatabaseConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTimeSeriesDatabaseConnectionResultOutput{})
}
