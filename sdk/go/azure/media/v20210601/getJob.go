


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:media/v20210601:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	AccountName       string `pulumi:"accountName"`
	JobName           string `pulumi:"jobName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TransformName     string `pulumi:"transformName"`
}


type LookupJobResult struct {
	CorrelationData map[string]string        `pulumi:"correlationData"`
	Created         string                   `pulumi:"created"`
	Description     *string                  `pulumi:"description"`
	EndTime         string                   `pulumi:"endTime"`
	Id              string                   `pulumi:"id"`
	Input           interface{}              `pulumi:"input"`
	LastModified    string                   `pulumi:"lastModified"`
	Name            string                   `pulumi:"name"`
	Outputs         []JobOutputAssetResponse `pulumi:"outputs"`
	Priority        *string                  `pulumi:"priority"`
	StartTime       string                   `pulumi:"startTime"`
	State           string                   `pulumi:"state"`
	SystemData      SystemDataResponse       `pulumi:"systemData"`
	Type            string                   `pulumi:"type"`
}
