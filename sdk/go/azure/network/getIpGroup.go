


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIpGroup(ctx *pulumi.Context, args *LookupIpGroupArgs, opts ...pulumi.InvokeOption) (*LookupIpGroupResult, error) {
	var rv LookupIpGroupResult
	err := ctx.Invoke("azure-native:network:getIpGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIpGroupArgs struct {
	Expand            *string `pulumi:"expand"`
	IpGroupsName      string  `pulumi:"ipGroupsName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupIpGroupResult struct {
	Etag              string                `pulumi:"etag"`
	FirewallPolicies  []SubResourceResponse `pulumi:"firewallPolicies"`
	Firewalls         []SubResourceResponse `pulumi:"firewalls"`
	Id                *string               `pulumi:"id"`
	IpAddresses       []string              `pulumi:"ipAddresses"`
	Location          *string               `pulumi:"location"`
	Name              string                `pulumi:"name"`
	ProvisioningState string                `pulumi:"provisioningState"`
	Tags              map[string]string     `pulumi:"tags"`
	Type              string                `pulumi:"type"`
}

func LookupIpGroupOutput(ctx *pulumi.Context, args LookupIpGroupOutputArgs, opts ...pulumi.InvokeOption) LookupIpGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpGroupResult, error) {
			args := v.(LookupIpGroupArgs)
			r, err := LookupIpGroup(ctx, &args, opts...)
			var s LookupIpGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpGroupResultOutput)
}

type LookupIpGroupOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	IpGroupsName      pulumi.StringInput    `pulumi:"ipGroupsName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupIpGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpGroupArgs)(nil)).Elem()
}


type LookupIpGroupResultOutput struct{ *pulumi.OutputState }

func (LookupIpGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpGroupResult)(nil)).Elem()
}

func (o LookupIpGroupResultOutput) ToLookupIpGroupResultOutput() LookupIpGroupResultOutput {
	return o
}

func (o LookupIpGroupResultOutput) ToLookupIpGroupResultOutputWithContext(ctx context.Context) LookupIpGroupResultOutput {
	return o
}

func (o LookupIpGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupIpGroupResultOutput) FirewallPolicies() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupIpGroupResult) []SubResourceResponse { return v.FirewallPolicies }).(SubResourceResponseArrayOutput)
}

func (o LookupIpGroupResultOutput) Firewalls() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupIpGroupResult) []SubResourceResponse { return v.Firewalls }).(SubResourceResponseArrayOutput)
}

func (o LookupIpGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupIpGroupResultOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIpGroupResult) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupIpGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIpGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIpGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupIpGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIpGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIpGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpGroupResultOutput{})
}
