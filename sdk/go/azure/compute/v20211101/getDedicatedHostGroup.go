


package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDedicatedHostGroup(ctx *pulumi.Context, args *LookupDedicatedHostGroupArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedHostGroupResult, error) {
	var rv LookupDedicatedHostGroupResult
	err := ctx.Invoke("azure-native:compute/v20211101:getDedicatedHostGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDedicatedHostGroupArgs struct {
	Expand            *string `pulumi:"expand"`
	HostGroupName     string  `pulumi:"hostGroupName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupDedicatedHostGroupResult struct {
	Hosts                     []SubResourceReadOnlyResponse          `pulumi:"hosts"`
	Id                        string                                 `pulumi:"id"`
	InstanceView              DedicatedHostGroupInstanceViewResponse `pulumi:"instanceView"`
	Location                  string                                 `pulumi:"location"`
	Name                      string                                 `pulumi:"name"`
	PlatformFaultDomainCount  int                                    `pulumi:"platformFaultDomainCount"`
	SupportAutomaticPlacement *bool                                  `pulumi:"supportAutomaticPlacement"`
	Tags                      map[string]string                      `pulumi:"tags"`
	Type                      string                                 `pulumi:"type"`
	Zones                     []string                               `pulumi:"zones"`
}

func LookupDedicatedHostGroupOutput(ctx *pulumi.Context, args LookupDedicatedHostGroupOutputArgs, opts ...pulumi.InvokeOption) LookupDedicatedHostGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDedicatedHostGroupResult, error) {
			args := v.(LookupDedicatedHostGroupArgs)
			r, err := LookupDedicatedHostGroup(ctx, &args, opts...)
			var s LookupDedicatedHostGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDedicatedHostGroupResultOutput)
}

type LookupDedicatedHostGroupOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	HostGroupName     pulumi.StringInput    `pulumi:"hostGroupName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupDedicatedHostGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedHostGroupArgs)(nil)).Elem()
}


type LookupDedicatedHostGroupResultOutput struct{ *pulumi.OutputState }

func (LookupDedicatedHostGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedHostGroupResult)(nil)).Elem()
}

func (o LookupDedicatedHostGroupResultOutput) ToLookupDedicatedHostGroupResultOutput() LookupDedicatedHostGroupResultOutput {
	return o
}

func (o LookupDedicatedHostGroupResultOutput) ToLookupDedicatedHostGroupResultOutputWithContext(ctx context.Context) LookupDedicatedHostGroupResultOutput {
	return o
}

func (o LookupDedicatedHostGroupResultOutput) Hosts() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) []SubResourceReadOnlyResponse { return v.Hosts }).(SubResourceReadOnlyResponseArrayOutput)
}

func (o LookupDedicatedHostGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDedicatedHostGroupResultOutput) InstanceView() DedicatedHostGroupInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) DedicatedHostGroupInstanceViewResponse { return v.InstanceView }).(DedicatedHostGroupInstanceViewResponseOutput)
}

func (o LookupDedicatedHostGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDedicatedHostGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDedicatedHostGroupResultOutput) PlatformFaultDomainCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) int { return v.PlatformFaultDomainCount }).(pulumi.IntOutput)
}

func (o LookupDedicatedHostGroupResultOutput) SupportAutomaticPlacement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) *bool { return v.SupportAutomaticPlacement }).(pulumi.BoolPtrOutput)
}

func (o LookupDedicatedHostGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDedicatedHostGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDedicatedHostGroupResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDedicatedHostGroupResultOutput{})
}
