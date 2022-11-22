


package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppSiteExtensionSlot(ctx *pulumi.Context, args *LookupWebAppSiteExtensionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSiteExtensionSlotResult, error) {
	var rv LookupWebAppSiteExtensionSlotResult
	err := ctx.Invoke("azure-native:web:getWebAppSiteExtensionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSiteExtensionSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SiteExtensionId   string `pulumi:"siteExtensionId"`
	Slot              string `pulumi:"slot"`
}


type LookupWebAppSiteExtensionSlotResult struct {
	Authors                    []string `pulumi:"authors"`
	Comment                    *string  `pulumi:"comment"`
	Description                *string  `pulumi:"description"`
	DownloadCount              *int     `pulumi:"downloadCount"`
	ExtensionId                *string  `pulumi:"extensionId"`
	ExtensionType              *string  `pulumi:"extensionType"`
	ExtensionUrl               *string  `pulumi:"extensionUrl"`
	FeedUrl                    *string  `pulumi:"feedUrl"`
	IconUrl                    *string  `pulumi:"iconUrl"`
	Id                         string   `pulumi:"id"`
	InstalledDateTime          *string  `pulumi:"installedDateTime"`
	InstallerCommandLineParams *string  `pulumi:"installerCommandLineParams"`
	Kind                       *string  `pulumi:"kind"`
	LicenseUrl                 *string  `pulumi:"licenseUrl"`
	LocalIsLatestVersion       *bool    `pulumi:"localIsLatestVersion"`
	LocalPath                  *string  `pulumi:"localPath"`
	Name                       string   `pulumi:"name"`
	ProjectUrl                 *string  `pulumi:"projectUrl"`
	ProvisioningState          *string  `pulumi:"provisioningState"`
	PublishedDateTime          *string  `pulumi:"publishedDateTime"`
	Summary                    *string  `pulumi:"summary"`
	Title                      *string  `pulumi:"title"`
	Type                       string   `pulumi:"type"`
	Version                    *string  `pulumi:"version"`
}

func LookupWebAppSiteExtensionSlotOutput(ctx *pulumi.Context, args LookupWebAppSiteExtensionSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppSiteExtensionSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppSiteExtensionSlotResult, error) {
			args := v.(LookupWebAppSiteExtensionSlotArgs)
			r, err := LookupWebAppSiteExtensionSlot(ctx, &args, opts...)
			var s LookupWebAppSiteExtensionSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppSiteExtensionSlotResultOutput)
}

type LookupWebAppSiteExtensionSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SiteExtensionId   pulumi.StringInput `pulumi:"siteExtensionId"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppSiteExtensionSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSiteExtensionSlotArgs)(nil)).Elem()
}


type LookupWebAppSiteExtensionSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppSiteExtensionSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSiteExtensionSlotResult)(nil)).Elem()
}

func (o LookupWebAppSiteExtensionSlotResultOutput) ToLookupWebAppSiteExtensionSlotResultOutput() LookupWebAppSiteExtensionSlotResultOutput {
	return o
}

func (o LookupWebAppSiteExtensionSlotResultOutput) ToLookupWebAppSiteExtensionSlotResultOutputWithContext(ctx context.Context) LookupWebAppSiteExtensionSlotResultOutput {
	return o
}

func (o LookupWebAppSiteExtensionSlotResultOutput) Authors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) []string { return v.Authors }).(pulumi.StringArrayOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) DownloadCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *int { return v.DownloadCount }).(pulumi.IntPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) ExtensionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.ExtensionId }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) ExtensionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.ExtensionType }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) ExtensionUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.ExtensionUrl }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) FeedUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.FeedUrl }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.IconUrl }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) InstalledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.InstalledDateTime }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) InstallerCommandLineParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.InstallerCommandLineParams }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) LicenseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.LicenseUrl }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) LocalIsLatestVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *bool { return v.LocalIsLatestVersion }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) LocalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.LocalPath }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) ProjectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.ProjectUrl }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) PublishedDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.PublishedDateTime }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) Summary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.Summary }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppSiteExtensionSlotResultOutput{})
}
