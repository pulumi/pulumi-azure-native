


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDicomService(ctx *pulumi.Context, args *LookupDicomServiceArgs, opts ...pulumi.InvokeOption) (*LookupDicomServiceResult, error) {
	var rv LookupDicomServiceResult
	err := ctx.Invoke("azure-native:healthcareapis/v20210601preview:getDicomService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDicomServiceArgs struct {
	DicomServiceName  string `pulumi:"dicomServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupDicomServiceResult struct {
	AuthenticationConfiguration *DicomServiceAuthenticationConfigurationResponse `pulumi:"authenticationConfiguration"`
	Etag                        *string                                          `pulumi:"etag"`
	Id                          string                                           `pulumi:"id"`
	Location                    *string                                          `pulumi:"location"`
	Name                        string                                           `pulumi:"name"`
	ProvisioningState           string                                           `pulumi:"provisioningState"`
	ServiceUrl                  string                                           `pulumi:"serviceUrl"`
	SystemData                  SystemDataResponse                               `pulumi:"systemData"`
	Tags                        map[string]string                                `pulumi:"tags"`
	Type                        string                                           `pulumi:"type"`
}
