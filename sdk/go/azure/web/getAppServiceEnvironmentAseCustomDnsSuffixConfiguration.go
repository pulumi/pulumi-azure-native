


package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppServiceEnvironmentAseCustomDnsSuffixConfiguration(ctx *pulumi.Context, args *LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResult, error) {
	var rv LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResult
	err := ctx.Invoke("azure-native:web:getAppServiceEnvironmentAseCustomDnsSuffixConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResult struct {
	CertificateUrl            *string `pulumi:"certificateUrl"`
	DnsSuffix                 *string `pulumi:"dnsSuffix"`
	Id                        string  `pulumi:"id"`
	KeyVaultReferenceIdentity *string `pulumi:"keyVaultReferenceIdentity"`
	Kind                      *string `pulumi:"kind"`
	Name                      string  `pulumi:"name"`
	ProvisioningDetails       string  `pulumi:"provisioningDetails"`
	ProvisioningState         string  `pulumi:"provisioningState"`
	Type                      string  `pulumi:"type"`
}

func LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput(ctx *pulumi.Context, args LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResult, error) {
			args := v.(LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationArgs)
			r, err := LookupAppServiceEnvironmentAseCustomDnsSuffixConfiguration(ctx, &args, opts...)
			var s LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput)
}

type LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationArgs)(nil)).Elem()
}


type LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResult)(nil)).Elem()
}

func (o LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput) ToLookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput() LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput {
	return o
}

func (o LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput) ToLookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutputWithContext(ctx context.Context) LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput {
	return o
}

func (o LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResult) *string {
		return v.CertificateUrl
	}).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResult) *string { return v.DnsSuffix }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResult) *string {
		return v.KeyVaultReferenceIdentity
	}).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput) ProvisioningDetails() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResult) string {
		return v.ProvisioningDetails
	}).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResult) string {
		return v.ProvisioningState
	}).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppServiceEnvironmentAseCustomDnsSuffixConfigurationResultOutput{})
}
