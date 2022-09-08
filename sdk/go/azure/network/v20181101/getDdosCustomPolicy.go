


package v20181101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDdosCustomPolicy(ctx *pulumi.Context, args *LookupDdosCustomPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDdosCustomPolicyResult, error) {
	var rv LookupDdosCustomPolicyResult
	err := ctx.Invoke("azure-native:network/v20181101:getDdosCustomPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDdosCustomPolicyArgs struct {
	DdosCustomPolicyName string `pulumi:"ddosCustomPolicyName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupDdosCustomPolicyResult struct {
	Etag                   string                                 `pulumi:"etag"`
	Id                     *string                                `pulumi:"id"`
	Location               *string                                `pulumi:"location"`
	Name                   string                                 `pulumi:"name"`
	ProtocolCustomSettings []ProtocolCustomSettingsFormatResponse `pulumi:"protocolCustomSettings"`
	ProvisioningState      string                                 `pulumi:"provisioningState"`
	PublicIPAddresses      []SubResourceResponse                  `pulumi:"publicIPAddresses"`
	ResourceGuid           string                                 `pulumi:"resourceGuid"`
	Tags                   map[string]string                      `pulumi:"tags"`
	Type                   string                                 `pulumi:"type"`
}

func LookupDdosCustomPolicyOutput(ctx *pulumi.Context, args LookupDdosCustomPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupDdosCustomPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDdosCustomPolicyResult, error) {
			args := v.(LookupDdosCustomPolicyArgs)
			r, err := LookupDdosCustomPolicy(ctx, &args, opts...)
			var s LookupDdosCustomPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDdosCustomPolicyResultOutput)
}

type LookupDdosCustomPolicyOutputArgs struct {
	DdosCustomPolicyName pulumi.StringInput `pulumi:"ddosCustomPolicyName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDdosCustomPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDdosCustomPolicyArgs)(nil)).Elem()
}


type LookupDdosCustomPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupDdosCustomPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDdosCustomPolicyResult)(nil)).Elem()
}

func (o LookupDdosCustomPolicyResultOutput) ToLookupDdosCustomPolicyResultOutput() LookupDdosCustomPolicyResultOutput {
	return o
}

func (o LookupDdosCustomPolicyResultOutput) ToLookupDdosCustomPolicyResultOutputWithContext(ctx context.Context) LookupDdosCustomPolicyResultOutput {
	return o
}

func (o LookupDdosCustomPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosCustomPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDdosCustomPolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDdosCustomPolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupDdosCustomPolicyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDdosCustomPolicyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDdosCustomPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosCustomPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDdosCustomPolicyResultOutput) ProtocolCustomSettings() ProtocolCustomSettingsFormatResponseArrayOutput {
	return o.ApplyT(func(v LookupDdosCustomPolicyResult) []ProtocolCustomSettingsFormatResponse {
		return v.ProtocolCustomSettings
	}).(ProtocolCustomSettingsFormatResponseArrayOutput)
}

func (o LookupDdosCustomPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosCustomPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDdosCustomPolicyResultOutput) PublicIPAddresses() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupDdosCustomPolicyResult) []SubResourceResponse { return v.PublicIPAddresses }).(SubResourceResponseArrayOutput)
}

func (o LookupDdosCustomPolicyResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosCustomPolicyResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupDdosCustomPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDdosCustomPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDdosCustomPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosCustomPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDdosCustomPolicyResultOutput{})
}
