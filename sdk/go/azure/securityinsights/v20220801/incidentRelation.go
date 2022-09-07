


package v20220801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IncidentRelation struct {
	pulumi.CustomResourceState

	Etag                pulumi.StringPtrOutput   `pulumi:"etag"`
	Name                pulumi.StringOutput      `pulumi:"name"`
	RelatedResourceId   pulumi.StringOutput      `pulumi:"relatedResourceId"`
	RelatedResourceKind pulumi.StringOutput      `pulumi:"relatedResourceKind"`
	RelatedResourceName pulumi.StringOutput      `pulumi:"relatedResourceName"`
	RelatedResourceType pulumi.StringOutput      `pulumi:"relatedResourceType"`
	SystemData          SystemDataResponseOutput `pulumi:"systemData"`
	Type                pulumi.StringOutput      `pulumi:"type"`
}


func NewIncidentRelation(ctx *pulumi.Context,
	name string, args *IncidentRelationArgs, opts ...pulumi.ResourceOption) (*IncidentRelation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IncidentId == nil {
		return nil, errors.New("invalid value for required argument 'IncidentId'")
	}
	if args.RelatedResourceId == nil {
		return nil, errors.New("invalid value for required argument 'RelatedResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210401:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:IncidentRelation"),
		},
	})
	opts = append(opts, aliases)
	var resource IncidentRelation
	err := ctx.RegisterResource("azure-native:securityinsights/v20220801:IncidentRelation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIncidentRelation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IncidentRelationState, opts ...pulumi.ResourceOption) (*IncidentRelation, error) {
	var resource IncidentRelation
	err := ctx.ReadResource("azure-native:securityinsights/v20220801:IncidentRelation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type incidentRelationState struct {
}

type IncidentRelationState struct {
}

func (IncidentRelationState) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentRelationState)(nil)).Elem()
}

type incidentRelationArgs struct {
	IncidentId        string  `pulumi:"incidentId"`
	RelatedResourceId string  `pulumi:"relatedResourceId"`
	RelationName      *string `pulumi:"relationName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type IncidentRelationArgs struct {
	IncidentId        pulumi.StringInput
	RelatedResourceId pulumi.StringInput
	RelationName      pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (IncidentRelationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentRelationArgs)(nil)).Elem()
}

type IncidentRelationInput interface {
	pulumi.Input

	ToIncidentRelationOutput() IncidentRelationOutput
	ToIncidentRelationOutputWithContext(ctx context.Context) IncidentRelationOutput
}

func (*IncidentRelation) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentRelation)(nil)).Elem()
}

func (i *IncidentRelation) ToIncidentRelationOutput() IncidentRelationOutput {
	return i.ToIncidentRelationOutputWithContext(context.Background())
}

func (i *IncidentRelation) ToIncidentRelationOutputWithContext(ctx context.Context) IncidentRelationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentRelationOutput)
}

type IncidentRelationOutput struct{ *pulumi.OutputState }

func (IncidentRelationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentRelation)(nil)).Elem()
}

func (o IncidentRelationOutput) ToIncidentRelationOutput() IncidentRelationOutput {
	return o
}

func (o IncidentRelationOutput) ToIncidentRelationOutputWithContext(ctx context.Context) IncidentRelationOutput {
	return o
}

func (o IncidentRelationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentRelation) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o IncidentRelationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentRelation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IncidentRelationOutput) RelatedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentRelation) pulumi.StringOutput { return v.RelatedResourceId }).(pulumi.StringOutput)
}

func (o IncidentRelationOutput) RelatedResourceKind() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentRelation) pulumi.StringOutput { return v.RelatedResourceKind }).(pulumi.StringOutput)
}

func (o IncidentRelationOutput) RelatedResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentRelation) pulumi.StringOutput { return v.RelatedResourceName }).(pulumi.StringOutput)
}

func (o IncidentRelationOutput) RelatedResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentRelation) pulumi.StringOutput { return v.RelatedResourceType }).(pulumi.StringOutput)
}

func (o IncidentRelationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IncidentRelation) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o IncidentRelationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentRelation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IncidentRelationOutput{})
}
