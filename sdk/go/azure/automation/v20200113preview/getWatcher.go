


package v20200113preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWatcher(ctx *pulumi.Context, args *LookupWatcherArgs, opts ...pulumi.InvokeOption) (*LookupWatcherResult, error) {
	var rv LookupWatcherResult
	err := ctx.Invoke("azure-native:automation/v20200113preview:getWatcher", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWatcherArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	WatcherName           string `pulumi:"watcherName"`
}


type LookupWatcherResult struct {
	CreationTime                string            `pulumi:"creationTime"`
	Description                 *string           `pulumi:"description"`
	Etag                        *string           `pulumi:"etag"`
	ExecutionFrequencyInSeconds *float64          `pulumi:"executionFrequencyInSeconds"`
	Id                          string            `pulumi:"id"`
	LastModifiedBy              string            `pulumi:"lastModifiedBy"`
	LastModifiedTime            string            `pulumi:"lastModifiedTime"`
	Location                    *string           `pulumi:"location"`
	Name                        string            `pulumi:"name"`
	ScriptName                  *string           `pulumi:"scriptName"`
	ScriptParameters            map[string]string `pulumi:"scriptParameters"`
	ScriptRunOn                 *string           `pulumi:"scriptRunOn"`
	Status                      string            `pulumi:"status"`
	Tags                        map[string]string `pulumi:"tags"`
	Type                        string            `pulumi:"type"`
}
