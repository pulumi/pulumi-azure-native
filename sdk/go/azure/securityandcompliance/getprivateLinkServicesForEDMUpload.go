


package securityandcompliance

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetprivateLinkServicesForEDMUpload(ctx *pulumi.Context, args *GetprivateLinkServicesForEDMUploadArgs, opts ...pulumi.InvokeOption) (*GetprivateLinkServicesForEDMUploadResult, error) {
	var rv GetprivateLinkServicesForEDMUploadResult
	err := ctx.Invoke("azure-native:securityandcompliance:getprivateLinkServicesForEDMUpload", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetprivateLinkServicesForEDMUploadArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type GetprivateLinkServicesForEDMUploadResult struct {
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

func GetprivateLinkServicesForEDMUploadOutput(ctx *pulumi.Context, args GetprivateLinkServicesForEDMUploadOutputArgs, opts ...pulumi.InvokeOption) GetprivateLinkServicesForEDMUploadResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetprivateLinkServicesForEDMUploadResult, error) {
			args := v.(GetprivateLinkServicesForEDMUploadArgs)
			r, err := GetprivateLinkServicesForEDMUpload(ctx, &args, opts...)
			var s GetprivateLinkServicesForEDMUploadResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetprivateLinkServicesForEDMUploadResultOutput)
}

type GetprivateLinkServicesForEDMUploadOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (GetprivateLinkServicesForEDMUploadOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetprivateLinkServicesForEDMUploadArgs)(nil)).Elem()
}


type GetprivateLinkServicesForEDMUploadResultOutput struct{ *pulumi.OutputState }

func (GetprivateLinkServicesForEDMUploadResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetprivateLinkServicesForEDMUploadResult)(nil)).Elem()
}

func (o GetprivateLinkServicesForEDMUploadResultOutput) ToGetprivateLinkServicesForEDMUploadResultOutput() GetprivateLinkServicesForEDMUploadResultOutput {
	return o
}

func (o GetprivateLinkServicesForEDMUploadResultOutput) ToGetprivateLinkServicesForEDMUploadResultOutputWithContext(ctx context.Context) GetprivateLinkServicesForEDMUploadResultOutput {
	return o
}

func (o GetprivateLinkServicesForEDMUploadResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForEDMUploadResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o GetprivateLinkServicesForEDMUploadResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForEDMUploadResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForEDMUploadResultOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForEDMUploadResult) *ServicesResourceResponseIdentity { return v.Identity }).(ServicesResourceResponseIdentityPtrOutput)
}

func (o GetprivateLinkServicesForEDMUploadResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForEDMUploadResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForEDMUploadResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForEDMUploadResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForEDMUploadResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForEDMUploadResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForEDMUploadResultOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForEDMUploadResult) ServicesPropertiesResponse { return v.Properties }).(ServicesPropertiesResponseOutput)
}

func (o GetprivateLinkServicesForEDMUploadResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForEDMUploadResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetprivateLinkServicesForEDMUploadResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForEDMUploadResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetprivateLinkServicesForEDMUploadResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForEDMUploadResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetprivateLinkServicesForEDMUploadResultOutput{})
}
