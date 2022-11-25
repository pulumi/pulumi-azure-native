


package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPartnerConfiguration(ctx *pulumi.Context, args *LookupPartnerConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupPartnerConfigurationResult, error) {
	var rv LookupPartnerConfigurationResult
	err := ctx.Invoke("azure-native:eventgrid:getPartnerConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerConfigurationArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPartnerConfigurationResult struct {
	Id                   string                        `pulumi:"id"`
	Location             *string                       `pulumi:"location"`
	Name                 string                        `pulumi:"name"`
	PartnerAuthorization *PartnerAuthorizationResponse `pulumi:"partnerAuthorization"`
	ProvisioningState    *string                       `pulumi:"provisioningState"`
	SystemData           SystemDataResponse            `pulumi:"systemData"`
	Tags                 map[string]string             `pulumi:"tags"`
	Type                 string                        `pulumi:"type"`
}

func LookupPartnerConfigurationOutput(ctx *pulumi.Context, args LookupPartnerConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupPartnerConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPartnerConfigurationResult, error) {
			args := v.(LookupPartnerConfigurationArgs)
			r, err := LookupPartnerConfiguration(ctx, &args, opts...)
			var s LookupPartnerConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPartnerConfigurationResultOutput)
}

type LookupPartnerConfigurationOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPartnerConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerConfigurationArgs)(nil)).Elem()
}


type LookupPartnerConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupPartnerConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerConfigurationResult)(nil)).Elem()
}

func (o LookupPartnerConfigurationResultOutput) ToLookupPartnerConfigurationResultOutput() LookupPartnerConfigurationResultOutput {
	return o
}

func (o LookupPartnerConfigurationResultOutput) ToLookupPartnerConfigurationResultOutputWithContext(ctx context.Context) LookupPartnerConfigurationResultOutput {
	return o
}

func (o LookupPartnerConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPartnerConfigurationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerConfigurationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPartnerConfigurationResultOutput) PartnerAuthorization() PartnerAuthorizationResponsePtrOutput {
	return o.ApplyT(func(v LookupPartnerConfigurationResult) *PartnerAuthorizationResponse { return v.PartnerAuthorization }).(PartnerAuthorizationResponsePtrOutput)
}

func (o LookupPartnerConfigurationResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerConfigurationResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPartnerConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPartnerConfigurationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPartnerConfigurationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPartnerConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartnerConfigurationResultOutput{})
}
