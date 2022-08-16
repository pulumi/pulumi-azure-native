


package v20170101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupRelationship(ctx *pulumi.Context, args *LookupRelationshipArgs, opts ...pulumi.InvokeOption) (*LookupRelationshipResult, error) {
	var rv LookupRelationshipResult
	err := ctx.Invoke("azure-native:customerinsights/v20170101:getRelationship", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRelationshipArgs struct {
	HubName           string `pulumi:"hubName"`
	RelationshipName  string `pulumi:"relationshipName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRelationshipResult struct {
	Cardinality        *string                           `pulumi:"cardinality"`
	Description        map[string]string                 `pulumi:"description"`
	DisplayName        map[string]string                 `pulumi:"displayName"`
	ExpiryDateTimeUtc  *string                           `pulumi:"expiryDateTimeUtc"`
	Fields             []PropertyDefinitionResponse      `pulumi:"fields"`
	Id                 string                            `pulumi:"id"`
	LookupMappings     []RelationshipTypeMappingResponse `pulumi:"lookupMappings"`
	Name               string                            `pulumi:"name"`
	ProfileType        string                            `pulumi:"profileType"`
	ProvisioningState  string                            `pulumi:"provisioningState"`
	RelatedProfileType string                            `pulumi:"relatedProfileType"`
	RelationshipGuidId string                            `pulumi:"relationshipGuidId"`
	RelationshipName   string                            `pulumi:"relationshipName"`
	TenantId           string                            `pulumi:"tenantId"`
	Type               string                            `pulumi:"type"`
}

func LookupRelationshipOutput(ctx *pulumi.Context, args LookupRelationshipOutputArgs, opts ...pulumi.InvokeOption) LookupRelationshipResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRelationshipResult, error) {
			args := v.(LookupRelationshipArgs)
			r, err := LookupRelationship(ctx, &args, opts...)
			var s LookupRelationshipResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRelationshipResultOutput)
}

type LookupRelationshipOutputArgs struct {
	HubName           pulumi.StringInput `pulumi:"hubName"`
	RelationshipName  pulumi.StringInput `pulumi:"relationshipName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRelationshipOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRelationshipArgs)(nil)).Elem()
}


type LookupRelationshipResultOutput struct{ *pulumi.OutputState }

func (LookupRelationshipResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRelationshipResult)(nil)).Elem()
}

func (o LookupRelationshipResultOutput) ToLookupRelationshipResultOutput() LookupRelationshipResultOutput {
	return o
}

func (o LookupRelationshipResultOutput) ToLookupRelationshipResultOutputWithContext(ctx context.Context) LookupRelationshipResultOutput {
	return o
}

func (o LookupRelationshipResultOutput) Cardinality() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRelationshipResult) *string { return v.Cardinality }).(pulumi.StringPtrOutput)
}

func (o LookupRelationshipResultOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRelationshipResult) map[string]string { return v.Description }).(pulumi.StringMapOutput)
}

func (o LookupRelationshipResultOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRelationshipResult) map[string]string { return v.DisplayName }).(pulumi.StringMapOutput)
}

func (o LookupRelationshipResultOutput) ExpiryDateTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRelationshipResult) *string { return v.ExpiryDateTimeUtc }).(pulumi.StringPtrOutput)
}

func (o LookupRelationshipResultOutput) Fields() PropertyDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupRelationshipResult) []PropertyDefinitionResponse { return v.Fields }).(PropertyDefinitionResponseArrayOutput)
}

func (o LookupRelationshipResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRelationshipResultOutput) LookupMappings() RelationshipTypeMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupRelationshipResult) []RelationshipTypeMappingResponse { return v.LookupMappings }).(RelationshipTypeMappingResponseArrayOutput)
}

func (o LookupRelationshipResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRelationshipResultOutput) ProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.ProfileType }).(pulumi.StringOutput)
}

func (o LookupRelationshipResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRelationshipResultOutput) RelatedProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.RelatedProfileType }).(pulumi.StringOutput)
}

func (o LookupRelationshipResultOutput) RelationshipGuidId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.RelationshipGuidId }).(pulumi.StringOutput)
}

func (o LookupRelationshipResultOutput) RelationshipName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.RelationshipName }).(pulumi.StringOutput)
}

func (o LookupRelationshipResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupRelationshipResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRelationshipResultOutput{})
}
