


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppSiteExtension(ctx *pulumi.Context, args *LookupWebAppSiteExtensionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSiteExtensionResult, error) {
	var rv LookupWebAppSiteExtensionResult
	err := ctx.Invoke("azure-native:web/v20210101:getWebAppSiteExtension", args, &rv, opts...)
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
