


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SiteMetadata struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringPtrOutput `pulumi:"kind"`
	Location   pulumi.StringOutput    `pulumi:"location"`
	Name       pulumi.StringPtrOutput `pulumi:"name"`
	Properties pulumi.StringMapOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput `pulumi:"tags"`
	Type       pulumi.StringPtrOutput `pulumi:"type"`
}


func NewSiteMetadata(ctx *pulumi.Context,
	name string, args *SiteMetadataArgs, opts ...pulumi.ResourceOption) (*SiteMetadata, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteMetadata"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:SiteMetadata"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteMetadata
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteMetadata", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteMetadata(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteMetadataState, opts ...pulumi.ResourceOption) (*SiteMetadata, error) {
	var resource SiteMetadata
	err := ctx.ReadResource("azure-native:web/v20150801:SiteMetadata", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteMetadataState struct {
}

type SiteMetadataState struct {
}

func (SiteMetadataState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteMetadataState)(nil)).Elem()
}

type siteMetadataArgs struct {
	Id                *string           `pulumi:"id"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	Properties        map[string]string `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
}


type SiteMetadataArgs struct {
	Id                pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringInput
	Properties        pulumi.StringMapInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
}

func (SiteMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteMetadataArgs)(nil)).Elem()
}

type SiteMetadataInput interface {
	pulumi.Input

	ToSiteMetadataOutput() SiteMetadataOutput
	ToSiteMetadataOutputWithContext(ctx context.Context) SiteMetadataOutput
}

func (*SiteMetadata) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteMetadata)(nil))
}

func (i *SiteMetadata) ToSiteMetadataOutput() SiteMetadataOutput {
	return i.ToSiteMetadataOutputWithContext(context.Background())
}

func (i *SiteMetadata) ToSiteMetadataOutputWithContext(ctx context.Context) SiteMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteMetadataOutput)
}

type SiteMetadataOutput struct{ *pulumi.OutputState }

func (SiteMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteMetadata)(nil))
}

func (o SiteMetadataOutput) ToSiteMetadataOutput() SiteMetadataOutput {
	return o
}

func (o SiteMetadataOutput) ToSiteMetadataOutputWithContext(ctx context.Context) SiteMetadataOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SiteMetadataInput)(nil)).Elem(), &SiteMetadata{})
	pulumi.RegisterOutputType(SiteMetadataOutput{})
}
