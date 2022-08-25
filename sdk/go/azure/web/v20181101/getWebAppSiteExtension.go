


package v20181101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppSiteExtension(ctx *pulumi.Context, args *LookupWebAppSiteExtensionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSiteExtensionResult, error) {
	var rv LookupWebAppSiteExtensionResult
	err := ctx.Invoke("azure-native:web/v20181101:getWebAppSiteExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSiteExtensionArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SiteExtensionId   string `pulumi:"siteExtensionId"`
}


type LookupWebAppSiteExtensionResult struct {
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

func LookupWebAppSiteExtensionOutput(ctx *pulumi.Context, args LookupWebAppSiteExtensionOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppSiteExtensionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppSiteExtensionResult, error) {
			args := v.(LookupWebAppSiteExtensionArgs)
			r, err := LookupWebAppSiteExtension(ctx, &args, opts...)
			var s LookupWebAppSiteExtensionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppSiteExtensionResultOutput)
}

type LookupWebAppSiteExtensionOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SiteExtensionId   pulumi.StringInput `pulumi:"siteExtensionId"`
}

func (LookupWebAppSiteExtensionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSiteExtensionArgs)(nil)).Elem()
}


type LookupWebAppSiteExtensionResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppSiteExtensionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSiteExtensionResult)(nil)).Elem()
}

func (o LookupWebAppSiteExtensionResultOutput) ToLookupWebAppSiteExtensionResultOutput() LookupWebAppSiteExtensionResultOutput {
	return o
}

func (o LookupWebAppSiteExtensionResultOutput) ToLookupWebAppSiteExtensionResultOutputWithContext(ctx context.Context) LookupWebAppSiteExtensionResultOutput {
	return o
}

func (o LookupWebAppSiteExtensionResultOutput) Authors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) []string { return v.Authors }).(pulumi.StringArrayOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) DownloadCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *int { return v.DownloadCount }).(pulumi.IntPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) ExtensionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.ExtensionId }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) ExtensionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.ExtensionType }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) ExtensionUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.ExtensionUrl }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) FeedUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.FeedUrl }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.IconUrl }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) InstalledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.InstalledDateTime }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) InstallerCommandLineParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.InstallerCommandLineParams }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) LicenseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.LicenseUrl }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) LocalIsLatestVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *bool { return v.LocalIsLatestVersion }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) LocalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.LocalPath }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) ProjectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.ProjectUrl }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) PublishedDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.PublishedDateTime }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) Summary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Summary }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppSiteExtensionResultOutput{})
}
