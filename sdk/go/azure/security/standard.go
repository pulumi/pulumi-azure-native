


package security

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Standard struct {
	pulumi.CustomResourceState

	Category     pulumi.StringPtrOutput                         `pulumi:"category"`
	Components   StandardComponentPropertiesResponseArrayOutput `pulumi:"components"`
	Description  pulumi.StringPtrOutput                         `pulumi:"description"`
	DisplayName  pulumi.StringPtrOutput                         `pulumi:"displayName"`
	Etag         pulumi.StringPtrOutput                         `pulumi:"etag"`
	Kind         pulumi.StringPtrOutput                         `pulumi:"kind"`
	Location     pulumi.StringPtrOutput                         `pulumi:"location"`
	Name         pulumi.StringOutput                            `pulumi:"name"`
	StandardType pulumi.StringOutput                            `pulumi:"standardType"`
	SystemData   SystemDataResponseOutput                       `pulumi:"systemData"`
	Tags         pulumi.StringMapOutput                         `pulumi:"tags"`
	Type         pulumi.StringOutput                            `pulumi:"type"`
}


func NewStandard(ctx *pulumi.Context,
	name string, args *StandardArgs, opts ...pulumi.ResourceOption) (*Standard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:security:Standard"),
		},
		{
			Type: pulumi.String("azure-native:security/v20210801preview:Standard"),
		},
		{
			Type: pulumi.String("azure-nextgen:security/v20210801preview:Standard"),
		},
	})
	opts = append(opts, aliases)
	var resource Standard
	err := ctx.RegisterResource("azure-native:security:Standard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStandard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StandardState, opts ...pulumi.ResourceOption) (*Standard, error) {
	var resource Standard
	err := ctx.ReadResource("azure-native:security:Standard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type standardState struct {
}

type StandardState struct {
}

func (StandardState) ElementType() reflect.Type {
	return reflect.TypeOf((*standardState)(nil)).Elem()
}

type standardArgs struct {
	Category          *string                       `pulumi:"category"`
	Components        []StandardComponentProperties `pulumi:"components"`
	Description       *string                       `pulumi:"description"`
	DisplayName       *string                       `pulumi:"displayName"`
	Etag              *string                       `pulumi:"etag"`
	Kind              *string                       `pulumi:"kind"`
	Location          *string                       `pulumi:"location"`
	ResourceGroupName string                        `pulumi:"resourceGroupName"`
	StandardId        *string                       `pulumi:"standardId"`
	Tags              map[string]string             `pulumi:"tags"`
}


type StandardArgs struct {
	Category          pulumi.StringPtrInput
	Components        StandardComponentPropertiesArrayInput
	Description       pulumi.StringPtrInput
	DisplayName       pulumi.StringPtrInput
	Etag              pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	StandardId        pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (StandardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*standardArgs)(nil)).Elem()
}

type StandardInput interface {
	pulumi.Input

	ToStandardOutput() StandardOutput
	ToStandardOutputWithContext(ctx context.Context) StandardOutput
}

func (*Standard) ElementType() reflect.Type {
	return reflect.TypeOf((*Standard)(nil))
}

func (i *Standard) ToStandardOutput() StandardOutput {
	return i.ToStandardOutputWithContext(context.Background())
}

func (i *Standard) ToStandardOutputWithContext(ctx context.Context) StandardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardOutput)
}

type StandardOutput struct{ *pulumi.OutputState }

func (StandardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Standard)(nil))
}

func (o StandardOutput) ToStandardOutput() StandardOutput {
	return o
}

func (o StandardOutput) ToStandardOutputWithContext(ctx context.Context) StandardOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StandardOutput{})
}
