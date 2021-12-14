


package v20191101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupArtifactSource(ctx *pulumi.Context, args *LookupArtifactSourceArgs, opts ...pulumi.InvokeOption) (*LookupArtifactSourceResult, error) {
	var rv LookupArtifactSourceResult
	err := ctx.Invoke("azure-native:deploymentmanager/v20191101preview:getArtifactSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArtifactSourceArgs struct {
	ArtifactSourceName string `pulumi:"artifactSourceName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupArtifactSourceResult struct {
	ArtifactRoot   *string                   `pulumi:"artifactRoot"`
	Authentication SasAuthenticationResponse `pulumi:"authentication"`
	Id             string                    `pulumi:"id"`
	Location       string                    `pulumi:"location"`
	Name           string                    `pulumi:"name"`
	SourceType     string                    `pulumi:"sourceType"`
	Tags           map[string]string         `pulumi:"tags"`
	Type           string                    `pulumi:"type"`
}
