


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SiteMetadataSlot struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringPtrOutput `pulumi:"kind"`
	Location   pulumi.StringOutput    `pulumi:"location"`
	Name       pulumi.StringPtrOutput `pulumi:"name"`
	Properties pulumi.StringMapOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput `pulumi:"tags"`
	Type       pulumi.StringPtrOutput `pulumi:"type"`
}


func NewSiteMetadataSlot(ctx *pulumi.Context,
	name string, args *SiteMetadataSlotArgs, opts ...pulumi.ResourceOption) (*SiteMetadataSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:SiteMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteMetadataSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteMetadataSlot
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteMetadataSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteMetadataSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteMetadataSlotState, opts ...pulumi.ResourceOption) (*SiteMetadataSlot, error) {
	var resource SiteMetadataSlot
	err := ctx.ReadResource("azure-native:web/v20150801:SiteMetadataSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteMetadataSlotState struct {
}

type SiteMetadataSlotState struct {
}

func (SiteMetadataSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteMetadataSlotState)(nil)).Elem()
}

type siteMetadataSlotArgs struct {
	Id                *string           `pulumi:"id"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	Properties        map[string]string `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Slot              string            `pulumi:"slot"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
}


type SiteMetadataSlotArgs struct {
	Id                pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringInput
	Properties        pulumi.StringMapInput
	ResourceGroupName pulumi.StringInput
	Slot              pulumi.StringInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
}

func (SiteMetadataSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteMetadataSlotArgs)(nil)).Elem()
}

type SiteMetadataSlotInput interface {
	pulumi.Input

	ToSiteMetadataSlotOutput() SiteMetadataSlotOutput
	ToSiteMetadataSlotOutputWithContext(ctx context.Context) SiteMetadataSlotOutput
}

func (*SiteMetadataSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteMetadataSlot)(nil))
}

func (i *SiteMetadataSlot) ToSiteMetadataSlotOutput() SiteMetadataSlotOutput {
	return i.ToSiteMetadataSlotOutputWithContext(context.Background())
}

func (i *SiteMetadataSlot) ToSiteMetadataSlotOutputWithContext(ctx context.Context) SiteMetadataSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteMetadataSlotOutput)
}

type SiteMetadataSlotOutput struct{ *pulumi.OutputState }

func (SiteMetadataSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteMetadataSlot)(nil))
}

func (o SiteMetadataSlotOutput) ToSiteMetadataSlotOutput() SiteMetadataSlotOutput {
	return o
}

func (o SiteMetadataSlotOutput) ToSiteMetadataSlotOutputWithContext(ctx context.Context) SiteMetadataSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SiteMetadataSlotOutput{})
}
