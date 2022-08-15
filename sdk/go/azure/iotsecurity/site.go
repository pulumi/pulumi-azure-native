


package iotsecurity

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Site struct {
	pulumi.CustomResourceState

	DisplayName pulumi.StringOutput      `pulumi:"displayName"`
	Name        pulumi.StringOutput      `pulumi:"name"`
	SystemData  SystemDataResponseOutput `pulumi:"systemData"`
	Tags        pulumi.StringMapOutput   `pulumi:"tags"`
	Type        pulumi.StringOutput      `pulumi:"type"`
}


func NewSite(ctx *pulumi.Context,
	name string, args *SiteArgs, opts ...pulumi.ResourceOption) (*Site, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:iotsecurity/v20210201preview:Site"),
		},
	})
	opts = append(opts, aliases)
	var resource Site
	err := ctx.RegisterResource("azure-native:iotsecurity:Site", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteState, opts ...pulumi.ResourceOption) (*Site, error) {
	var resource Site
	err := ctx.ReadResource("azure-native:iotsecurity:Site", name, id, state, &resource, opts...)
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
	DisplayName string            `pulumi:"displayName"`
	Scope       string            `pulumi:"scope"`
	Tags        map[string]string `pulumi:"tags"`
}


type SiteArgs struct {
	DisplayName pulumi.StringInput
	Scope       pulumi.StringInput
	Tags        pulumi.StringMapInput
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

func (o SiteOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o SiteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SiteOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Site) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
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
