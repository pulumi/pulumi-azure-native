


package v20210712

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20210712:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	ApplicationGroupName string `pulumi:"applicationGroupName"`
	ApplicationName      string `pulumi:"applicationName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupApplicationResult struct {
	ApplicationType          *string `pulumi:"applicationType"`
	CommandLineArguments     *string `pulumi:"commandLineArguments"`
	CommandLineSetting       string  `pulumi:"commandLineSetting"`
	Description              *string `pulumi:"description"`
	FilePath                 *string `pulumi:"filePath"`
	FriendlyName             *string `pulumi:"friendlyName"`
	IconContent              string  `pulumi:"iconContent"`
	IconHash                 string  `pulumi:"iconHash"`
	IconIndex                *int    `pulumi:"iconIndex"`
	IconPath                 *string `pulumi:"iconPath"`
	Id                       string  `pulumi:"id"`
	MsixPackageApplicationId *string `pulumi:"msixPackageApplicationId"`
	MsixPackageFamilyName    *string `pulumi:"msixPackageFamilyName"`
	Name                     string  `pulumi:"name"`
	ObjectId                 string  `pulumi:"objectId"`
	ShowInPortal             *bool   `pulumi:"showInPortal"`
	Type                     string  `pulumi:"type"`
}
