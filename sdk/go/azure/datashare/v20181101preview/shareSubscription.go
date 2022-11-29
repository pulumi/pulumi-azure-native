


package v20181101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ShareSubscription struct {
	pulumi.CustomResourceState

	CreatedAt               pulumi.StringOutput `pulumi:"createdAt"`
	InvitationId            pulumi.StringOutput `pulumi:"invitationId"`
	Name                    pulumi.StringOutput `pulumi:"name"`
	ProviderEmail           pulumi.StringOutput `pulumi:"providerEmail"`
	ProviderName            pulumi.StringOutput `pulumi:"providerName"`
	ProviderTenantName      pulumi.StringOutput `pulumi:"providerTenantName"`
	ProvisioningState       pulumi.StringOutput `pulumi:"provisioningState"`
	ShareDescription        pulumi.StringOutput `pulumi:"shareDescription"`
	ShareKind               pulumi.StringOutput `pulumi:"shareKind"`
	ShareName               pulumi.StringOutput `pulumi:"shareName"`
	ShareSubscriptionStatus pulumi.StringOutput `pulumi:"shareSubscriptionStatus"`
	ShareTerms              pulumi.StringOutput `pulumi:"shareTerms"`
	Type                    pulumi.StringOutput `pulumi:"type"`
	UserEmail               pulumi.StringOutput `pulumi:"userEmail"`
	UserName                pulumi.StringOutput `pulumi:"userName"`
}


func NewShareSubscription(ctx *pulumi.Context,
	name string, args *ShareSubscriptionArgs, opts ...pulumi.ResourceOption) (*ShareSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.InvitationId == nil {
		return nil, errors.New("invalid value for required argument 'InvitationId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:ShareSubscription"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ShareSubscription"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ShareSubscription"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:ShareSubscription"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:ShareSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource ShareSubscription
	err := ctx.RegisterResource("azure-native:datashare/v20181101preview:ShareSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetShareSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShareSubscriptionState, opts ...pulumi.ResourceOption) (*ShareSubscription, error) {
	var resource ShareSubscription
	err := ctx.ReadResource("azure-native:datashare/v20181101preview:ShareSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type shareSubscriptionState struct {
}

type ShareSubscriptionState struct {
}

func (ShareSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*shareSubscriptionState)(nil)).Elem()
}

type shareSubscriptionArgs struct {
	AccountName           string  `pulumi:"accountName"`
	InvitationId          string  `pulumi:"invitationId"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName *string `pulumi:"shareSubscriptionName"`
}


type ShareSubscriptionArgs struct {
	AccountName           pulumi.StringInput
	InvitationId          pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	ShareSubscriptionName pulumi.StringPtrInput
}

func (ShareSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shareSubscriptionArgs)(nil)).Elem()
}

type ShareSubscriptionInput interface {
	pulumi.Input

	ToShareSubscriptionOutput() ShareSubscriptionOutput
	ToShareSubscriptionOutputWithContext(ctx context.Context) ShareSubscriptionOutput
}

func (*ShareSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**ShareSubscription)(nil)).Elem()
}

func (i *ShareSubscription) ToShareSubscriptionOutput() ShareSubscriptionOutput {
	return i.ToShareSubscriptionOutputWithContext(context.Background())
}

func (i *ShareSubscription) ToShareSubscriptionOutputWithContext(ctx context.Context) ShareSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareSubscriptionOutput)
}

type ShareSubscriptionOutput struct{ *pulumi.OutputState }

func (ShareSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShareSubscription)(nil)).Elem()
}

func (o ShareSubscriptionOutput) ToShareSubscriptionOutput() ShareSubscriptionOutput {
	return o
}

func (o ShareSubscriptionOutput) ToShareSubscriptionOutputWithContext(ctx context.Context) ShareSubscriptionOutput {
	return o
}

func (o ShareSubscriptionOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareSubscription) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o ShareSubscriptionOutput) InvitationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareSubscription) pulumi.StringOutput { return v.InvitationId }).(pulumi.StringOutput)
}

func (o ShareSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ShareSubscriptionOutput) ProviderEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareSubscription) pulumi.StringOutput { return v.ProviderEmail }).(pulumi.StringOutput)
}

func (o ShareSubscriptionOutput) ProviderName() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareSubscription) pulumi.StringOutput { return v.ProviderName }).(pulumi.StringOutput)
}

func (o ShareSubscriptionOutput) ProviderTenantName() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareSubscription) pulumi.StringOutput { return v.ProviderTenantName }).(pulumi.StringOutput)
}

func (o ShareSubscriptionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareSubscription) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ShareSubscriptionOutput) ShareDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareSubscription) pulumi.StringOutput { return v.ShareDescription }).(pulumi.StringOutput)
}

func (o ShareSubscriptionOutput) ShareKind() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareSubscription) pulumi.StringOutput { return v.ShareKind }).(pulumi.StringOutput)
}

func (o ShareSubscriptionOutput) ShareName() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareSubscription) pulumi.StringOutput { return v.ShareName }).(pulumi.StringOutput)
}

func (o ShareSubscriptionOutput) ShareSubscriptionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareSubscription) pulumi.StringOutput { return v.ShareSubscriptionStatus }).(pulumi.StringOutput)
}

func (o ShareSubscriptionOutput) ShareTerms() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareSubscription) pulumi.StringOutput { return v.ShareTerms }).(pulumi.StringOutput)
}

func (o ShareSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ShareSubscriptionOutput) UserEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareSubscription) pulumi.StringOutput { return v.UserEmail }).(pulumi.StringOutput)
}

func (o ShareSubscriptionOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareSubscription) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ShareSubscriptionOutput{})
}
