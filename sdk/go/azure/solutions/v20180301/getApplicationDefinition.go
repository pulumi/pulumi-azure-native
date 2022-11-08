


package v20180301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupApplicationDefinition(ctx *pulumi.Context, args *LookupApplicationDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupApplicationDefinitionResult, error) {
	var rv LookupApplicationDefinitionResult
	err := ctx.Invoke("azure-native:solutions/v20180301:getApplicationDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationDefinitionArgs struct {
	ApplicationDefinitionName string `pulumi:"applicationDefinitionName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
}


type LookupApplicationDefinitionResult struct {
	Artifacts          []ApplicationDefinitionArtifactResponse `pulumi:"artifacts"`
	Authorizations     []ApplicationAuthorizationResponse      `pulumi:"authorizations"`
	CreateUiDefinition interface{}                             `pulumi:"createUiDefinition"`
	Description        *string                                 `pulumi:"description"`
	DisplayName        *string                                 `pulumi:"displayName"`
	Id                 string                                  `pulumi:"id"`
	IsEnabled          *bool                                   `pulumi:"isEnabled"`
	Location           *string                                 `pulumi:"location"`
	LockLevel          string                                  `pulumi:"lockLevel"`
	MainTemplate       interface{}                             `pulumi:"mainTemplate"`
	ManagedBy          *string                                 `pulumi:"managedBy"`
	Name               string                                  `pulumi:"name"`
	PackageFileUri     *string                                 `pulumi:"packageFileUri"`
	Policies           []ApplicationPolicyResponse             `pulumi:"policies"`
	Sku                *SkuResponse                            `pulumi:"sku"`
	Tags               map[string]string                       `pulumi:"tags"`
	Type               string                                  `pulumi:"type"`
}

func LookupApplicationDefinitionOutput(ctx *pulumi.Context, args LookupApplicationDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationDefinitionResult, error) {
			args := v.(LookupApplicationDefinitionArgs)
			r, err := LookupApplicationDefinition(ctx, &args, opts...)
			var s LookupApplicationDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationDefinitionResultOutput)
}

type LookupApplicationDefinitionOutputArgs struct {
	ApplicationDefinitionName pulumi.StringInput `pulumi:"applicationDefinitionName"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationDefinitionArgs)(nil)).Elem()
}


type LookupApplicationDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationDefinitionResult)(nil)).Elem()
}

func (o LookupApplicationDefinitionResultOutput) ToLookupApplicationDefinitionResultOutput() LookupApplicationDefinitionResultOutput {
	return o
}

func (o LookupApplicationDefinitionResultOutput) ToLookupApplicationDefinitionResultOutputWithContext(ctx context.Context) LookupApplicationDefinitionResultOutput {
	return o
}

func (o LookupApplicationDefinitionResultOutput) Artifacts() ApplicationDefinitionArtifactResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) []ApplicationDefinitionArtifactResponse { return v.Artifacts }).(ApplicationDefinitionArtifactResponseArrayOutput)
}

func (o LookupApplicationDefinitionResultOutput) Authorizations() ApplicationAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) []ApplicationAuthorizationResponse { return v.Authorizations }).(ApplicationAuthorizationResponseArrayOutput)
}

func (o LookupApplicationDefinitionResultOutput) CreateUiDefinition() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) interface{} { return v.CreateUiDefinition }).(pulumi.AnyOutput)
}

func (o LookupApplicationDefinitionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationDefinitionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationDefinitionResultOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupApplicationDefinitionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationDefinitionResultOutput) LockLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) string { return v.LockLevel }).(pulumi.StringOutput)
}

func (o LookupApplicationDefinitionResultOutput) MainTemplate() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) interface{} { return v.MainTemplate }).(pulumi.AnyOutput)
}

func (o LookupApplicationDefinitionResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationDefinitionResultOutput) PackageFileUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *string { return v.PackageFileUri }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationDefinitionResultOutput) Policies() ApplicationPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) []ApplicationPolicyResponse { return v.Policies }).(ApplicationPolicyResponseArrayOutput)
}

func (o LookupApplicationDefinitionResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupApplicationDefinitionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupApplicationDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationDefinitionResultOutput{})
}
