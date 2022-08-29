


package powerplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnterprisePolicy(ctx *pulumi.Context, args *LookupEnterprisePolicyArgs, opts ...pulumi.InvokeOption) (*LookupEnterprisePolicyResult, error) {
	var rv LookupEnterprisePolicyResult
	err := ctx.Invoke("azure-native:powerplatform:getEnterprisePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnterprisePolicyArgs struct {
	EnterprisePolicyName string `pulumi:"enterprisePolicyName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupEnterprisePolicyResult struct {
	Encryption       *PropertiesResponseEncryption       `pulumi:"encryption"`
	Id               string                              `pulumi:"id"`
	Identity         *EnterprisePolicyIdentityResponse   `pulumi:"identity"`
	Kind             string                              `pulumi:"kind"`
	Location         string                              `pulumi:"location"`
	Lockbox          *PropertiesResponseLockbox          `pulumi:"lockbox"`
	Name             string                              `pulumi:"name"`
	NetworkInjection *PropertiesResponseNetworkInjection `pulumi:"networkInjection"`
	SystemData       SystemDataResponse                  `pulumi:"systemData"`
	SystemId         string                              `pulumi:"systemId"`
	Tags             map[string]string                   `pulumi:"tags"`
	Type             string                              `pulumi:"type"`
}

func LookupEnterprisePolicyOutput(ctx *pulumi.Context, args LookupEnterprisePolicyOutputArgs, opts ...pulumi.InvokeOption) LookupEnterprisePolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnterprisePolicyResult, error) {
			args := v.(LookupEnterprisePolicyArgs)
			r, err := LookupEnterprisePolicy(ctx, &args, opts...)
			var s LookupEnterprisePolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnterprisePolicyResultOutput)
}

type LookupEnterprisePolicyOutputArgs struct {
	EnterprisePolicyName pulumi.StringInput `pulumi:"enterprisePolicyName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEnterprisePolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnterprisePolicyArgs)(nil)).Elem()
}


type LookupEnterprisePolicyResultOutput struct{ *pulumi.OutputState }

func (LookupEnterprisePolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnterprisePolicyResult)(nil)).Elem()
}

func (o LookupEnterprisePolicyResultOutput) ToLookupEnterprisePolicyResultOutput() LookupEnterprisePolicyResultOutput {
	return o
}

func (o LookupEnterprisePolicyResultOutput) ToLookupEnterprisePolicyResultOutputWithContext(ctx context.Context) LookupEnterprisePolicyResultOutput {
	return o
}

func (o LookupEnterprisePolicyResultOutput) Encryption() PropertiesResponseEncryptionPtrOutput {
	return o.ApplyT(func(v LookupEnterprisePolicyResult) *PropertiesResponseEncryption { return v.Encryption }).(PropertiesResponseEncryptionPtrOutput)
}

func (o LookupEnterprisePolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterprisePolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEnterprisePolicyResultOutput) Identity() EnterprisePolicyIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupEnterprisePolicyResult) *EnterprisePolicyIdentityResponse { return v.Identity }).(EnterprisePolicyIdentityResponsePtrOutput)
}

func (o LookupEnterprisePolicyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterprisePolicyResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupEnterprisePolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterprisePolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupEnterprisePolicyResultOutput) Lockbox() PropertiesResponseLockboxPtrOutput {
	return o.ApplyT(func(v LookupEnterprisePolicyResult) *PropertiesResponseLockbox { return v.Lockbox }).(PropertiesResponseLockboxPtrOutput)
}

func (o LookupEnterprisePolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterprisePolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnterprisePolicyResultOutput) NetworkInjection() PropertiesResponseNetworkInjectionPtrOutput {
	return o.ApplyT(func(v LookupEnterprisePolicyResult) *PropertiesResponseNetworkInjection { return v.NetworkInjection }).(PropertiesResponseNetworkInjectionPtrOutput)
}

func (o LookupEnterprisePolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEnterprisePolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupEnterprisePolicyResultOutput) SystemId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterprisePolicyResult) string { return v.SystemId }).(pulumi.StringOutput)
}

func (o LookupEnterprisePolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnterprisePolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupEnterprisePolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterprisePolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnterprisePolicyResultOutput{})
}
