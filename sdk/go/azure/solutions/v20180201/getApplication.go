


package v20180201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:solutions/v20180201:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	ApplicationName   string `pulumi:"applicationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupApplicationResult struct {
	ApplicationDefinitionId *string                                     `pulumi:"applicationDefinitionId"`
	Artifacts               []ApplicationArtifactResponse               `pulumi:"artifacts"`
	Authorizations          []ApplicationAuthorizationResponse          `pulumi:"authorizations"`
	BillingDetails          ApplicationBillingDetailsDefinitionResponse `pulumi:"billingDetails"`
	CreatedBy               ApplicationClientDetailsResponse            `pulumi:"createdBy"`
	CustomerSupport         ApplicationPackageContactResponse           `pulumi:"customerSupport"`
	Id                      string                                      `pulumi:"id"`
	Identity                *IdentityResponse                           `pulumi:"identity"`
	Kind                    string                                      `pulumi:"kind"`
	Location                *string                                     `pulumi:"location"`
	ManagedBy               *string                                     `pulumi:"managedBy"`
	ManagedResourceGroupId  *string                                     `pulumi:"managedResourceGroupId"`
	Name                    string                                      `pulumi:"name"`
	Outputs                 interface{}                                 `pulumi:"outputs"`
	Parameters              interface{}                                 `pulumi:"parameters"`
	Plan                    *PlanResponse                               `pulumi:"plan"`
	ProvisioningState       string                                      `pulumi:"provisioningState"`
	PublisherPackageId      *string                                     `pulumi:"publisherPackageId"`
	PublisherTenantId       string                                      `pulumi:"publisherTenantId"`
	Sku                     *SkuResponse                                `pulumi:"sku"`
	SupportUrls             ApplicationPackageSupportUrlsResponse       `pulumi:"supportUrls"`
	Tags                    map[string]string                           `pulumi:"tags"`
	Type                    string                                      `pulumi:"type"`
	UpdatedBy               ApplicationClientDetailsResponse            `pulumi:"updatedBy"`
}

func LookupApplicationOutput(ctx *pulumi.Context, args LookupApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationResult, error) {
			args := v.(LookupApplicationArgs)
			r, err := LookupApplication(ctx, &args, opts...)
			var s LookupApplicationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationResultOutput)
}

type LookupApplicationOutputArgs struct {
	ApplicationName   pulumi.StringInput `pulumi:"applicationName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}


type LookupApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResult)(nil)).Elem()
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutput() LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutputWithContext(ctx context.Context) LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ApplicationDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.ApplicationDefinitionId }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) Artifacts() ApplicationArtifactResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []ApplicationArtifactResponse { return v.Artifacts }).(ApplicationArtifactResponseArrayOutput)
}

func (o LookupApplicationResultOutput) Authorizations() ApplicationAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []ApplicationAuthorizationResponse { return v.Authorizations }).(ApplicationAuthorizationResponseArrayOutput)
}

func (o LookupApplicationResultOutput) BillingDetails() ApplicationBillingDetailsDefinitionResponseOutput {
	return o.ApplyT(func(v LookupApplicationResult) ApplicationBillingDetailsDefinitionResponse { return v.BillingDetails }).(ApplicationBillingDetailsDefinitionResponseOutput)
}

func (o LookupApplicationResultOutput) CreatedBy() ApplicationClientDetailsResponseOutput {
	return o.ApplyT(func(v LookupApplicationResult) ApplicationClientDetailsResponse { return v.CreatedBy }).(ApplicationClientDetailsResponseOutput)
}

func (o LookupApplicationResultOutput) CustomerSupport() ApplicationPackageContactResponseOutput {
	return o.ApplyT(func(v LookupApplicationResult) ApplicationPackageContactResponse { return v.CustomerSupport }).(ApplicationPackageContactResponseOutput)
}

func (o LookupApplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupApplicationResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) ManagedResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.ManagedResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Outputs() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupApplicationResult) interface{} { return v.Outputs }).(pulumi.AnyOutput)
}

func (o LookupApplicationResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupApplicationResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o LookupApplicationResultOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *PlanResponse { return v.Plan }).(PlanResponsePtrOutput)
}

func (o LookupApplicationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) PublisherPackageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.PublisherPackageId }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) PublisherTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.PublisherTenantId }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupApplicationResultOutput) SupportUrls() ApplicationPackageSupportUrlsResponseOutput {
	return o.ApplyT(func(v LookupApplicationResult) ApplicationPackageSupportUrlsResponse { return v.SupportUrls }).(ApplicationPackageSupportUrlsResponseOutput)
}

func (o LookupApplicationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupApplicationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) UpdatedBy() ApplicationClientDetailsResponseOutput {
	return o.ApplyT(func(v LookupApplicationResult) ApplicationClientDetailsResponse { return v.UpdatedBy }).(ApplicationClientDetailsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}
