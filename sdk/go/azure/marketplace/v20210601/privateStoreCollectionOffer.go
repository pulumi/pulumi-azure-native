


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateStoreCollectionOffer struct {
	pulumi.CustomResourceState

	CreatedAt                      pulumi.StringOutput      `pulumi:"createdAt"`
	ETag                           pulumi.StringPtrOutput   `pulumi:"eTag"`
	IconFileUris                   pulumi.StringMapOutput   `pulumi:"iconFileUris"`
	ModifiedAt                     pulumi.StringOutput      `pulumi:"modifiedAt"`
	Name                           pulumi.StringOutput      `pulumi:"name"`
	OfferDisplayName               pulumi.StringOutput      `pulumi:"offerDisplayName"`
	Plans                          PlanResponseArrayOutput  `pulumi:"plans"`
	PrivateStoreId                 pulumi.StringOutput      `pulumi:"privateStoreId"`
	PublisherDisplayName           pulumi.StringOutput      `pulumi:"publisherDisplayName"`
	SpecificPlanIdsLimitation      pulumi.StringArrayOutput `pulumi:"specificPlanIdsLimitation"`
	SystemData                     SystemDataResponseOutput `pulumi:"systemData"`
	Type                           pulumi.StringOutput      `pulumi:"type"`
	UniqueOfferId                  pulumi.StringOutput      `pulumi:"uniqueOfferId"`
	UpdateSuppressedDueIdempotence pulumi.BoolPtrOutput     `pulumi:"updateSuppressedDueIdempotence"`
}


func NewPrivateStoreCollectionOffer(ctx *pulumi.Context,
	name string, args *PrivateStoreCollectionOfferArgs, opts ...pulumi.ResourceOption) (*PrivateStoreCollectionOffer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CollectionId == nil {
		return nil, errors.New("invalid value for required argument 'CollectionId'")
	}
	if args.PrivateStoreId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateStoreId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:marketplace/v20210601:PrivateStoreCollectionOffer"),
		},
		{
			Type: pulumi.String("azure-native:marketplace:PrivateStoreCollectionOffer"),
		},
		{
			Type: pulumi.String("azure-nextgen:marketplace:PrivateStoreCollectionOffer"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateStoreCollectionOffer
	err := ctx.RegisterResource("azure-native:marketplace/v20210601:PrivateStoreCollectionOffer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateStoreCollectionOffer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateStoreCollectionOfferState, opts ...pulumi.ResourceOption) (*PrivateStoreCollectionOffer, error) {
	var resource PrivateStoreCollectionOffer
	err := ctx.ReadResource("azure-native:marketplace/v20210601:PrivateStoreCollectionOffer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateStoreCollectionOfferState struct {
}

type PrivateStoreCollectionOfferState struct {
}

func (PrivateStoreCollectionOfferState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateStoreCollectionOfferState)(nil)).Elem()
}

type privateStoreCollectionOfferArgs struct {
	CollectionId                   string            `pulumi:"collectionId"`
	ETag                           *string           `pulumi:"eTag"`
	IconFileUris                   map[string]string `pulumi:"iconFileUris"`
	OfferId                        *string           `pulumi:"offerId"`
	Plans                          []Plan            `pulumi:"plans"`
	PrivateStoreId                 string            `pulumi:"privateStoreId"`
	SpecificPlanIdsLimitation      []string          `pulumi:"specificPlanIdsLimitation"`
	UpdateSuppressedDueIdempotence *bool             `pulumi:"updateSuppressedDueIdempotence"`
}


type PrivateStoreCollectionOfferArgs struct {
	CollectionId                   pulumi.StringInput
	ETag                           pulumi.StringPtrInput
	IconFileUris                   pulumi.StringMapInput
	OfferId                        pulumi.StringPtrInput
	Plans                          PlanArrayInput
	PrivateStoreId                 pulumi.StringInput
	SpecificPlanIdsLimitation      pulumi.StringArrayInput
	UpdateSuppressedDueIdempotence pulumi.BoolPtrInput
}

func (PrivateStoreCollectionOfferArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateStoreCollectionOfferArgs)(nil)).Elem()
}

type PrivateStoreCollectionOfferInput interface {
	pulumi.Input

	ToPrivateStoreCollectionOfferOutput() PrivateStoreCollectionOfferOutput
	ToPrivateStoreCollectionOfferOutputWithContext(ctx context.Context) PrivateStoreCollectionOfferOutput
}

func (*PrivateStoreCollectionOffer) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateStoreCollectionOffer)(nil))
}

func (i *PrivateStoreCollectionOffer) ToPrivateStoreCollectionOfferOutput() PrivateStoreCollectionOfferOutput {
	return i.ToPrivateStoreCollectionOfferOutputWithContext(context.Background())
}

func (i *PrivateStoreCollectionOffer) ToPrivateStoreCollectionOfferOutputWithContext(ctx context.Context) PrivateStoreCollectionOfferOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateStoreCollectionOfferOutput)
}

type PrivateStoreCollectionOfferOutput struct{ *pulumi.OutputState }

func (PrivateStoreCollectionOfferOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateStoreCollectionOffer)(nil))
}

func (o PrivateStoreCollectionOfferOutput) ToPrivateStoreCollectionOfferOutput() PrivateStoreCollectionOfferOutput {
	return o
}

func (o PrivateStoreCollectionOfferOutput) ToPrivateStoreCollectionOfferOutputWithContext(ctx context.Context) PrivateStoreCollectionOfferOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateStoreCollectionOfferInput)(nil)).Elem(), &PrivateStoreCollectionOffer{})
	pulumi.RegisterOutputType(PrivateStoreCollectionOfferOutput{})
}
