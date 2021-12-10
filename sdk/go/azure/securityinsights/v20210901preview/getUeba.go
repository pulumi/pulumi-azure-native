


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUeba(ctx *pulumi.Context, args *LookupUebaArgs, opts ...pulumi.InvokeOption) (*LookupUebaResult, error) {
	var rv LookupUebaResult
	err := ctx.Invoke("azure-native:securityinsights/v20210901preview:getUeba", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUebaArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SettingsName      string `pulumi:"settingsName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupUebaResult struct {
	DataSources []string           `pulumi:"dataSources"`
	Etag        *string            `pulumi:"etag"`
	Id          string             `pulumi:"id"`
	Kind        string             `pulumi:"kind"`
	Name        string             `pulumi:"name"`
	SystemData  SystemDataResponse `pulumi:"systemData"`
	Type        string             `pulumi:"type"`
}
