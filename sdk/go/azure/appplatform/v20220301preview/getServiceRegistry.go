


package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceRegistry(ctx *pulumi.Context, args *LookupServiceRegistryArgs, opts ...pulumi.InvokeOption) (*LookupServiceRegistryResult, error) {
	var rv LookupServiceRegistryResult
	err := ctx.Invoke("azure-native:appplatform/v20220301preview:getServiceRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceRegistryArgs struct {
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	ServiceName         string `pulumi:"serviceName"`
	ServiceRegistryName string `pulumi:"serviceRegistryName"`
}


type LookupServiceRegistryResult struct {
	Id         string                            `pulumi:"id"`
	Name       string                            `pulumi:"name"`
	Properties ServiceRegistryPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                `pulumi:"systemData"`
	Type       string                            `pulumi:"type"`
}

func LookupServiceRegistryOutput(ctx *pulumi.Context, args LookupServiceRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupServiceRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceRegistryResult, error) {
			args := v.(LookupServiceRegistryArgs)
			r, err := LookupServiceRegistry(ctx, &args, opts...)
			var s LookupServiceRegistryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceRegistryResultOutput)
}

type LookupServiceRegistryOutputArgs struct {
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName         pulumi.StringInput `pulumi:"serviceName"`
	ServiceRegistryName pulumi.StringInput `pulumi:"serviceRegistryName"`
}

func (LookupServiceRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceRegistryArgs)(nil)).Elem()
}


type LookupServiceRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupServiceRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceRegistryResult)(nil)).Elem()
}

func (o LookupServiceRegistryResultOutput) ToLookupServiceRegistryResultOutput() LookupServiceRegistryResultOutput {
	return o
}

func (o LookupServiceRegistryResultOutput) ToLookupServiceRegistryResultOutputWithContext(ctx context.Context) LookupServiceRegistryResultOutput {
	return o
}

func (o LookupServiceRegistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceRegistryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceRegistryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceRegistryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceRegistryResultOutput) Properties() ServiceRegistryPropertiesResponseOutput {
	return o.ApplyT(func(v LookupServiceRegistryResult) ServiceRegistryPropertiesResponse { return v.Properties }).(ServiceRegistryPropertiesResponseOutput)
}

func (o LookupServiceRegistryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupServiceRegistryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupServiceRegistryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceRegistryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceRegistryResultOutput{})
}
