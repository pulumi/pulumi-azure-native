


package migrate

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Assessment struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput             `pulumi:"eTag"`
	Name       pulumi.StringOutput                `pulumi:"name"`
	Properties AssessmentPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                `pulumi:"type"`
}


func NewAssessment(ctx *pulumi.Context,
	name string, args *AssessmentArgs, opts ...pulumi.ResourceOption) (*Assessment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20191001:Assessment"),
		},
	})
	opts = append(opts, aliases)
	var resource Assessment
	err := ctx.RegisterResource("azure-native:migrate:Assessment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAssessment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssessmentState, opts ...pulumi.ResourceOption) (*Assessment, error) {
	var resource Assessment
	err := ctx.ReadResource("azure-native:migrate:Assessment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type assessmentState struct {
}

type AssessmentState struct {
}

func (AssessmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentState)(nil)).Elem()
}

type assessmentArgs struct {
	AssessmentName    *string              `pulumi:"assessmentName"`
	ETag              *string              `pulumi:"eTag"`
	GroupName         string               `pulumi:"groupName"`
	ProjectName       string               `pulumi:"projectName"`
	Properties        AssessmentProperties `pulumi:"properties"`
	ResourceGroupName string               `pulumi:"resourceGroupName"`
}


type AssessmentArgs struct {
	AssessmentName    pulumi.StringPtrInput
	ETag              pulumi.StringPtrInput
	GroupName         pulumi.StringInput
	ProjectName       pulumi.StringInput
	Properties        AssessmentPropertiesInput
	ResourceGroupName pulumi.StringInput
}

func (AssessmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentArgs)(nil)).Elem()
}

type AssessmentInput interface {
	pulumi.Input

	ToAssessmentOutput() AssessmentOutput
	ToAssessmentOutputWithContext(ctx context.Context) AssessmentOutput
}

func (*Assessment) ElementType() reflect.Type {
	return reflect.TypeOf((**Assessment)(nil)).Elem()
}

func (i *Assessment) ToAssessmentOutput() AssessmentOutput {
	return i.ToAssessmentOutputWithContext(context.Background())
}

func (i *Assessment) ToAssessmentOutputWithContext(ctx context.Context) AssessmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentOutput)
}

type AssessmentOutput struct{ *pulumi.OutputState }

func (AssessmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Assessment)(nil)).Elem()
}

func (o AssessmentOutput) ToAssessmentOutput() AssessmentOutput {
	return o
}

func (o AssessmentOutput) ToAssessmentOutputWithContext(ctx context.Context) AssessmentOutput {
	return o
}

func (o AssessmentOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o AssessmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AssessmentOutput) Properties() AssessmentPropertiesResponseOutput {
	return o.ApplyT(func(v *Assessment) AssessmentPropertiesResponseOutput { return v.Properties }).(AssessmentPropertiesResponseOutput)
}

func (o AssessmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AssessmentOutput{})
}
