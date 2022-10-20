


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIncidentRelation(ctx *pulumi.Context, args *LookupIncidentRelationArgs, opts ...pulumi.InvokeOption) (*LookupIncidentRelationResult, error) {
	var rv LookupIncidentRelationResult
	err := ctx.Invoke("azure-native:securityinsights/v20220901preview:getIncidentRelation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIncidentRelationArgs struct {
	IncidentId        string `pulumi:"incidentId"`
	RelationName      string `pulumi:"relationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupIncidentRelationResult struct {
	Etag                *string            `pulumi:"etag"`
	Id                  string             `pulumi:"id"`
	Name                string             `pulumi:"name"`
	RelatedResourceId   string             `pulumi:"relatedResourceId"`
	RelatedResourceKind string             `pulumi:"relatedResourceKind"`
	RelatedResourceName string             `pulumi:"relatedResourceName"`
	RelatedResourceType string             `pulumi:"relatedResourceType"`
	SystemData          SystemDataResponse `pulumi:"systemData"`
	Type                string             `pulumi:"type"`
}

func LookupIncidentRelationOutput(ctx *pulumi.Context, args LookupIncidentRelationOutputArgs, opts ...pulumi.InvokeOption) LookupIncidentRelationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIncidentRelationResult, error) {
			args := v.(LookupIncidentRelationArgs)
			r, err := LookupIncidentRelation(ctx, &args, opts...)
			var s LookupIncidentRelationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIncidentRelationResultOutput)
}

type LookupIncidentRelationOutputArgs struct {
	IncidentId        pulumi.StringInput `pulumi:"incidentId"`
	RelationName      pulumi.StringInput `pulumi:"relationName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIncidentRelationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIncidentRelationArgs)(nil)).Elem()
}


type LookupIncidentRelationResultOutput struct{ *pulumi.OutputState }

func (LookupIncidentRelationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIncidentRelationResult)(nil)).Elem()
}

func (o LookupIncidentRelationResultOutput) ToLookupIncidentRelationResultOutput() LookupIncidentRelationResultOutput {
	return o
}

func (o LookupIncidentRelationResultOutput) ToLookupIncidentRelationResultOutputWithContext(ctx context.Context) LookupIncidentRelationResultOutput {
	return o
}

func (o LookupIncidentRelationResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupIncidentRelationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIncidentRelationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIncidentRelationResultOutput) RelatedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) string { return v.RelatedResourceId }).(pulumi.StringOutput)
}

func (o LookupIncidentRelationResultOutput) RelatedResourceKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) string { return v.RelatedResourceKind }).(pulumi.StringOutput)
}

func (o LookupIncidentRelationResultOutput) RelatedResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) string { return v.RelatedResourceName }).(pulumi.StringOutput)
}

func (o LookupIncidentRelationResultOutput) RelatedResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) string { return v.RelatedResourceType }).(pulumi.StringOutput)
}

func (o LookupIncidentRelationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupIncidentRelationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIncidentRelationResultOutput{})
}
