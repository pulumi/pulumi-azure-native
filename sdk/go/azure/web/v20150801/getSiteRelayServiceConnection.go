


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSiteRelayServiceConnection(ctx *pulumi.Context, args *LookupSiteRelayServiceConnectionArgs, opts ...pulumi.InvokeOption) (*LookupSiteRelayServiceConnectionResult, error) {
	var rv LookupSiteRelayServiceConnectionResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteRelayServiceConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteRelayServiceConnectionArgs struct {
	EntityName        string `pulumi:"entityName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSiteRelayServiceConnectionResult struct {
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

func LookupSiteRelayServiceConnectionOutput(ctx *pulumi.Context, args LookupSiteRelayServiceConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupSiteRelayServiceConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteRelayServiceConnectionResult, error) {
			args := v.(LookupSiteRelayServiceConnectionArgs)
			r, err := LookupSiteRelayServiceConnection(ctx, &args, opts...)
			var s LookupSiteRelayServiceConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteRelayServiceConnectionResultOutput)
}

type LookupSiteRelayServiceConnectionOutputArgs struct {
	EntityName        pulumi.StringInput `pulumi:"entityName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSiteRelayServiceConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteRelayServiceConnectionArgs)(nil)).Elem()
}


type LookupSiteRelayServiceConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupSiteRelayServiceConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteRelayServiceConnectionResult)(nil)).Elem()
}

func (o LookupSiteRelayServiceConnectionResultOutput) ToLookupSiteRelayServiceConnectionResultOutput() LookupSiteRelayServiceConnectionResultOutput {
	return o
}

func (o LookupSiteRelayServiceConnectionResultOutput) ToLookupSiteRelayServiceConnectionResultOutputWithContext(ctx context.Context) LookupSiteRelayServiceConnectionResultOutput {
	return o
}

func (o LookupSiteRelayServiceConnectionResultOutput) BiztalkUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionResult) *string { return v.BiztalkUri }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionResultOutput) EntityConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionResult) *string { return v.EntityConnectionString }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionResultOutput) EntityName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionResult) *string { return v.EntityName }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionResultOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionResult) *string { return v.Hostname }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteRelayServiceConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o LookupSiteRelayServiceConnectionResultOutput) ResourceConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionResult) *string { return v.ResourceConnectionString }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionResultOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionResult) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o LookupSiteRelayServiceConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteRelayServiceConnectionResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteRelayServiceConnectionResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteRelayServiceConnectionResultOutput{})
}
