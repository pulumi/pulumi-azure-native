


package securityandcompliance

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetprivateLinkServicesForMIPPolicySync(ctx *pulumi.Context, args *GetprivateLinkServicesForMIPPolicySyncArgs, opts ...pulumi.InvokeOption) (*GetprivateLinkServicesForMIPPolicySyncResult, error) {
	var rv GetprivateLinkServicesForMIPPolicySyncResult
	err := ctx.Invoke("azure-native:securityandcompliance:getprivateLinkServicesForMIPPolicySync", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetprivateLinkServicesForMIPPolicySyncArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type GetprivateLinkServicesForMIPPolicySyncResult struct {
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

func GetprivateLinkServicesForMIPPolicySyncOutput(ctx *pulumi.Context, args GetprivateLinkServicesForMIPPolicySyncOutputArgs, opts ...pulumi.InvokeOption) GetprivateLinkServicesForMIPPolicySyncResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetprivateLinkServicesForMIPPolicySyncResult, error) {
			args := v.(GetprivateLinkServicesForMIPPolicySyncArgs)
			r, err := GetprivateLinkServicesForMIPPolicySync(ctx, &args, opts...)
			var s GetprivateLinkServicesForMIPPolicySyncResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetprivateLinkServicesForMIPPolicySyncResultOutput)
}

type GetprivateLinkServicesForMIPPolicySyncOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (GetprivateLinkServicesForMIPPolicySyncOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetprivateLinkServicesForMIPPolicySyncArgs)(nil)).Elem()
}


type GetprivateLinkServicesForMIPPolicySyncResultOutput struct{ *pulumi.OutputState }

func (GetprivateLinkServicesForMIPPolicySyncResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetprivateLinkServicesForMIPPolicySyncResult)(nil)).Elem()
}

func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) ToGetprivateLinkServicesForMIPPolicySyncResultOutput() GetprivateLinkServicesForMIPPolicySyncResultOutput {
	return o
}

func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) ToGetprivateLinkServicesForMIPPolicySyncResultOutputWithContext(ctx context.Context) GetprivateLinkServicesForMIPPolicySyncResultOutput {
	return o
}

func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) *ServicesResourceResponseIdentity {
		return v.Identity
	}).(ServicesResourceResponseIdentityPtrOutput)
}

func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) ServicesPropertiesResponse { return v.Properties }).(ServicesPropertiesResponseOutput)
}

func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetprivateLinkServicesForMIPPolicySyncResultOutput{})
}
