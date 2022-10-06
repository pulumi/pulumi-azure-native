


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApp(ctx *pulumi.Context, args *LookupAppArgs, opts ...pulumi.InvokeOption) (*LookupAppResult, error) {
	var rv LookupAppResult
	err := ctx.Invoke("azure-native:iotcentral/v20210601:getApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupAppResult struct {
	ApplicationId string                                 `pulumi:"applicationId"`
	DisplayName   *string                                `pulumi:"displayName"`
	Id            string                                 `pulumi:"id"`
	Identity      *SystemAssignedServiceIdentityResponse `pulumi:"identity"`
	Location      string                                 `pulumi:"location"`
	Name          string                                 `pulumi:"name"`
	Sku           AppSkuInfoResponse                     `pulumi:"sku"`
	State         string                                 `pulumi:"state"`
	Subdomain     *string                                `pulumi:"subdomain"`
	Tags          map[string]string                      `pulumi:"tags"`
	Template      *string                                `pulumi:"template"`
	Type          string                                 `pulumi:"type"`
}

func LookupAppOutput(ctx *pulumi.Context, args LookupAppOutputArgs, opts ...pulumi.InvokeOption) LookupAppResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppResult, error) {
			args := v.(LookupAppArgs)
			r, err := LookupApp(ctx, &args, opts...)
			var s LookupAppResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppResultOutput)
}

type LookupAppOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupAppOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppArgs)(nil)).Elem()
}


type LookupAppResultOutput struct{ *pulumi.OutputState }

func (LookupAppResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppResult)(nil)).Elem()
}

func (o LookupAppResultOutput) ToLookupAppResultOutput() LookupAppResultOutput {
	return o
}

func (o LookupAppResultOutput) ToLookupAppResultOutputWithContext(ctx context.Context) LookupAppResultOutput {
	return o
}

func (o LookupAppResultOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.ApplicationId }).(pulumi.StringOutput)
}

func (o LookupAppResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAppResultOutput) Identity() SystemAssignedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupAppResult) *SystemAssignedServiceIdentityResponse { return v.Identity }).(SystemAssignedServiceIdentityResponsePtrOutput)
}

func (o LookupAppResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAppResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAppResultOutput) Sku() AppSkuInfoResponseOutput {
	return o.ApplyT(func(v LookupAppResult) AppSkuInfoResponse { return v.Sku }).(AppSkuInfoResponseOutput)
}

func (o LookupAppResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupAppResultOutput) Subdomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.Subdomain }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAppResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAppResultOutput) Template() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.Template }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppResultOutput{})
}
