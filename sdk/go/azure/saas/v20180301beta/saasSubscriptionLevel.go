


package v20180301beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SaasSubscriptionLevel struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                  `pulumi:"name"`
	Properties SaasResourceResponsePropertiesOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput               `pulumi:"tags"`
	Type       pulumi.StringOutput                  `pulumi:"type"`
}


func NewSaasSubscriptionLevel(ctx *pulumi.Context,
	name string, args *SaasSubscriptionLevelArgs, opts ...pulumi.ResourceOption) (*SaasSubscriptionLevel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:saas:SaasSubscriptionLevel"),
		},
	})
	opts = append(opts, aliases)
	var resource SaasSubscriptionLevel
	err := ctx.RegisterResource("azure-native:saas/v20180301beta:SaasSubscriptionLevel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSaasSubscriptionLevel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SaasSubscriptionLevelState, opts ...pulumi.ResourceOption) (*SaasSubscriptionLevel, error) {
	var resource SaasSubscriptionLevel
	err := ctx.ReadResource("azure-native:saas/v20180301beta:SaasSubscriptionLevel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type saasSubscriptionLevelState struct {
}

type SaasSubscriptionLevelState struct {
}

func (SaasSubscriptionLevelState) ElementType() reflect.Type {
	return reflect.TypeOf((*saasSubscriptionLevelState)(nil)).Elem()
}

type saasSubscriptionLevelArgs struct {
	Location          *string                 `pulumi:"location"`
	Name              *string                 `pulumi:"name"`
	Properties        *SaasCreationProperties `pulumi:"properties"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	ResourceName      *string                 `pulumi:"resourceName"`
	Tags              map[string]string       `pulumi:"tags"`
}


type SaasSubscriptionLevelArgs struct {
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	Properties        SaasCreationPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (SaasSubscriptionLevelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*saasSubscriptionLevelArgs)(nil)).Elem()
}

type SaasSubscriptionLevelInput interface {
	pulumi.Input

	ToSaasSubscriptionLevelOutput() SaasSubscriptionLevelOutput
	ToSaasSubscriptionLevelOutputWithContext(ctx context.Context) SaasSubscriptionLevelOutput
}

func (*SaasSubscriptionLevel) ElementType() reflect.Type {
	return reflect.TypeOf((**SaasSubscriptionLevel)(nil)).Elem()
}

func (i *SaasSubscriptionLevel) ToSaasSubscriptionLevelOutput() SaasSubscriptionLevelOutput {
	return i.ToSaasSubscriptionLevelOutputWithContext(context.Background())
}

func (i *SaasSubscriptionLevel) ToSaasSubscriptionLevelOutputWithContext(ctx context.Context) SaasSubscriptionLevelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SaasSubscriptionLevelOutput)
}

type SaasSubscriptionLevelOutput struct{ *pulumi.OutputState }

func (SaasSubscriptionLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SaasSubscriptionLevel)(nil)).Elem()
}

func (o SaasSubscriptionLevelOutput) ToSaasSubscriptionLevelOutput() SaasSubscriptionLevelOutput {
	return o
}

func (o SaasSubscriptionLevelOutput) ToSaasSubscriptionLevelOutputWithContext(ctx context.Context) SaasSubscriptionLevelOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SaasSubscriptionLevelOutput{})
}
