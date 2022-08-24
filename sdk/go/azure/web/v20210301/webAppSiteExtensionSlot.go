


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppSiteExtensionSlot struct {
	pulumi.CustomResourceState

	Authors                    pulumi.StringArrayOutput `pulumi:"authors"`
	Comment                    pulumi.StringPtrOutput   `pulumi:"comment"`
	Description                pulumi.StringPtrOutput   `pulumi:"description"`
	DownloadCount              pulumi.IntPtrOutput      `pulumi:"downloadCount"`
	ExtensionId                pulumi.StringPtrOutput   `pulumi:"extensionId"`
	ExtensionType              pulumi.StringPtrOutput   `pulumi:"extensionType"`
	ExtensionUrl               pulumi.StringPtrOutput   `pulumi:"extensionUrl"`
	FeedUrl                    pulumi.StringPtrOutput   `pulumi:"feedUrl"`
	IconUrl                    pulumi.StringPtrOutput   `pulumi:"iconUrl"`
	InstalledDateTime          pulumi.StringPtrOutput   `pulumi:"installedDateTime"`
	InstallerCommandLineParams pulumi.StringPtrOutput   `pulumi:"installerCommandLineParams"`
	Kind                       pulumi.StringPtrOutput   `pulumi:"kind"`
	LicenseUrl                 pulumi.StringPtrOutput   `pulumi:"licenseUrl"`
	LocalIsLatestVersion       pulumi.BoolPtrOutput     `pulumi:"localIsLatestVersion"`
	LocalPath                  pulumi.StringPtrOutput   `pulumi:"localPath"`
	Name                       pulumi.StringOutput      `pulumi:"name"`
	ProjectUrl                 pulumi.StringPtrOutput   `pulumi:"projectUrl"`
	ProvisioningState          pulumi.StringPtrOutput   `pulumi:"provisioningState"`
	PublishedDateTime          pulumi.StringPtrOutput   `pulumi:"publishedDateTime"`
	Summary                    pulumi.StringPtrOutput   `pulumi:"summary"`
	Title                      pulumi.StringPtrOutput   `pulumi:"title"`
	Type                       pulumi.StringOutput      `pulumi:"type"`
	Version                    pulumi.StringPtrOutput   `pulumi:"version"`
}


func NewWebAppSiteExtensionSlot(ctx *pulumi.Context,
	name string, args *WebAppSiteExtensionSlotArgs, opts ...pulumi.ResourceOption) (*WebAppSiteExtensionSlot, error) {
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
			Type: pulumi.String("azure-native:web:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppSiteExtensionSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppSiteExtensionSlot
	err := ctx.RegisterResource("azure-native:web/v20210301:WebAppSiteExtensionSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppSiteExtensionSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSiteExtensionSlotState, opts ...pulumi.ResourceOption) (*WebAppSiteExtensionSlot, error) {
	var resource WebAppSiteExtensionSlot
	err := ctx.ReadResource("azure-native:web/v20210301:WebAppSiteExtensionSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppSiteExtensionSlotState struct {
}

type WebAppSiteExtensionSlotState struct {
}

func (WebAppSiteExtensionSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSiteExtensionSlotState)(nil)).Elem()
}

type webAppSiteExtensionSlotArgs struct {
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SiteExtensionId   *string `pulumi:"siteExtensionId"`
	Slot              string  `pulumi:"slot"`
}


type WebAppSiteExtensionSlotArgs struct {
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SiteExtensionId   pulumi.StringPtrInput
	Slot              pulumi.StringInput
}

func (WebAppSiteExtensionSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSiteExtensionSlotArgs)(nil)).Elem()
}

type WebAppSiteExtensionSlotInput interface {
	pulumi.Input

	ToWebAppSiteExtensionSlotOutput() WebAppSiteExtensionSlotOutput
	ToWebAppSiteExtensionSlotOutputWithContext(ctx context.Context) WebAppSiteExtensionSlotOutput
}

func (*WebAppSiteExtensionSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSiteExtensionSlot)(nil)).Elem()
}

func (i *WebAppSiteExtensionSlot) ToWebAppSiteExtensionSlotOutput() WebAppSiteExtensionSlotOutput {
	return i.ToWebAppSiteExtensionSlotOutputWithContext(context.Background())
}

func (i *WebAppSiteExtensionSlot) ToWebAppSiteExtensionSlotOutputWithContext(ctx context.Context) WebAppSiteExtensionSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSiteExtensionSlotOutput)
}

type WebAppSiteExtensionSlotOutput struct{ *pulumi.OutputState }

func (WebAppSiteExtensionSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSiteExtensionSlot)(nil)).Elem()
}

func (o WebAppSiteExtensionSlotOutput) ToWebAppSiteExtensionSlotOutput() WebAppSiteExtensionSlotOutput {
	return o
}

func (o WebAppSiteExtensionSlotOutput) ToWebAppSiteExtensionSlotOutputWithContext(ctx context.Context) WebAppSiteExtensionSlotOutput {
	return o
}

func (o WebAppSiteExtensionSlotOutput) Authors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringArrayOutput { return v.Authors }).(pulumi.StringArrayOutput)
}

func (o WebAppSiteExtensionSlotOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) DownloadCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.IntPtrOutput { return v.DownloadCount }).(pulumi.IntPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) ExtensionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.ExtensionId }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) ExtensionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.ExtensionType }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) ExtensionUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.ExtensionUrl }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) FeedUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.FeedUrl }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.IconUrl }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) InstalledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.InstalledDateTime }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) InstallerCommandLineParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.InstallerCommandLineParams }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) LicenseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.LicenseUrl }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) LocalIsLatestVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.BoolPtrOutput { return v.LocalIsLatestVersion }).(pulumi.BoolPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) LocalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.LocalPath }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppSiteExtensionSlotOutput) ProjectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.ProjectUrl }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) PublishedDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.PublishedDateTime }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) Summary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.Summary }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.Title }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WebAppSiteExtensionSlotOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtensionSlot) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppSiteExtensionSlotOutput{})
}
