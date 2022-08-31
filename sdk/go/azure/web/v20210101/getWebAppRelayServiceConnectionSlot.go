


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppRelayServiceConnectionSlot(ctx *pulumi.Context, args *LookupWebAppRelayServiceConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppRelayServiceConnectionSlotResult, error) {
	var rv LookupWebAppRelayServiceConnectionSlotResult
	err := ctx.Invoke("azure-native:web/v20210101:getWebAppRelayServiceConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppRelayServiceConnectionSlotArgs struct {
	EntityName        string `pulumi:"entityName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupWebAppRelayServiceConnectionSlotResult struct {
	BiztalkUri               *string `pulumi:"biztalkUri"`
	EntityConnectionString   *string `pulumi:"entityConnectionString"`
	EntityName               *string `pulumi:"entityName"`
	Hostname                 *string `pulumi:"hostname"`
	Id                       string  `pulumi:"id"`
	Kind                     *string `pulumi:"kind"`
	Name                     string  `pulumi:"name"`
	Port                     *int    `pulumi:"port"`
	ResourceConnectionString *string `pulumi:"resourceConnectionString"`
	ResourceType             *string `pulumi:"resourceType"`
	Type                     string  `pulumi:"type"`
}

func LookupWebAppRelayServiceConnectionSlotOutput(ctx *pulumi.Context, args LookupWebAppRelayServiceConnectionSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppRelayServiceConnectionSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppRelayServiceConnectionSlotResult, error) {
			args := v.(LookupWebAppRelayServiceConnectionSlotArgs)
			r, err := LookupWebAppRelayServiceConnectionSlot(ctx, &args, opts...)
			var s LookupWebAppRelayServiceConnectionSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppRelayServiceConnectionSlotResultOutput)
}

type LookupWebAppRelayServiceConnectionSlotOutputArgs struct {
	EntityName        pulumi.StringInput `pulumi:"entityName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppRelayServiceConnectionSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppRelayServiceConnectionSlotArgs)(nil)).Elem()
}


type LookupWebAppRelayServiceConnectionSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppRelayServiceConnectionSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppRelayServiceConnectionSlotResult)(nil)).Elem()
}

func (o LookupWebAppRelayServiceConnectionSlotResultOutput) ToLookupWebAppRelayServiceConnectionSlotResultOutput() LookupWebAppRelayServiceConnectionSlotResultOutput {
	return o
}

func (o LookupWebAppRelayServiceConnectionSlotResultOutput) ToLookupWebAppRelayServiceConnectionSlotResultOutputWithContext(ctx context.Context) LookupWebAppRelayServiceConnectionSlotResultOutput {
	return o
}

func (o LookupWebAppRelayServiceConnectionSlotResultOutput) BiztalkUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionSlotResult) *string { return v.BiztalkUri }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionSlotResultOutput) EntityConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionSlotResult) *string { return v.EntityConnectionString }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionSlotResultOutput) EntityName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionSlotResult) *string { return v.EntityName }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionSlotResultOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionSlotResult) *string { return v.Hostname }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppRelayServiceConnectionSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppRelayServiceConnectionSlotResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionSlotResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionSlotResultOutput) ResourceConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionSlotResult) *string { return v.ResourceConnectionString }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionSlotResultOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionSlotResult) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppRelayServiceConnectionSlotResultOutput{})
}
