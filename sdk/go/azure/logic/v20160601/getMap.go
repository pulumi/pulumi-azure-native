


package v20160601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMap(ctx *pulumi.Context, args *LookupMapArgs, opts ...pulumi.InvokeOption) (*LookupMapResult, error) {
	var rv LookupMapResult
	err := ctx.Invoke("azure-native:logic/v20160601:getMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMapArgs struct {
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	MapName                string `pulumi:"mapName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupMapResult struct {
	ChangedTime      string                                                   `pulumi:"changedTime"`
	Content          *string                                                  `pulumi:"content"`
	ContentLink      ContentLinkResponse                                      `pulumi:"contentLink"`
	ContentType      *string                                                  `pulumi:"contentType"`
	CreatedTime      string                                                   `pulumi:"createdTime"`
	Id               string                                                   `pulumi:"id"`
	Location         *string                                                  `pulumi:"location"`
	MapType          string                                                   `pulumi:"mapType"`
	Metadata         interface{}                                              `pulumi:"metadata"`
	Name             string                                                   `pulumi:"name"`
	ParametersSchema *IntegrationAccountMapPropertiesResponseParametersSchema `pulumi:"parametersSchema"`
	Tags             map[string]string                                        `pulumi:"tags"`
	Type             string                                                   `pulumi:"type"`
}
