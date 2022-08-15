


package v20210308

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetprivateLinkServicesForSCCPowershell(ctx *pulumi.Context, args *GetprivateLinkServicesForSCCPowershellArgs, opts ...pulumi.InvokeOption) (*GetprivateLinkServicesForSCCPowershellResult, error) {
	var rv GetprivateLinkServicesForSCCPowershellResult
	err := ctx.Invoke("azure-native:securityandcompliance/v20210308:getprivateLinkServicesForSCCPowershell", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetprivateLinkServicesForSCCPowershellArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type GetprivateLinkServicesForSCCPowershellResult struct {
	Etag       *string                           `pulumi:"etag"`
	Id         string                            `pulumi:"id"`
	Identity   *ServicesResourceResponseIdentity `pulumi:"identity"`
	Kind       string                            `pulumi:"kind"`
	Location   string                            `pulumi:"location"`
	Name       string                            `pulumi:"name"`
	Properties ServicesPropertiesResponse        `pulumi:"properties"`
	SystemData SystemDataResponse                `pulumi:"systemData"`
	Tags       map[string]string                 `pulumi:"tags"`
	Type       string                            `pulumi:"type"`
}

func GetprivateLinkServicesForSCCPowershellOutput(ctx *pulumi.Context, args GetprivateLinkServicesForSCCPowershellOutputArgs, opts ...pulumi.InvokeOption) GetprivateLinkServicesForSCCPowershellResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetprivateLinkServicesForSCCPowershellResult, error) {
			args := v.(GetprivateLinkServicesForSCCPowershellArgs)
			r, err := GetprivateLinkServicesForSCCPowershell(ctx, &args, opts...)
			var s GetprivateLinkServicesForSCCPowershellResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetprivateLinkServicesForSCCPowershellResultOutput)
}

type GetprivateLinkServicesForSCCPowershellOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (GetprivateLinkServicesForSCCPowershellOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetprivateLinkServicesForSCCPowershellArgs)(nil)).Elem()
}


type GetprivateLinkServicesForSCCPowershellResultOutput struct{ *pulumi.OutputState }

func (GetprivateLinkServicesForSCCPowershellResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetprivateLinkServicesForSCCPowershellResult)(nil)).Elem()
}

func (o GetprivateLinkServicesForSCCPowershellResultOutput) ToGetprivateLinkServicesForSCCPowershellResultOutput() GetprivateLinkServicesForSCCPowershellResultOutput {
	return o
}

func (o GetprivateLinkServicesForSCCPowershellResultOutput) ToGetprivateLinkServicesForSCCPowershellResultOutputWithContext(ctx context.Context) GetprivateLinkServicesForSCCPowershellResultOutput {
	return o
}

func (o GetprivateLinkServicesForSCCPowershellResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForSCCPowershellResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o GetprivateLinkServicesForSCCPowershellResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForSCCPowershellResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForSCCPowershellResultOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForSCCPowershellResult) *ServicesResourceResponseIdentity {
		return v.Identity
	}).(ServicesResourceResponseIdentityPtrOutput)
}

func (o GetprivateLinkServicesForSCCPowershellResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForSCCPowershellResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForSCCPowershellResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForSCCPowershellResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForSCCPowershellResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForSCCPowershellResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForSCCPowershellResultOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForSCCPowershellResult) ServicesPropertiesResponse { return v.Properties }).(ServicesPropertiesResponseOutput)
}

func (o GetprivateLinkServicesForSCCPowershellResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForSCCPowershellResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetprivateLinkServicesForSCCPowershellResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForSCCPowershellResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetprivateLinkServicesForSCCPowershellResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForSCCPowershellResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetprivateLinkServicesForSCCPowershellResultOutput{})
}
