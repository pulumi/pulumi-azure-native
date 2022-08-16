


package v20160801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppRelayServiceConnection(ctx *pulumi.Context, args *LookupWebAppRelayServiceConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppRelayServiceConnectionResult, error) {
	var rv LookupWebAppRelayServiceConnectionResult
	err := ctx.Invoke("azure-native:web/v20160801:getWebAppRelayServiceConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppRelayServiceConnectionArgs struct {
	EntityName        string `pulumi:"entityName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppRelayServiceConnectionResult struct {
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

func LookupWebAppRelayServiceConnectionOutput(ctx *pulumi.Context, args LookupWebAppRelayServiceConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppRelayServiceConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppRelayServiceConnectionResult, error) {
			args := v.(LookupWebAppRelayServiceConnectionArgs)
			r, err := LookupWebAppRelayServiceConnection(ctx, &args, opts...)
			var s LookupWebAppRelayServiceConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppRelayServiceConnectionResultOutput)
}

type LookupWebAppRelayServiceConnectionOutputArgs struct {
	EntityName        pulumi.StringInput `pulumi:"entityName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppRelayServiceConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppRelayServiceConnectionArgs)(nil)).Elem()
}


type LookupWebAppRelayServiceConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppRelayServiceConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppRelayServiceConnectionResult)(nil)).Elem()
}

func (o LookupWebAppRelayServiceConnectionResultOutput) ToLookupWebAppRelayServiceConnectionResultOutput() LookupWebAppRelayServiceConnectionResultOutput {
	return o
}

func (o LookupWebAppRelayServiceConnectionResultOutput) ToLookupWebAppRelayServiceConnectionResultOutputWithContext(ctx context.Context) LookupWebAppRelayServiceConnectionResultOutput {
	return o
}

func (o LookupWebAppRelayServiceConnectionResultOutput) BiztalkUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) *string { return v.BiztalkUri }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionResultOutput) EntityConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) *string { return v.EntityConnectionString }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionResultOutput) EntityName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) *string { return v.EntityName }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionResultOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) *string { return v.Hostname }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppRelayServiceConnectionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppRelayServiceConnectionResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionResultOutput) ResourceConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) *string { return v.ResourceConnectionString }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionResultOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppRelayServiceConnectionResultOutput{})
}
