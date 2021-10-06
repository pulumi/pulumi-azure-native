


package testbase

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPackage(ctx *pulumi.Context, args *LookupPackageArgs, opts ...pulumi.InvokeOption) (*LookupPackageResult, error) {
	var rv LookupPackageResult
	err := ctx.Invoke("azure-native:testbase:getPackage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPackageArgs struct {
	PackageName         string `pulumi:"packageName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}


type LookupPackageResult struct {
	ApplicationName   string                            `pulumi:"applicationName"`
	BlobPath          string                            `pulumi:"blobPath"`
	Etag              string                            `pulumi:"etag"`
	FlightingRing     string                            `pulumi:"flightingRing"`
	Id                string                            `pulumi:"id"`
	IsEnabled         bool                              `pulumi:"isEnabled"`
	LastModifiedTime  string                            `pulumi:"lastModifiedTime"`
	Location          string                            `pulumi:"location"`
	Name              string                            `pulumi:"name"`
	PackageStatus     string                            `pulumi:"packageStatus"`
	ProvisioningState string                            `pulumi:"provisioningState"`
	SystemData        SystemDataResponse                `pulumi:"systemData"`
	Tags              map[string]string                 `pulumi:"tags"`
	TargetOSList      []TargetOSInfoResponse            `pulumi:"targetOSList"`
	TestTypes         []string                          `pulumi:"testTypes"`
	Tests             []TestResponse                    `pulumi:"tests"`
	Type              string                            `pulumi:"type"`
	ValidationResults []PackageValidationResultResponse `pulumi:"validationResults"`
	Version           string                            `pulumi:"version"`
}
