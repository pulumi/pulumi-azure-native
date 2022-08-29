


package v20170101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupLink(ctx *pulumi.Context, args *LookupLinkArgs, opts ...pulumi.InvokeOption) (*LookupLinkResult, error) {
	var rv LookupLinkResult
	err := ctx.Invoke("azure-native:customerinsights/v20170101:getLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkArgs struct {
	HubName           string `pulumi:"hubName"`
	LinkName          string `pulumi:"linkName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLinkResult struct {
	Description                   map[string]string                      `pulumi:"description"`
	DisplayName                   map[string]string                      `pulumi:"displayName"`
	Id                            string                                 `pulumi:"id"`
	LinkName                      string                                 `pulumi:"linkName"`
	Mappings                      []TypePropertiesMappingResponse        `pulumi:"mappings"`
	Name                          string                                 `pulumi:"name"`
	OperationType                 *string                                `pulumi:"operationType"`
	ParticipantPropertyReferences []ParticipantPropertyReferenceResponse `pulumi:"participantPropertyReferences"`
	ProvisioningState             string                                 `pulumi:"provisioningState"`
	ReferenceOnly                 *bool                                  `pulumi:"referenceOnly"`
	SourceInteractionType         string                                 `pulumi:"sourceInteractionType"`
	TargetProfileType             string                                 `pulumi:"targetProfileType"`
	TenantId                      string                                 `pulumi:"tenantId"`
	Type                          string                                 `pulumi:"type"`
}

func LookupLinkOutput(ctx *pulumi.Context, args LookupLinkOutputArgs, opts ...pulumi.InvokeOption) LookupLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLinkResult, error) {
			args := v.(LookupLinkArgs)
			r, err := LookupLink(ctx, &args, opts...)
			var s LookupLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLinkResultOutput)
}

type LookupLinkOutputArgs struct {
	HubName           pulumi.StringInput `pulumi:"hubName"`
	LinkName          pulumi.StringInput `pulumi:"linkName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkArgs)(nil)).Elem()
}


type LookupLinkResultOutput struct{ *pulumi.OutputState }

func (LookupLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkResult)(nil)).Elem()
}

func (o LookupLinkResultOutput) ToLookupLinkResultOutput() LookupLinkResultOutput {
	return o
}

func (o LookupLinkResultOutput) ToLookupLinkResultOutputWithContext(ctx context.Context) LookupLinkResultOutput {
	return o
}

func (o LookupLinkResultOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLinkResult) map[string]string { return v.Description }).(pulumi.StringMapOutput)
}

func (o LookupLinkResultOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLinkResult) map[string]string { return v.DisplayName }).(pulumi.StringMapOutput)
}

func (o LookupLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLinkResultOutput) LinkName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkResult) string { return v.LinkName }).(pulumi.StringOutput)
}

func (o LookupLinkResultOutput) Mappings() TypePropertiesMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupLinkResult) []TypePropertiesMappingResponse { return v.Mappings }).(TypePropertiesMappingResponseArrayOutput)
}

func (o LookupLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLinkResultOutput) OperationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLinkResult) *string { return v.OperationType }).(pulumi.StringPtrOutput)
}

func (o LookupLinkResultOutput) ParticipantPropertyReferences() ParticipantPropertyReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupLinkResult) []ParticipantPropertyReferenceResponse {
		return v.ParticipantPropertyReferences
	}).(ParticipantPropertyReferenceResponseArrayOutput)
}

func (o LookupLinkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupLinkResultOutput) ReferenceOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLinkResult) *bool { return v.ReferenceOnly }).(pulumi.BoolPtrOutput)
}

func (o LookupLinkResultOutput) SourceInteractionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkResult) string { return v.SourceInteractionType }).(pulumi.StringOutput)
}

func (o LookupLinkResultOutput) TargetProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkResult) string { return v.TargetProfileType }).(pulumi.StringOutput)
}

func (o LookupLinkResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLinkResultOutput{})
}
