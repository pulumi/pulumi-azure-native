


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("azure-native:apimanagement/v20201201:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGroupArgs struct {
	GroupId           string `pulumi:"groupId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupGroupResult struct {
	BuiltIn     bool    `pulumi:"builtIn"`
	Description *string `pulumi:"description"`
	DisplayName string  `pulumi:"displayName"`
	ExternalId  *string `pulumi:"externalId"`
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	Type        string  `pulumi:"type"`
}
