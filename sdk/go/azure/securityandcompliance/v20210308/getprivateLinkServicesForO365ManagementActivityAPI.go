


package v20210308

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetprivateLinkServicesForO365ManagementActivityAPI(ctx *pulumi.Context, args *GetprivateLinkServicesForO365ManagementActivityAPIArgs, opts ...pulumi.InvokeOption) (*GetprivateLinkServicesForO365ManagementActivityAPIResult, error) {
	var rv GetprivateLinkServicesForO365ManagementActivityAPIResult
	err := ctx.Invoke("azure-native:securityandcompliance/v20210308:getprivateLinkServicesForO365ManagementActivityAPI", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetprivateLinkServicesForO365ManagementActivityAPIArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type GetprivateLinkServicesForO365ManagementActivityAPIResult struct {
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

func GetprivateLinkServicesForO365ManagementActivityAPIOutput(ctx *pulumi.Context, args GetprivateLinkServicesForO365ManagementActivityAPIOutputArgs, opts ...pulumi.InvokeOption) GetprivateLinkServicesForO365ManagementActivityAPIResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetprivateLinkServicesForO365ManagementActivityAPIResult, error) {
			args := v.(GetprivateLinkServicesForO365ManagementActivityAPIArgs)
			r, err := GetprivateLinkServicesForO365ManagementActivityAPI(ctx, &args, opts...)
			var s GetprivateLinkServicesForO365ManagementActivityAPIResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetprivateLinkServicesForO365ManagementActivityAPIResultOutput)
}

type GetprivateLinkServicesForO365ManagementActivityAPIOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (GetprivateLinkServicesForO365ManagementActivityAPIOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetprivateLinkServicesForO365ManagementActivityAPIArgs)(nil)).Elem()
}


type GetprivateLinkServicesForO365ManagementActivityAPIResultOutput struct{ *pulumi.OutputState }

func (GetprivateLinkServicesForO365ManagementActivityAPIResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetprivateLinkServicesForO365ManagementActivityAPIResult)(nil)).Elem()
}

func (o GetprivateLinkServicesForO365ManagementActivityAPIResultOutput) ToGetprivateLinkServicesForO365ManagementActivityAPIResultOutput() GetprivateLinkServicesForO365ManagementActivityAPIResultOutput {
	return o
}

func (o GetprivateLinkServicesForO365ManagementActivityAPIResultOutput) ToGetprivateLinkServicesForO365ManagementActivityAPIResultOutputWithContext(ctx context.Context) GetprivateLinkServicesForO365ManagementActivityAPIResultOutput {
	return o
}

func (o GetprivateLinkServicesForO365ManagementActivityAPIResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForO365ManagementActivityAPIResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o GetprivateLinkServicesForO365ManagementActivityAPIResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForO365ManagementActivityAPIResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForO365ManagementActivityAPIResultOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForO365ManagementActivityAPIResult) *ServicesResourceResponseIdentity {
		return v.Identity
	}).(ServicesResourceResponseIdentityPtrOutput)
}

func (o GetprivateLinkServicesForO365ManagementActivityAPIResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForO365ManagementActivityAPIResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForO365ManagementActivityAPIResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForO365ManagementActivityAPIResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForO365ManagementActivityAPIResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForO365ManagementActivityAPIResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForO365ManagementActivityAPIResultOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForO365ManagementActivityAPIResult) ServicesPropertiesResponse {
		return v.Properties
	}).(ServicesPropertiesResponseOutput)
}

func (o GetprivateLinkServicesForO365ManagementActivityAPIResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForO365ManagementActivityAPIResult) SystemDataResponse {
		return v.SystemData
	}).(SystemDataResponseOutput)
}

func (o GetprivateLinkServicesForO365ManagementActivityAPIResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForO365ManagementActivityAPIResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetprivateLinkServicesForO365ManagementActivityAPIResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForO365ManagementActivityAPIResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetprivateLinkServicesForO365ManagementActivityAPIResultOutput{})
}
