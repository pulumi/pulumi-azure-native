


package v20210325preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetprivateLinkServicesForM365ComplianceCenter(ctx *pulumi.Context, args *GetprivateLinkServicesForM365ComplianceCenterArgs, opts ...pulumi.InvokeOption) (*GetprivateLinkServicesForM365ComplianceCenterResult, error) {
	var rv GetprivateLinkServicesForM365ComplianceCenterResult
	err := ctx.Invoke("azure-native:m365securityandcompliance/v20210325preview:getprivateLinkServicesForM365ComplianceCenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetprivateLinkServicesForM365ComplianceCenterArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type GetprivateLinkServicesForM365ComplianceCenterResult struct {
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

func GetprivateLinkServicesForM365ComplianceCenterOutput(ctx *pulumi.Context, args GetprivateLinkServicesForM365ComplianceCenterOutputArgs, opts ...pulumi.InvokeOption) GetprivateLinkServicesForM365ComplianceCenterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetprivateLinkServicesForM365ComplianceCenterResult, error) {
			args := v.(GetprivateLinkServicesForM365ComplianceCenterArgs)
			r, err := GetprivateLinkServicesForM365ComplianceCenter(ctx, &args, opts...)
			var s GetprivateLinkServicesForM365ComplianceCenterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetprivateLinkServicesForM365ComplianceCenterResultOutput)
}

type GetprivateLinkServicesForM365ComplianceCenterOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (GetprivateLinkServicesForM365ComplianceCenterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetprivateLinkServicesForM365ComplianceCenterArgs)(nil)).Elem()
}


type GetprivateLinkServicesForM365ComplianceCenterResultOutput struct{ *pulumi.OutputState }

func (GetprivateLinkServicesForM365ComplianceCenterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetprivateLinkServicesForM365ComplianceCenterResult)(nil)).Elem()
}

func (o GetprivateLinkServicesForM365ComplianceCenterResultOutput) ToGetprivateLinkServicesForM365ComplianceCenterResultOutput() GetprivateLinkServicesForM365ComplianceCenterResultOutput {
	return o
}

func (o GetprivateLinkServicesForM365ComplianceCenterResultOutput) ToGetprivateLinkServicesForM365ComplianceCenterResultOutputWithContext(ctx context.Context) GetprivateLinkServicesForM365ComplianceCenterResultOutput {
	return o
}

func (o GetprivateLinkServicesForM365ComplianceCenterResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365ComplianceCenterResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o GetprivateLinkServicesForM365ComplianceCenterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365ComplianceCenterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForM365ComplianceCenterResultOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365ComplianceCenterResult) *ServicesResourceResponseIdentity {
		return v.Identity
	}).(ServicesResourceResponseIdentityPtrOutput)
}

func (o GetprivateLinkServicesForM365ComplianceCenterResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365ComplianceCenterResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForM365ComplianceCenterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365ComplianceCenterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForM365ComplianceCenterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365ComplianceCenterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForM365ComplianceCenterResultOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365ComplianceCenterResult) ServicesPropertiesResponse {
		return v.Properties
	}).(ServicesPropertiesResponseOutput)
}

func (o GetprivateLinkServicesForM365ComplianceCenterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365ComplianceCenterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetprivateLinkServicesForM365ComplianceCenterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365ComplianceCenterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetprivateLinkServicesForM365ComplianceCenterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365ComplianceCenterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetprivateLinkServicesForM365ComplianceCenterResultOutput{})
}
