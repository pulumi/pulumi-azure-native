


package v20200707

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Site struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput       `pulumi:"eTag"`
	Location   pulumi.StringPtrOutput       `pulumi:"location"`
	Name       pulumi.StringPtrOutput       `pulumi:"name"`
	Properties SitePropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput     `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput       `pulumi:"tags"`
	Type       pulumi.StringOutput          `pulumi:"type"`
}


func NewSite(ctx *pulumi.Context,
	name string, args *SiteArgs, opts ...pulumi.ResourceOption) (*Site, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:offazure/v20200707:Site"),
		},
		{
			Type: pulumi.String("azure-native:offazure:Site"),
		},
		{
			Type: pulumi.String("azure-nextgen:offazure:Site"),
		},
		{
			Type: pulumi.String("azure-native:offazure/v20200101:Site"),
		},
		{
			Type: pulumi.String("azure-nextgen:offazure/v20200101:Site"),
		},
	})
	opts = append(opts, aliases)
	var resource Site
	err := ctx.RegisterResource("azure-native:offazure/v20200707:Site", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteState, opts ...pulumi.ResourceOption) (*Site, error) {
	var resource Site
	err := ctx.ReadResource("azure-native:offazure/v20200707:Site", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteState struct {
}

type SiteState struct {
}

func (SiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteState)(nil)).Elem()
}

type siteArgs struct {
	ETag              *string           `pulumi:"eTag"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	Properties        *SiteProperties   `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	SiteName          *string           `pulumi:"siteName"`
	Tags              map[string]string `pulumi:"tags"`
}


type SiteArgs struct {
	ETag              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	Properties        SitePropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	SiteName          pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (SiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteArgs)(nil)).Elem()
}

type SiteInput interface {
	pulumi.Input

	ToSiteOutput() SiteOutput
	ToSiteOutputWithContext(ctx context.Context) SiteOutput
}

func (*Site) ElementType() reflect.Type {
	return reflect.TypeOf((*Site)(nil))
}

func (i *Site) ToSiteOutput() SiteOutput {
	return i.ToSiteOutputWithContext(context.Background())
}

func (i *Site) ToSiteOutputWithContext(ctx context.Context) SiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteOutput)
}

type SiteOutput struct{ *pulumi.OutputState }

func (SiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Site)(nil))
}

func (o SiteOutput) ToSiteOutput() SiteOutput {
	return o
}

func (o SiteOutput) ToSiteOutputWithContext(ctx context.Context) SiteOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SiteInput)(nil)).Elem(), &Site{})
	pulumi.RegisterOutputType(SiteOutput{})
}
