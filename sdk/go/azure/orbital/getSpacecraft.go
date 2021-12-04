


package orbital

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSpacecraft(ctx *pulumi.Context, args *LookupSpacecraftArgs, opts ...pulumi.InvokeOption) (*LookupSpacecraftResult, error) {
	var rv LookupSpacecraftResult
	err := ctx.Invoke("azure-native:orbital:getSpacecraft", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSpacecraftArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SpacecraftName    string `pulumi:"spacecraftName"`
}


type LookupSpacecraftResult struct {
	AuthorizationStatus         string                   `pulumi:"authorizationStatus"`
	AuthorizationStatusExtended string                   `pulumi:"authorizationStatusExtended"`
	Etag                        string                   `pulumi:"etag"`
	Id                          string                   `pulumi:"id"`
	Links                       []SpacecraftLinkResponse `pulumi:"links"`
	Location                    string                   `pulumi:"location"`
	Name                        string                   `pulumi:"name"`
	NoradId                     string                   `pulumi:"noradId"`
	SystemData                  SystemDataResponse       `pulumi:"systemData"`
	Tags                        map[string]string        `pulumi:"tags"`
	TitleLine                   *string                  `pulumi:"titleLine"`
	TleLine1                    *string                  `pulumi:"tleLine1"`
	TleLine2                    *string                  `pulumi:"tleLine2"`
	Type                        string                   `pulumi:"type"`
}
