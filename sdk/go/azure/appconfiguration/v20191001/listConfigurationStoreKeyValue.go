


package v20191001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConfigurationStoreKeyValue(ctx *pulumi.Context, args *ListConfigurationStoreKeyValueArgs, opts ...pulumi.InvokeOption) (*ListConfigurationStoreKeyValueResult, error) {
	var rv ListConfigurationStoreKeyValueResult
	err := ctx.Invoke("azure-native:appconfiguration/v20191001:listConfigurationStoreKeyValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConfigurationStoreKeyValueArgs struct {
	ConfigStoreName   string  `pulumi:"configStoreName"`
	Key               string  `pulumi:"key"`
	Label             *string `pulumi:"label"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type ListConfigurationStoreKeyValueResult struct {
	ContentType  string            `pulumi:"contentType"`
	ETag         string            `pulumi:"eTag"`
	Key          string            `pulumi:"key"`
	Label        string            `pulumi:"label"`
	LastModified string            `pulumi:"lastModified"`
	Locked       bool              `pulumi:"locked"`
	Tags         map[string]string `pulumi:"tags"`
	Value        string            `pulumi:"value"`
}
