


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Site struct {
	pulumi.CustomResourceState

	CreatedAt          pulumi.StringPtrOutput         `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrOutput         `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrOutput         `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrOutput         `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrOutput         `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrOutput         `pulumi:"lastModifiedByType"`
	Location           pulumi.StringOutput            `pulumi:"location"`
	Name               pulumi.StringOutput            `pulumi:"name"`
	NetworkFunctions   SubResourceResponseArrayOutput `pulumi:"networkFunctions"`
	ProvisioningState  pulumi.StringOutput            `pulumi:"provisioningState"`
	Tags               pulumi.StringMapOutput         `pulumi:"tags"`
	Type               pulumi.StringOutput            `pulumi:"type"`
}


func NewSite(ctx *pulumi.Context,
	name string, args *SiteArgs, opts ...pulumi.ResourceOption) (*Site, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MobileNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'MobileNetworkName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork:Site"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220301preview:Site"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220401preview:Site"),
		},
	})
	opts = append(opts, aliases)
	var resource Site
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20220101preview:Site", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteState, opts ...pulumi.ResourceOption) (*Site, error) {
	var resource Site
	err := ctx.ReadResource("azure-native:mobilenetwork/v20220101preview:Site", name, id, state, &resource, opts...)
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
	CreatedAt          *string           `pulumi:"createdAt"`
	CreatedBy          *string           `pulumi:"createdBy"`
	CreatedByType      *string           `pulumi:"createdByType"`
	LastModifiedAt     *string           `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string           `pulumi:"lastModifiedBy"`
	LastModifiedByType *string           `pulumi:"lastModifiedByType"`
	Location           *string           `pulumi:"location"`
	MobileNetworkName  string            `pulumi:"mobileNetworkName"`
	NetworkFunctions   []SubResource     `pulumi:"networkFunctions"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	SiteName           *string           `pulumi:"siteName"`
	Tags               map[string]string `pulumi:"tags"`
}


type SiteArgs struct {
	CreatedAt          pulumi.StringPtrInput
	CreatedBy          pulumi.StringPtrInput
	CreatedByType      pulumi.StringPtrInput
	LastModifiedAt     pulumi.StringPtrInput
	LastModifiedBy     pulumi.StringPtrInput
	LastModifiedByType pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	MobileNetworkName  pulumi.StringInput
	NetworkFunctions   SubResourceArrayInput
	ResourceGroupName  pulumi.StringInput
	SiteName           pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
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
	return reflect.TypeOf((**Site)(nil)).Elem()
}

func (i *Site) ToSiteOutput() SiteOutput {
	return i.ToSiteOutputWithContext(context.Background())
}

func (i *Site) ToSiteOutputWithContext(ctx context.Context) SiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteOutput)
}

type SiteOutput struct{ *pulumi.OutputState }

func (SiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Site)(nil)).Elem()
}

func (o SiteOutput) ToSiteOutput() SiteOutput {
	return o
}

func (o SiteOutput) ToSiteOutputWithContext(ctx context.Context) SiteOutput {
	return o
}

func (o SiteOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SiteOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SiteOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SiteOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SiteOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SiteOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o SiteOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SiteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SiteOutput) NetworkFunctions() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *Site) SubResourceResponseArrayOutput { return v.NetworkFunctions }).(SubResourceResponseArrayOutput)
}

func (o SiteOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SiteOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Site) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SiteOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteOutput{})
}
