


package v20200601

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
			Type: pulumi.String("azure-nextgen:web/v20200601:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSiteExtensionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:WebAppSiteExtensionSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppSiteExtensionSlot
	err := ctx.RegisterResource("azure-native:web/v20200601:WebAppSiteExtensionSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppSiteExtensionSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSiteExtensionSlotState, opts ...pulumi.ResourceOption) (*WebAppSiteExtensionSlot, error) {
	var resource WebAppSiteExtensionSlot
	err := ctx.ReadResource("azure-native:web/v20200601:WebAppSiteExtensionSlot", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*WebAppSiteExtensionSlot)(nil))
}

func (i *WebAppSiteExtensionSlot) ToWebAppSiteExtensionSlotOutput() WebAppSiteExtensionSlotOutput {
	return i.ToWebAppSiteExtensionSlotOutputWithContext(context.Background())
}

func (i *WebAppSiteExtensionSlot) ToWebAppSiteExtensionSlotOutputWithContext(ctx context.Context) WebAppSiteExtensionSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSiteExtensionSlotOutput)
}

type WebAppSiteExtensionSlotOutput struct{ *pulumi.OutputState }

func (WebAppSiteExtensionSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppSiteExtensionSlot)(nil))
}

func (o WebAppSiteExtensionSlotOutput) ToWebAppSiteExtensionSlotOutput() WebAppSiteExtensionSlotOutput {
	return o
}

func (o WebAppSiteExtensionSlotOutput) ToWebAppSiteExtensionSlotOutputWithContext(ctx context.Context) WebAppSiteExtensionSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppSiteExtensionSlotOutput{})
}
