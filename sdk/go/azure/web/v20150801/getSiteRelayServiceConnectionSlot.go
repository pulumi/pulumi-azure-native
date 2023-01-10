


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSiteRelayServiceConnectionSlot(ctx *pulumi.Context, args *LookupSiteRelayServiceConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupSiteRelayServiceConnectionSlotResult, error) {
	var rv LookupSiteRelayServiceConnectionSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteRelayServiceConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteRelayServiceConnectionSlotArgs struct {
	EntityName        string `pulumi:"entityName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupSiteRelayServiceConnectionSlotResult struct {
	BiztalkUri               *string           `pulumi:"biztalkUri"`
	EntityConnectionString   *string           `pulumi:"entityConnectionString"`
	EntityName               *string           `pulumi:"entityName"`
	Hostname                 *string           `pulumi:"hostname"`
	Id                       *string           `pulumi:"id"`
	Kind                     *string           `pulumi:"kind"`
	Location                 string            `pulumi:"location"`
	Name                     *string           `pulumi:"name"`
	Port                     *int              `pulumi:"port"`
	ResourceConnectionString *string           `pulumi:"resourceConnectionString"`
	ResourceType             *string           `pulumi:"resourceType"`
	Tags                     map[string]string `pulumi:"tags"`
	Type                     *string           `pulumi:"type"`
}

func LookupSiteRelayServiceConnectionSlotOutput(ctx *pulumi.Context, args LookupSiteRelayServiceConnectionSlotOutputArgs, opts ...pulumi.InvokeOption) LookupSiteRelayServiceConnectionSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteRelayServiceConnectionSlotResult, error) {
			args := v.(LookupSiteRelayServiceConnectionSlotArgs)
			r, err := LookupSiteRelayServiceConnectionSlot(ctx, &args, opts...)
			var s LookupSiteRelayServiceConnectionSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteRelayServiceConnectionSlotResultOutput)
}

type LookupSiteRelayServiceConnectionSlotOutputArgs struct {
	EntityName        pulumi.StringInput `pulumi:"entityName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (LookupSiteRelayServiceConnectionSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteRelayServiceConnectionSlotArgs)(nil)).Elem()
}


type LookupSiteRelayServiceConnectionSlotResultOutput struct{ *pulumi.OutputState }

func (LookupSiteRelayServiceConnectionSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteRelayServiceConnectionSlotResult)(nil)).Elem()
}

func (o LookupSiteRelayServiceConnectionSlotResultOutput) ToLookupSiteRelayServiceConnectionSlotResultOutput() LookupSiteRelayServiceConnectionSlotResultOutput {
	return o
}

func (o LookupSiteRelayServiceConnectionSlotResultOutput) ToLookupSiteRelayServiceConnectionSlotResultOutputWithContext(ctx context.Context) LookupSiteRelayServiceConnectionSlotResultOutput {
	return o
}

func (o LookupSiteRelayServiceConnectionSlotResultOutput) BiztalkUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionSlotResult) *string { return v.BiztalkUri }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionSlotResultOutput) EntityConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionSlotResult) *string { return v.EntityConnectionString }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionSlotResultOutput) EntityName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionSlotResult) *string { return v.EntityName }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionSlotResultOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionSlotResult) *string { return v.Hostname }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionSlotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionSlotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteRelayServiceConnectionSlotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionSlotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionSlotResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionSlotResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o LookupSiteRelayServiceConnectionSlotResultOutput) ResourceConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionSlotResult) *string { return v.ResourceConnectionString }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionSlotResultOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionSlotResult) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteRelayServiceConnectionSlotResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionSlotResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteRelayServiceConnectionSlotResultOutput{})
}
