


package v20210115

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
			Type: pulumi.String("azure-nextgen:web/v20210115:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:WebAppSiteExtension"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppSiteExtension
	err := ctx.RegisterResource("azure-native:web/v20210115:WebAppSiteExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppSiteExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSiteExtensionState, opts ...pulumi.ResourceOption) (*WebAppSiteExtension, error) {
	var resource WebAppSiteExtension
	err := ctx.ReadResource("azure-native:web/v20210115:WebAppSiteExtension", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*WebAppSiteExtension)(nil))
}

func (i *WebAppSiteExtension) ToWebAppSiteExtensionOutput() WebAppSiteExtensionOutput {
	return i.ToWebAppSiteExtensionOutputWithContext(context.Background())
}

func (i *WebAppSiteExtension) ToWebAppSiteExtensionOutputWithContext(ctx context.Context) WebAppSiteExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSiteExtensionOutput)
}

type WebAppSiteExtensionOutput struct{ *pulumi.OutputState }

func (WebAppSiteExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppSiteExtension)(nil))
}

func (o WebAppSiteExtensionOutput) ToWebAppSiteExtensionOutput() WebAppSiteExtensionOutput {
	return o
}

func (o WebAppSiteExtensionOutput) ToWebAppSiteExtensionOutputWithContext(ctx context.Context) WebAppSiteExtensionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebAppSiteExtensionInput)(nil)).Elem(), &WebAppSiteExtension{})
	pulumi.RegisterOutputType(WebAppSiteExtensionOutput{})
}
