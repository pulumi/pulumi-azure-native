


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IncidentComment struct {
	pulumi.CustomResourceState

	Author              ClientInfoResponseOutput `pulumi:"author"`
	CreatedTimeUtc      pulumi.StringOutput      `pulumi:"createdTimeUtc"`
	Etag                pulumi.StringPtrOutput   `pulumi:"etag"`
	LastModifiedTimeUtc pulumi.StringOutput      `pulumi:"lastModifiedTimeUtc"`
	Message             pulumi.StringOutput      `pulumi:"message"`
	Name                pulumi.StringOutput      `pulumi:"name"`
	SystemData          SystemDataResponseOutput `pulumi:"systemData"`
	Type                pulumi.StringOutput      `pulumi:"type"`
}


func NewIncidentComment(ctx *pulumi.Context,
	name string, args *IncidentCommentArgs, opts ...pulumi.ResourceOption) (*IncidentComment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IncidentId == nil {
		return nil, errors.New("invalid value for required argument 'IncidentId'")
	}
	if args.Message == nil {
		return nil, errors.New("invalid value for required argument 'Message'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210401:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:IncidentComment"),
		},
	})
	opts = append(opts, aliases)
	var resource IncidentComment
	err := ctx.RegisterResource("azure-native:securityinsights/v20221001preview:IncidentComment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIncidentComment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IncidentCommentState, opts ...pulumi.ResourceOption) (*IncidentComment, error) {
	var resource IncidentComment
	err := ctx.ReadResource("azure-native:securityinsights/v20221001preview:IncidentComment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type incidentCommentState struct {
}

type IncidentCommentState struct {
}

func (IncidentCommentState) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentCommentState)(nil)).Elem()
}

type incidentCommentArgs struct {
	IncidentCommentId *string `pulumi:"incidentCommentId"`
	IncidentId        string  `pulumi:"incidentId"`
	Message           string  `pulumi:"message"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type IncidentCommentArgs struct {
	IncidentCommentId pulumi.StringPtrInput
	IncidentId        pulumi.StringInput
	Message           pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (IncidentCommentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentCommentArgs)(nil)).Elem()
}

type IncidentCommentInput interface {
	pulumi.Input

	ToIncidentCommentOutput() IncidentCommentOutput
	ToIncidentCommentOutputWithContext(ctx context.Context) IncidentCommentOutput
}

func (*IncidentComment) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentComment)(nil)).Elem()
}

func (i *IncidentComment) ToIncidentCommentOutput() IncidentCommentOutput {
	return i.ToIncidentCommentOutputWithContext(context.Background())
}

func (i *IncidentComment) ToIncidentCommentOutputWithContext(ctx context.Context) IncidentCommentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentCommentOutput)
}

type IncidentCommentOutput struct{ *pulumi.OutputState }

func (IncidentCommentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentComment)(nil)).Elem()
}

func (o IncidentCommentOutput) ToIncidentCommentOutput() IncidentCommentOutput {
	return o
}

func (o IncidentCommentOutput) ToIncidentCommentOutputWithContext(ctx context.Context) IncidentCommentOutput {
	return o
}

func (o IncidentCommentOutput) Author() ClientInfoResponseOutput {
	return o.ApplyT(func(v *IncidentComment) ClientInfoResponseOutput { return v.Author }).(ClientInfoResponseOutput)
}

func (o IncidentCommentOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentComment) pulumi.StringOutput { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

func (o IncidentCommentOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentComment) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o IncidentCommentOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentComment) pulumi.StringOutput { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

func (o IncidentCommentOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentComment) pulumi.StringOutput { return v.Message }).(pulumi.StringOutput)
}

func (o IncidentCommentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentComment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IncidentCommentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IncidentComment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o IncidentCommentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentComment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IncidentCommentOutput{})
}
