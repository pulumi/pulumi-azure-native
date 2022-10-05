


package v20201001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppSiteExtension struct {
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
	SystemData                 SystemDataResponseOutput `pulumi:"systemData"`
	Title                      pulumi.StringPtrOutput   `pulumi:"title"`
	Type                       pulumi.StringOutput      `pulumi:"type"`
	Version                    pulumi.StringPtrOutput   `pulumi:"version"`
}


func NewWebAppSiteExtension(ctx *pulumi.Context,
	name string, args *WebAppSiteExtensionArgs, opts ...pulumi.ResourceOption) (*WebAppSiteExtension, error) {
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
			Type: pulumi.String("azure-native:web:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppSiteExtension"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppSiteExtension
	err := ctx.RegisterResource("azure-native:web/v20201001:WebAppSiteExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppSiteExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSiteExtensionState, opts ...pulumi.ResourceOption) (*WebAppSiteExtension, error) {
	var resource WebAppSiteExtension
	err := ctx.ReadResource("azure-native:web/v20201001:WebAppSiteExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppSiteExtensionState struct {
}

type WebAppSiteExtensionState struct {
}

func (WebAppSiteExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSiteExtensionState)(nil)).Elem()
}

type webAppSiteExtensionArgs struct {
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SiteExtensionId   *string `pulumi:"siteExtensionId"`
}


type WebAppSiteExtensionArgs struct {
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SiteExtensionId   pulumi.StringPtrInput
}

func (WebAppSiteExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSiteExtensionArgs)(nil)).Elem()
}

type WebAppSiteExtensionInput interface {
	pulumi.Input

	ToWebAppSiteExtensionOutput() WebAppSiteExtensionOutput
	ToWebAppSiteExtensionOutputWithContext(ctx context.Context) WebAppSiteExtensionOutput
}

func (*WebAppSiteExtension) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSiteExtension)(nil)).Elem()
}

func (i *WebAppSiteExtension) ToWebAppSiteExtensionOutput() WebAppSiteExtensionOutput {
	return i.ToWebAppSiteExtensionOutputWithContext(context.Background())
}

func (i *WebAppSiteExtension) ToWebAppSiteExtensionOutputWithContext(ctx context.Context) WebAppSiteExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSiteExtensionOutput)
}

type WebAppSiteExtensionOutput struct{ *pulumi.OutputState }

func (WebAppSiteExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSiteExtension)(nil)).Elem()
}

func (o WebAppSiteExtensionOutput) ToWebAppSiteExtensionOutput() WebAppSiteExtensionOutput {
	return o
}

func (o WebAppSiteExtensionOutput) ToWebAppSiteExtensionOutputWithContext(ctx context.Context) WebAppSiteExtensionOutput {
	return o
}

func (o WebAppSiteExtensionOutput) Authors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringArrayOutput { return v.Authors }).(pulumi.StringArrayOutput)
}

func (o WebAppSiteExtensionOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) DownloadCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.IntPtrOutput { return v.DownloadCount }).(pulumi.IntPtrOutput)
}

func (o WebAppSiteExtensionOutput) ExtensionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.ExtensionId }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) ExtensionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.ExtensionType }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) ExtensionUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.ExtensionUrl }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) FeedUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.FeedUrl }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.IconUrl }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) InstalledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.InstalledDateTime }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) InstallerCommandLineParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.InstallerCommandLineParams }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) LicenseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.LicenseUrl }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) LocalIsLatestVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.BoolPtrOutput { return v.LocalIsLatestVersion }).(pulumi.BoolPtrOutput)
}

func (o WebAppSiteExtensionOutput) LocalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.LocalPath }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppSiteExtensionOutput) ProjectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.ProjectUrl }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) PublishedDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.PublishedDateTime }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) Summary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.Summary }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WebAppSiteExtensionOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.Title }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WebAppSiteExtensionOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppSiteExtensionOutput{})
}
