


package v20160901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupApplianceDefinition(ctx *pulumi.Context, args *LookupApplianceDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupApplianceDefinitionResult, error) {
	var rv LookupApplianceDefinitionResult
	err := ctx.Invoke("azure-native:solutions/v20160901preview:getApplianceDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplianceDefinitionArgs struct {
	ApplianceDefinitionName string `pulumi:"applianceDefinitionName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupApplianceDefinitionResult struct {
	Artifacts      []ApplianceArtifactResponse              `pulumi:"artifacts"`
	Authorizations []ApplianceProviderAuthorizationResponse `pulumi:"authorizations"`
	Description    *string                                  `pulumi:"description"`
	DisplayName    *string                                  `pulumi:"displayName"`
	Id             string                                   `pulumi:"id"`
	Identity       *IdentityResponse                        `pulumi:"identity"`
	Location       *string                                  `pulumi:"location"`
	LockLevel      string                                   `pulumi:"lockLevel"`
	ManagedBy      *string                                  `pulumi:"managedBy"`
	Name           string                                   `pulumi:"name"`
	PackageFileUri string                                   `pulumi:"packageFileUri"`
	Sku            *SkuResponse                             `pulumi:"sku"`
	Tags           map[string]string                        `pulumi:"tags"`
	Type           string                                   `pulumi:"type"`
}

func LookupApplianceDefinitionOutput(ctx *pulumi.Context, args LookupApplianceDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupApplianceDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplianceDefinitionResult, error) {
			args := v.(LookupApplianceDefinitionArgs)
			r, err := LookupApplianceDefinition(ctx, &args, opts...)
			var s LookupApplianceDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplianceDefinitionResultOutput)
}

type LookupApplianceDefinitionOutputArgs struct {
	ApplianceDefinitionName pulumi.StringInput `pulumi:"applianceDefinitionName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplianceDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplianceDefinitionArgs)(nil)).Elem()
}


type LookupApplianceDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupApplianceDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplianceDefinitionResult)(nil)).Elem()
}

func (o LookupApplianceDefinitionResultOutput) ToLookupApplianceDefinitionResultOutput() LookupApplianceDefinitionResultOutput {
	return o
}

func (o LookupApplianceDefinitionResultOutput) ToLookupApplianceDefinitionResultOutputWithContext(ctx context.Context) LookupApplianceDefinitionResultOutput {
	return o
}

func (o LookupApplianceDefinitionResultOutput) Artifacts() ApplianceArtifactResponseArrayOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) []ApplianceArtifactResponse { return v.Artifacts }).(ApplianceArtifactResponseArrayOutput)
}

func (o LookupApplianceDefinitionResultOutput) Authorizations() ApplianceProviderAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) []ApplianceProviderAuthorizationResponse {
		return v.Authorizations
	}).(ApplianceProviderAuthorizationResponseArrayOutput)
}

func (o LookupApplianceDefinitionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupApplianceDefinitionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupApplianceDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplianceDefinitionResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupApplianceDefinitionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupApplianceDefinitionResultOutput) LockLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) string { return v.LockLevel }).(pulumi.StringOutput)
}

func (o LookupApplianceDefinitionResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o LookupApplianceDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplianceDefinitionResultOutput) PackageFileUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) string { return v.PackageFileUri }).(pulumi.StringOutput)
}

func (o LookupApplianceDefinitionResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupApplianceDefinitionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupApplianceDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplianceDefinitionResultOutput{})
}
