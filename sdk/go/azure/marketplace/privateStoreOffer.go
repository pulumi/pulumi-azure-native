


package marketplace

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateStoreOffer struct {
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
	Type                           pulumi.StringOutput      `pulumi:"type"`
	UniqueOfferId                  pulumi.StringOutput      `pulumi:"uniqueOfferId"`
	UpdateSuppressedDueIdempotence pulumi.BoolPtrOutput     `pulumi:"updateSuppressedDueIdempotence"`
}


func NewPrivateStoreOffer(ctx *pulumi.Context,
	name string, args *PrivateStoreOfferArgs, opts ...pulumi.ResourceOption) (*PrivateStoreOffer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateStoreId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateStoreId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:marketplace/v20200101:PrivateStoreOffer"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateStoreOffer
	err := ctx.RegisterResource("azure-native:marketplace:PrivateStoreOffer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateStoreOffer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateStoreOfferState, opts ...pulumi.ResourceOption) (*PrivateStoreOffer, error) {
	var resource PrivateStoreOffer
	err := ctx.ReadResource("azure-native:marketplace:PrivateStoreOffer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateStoreOfferState struct {
}

type PrivateStoreOfferState struct {
}

func (PrivateStoreOfferState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateStoreOfferState)(nil)).Elem()
}

type privateStoreOfferArgs struct {
	ETag                           *string           `pulumi:"eTag"`
	IconFileUris                   map[string]string `pulumi:"iconFileUris"`
	OfferId                        *string           `pulumi:"offerId"`
	Plans                          []Plan            `pulumi:"plans"`
	PrivateStoreId                 string            `pulumi:"privateStoreId"`
	SpecificPlanIdsLimitation      []string          `pulumi:"specificPlanIdsLimitation"`
	UpdateSuppressedDueIdempotence *bool             `pulumi:"updateSuppressedDueIdempotence"`
}


type PrivateStoreOfferArgs struct {
	ETag                           pulumi.StringPtrInput
	IconFileUris                   pulumi.StringMapInput
	OfferId                        pulumi.StringPtrInput
	Plans                          PlanArrayInput
	PrivateStoreId                 pulumi.StringInput
	SpecificPlanIdsLimitation      pulumi.StringArrayInput
	UpdateSuppressedDueIdempotence pulumi.BoolPtrInput
}

func (PrivateStoreOfferArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateStoreOfferArgs)(nil)).Elem()
}

type PrivateStoreOfferInput interface {
	pulumi.Input

	ToPrivateStoreOfferOutput() PrivateStoreOfferOutput
	ToPrivateStoreOfferOutputWithContext(ctx context.Context) PrivateStoreOfferOutput
}

func (*PrivateStoreOffer) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateStoreOffer)(nil))
}

func (i *PrivateStoreOffer) ToPrivateStoreOfferOutput() PrivateStoreOfferOutput {
	return i.ToPrivateStoreOfferOutputWithContext(context.Background())
}

func (i *PrivateStoreOffer) ToPrivateStoreOfferOutputWithContext(ctx context.Context) PrivateStoreOfferOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateStoreOfferOutput)
}

type PrivateStoreOfferOutput struct{ *pulumi.OutputState }

func (PrivateStoreOfferOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateStoreOffer)(nil))
}

func (o PrivateStoreOfferOutput) ToPrivateStoreOfferOutput() PrivateStoreOfferOutput {
	return o
}

func (o PrivateStoreOfferOutput) ToPrivateStoreOfferOutputWithContext(ctx context.Context) PrivateStoreOfferOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateStoreOfferOutput{})
}
