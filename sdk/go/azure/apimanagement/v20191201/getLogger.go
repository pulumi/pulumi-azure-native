


package v20191201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLogger(ctx *pulumi.Context, args *LookupLoggerArgs, opts ...pulumi.InvokeOption) (*LookupLoggerResult, error) {
	var rv LookupLoggerResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201:getLogger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoggerArgs struct {
	LoggerId          string `pulumi:"loggerId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupLoggerResult struct {
	Credentials map[string]string `pulumi:"credentials"`
	Description *string           `pulumi:"description"`
	Id          string            `pulumi:"id"`
	IsBuffered  *bool             `pulumi:"isBuffered"`
	LoggerType  string            `pulumi:"loggerType"`
	Name        string            `pulumi:"name"`
	ResourceId  *string           `pulumi:"resourceId"`
	Type        string            `pulumi:"type"`
}
