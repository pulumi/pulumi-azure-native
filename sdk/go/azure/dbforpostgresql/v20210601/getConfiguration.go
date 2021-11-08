


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfiguration(ctx *pulumi.Context, args *LookupConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationResult, error) {
	var rv LookupConfigurationResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20210601:getConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationArgs struct {
	ConfigurationName string `pulumi:"configurationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupConfigurationResult struct {
	AllowedValues string             `pulumi:"allowedValues"`
	DataType      string             `pulumi:"dataType"`
	DefaultValue  string             `pulumi:"defaultValue"`
	Description   string             `pulumi:"description"`
	Id            string             `pulumi:"id"`
	Name          string             `pulumi:"name"`
	Source        *string            `pulumi:"source"`
	SystemData    SystemDataResponse `pulumi:"systemData"`
	Type          string             `pulumi:"type"`
	Value         *string            `pulumi:"value"`
}
