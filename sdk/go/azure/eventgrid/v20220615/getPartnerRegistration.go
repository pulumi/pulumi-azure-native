


package v20220615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPartnerRegistration(ctx *pulumi.Context, args *LookupPartnerRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupPartnerRegistrationResult, error) {
	var rv LookupPartnerRegistrationResult
	err := ctx.Invoke("azure-native:eventgrid/v20220615:getPartnerRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerRegistrationArgs struct {
	PartnerRegistrationName string `pulumi:"partnerRegistrationName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupPartnerRegistrationResult struct {
	Id                             string             `pulumi:"id"`
	Location                       string             `pulumi:"location"`
	Name                           string             `pulumi:"name"`
	PartnerRegistrationImmutableId *string            `pulumi:"partnerRegistrationImmutableId"`
	ProvisioningState              string             `pulumi:"provisioningState"`
	SystemData                     SystemDataResponse `pulumi:"systemData"`
	Tags                           map[string]string  `pulumi:"tags"`
	Type                           string             `pulumi:"type"`
}

func LookupPartnerRegistrationOutput(ctx *pulumi.Context, args LookupPartnerRegistrationOutputArgs, opts ...pulumi.InvokeOption) LookupPartnerRegistrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPartnerRegistrationResult, error) {
			args := v.(LookupPartnerRegistrationArgs)
			r, err := LookupPartnerRegistration(ctx, &args, opts...)
			var s LookupPartnerRegistrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPartnerRegistrationResultOutput)
}

type LookupPartnerRegistrationOutputArgs struct {
	PartnerRegistrationName pulumi.StringInput `pulumi:"partnerRegistrationName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPartnerRegistrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerRegistrationArgs)(nil)).Elem()
}


type LookupPartnerRegistrationResultOutput struct{ *pulumi.OutputState }

func (LookupPartnerRegistrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerRegistrationResult)(nil)).Elem()
}

func (o LookupPartnerRegistrationResultOutput) ToLookupPartnerRegistrationResultOutput() LookupPartnerRegistrationResultOutput {
	return o
}

func (o LookupPartnerRegistrationResultOutput) ToLookupPartnerRegistrationResultOutputWithContext(ctx context.Context) LookupPartnerRegistrationResultOutput {
	return o
}

func (o LookupPartnerRegistrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPartnerRegistrationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPartnerRegistrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPartnerRegistrationResultOutput) PartnerRegistrationImmutableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) *string { return v.PartnerRegistrationImmutableId }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerRegistrationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPartnerRegistrationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPartnerRegistrationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPartnerRegistrationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartnerRegistrationResultOutput{})
}
