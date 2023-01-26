


package storagemover

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEndpoint(ctx *pulumi.Context, args *LookupEndpointArgs, opts ...pulumi.InvokeOption) (*LookupEndpointResult, error) {
	var rv LookupEndpointResult
	err := ctx.Invoke("azure-native:storagemover:getEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StorageMoverName  string `pulumi:"storageMoverName"`
}


type LookupEndpointResult struct {
	Id         string             `pulumi:"id"`
	Name       string             `pulumi:"name"`
	Properties interface{}        `pulumi:"properties"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupEndpointOutput(ctx *pulumi.Context, args LookupEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointResult, error) {
			args := v.(LookupEndpointArgs)
			r, err := LookupEndpoint(ctx, &args, opts...)
			var s LookupEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEndpointResultOutput)
}

type LookupEndpointOutputArgs struct {
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageMoverName  pulumi.StringInput `pulumi:"storageMoverName"`
}

func (LookupEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointArgs)(nil)).Elem()
}


type LookupEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointResult)(nil)).Elem()
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutput() LookupEndpointResultOutput {
	return o
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutputWithContext(ctx context.Context) LookupEndpointResultOutput {
	return o
}

func (o LookupEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupEndpointResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointResultOutput{})
}
