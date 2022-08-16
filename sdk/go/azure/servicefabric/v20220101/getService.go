


package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure-native:servicefabric/v20220101:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	ApplicationName   string `pulumi:"applicationName"`
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupServiceResult struct {
	Id         string             `pulumi:"id"`
	Location   *string            `pulumi:"location"`
	Name       string             `pulumi:"name"`
	Properties interface{}        `pulumi:"properties"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Tags       map[string]string  `pulumi:"tags"`
	Type       string             `pulumi:"type"`
}

func LookupServiceOutput(ctx *pulumi.Context, args LookupServiceOutputArgs, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceResult, error) {
			args := v.(LookupServiceArgs)
			r, err := LookupService(ctx, &args, opts...)
			var s LookupServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceResultOutput)
}

type LookupServiceOutputArgs struct {
	ApplicationName   pulumi.StringInput `pulumi:"applicationName"`
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}


type LookupServiceResultOutput struct{ *pulumi.OutputState }

func (LookupServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceResult)(nil)).Elem()
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutput() LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutputWithContext(ctx context.Context) LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupServiceResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}
