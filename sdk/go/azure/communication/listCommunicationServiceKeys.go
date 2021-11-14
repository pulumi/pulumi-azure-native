


package communication

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListCommunicationServiceKeys(ctx *pulumi.Context, args *ListCommunicationServiceKeysArgs, opts ...pulumi.InvokeOption) (*ListCommunicationServiceKeysResult, error) {
	var rv ListCommunicationServiceKeysResult
	err := ctx.Invoke("azure-native:communication:listCommunicationServiceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListCommunicationServiceKeysArgs struct {
	CommunicationServiceName string `pulumi:"communicationServiceName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
}


type ListCommunicationServiceKeysResult struct {
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	PrimaryKey                *string `pulumi:"primaryKey"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	SecondaryKey              *string `pulumi:"secondaryKey"`
}
