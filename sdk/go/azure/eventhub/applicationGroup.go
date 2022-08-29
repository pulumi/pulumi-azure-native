


package eventhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationGroup struct {
	pulumi.CustomResourceState

	ClientAppGroupIdentifier pulumi.StringOutput                 `pulumi:"clientAppGroupIdentifier"`
	IsEnabled                pulumi.BoolPtrOutput                `pulumi:"isEnabled"`
	Location                 pulumi.StringOutput                 `pulumi:"location"`
	Name                     pulumi.StringOutput                 `pulumi:"name"`
	Policies                 ThrottlingPolicyResponseArrayOutput `pulumi:"policies"`
	SystemData               SystemDataResponseOutput            `pulumi:"systemData"`
	Type                     pulumi.StringOutput                 `pulumi:"type"`
}


func NewApplicationGroup(ctx *pulumi.Context,
	name string, args *ApplicationGroupArgs, opts ...pulumi.ResourceOption) (*ApplicationGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientAppGroupIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'ClientAppGroupIdentifier'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventhub/v20220101preview:ApplicationGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationGroup
	err := ctx.RegisterResource("azure-native:eventhub:ApplicationGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplicationGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationGroupState, opts ...pulumi.ResourceOption) (*ApplicationGroup, error) {
	var resource ApplicationGroup
	err := ctx.ReadResource("azure-native:eventhub:ApplicationGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationGroupState struct {
}

type ApplicationGroupState struct {
}

func (ApplicationGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGroupState)(nil)).Elem()
}

type applicationGroupArgs struct {
	ApplicationGroupName     *string            `pulumi:"applicationGroupName"`
	ClientAppGroupIdentifier string             `pulumi:"clientAppGroupIdentifier"`
	IsEnabled                *bool              `pulumi:"isEnabled"`
	NamespaceName            string             `pulumi:"namespaceName"`
	Policies                 []ThrottlingPolicy `pulumi:"policies"`
	ResourceGroupName        string             `pulumi:"resourceGroupName"`
}


type ApplicationGroupArgs struct {
	ApplicationGroupName     pulumi.StringPtrInput
	ClientAppGroupIdentifier pulumi.StringInput
	IsEnabled                pulumi.BoolPtrInput
	NamespaceName            pulumi.StringInput
	Policies                 ThrottlingPolicyArrayInput
	ResourceGroupName        pulumi.StringInput
}

func (ApplicationGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGroupArgs)(nil)).Elem()
}

type ApplicationGroupInput interface {
	pulumi.Input

	ToApplicationGroupOutput() ApplicationGroupOutput
	ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput
}

func (*ApplicationGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroup)(nil)).Elem()
}

func (i *ApplicationGroup) ToApplicationGroupOutput() ApplicationGroupOutput {
	return i.ToApplicationGroupOutputWithContext(context.Background())
}

func (i *ApplicationGroup) ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGroupOutput)
}

type ApplicationGroupOutput struct{ *pulumi.OutputState }

func (ApplicationGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroup)(nil)).Elem()
}

func (o ApplicationGroupOutput) ToApplicationGroupOutput() ApplicationGroupOutput {
	return o
}

func (o ApplicationGroupOutput) ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput {
	return o
}

func (o ApplicationGroupOutput) ClientAppGroupIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.ClientAppGroupIdentifier }).(pulumi.StringOutput)
}

func (o ApplicationGroupOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.BoolPtrOutput { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ApplicationGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationGroupOutput) Policies() ThrottlingPolicyResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGroup) ThrottlingPolicyResponseArrayOutput { return v.Policies }).(ThrottlingPolicyResponseArrayOutput)
}

func (o ApplicationGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApplicationGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ApplicationGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationGroupOutput{})
}
