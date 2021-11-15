


package v20200101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HyperVSite struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput       `pulumi:"eTag"`
	Location   pulumi.StringPtrOutput       `pulumi:"location"`
	Name       pulumi.StringPtrOutput       `pulumi:"name"`
	Properties SitePropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput       `pulumi:"tags"`
	Type       pulumi.StringOutput          `pulumi:"type"`
}


func NewHyperVSite(ctx *pulumi.Context,
	name string, args *HyperVSiteArgs, opts ...pulumi.ResourceOption) (*HyperVSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:offazure:HyperVSite"),
		},
		{
			Type: pulumi.String("azure-native:offazure/v20200707:HyperVSite"),
		},
	})
	opts = append(opts, aliases)
	var resource HyperVSite
	err := ctx.RegisterResource("azure-native:offazure/v20200101:HyperVSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHyperVSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HyperVSiteState, opts ...pulumi.ResourceOption) (*HyperVSite, error) {
	var resource HyperVSite
	err := ctx.ReadResource("azure-native:offazure/v20200101:HyperVSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hyperVSiteState struct {
}

type HyperVSiteState struct {
}

func (HyperVSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*hyperVSiteState)(nil)).Elem()
}

type hyperVSiteArgs struct {
	ETag              *string           `pulumi:"eTag"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	Properties        *SiteProperties   `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	SiteName          *string           `pulumi:"siteName"`
	Tags              map[string]string `pulumi:"tags"`
}


type HyperVSiteArgs struct {
	ETag              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	Properties        SitePropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	SiteName          pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (HyperVSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hyperVSiteArgs)(nil)).Elem()
}

type HyperVSiteInput interface {
	pulumi.Input

	ToHyperVSiteOutput() HyperVSiteOutput
	ToHyperVSiteOutputWithContext(ctx context.Context) HyperVSiteOutput
}

func (*HyperVSite) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVSite)(nil))
}

func (i *HyperVSite) ToHyperVSiteOutput() HyperVSiteOutput {
	return i.ToHyperVSiteOutputWithContext(context.Background())
}

func (i *HyperVSite) ToHyperVSiteOutputWithContext(ctx context.Context) HyperVSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVSiteOutput)
}

type HyperVSiteOutput struct{ *pulumi.OutputState }

func (HyperVSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVSite)(nil))
}

func (o HyperVSiteOutput) ToHyperVSiteOutput() HyperVSiteOutput {
	return o
}

func (o HyperVSiteOutput) ToHyperVSiteOutputWithContext(ctx context.Context) HyperVSiteOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HyperVSiteOutput{})
}
