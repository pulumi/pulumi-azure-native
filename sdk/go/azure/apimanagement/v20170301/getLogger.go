


package v20170301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLogger(ctx *pulumi.Context, args *LookupLoggerArgs, opts ...pulumi.InvokeOption) (*LookupLoggerResult, error) {
	var rv LookupLoggerResult
	err := ctx.Invoke("azure-native:apimanagement/v20170301:getLogger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoggerArgs struct {
	Loggerid          string `pulumi:"loggerid"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupLoggerResult struct {
	Credentials map[string]string               `pulumi:"credentials"`
	Description *string                         `pulumi:"description"`
	Id          string                          `pulumi:"id"`
	IsBuffered  *bool                           `pulumi:"isBuffered"`
	LoggerType  string                          `pulumi:"loggerType"`
	Name        string                          `pulumi:"name"`
	Sampling    *LoggerSamplingContractResponse `pulumi:"sampling"`
	Type        string                          `pulumi:"type"`
}
