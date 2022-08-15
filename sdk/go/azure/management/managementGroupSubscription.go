


package management

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagementGroupSubscription struct {
	pulumi.CustomResourceState

	DisplayName pulumi.StringPtrOutput                     `pulumi:"displayName"`
	Name        pulumi.StringOutput                        `pulumi:"name"`
	Parent      DescendantParentGroupInfoResponsePtrOutput `pulumi:"parent"`
	State       pulumi.StringPtrOutput                     `pulumi:"state"`
	Tenant      pulumi.StringPtrOutput                     `pulumi:"tenant"`
	Type        pulumi.StringOutput                        `pulumi:"type"`
}


func NewManagementGroupSubscription(ctx *pulumi.Context,
	name string, args *ManagementGroupSubscriptionArgs, opts ...pulumi.ResourceOption) (*ManagementGroupSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:management/v20200501:ManagementGroupSubscription"),
		},
		{
			Type: pulumi.String("azure-native:management/v20201001:ManagementGroupSubscription"),
		},
		{
			Type: pulumi.String("azure-native:management/v20210401:ManagementGroupSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementGroupSubscription
	err := ctx.RegisterResource("azure-native:management:ManagementGroupSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagementGroupSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementGroupSubscriptionState, opts ...pulumi.ResourceOption) (*ManagementGroupSubscription, error) {
	var resource ManagementGroupSubscription
	err := ctx.ReadResource("azure-native:management:ManagementGroupSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managementGroupSubscriptionState struct {
}

type ManagementGroupSubscriptionState struct {
}

func (ManagementGroupSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupSubscriptionState)(nil)).Elem()
}

type managementGroupSubscriptionArgs struct {
	GroupId        string  `pulumi:"groupId"`
	SubscriptionId *string `pulumi:"subscriptionId"`
}


type ManagementGroupSubscriptionArgs struct {
	GroupId        pulumi.StringInput
	SubscriptionId pulumi.StringPtrInput
}

func (ManagementGroupSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupSubscriptionArgs)(nil)).Elem()
}

type ManagementGroupSubscriptionInput interface {
	pulumi.Input

	ToManagementGroupSubscriptionOutput() ManagementGroupSubscriptionOutput
	ToManagementGroupSubscriptionOutputWithContext(ctx context.Context) ManagementGroupSubscriptionOutput
}

func (*ManagementGroupSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroupSubscription)(nil)).Elem()
}

func (i *ManagementGroupSubscription) ToManagementGroupSubscriptionOutput() ManagementGroupSubscriptionOutput {
	return i.ToManagementGroupSubscriptionOutputWithContext(context.Background())
}

func (i *ManagementGroupSubscription) ToManagementGroupSubscriptionOutputWithContext(ctx context.Context) ManagementGroupSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupSubscriptionOutput)
}

type ManagementGroupSubscriptionOutput struct{ *pulumi.OutputState }

func (ManagementGroupSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroupSubscription)(nil)).Elem()
}

func (o ManagementGroupSubscriptionOutput) ToManagementGroupSubscriptionOutput() ManagementGroupSubscriptionOutput {
	return o
}

func (o ManagementGroupSubscriptionOutput) ToManagementGroupSubscriptionOutputWithContext(ctx context.Context) ManagementGroupSubscriptionOutput {
	return o
}

func (o ManagementGroupSubscriptionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupSubscription) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagementGroupSubscriptionOutput) Parent() DescendantParentGroupInfoResponsePtrOutput {
	return o.ApplyT(func(v *ManagementGroupSubscription) DescendantParentGroupInfoResponsePtrOutput { return v.Parent }).(DescendantParentGroupInfoResponsePtrOutput)
}

func (o ManagementGroupSubscriptionOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupSubscription) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupSubscriptionOutput) Tenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupSubscription) pulumi.StringPtrOutput { return v.Tenant }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementGroupSubscriptionOutput{})
}
