


package v20201015preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPartnerRegistration(ctx *pulumi.Context, args *LookupPartnerRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupPartnerRegistrationResult, error) {
	var rv LookupPartnerRegistrationResult
	err := ctx.Invoke("azure-native:eventgrid/v20201015preview:getPartnerRegistration", args, &rv, opts...)
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
	AuthorizedAzureSubscriptionIds  []string           `pulumi:"authorizedAzureSubscriptionIds"`
	CustomerServiceUri              *string            `pulumi:"customerServiceUri"`
	Id                              string             `pulumi:"id"`
	Location                        string             `pulumi:"location"`
	LogoUri                         *string            `pulumi:"logoUri"`
	LongDescription                 *string            `pulumi:"longDescription"`
	Name                            string             `pulumi:"name"`
	PartnerCustomerServiceExtension *string            `pulumi:"partnerCustomerServiceExtension"`
	PartnerCustomerServiceNumber    *string            `pulumi:"partnerCustomerServiceNumber"`
	PartnerName                     *string            `pulumi:"partnerName"`
	PartnerResourceTypeDescription  *string            `pulumi:"partnerResourceTypeDescription"`
	PartnerResourceTypeDisplayName  *string            `pulumi:"partnerResourceTypeDisplayName"`
	PartnerResourceTypeName         *string            `pulumi:"partnerResourceTypeName"`
	ProvisioningState               string             `pulumi:"provisioningState"`
	SetupUri                        *string            `pulumi:"setupUri"`
	SystemData                      SystemDataResponse `pulumi:"systemData"`
	Tags                            map[string]string  `pulumi:"tags"`
	Type                            string             `pulumi:"type"`
	VisibilityState                 *string            `pulumi:"visibilityState"`
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

func (o LookupPartnerRegistrationResultOutput) AuthorizedAzureSubscriptionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) []string { return v.AuthorizedAzureSubscriptionIds }).(pulumi.StringArrayOutput)
}

func (o LookupPartnerRegistrationResultOutput) CustomerServiceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) *string { return v.CustomerServiceUri }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerRegistrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPartnerRegistrationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPartnerRegistrationResultOutput) LogoUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) *string { return v.LogoUri }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerRegistrationResultOutput) LongDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) *string { return v.LongDescription }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerRegistrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPartnerRegistrationResultOutput) PartnerCustomerServiceExtension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) *string { return v.PartnerCustomerServiceExtension }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerRegistrationResultOutput) PartnerCustomerServiceNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) *string { return v.PartnerCustomerServiceNumber }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerRegistrationResultOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) *string { return v.PartnerName }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerRegistrationResultOutput) PartnerResourceTypeDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) *string { return v.PartnerResourceTypeDescription }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerRegistrationResultOutput) PartnerResourceTypeDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) *string { return v.PartnerResourceTypeDisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerRegistrationResultOutput) PartnerResourceTypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) *string { return v.PartnerResourceTypeName }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerRegistrationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPartnerRegistrationResultOutput) SetupUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) *string { return v.SetupUri }).(pulumi.StringPtrOutput)
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

func (o LookupPartnerRegistrationResultOutput) VisibilityState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerRegistrationResult) *string { return v.VisibilityState }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartnerRegistrationResultOutput{})
}
