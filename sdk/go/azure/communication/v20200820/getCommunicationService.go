


package v20200820

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCommunicationService(ctx *pulumi.Context, args *LookupCommunicationServiceArgs, opts ...pulumi.InvokeOption) (*LookupCommunicationServiceResult, error) {
	var rv LookupCommunicationServiceResult
	err := ctx.Invoke("azure-native:communication/v20200820:getCommunicationService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCommunicationServiceArgs struct {
	CommunicationServiceName string `pulumi:"communicationServiceName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
}


type LookupCommunicationServiceResult struct {
	DataLocation        string             `pulumi:"dataLocation"`
	HostName            string             `pulumi:"hostName"`
	Id                  string             `pulumi:"id"`
	ImmutableResourceId string             `pulumi:"immutableResourceId"`
	Location            *string            `pulumi:"location"`
	Name                string             `pulumi:"name"`
	NotificationHubId   string             `pulumi:"notificationHubId"`
	ProvisioningState   string             `pulumi:"provisioningState"`
	SystemData          SystemDataResponse `pulumi:"systemData"`
	Tags                map[string]string  `pulumi:"tags"`
	Type                string             `pulumi:"type"`
	Version             string             `pulumi:"version"`
}
