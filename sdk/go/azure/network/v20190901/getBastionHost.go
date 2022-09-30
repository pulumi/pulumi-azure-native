


package v20190901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBastionHost(ctx *pulumi.Context, args *LookupBastionHostArgs, opts ...pulumi.InvokeOption) (*LookupBastionHostResult, error) {
	var rv LookupBastionHostResult
	err := ctx.Invoke("azure-native:network/v20190901:getBastionHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBastionHostArgs struct {
	BastionHostName   string `pulumi:"bastionHostName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBastionHostResult struct {
	DnsName           *string                              `pulumi:"dnsName"`
	Etag              string                               `pulumi:"etag"`
	Id                *string                              `pulumi:"id"`
	IpConfigurations  []BastionHostIPConfigurationResponse `pulumi:"ipConfigurations"`
	Location          *string                              `pulumi:"location"`
	Name              string                               `pulumi:"name"`
	ProvisioningState string                               `pulumi:"provisioningState"`
	Tags              map[string]string                    `pulumi:"tags"`
	Type              string                               `pulumi:"type"`
}

func LookupBastionHostOutput(ctx *pulumi.Context, args LookupBastionHostOutputArgs, opts ...pulumi.InvokeOption) LookupBastionHostResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBastionHostResult, error) {
			args := v.(LookupBastionHostArgs)
			r, err := LookupBastionHost(ctx, &args, opts...)
			var s LookupBastionHostResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBastionHostResultOutput)
}

type LookupBastionHostOutputArgs struct {
	BastionHostName   pulumi.StringInput `pulumi:"bastionHostName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBastionHostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBastionHostArgs)(nil)).Elem()
}


type LookupBastionHostResultOutput struct{ *pulumi.OutputState }

func (LookupBastionHostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBastionHostResult)(nil)).Elem()
}

func (o LookupBastionHostResultOutput) ToLookupBastionHostResultOutput() LookupBastionHostResultOutput {
	return o
}

func (o LookupBastionHostResultOutput) ToLookupBastionHostResultOutputWithContext(ctx context.Context) LookupBastionHostResultOutput {
	return o
}

func (o LookupBastionHostResultOutput) DnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBastionHostResult) *string { return v.DnsName }).(pulumi.StringPtrOutput)
}

func (o LookupBastionHostResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBastionHostResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupBastionHostResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBastionHostResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupBastionHostResultOutput) IpConfigurations() BastionHostIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupBastionHostResult) []BastionHostIPConfigurationResponse { return v.IpConfigurations }).(BastionHostIPConfigurationResponseArrayOutput)
}

func (o LookupBastionHostResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBastionHostResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupBastionHostResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBastionHostResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBastionHostResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBastionHostResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupBastionHostResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBastionHostResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupBastionHostResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBastionHostResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBastionHostResultOutput{})
}
