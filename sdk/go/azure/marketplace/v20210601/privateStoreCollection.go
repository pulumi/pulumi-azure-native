


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateStoreCollection struct {
	pulumi.CustomResourceState

	AllSubscriptions  pulumi.BoolPtrOutput     `pulumi:"allSubscriptions"`
	Claim             pulumi.StringPtrOutput   `pulumi:"claim"`
	CollectionId      pulumi.StringOutput      `pulumi:"collectionId"`
	CollectionName    pulumi.StringPtrOutput   `pulumi:"collectionName"`
	Enabled           pulumi.BoolPtrOutput     `pulumi:"enabled"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	NumberOfOffers    pulumi.Float64Output     `pulumi:"numberOfOffers"`
	SubscriptionsList pulumi.StringArrayOutput `pulumi:"subscriptionsList"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewPrivateStoreCollection(ctx *pulumi.Context,
	name string, args *PrivateStoreCollectionArgs, opts ...pulumi.ResourceOption) (*PrivateStoreCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateStoreId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateStoreId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:marketplace:PrivateStoreCollection"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateStoreCollection
	err := ctx.RegisterResource("azure-native:marketplace/v20210601:PrivateStoreCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateStoreCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateStoreCollectionState, opts ...pulumi.ResourceOption) (*PrivateStoreCollection, error) {
	var resource PrivateStoreCollection
	err := ctx.ReadResource("azure-native:marketplace/v20210601:PrivateStoreCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateStoreCollectionState struct {
}

type PrivateStoreCollectionState struct {
}

func (PrivateStoreCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateStoreCollectionState)(nil)).Elem()
}

type privateStoreCollectionArgs struct {
	AllSubscriptions  *bool    `pulumi:"allSubscriptions"`
	Claim             *string  `pulumi:"claim"`
	CollectionId      *string  `pulumi:"collectionId"`
	CollectionName    *string  `pulumi:"collectionName"`
	Enabled           *bool    `pulumi:"enabled"`
	PrivateStoreId    string   `pulumi:"privateStoreId"`
	SubscriptionsList []string `pulumi:"subscriptionsList"`
}


type PrivateStoreCollectionArgs struct {
	AllSubscriptions  pulumi.BoolPtrInput
	Claim             pulumi.StringPtrInput
	CollectionId      pulumi.StringPtrInput
	CollectionName    pulumi.StringPtrInput
	Enabled           pulumi.BoolPtrInput
	PrivateStoreId    pulumi.StringInput
	SubscriptionsList pulumi.StringArrayInput
}

func (PrivateStoreCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateStoreCollectionArgs)(nil)).Elem()
}

type PrivateStoreCollectionInput interface {
	pulumi.Input

	ToPrivateStoreCollectionOutput() PrivateStoreCollectionOutput
	ToPrivateStoreCollectionOutputWithContext(ctx context.Context) PrivateStoreCollectionOutput
}

func (*PrivateStoreCollection) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateStoreCollection)(nil))
}

func (i *PrivateStoreCollection) ToPrivateStoreCollectionOutput() PrivateStoreCollectionOutput {
	return i.ToPrivateStoreCollectionOutputWithContext(context.Background())
}

func (i *PrivateStoreCollection) ToPrivateStoreCollectionOutputWithContext(ctx context.Context) PrivateStoreCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateStoreCollectionOutput)
}

type PrivateStoreCollectionOutput struct{ *pulumi.OutputState }

func (PrivateStoreCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateStoreCollection)(nil))
}

func (o PrivateStoreCollectionOutput) ToPrivateStoreCollectionOutput() PrivateStoreCollectionOutput {
	return o
}

func (o PrivateStoreCollectionOutput) ToPrivateStoreCollectionOutputWithContext(ctx context.Context) PrivateStoreCollectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateStoreCollectionOutput{})
}
