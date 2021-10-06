


package v20201001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppSiteExtensionSlot(ctx *pulumi.Context, args *LookupWebAppSiteExtensionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSiteExtensionSlotResult, error) {
	var rv LookupWebAppSiteExtensionSlotResult
	err := ctx.Invoke("azure-native:web/v20201001:getWebAppSiteExtensionSlot", args, &rv, opts...)
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
	Authors                    []string           `pulumi:"authors"`
	Comment                    *string            `pulumi:"comment"`
	Description                *string            `pulumi:"description"`
	DownloadCount              *int               `pulumi:"downloadCount"`
	ExtensionId                *string            `pulumi:"extensionId"`
	ExtensionType              *string            `pulumi:"extensionType"`
	ExtensionUrl               *string            `pulumi:"extensionUrl"`
	FeedUrl                    *string            `pulumi:"feedUrl"`
	IconUrl                    *string            `pulumi:"iconUrl"`
	Id                         string             `pulumi:"id"`
	InstalledDateTime          *string            `pulumi:"installedDateTime"`
	InstallerCommandLineParams *string            `pulumi:"installerCommandLineParams"`
	Kind                       *string            `pulumi:"kind"`
	LicenseUrl                 *string            `pulumi:"licenseUrl"`
	LocalIsLatestVersion       *bool              `pulumi:"localIsLatestVersion"`
	LocalPath                  *string            `pulumi:"localPath"`
	Name                       string             `pulumi:"name"`
	ProjectUrl                 *string            `pulumi:"projectUrl"`
	ProvisioningState          *string            `pulumi:"provisioningState"`
	PublishedDateTime          *string            `pulumi:"publishedDateTime"`
	Summary                    *string            `pulumi:"summary"`
	SystemData                 SystemDataResponse `pulumi:"systemData"`
	Title                      *string            `pulumi:"title"`
	Type                       string             `pulumi:"type"`
	Version                    *string            `pulumi:"version"`
}
