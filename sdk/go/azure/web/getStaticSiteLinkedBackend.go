


package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStaticSiteLinkedBackend(ctx *pulumi.Context, args *LookupStaticSiteLinkedBackendArgs, opts ...pulumi.InvokeOption) (*LookupStaticSiteLinkedBackendResult, error) {
	var rv LookupStaticSiteLinkedBackendResult
	err := ctx.Invoke("azure-native:web:getStaticSiteLinkedBackend", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticSiteLinkedBackendArgs struct {
	LinkedBackendName string `pulumi:"linkedBackendName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupStaticSiteLinkedBackendResult struct {
	BackendResourceId *string `pulumi:"backendResourceId"`
	CreatedOn         string  `pulumi:"createdOn"`
	Id                string  `pulumi:"id"`
	Kind              *string `pulumi:"kind"`
	Name              string  `pulumi:"name"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Region            *string `pulumi:"region"`
	Type              string  `pulumi:"type"`
}

func LookupStaticSiteLinkedBackendOutput(ctx *pulumi.Context, args LookupStaticSiteLinkedBackendOutputArgs, opts ...pulumi.InvokeOption) LookupStaticSiteLinkedBackendResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticSiteLinkedBackendResult, error) {
			args := v.(LookupStaticSiteLinkedBackendArgs)
			r, err := LookupStaticSiteLinkedBackend(ctx, &args, opts...)
			var s LookupStaticSiteLinkedBackendResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticSiteLinkedBackendResultOutput)
}

type LookupStaticSiteLinkedBackendOutputArgs struct {
	LinkedBackendName pulumi.StringInput `pulumi:"linkedBackendName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStaticSiteLinkedBackendOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteLinkedBackendArgs)(nil)).Elem()
}


type LookupStaticSiteLinkedBackendResultOutput struct{ *pulumi.OutputState }

func (LookupStaticSiteLinkedBackendResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteLinkedBackendResult)(nil)).Elem()
}

func (o LookupStaticSiteLinkedBackendResultOutput) ToLookupStaticSiteLinkedBackendResultOutput() LookupStaticSiteLinkedBackendResultOutput {
	return o
}

func (o LookupStaticSiteLinkedBackendResultOutput) ToLookupStaticSiteLinkedBackendResultOutputWithContext(ctx context.Context) LookupStaticSiteLinkedBackendResultOutput {
	return o
}

func (o LookupStaticSiteLinkedBackendResultOutput) BackendResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteLinkedBackendResult) *string { return v.BackendResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteLinkedBackendResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteLinkedBackendResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupStaticSiteLinkedBackendResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteLinkedBackendResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStaticSiteLinkedBackendResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteLinkedBackendResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteLinkedBackendResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteLinkedBackendResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStaticSiteLinkedBackendResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteLinkedBackendResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupStaticSiteLinkedBackendResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteLinkedBackendResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteLinkedBackendResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteLinkedBackendResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticSiteLinkedBackendResultOutput{})
}
