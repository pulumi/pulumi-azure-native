


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOrganization(ctx *pulumi.Context, args *LookupOrganizationArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationResult, error) {
	var rv LookupOrganizationResult
	err := ctx.Invoke("azure-native:confluent/v20210301preview:getOrganization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrganizationArgs struct {
	OrganizationName  string `pulumi:"organizationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupOrganizationResult struct {
	CreatedTime       string              `pulumi:"createdTime"`
	Id                string              `pulumi:"id"`
	Location          *string             `pulumi:"location"`
	Name              string              `pulumi:"name"`
	OfferDetail       OfferDetailResponse `pulumi:"offerDetail"`
	OrganizationId    string              `pulumi:"organizationId"`
	ProvisioningState string              `pulumi:"provisioningState"`
	SsoUrl            string              `pulumi:"ssoUrl"`
	SystemData        SystemDataResponse  `pulumi:"systemData"`
	Tags              map[string]string   `pulumi:"tags"`
	Type              string              `pulumi:"type"`
	UserDetail        UserDetailResponse  `pulumi:"userDetail"`
}

func LookupOrganizationOutput(ctx *pulumi.Context, args LookupOrganizationOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrganizationResult, error) {
			args := v.(LookupOrganizationArgs)
			r, err := LookupOrganization(ctx, &args, opts...)
			var s LookupOrganizationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrganizationResultOutput)
}

type LookupOrganizationOutputArgs struct {
	OrganizationName  pulumi.StringInput `pulumi:"organizationName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOrganizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationArgs)(nil)).Elem()
}


type LookupOrganizationResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationResult)(nil)).Elem()
}

func (o LookupOrganizationResultOutput) ToLookupOrganizationResultOutput() LookupOrganizationResultOutput {
	return o
}

func (o LookupOrganizationResultOutput) ToLookupOrganizationResultOutputWithContext(ctx context.Context) LookupOrganizationResultOutput {
	return o
}

func (o LookupOrganizationResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupOrganizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOrganizationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrganizationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupOrganizationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOrganizationResultOutput) OfferDetail() OfferDetailResponseOutput {
	return o.ApplyT(func(v LookupOrganizationResult) OfferDetailResponse { return v.OfferDetail }).(OfferDetailResponseOutput)
}

func (o LookupOrganizationResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupOrganizationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupOrganizationResultOutput) SsoUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.SsoUrl }).(pulumi.StringOutput)
}

func (o LookupOrganizationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOrganizationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOrganizationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOrganizationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupOrganizationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupOrganizationResultOutput) UserDetail() UserDetailResponseOutput {
	return o.ApplyT(func(v LookupOrganizationResult) UserDetailResponse { return v.UserDetail }).(UserDetailResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationResultOutput{})
}
