


package v20200701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKeyValue(ctx *pulumi.Context, args *LookupKeyValueArgs, opts ...pulumi.InvokeOption) (*LookupKeyValueResult, error) {
	var rv LookupKeyValueResult
	err := ctx.Invoke("azure-native:appconfiguration/v20200701preview:getKeyValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKeyValueArgs struct {
	ConfigStoreName   string `pulumi:"configStoreName"`
	KeyValueName      string `pulumi:"keyValueName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupKeyValueResult struct {
	ContentType  *string           `pulumi:"contentType"`
	ETag         string            `pulumi:"eTag"`
	Id           string            `pulumi:"id"`
	Key          string            `pulumi:"key"`
	Label        string            `pulumi:"label"`
	LastModified string            `pulumi:"lastModified"`
	Locked       bool              `pulumi:"locked"`
	Name         string            `pulumi:"name"`
	Tags         map[string]string `pulumi:"tags"`
	Type         string            `pulumi:"type"`
	Value        *string           `pulumi:"value"`
}
