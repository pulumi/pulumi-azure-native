


package v20210111

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetprivateLinkServicesForM365SecurityCenter(ctx *pulumi.Context, args *GetprivateLinkServicesForM365SecurityCenterArgs, opts ...pulumi.InvokeOption) (*GetprivateLinkServicesForM365SecurityCenterResult, error) {
	var rv GetprivateLinkServicesForM365SecurityCenterResult
	err := ctx.Invoke("azure-native:securityandcompliance/v20210111:getprivateLinkServicesForM365SecurityCenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetprivateLinkServicesForM365SecurityCenterArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type GetprivateLinkServicesForM365SecurityCenterResult struct {
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

func GetprivateLinkServicesForM365SecurityCenterOutput(ctx *pulumi.Context, args GetprivateLinkServicesForM365SecurityCenterOutputArgs, opts ...pulumi.InvokeOption) GetprivateLinkServicesForM365SecurityCenterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetprivateLinkServicesForM365SecurityCenterResult, error) {
			args := v.(GetprivateLinkServicesForM365SecurityCenterArgs)
			r, err := GetprivateLinkServicesForM365SecurityCenter(ctx, &args, opts...)
			var s GetprivateLinkServicesForM365SecurityCenterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetprivateLinkServicesForM365SecurityCenterResultOutput)
}

type GetprivateLinkServicesForM365SecurityCenterOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (GetprivateLinkServicesForM365SecurityCenterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetprivateLinkServicesForM365SecurityCenterArgs)(nil)).Elem()
}


type GetprivateLinkServicesForM365SecurityCenterResultOutput struct{ *pulumi.OutputState }

func (GetprivateLinkServicesForM365SecurityCenterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetprivateLinkServicesForM365SecurityCenterResult)(nil)).Elem()
}

func (o GetprivateLinkServicesForM365SecurityCenterResultOutput) ToGetprivateLinkServicesForM365SecurityCenterResultOutput() GetprivateLinkServicesForM365SecurityCenterResultOutput {
	return o
}

func (o GetprivateLinkServicesForM365SecurityCenterResultOutput) ToGetprivateLinkServicesForM365SecurityCenterResultOutputWithContext(ctx context.Context) GetprivateLinkServicesForM365SecurityCenterResultOutput {
	return o
}

func (o GetprivateLinkServicesForM365SecurityCenterResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365SecurityCenterResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o GetprivateLinkServicesForM365SecurityCenterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365SecurityCenterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForM365SecurityCenterResultOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365SecurityCenterResult) *ServicesResourceResponseIdentity {
		return v.Identity
	}).(ServicesResourceResponseIdentityPtrOutput)
}

func (o GetprivateLinkServicesForM365SecurityCenterResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365SecurityCenterResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForM365SecurityCenterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365SecurityCenterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForM365SecurityCenterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365SecurityCenterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForM365SecurityCenterResultOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365SecurityCenterResult) ServicesPropertiesResponse {
		return v.Properties
	}).(ServicesPropertiesResponseOutput)
}

func (o GetprivateLinkServicesForM365SecurityCenterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365SecurityCenterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetprivateLinkServicesForM365SecurityCenterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365SecurityCenterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetprivateLinkServicesForM365SecurityCenterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForM365SecurityCenterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetprivateLinkServicesForM365SecurityCenterResultOutput{})
}
