


package v20170101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupRelationshipLink(ctx *pulumi.Context, args *LookupRelationshipLinkArgs, opts ...pulumi.InvokeOption) (*LookupRelationshipLinkResult, error) {
	var rv LookupRelationshipLinkResult
	err := ctx.Invoke("azure-native:customerinsights/v20170101:getRelationshipLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRelationshipLinkArgs struct {
	HubName              string `pulumi:"hubName"`
	RelationshipLinkName string `pulumi:"relationshipLinkName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupRelationshipLinkResult struct {
	Description                      map[string]string                      `pulumi:"description"`
	DisplayName                      map[string]string                      `pulumi:"displayName"`
	Id                               string                                 `pulumi:"id"`
	InteractionType                  string                                 `pulumi:"interactionType"`
	LinkName                         string                                 `pulumi:"linkName"`
	Mappings                         []RelationshipLinkFieldMappingResponse `pulumi:"mappings"`
	Name                             string                                 `pulumi:"name"`
	ProfilePropertyReferences        []ParticipantPropertyReferenceResponse `pulumi:"profilePropertyReferences"`
	ProvisioningState                string                                 `pulumi:"provisioningState"`
	RelatedProfilePropertyReferences []ParticipantPropertyReferenceResponse `pulumi:"relatedProfilePropertyReferences"`
	RelationshipGuidId               string                                 `pulumi:"relationshipGuidId"`
	RelationshipName                 string                                 `pulumi:"relationshipName"`
	TenantId                         string                                 `pulumi:"tenantId"`
	Type                             string                                 `pulumi:"type"`
}

func LookupRelationshipLinkOutput(ctx *pulumi.Context, args LookupRelationshipLinkOutputArgs, opts ...pulumi.InvokeOption) LookupRelationshipLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRelationshipLinkResult, error) {
			args := v.(LookupRelationshipLinkArgs)
			r, err := LookupRelationshipLink(ctx, &args, opts...)
			var s LookupRelationshipLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRelationshipLinkResultOutput)
}

type LookupRelationshipLinkOutputArgs struct {
	HubName              pulumi.StringInput `pulumi:"hubName"`
	RelationshipLinkName pulumi.StringInput `pulumi:"relationshipLinkName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRelationshipLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRelationshipLinkArgs)(nil)).Elem()
}


type LookupRelationshipLinkResultOutput struct{ *pulumi.OutputState }

func (LookupRelationshipLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRelationshipLinkResult)(nil)).Elem()
}

func (o LookupRelationshipLinkResultOutput) ToLookupRelationshipLinkResultOutput() LookupRelationshipLinkResultOutput {
	return o
}

func (o LookupRelationshipLinkResultOutput) ToLookupRelationshipLinkResultOutputWithContext(ctx context.Context) LookupRelationshipLinkResultOutput {
	return o
}

func (o LookupRelationshipLinkResultOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) map[string]string { return v.Description }).(pulumi.StringMapOutput)
}

func (o LookupRelationshipLinkResultOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) map[string]string { return v.DisplayName }).(pulumi.StringMapOutput)
}

func (o LookupRelationshipLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRelationshipLinkResultOutput) InteractionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.InteractionType }).(pulumi.StringOutput)
}

func (o LookupRelationshipLinkResultOutput) LinkName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.LinkName }).(pulumi.StringOutput)
}

func (o LookupRelationshipLinkResultOutput) Mappings() RelationshipLinkFieldMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) []RelationshipLinkFieldMappingResponse { return v.Mappings }).(RelationshipLinkFieldMappingResponseArrayOutput)
}

func (o LookupRelationshipLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRelationshipLinkResultOutput) ProfilePropertyReferences() ParticipantPropertyReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) []ParticipantPropertyReferenceResponse {
		return v.ProfilePropertyReferences
	}).(ParticipantPropertyReferenceResponseArrayOutput)
}

func (o LookupRelationshipLinkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRelationshipLinkResultOutput) RelatedProfilePropertyReferences() ParticipantPropertyReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) []ParticipantPropertyReferenceResponse {
		return v.RelatedProfilePropertyReferences
	}).(ParticipantPropertyReferenceResponseArrayOutput)
}

func (o LookupRelationshipLinkResultOutput) RelationshipGuidId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.RelationshipGuidId }).(pulumi.StringOutput)
}

func (o LookupRelationshipLinkResultOutput) RelationshipName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.RelationshipName }).(pulumi.StringOutput)
}

func (o LookupRelationshipLinkResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupRelationshipLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRelationshipLinkResultOutput{})
}
