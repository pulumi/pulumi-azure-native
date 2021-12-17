


package costmanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudConnector struct {
	pulumi.CustomResourceState

	BillingModel                      pulumi.StringPtrOutput                `pulumi:"billingModel"`
	CollectionInfo                    ConnectorCollectionInfoResponseOutput `pulumi:"collectionInfo"`
	CreatedOn                         pulumi.StringOutput                   `pulumi:"createdOn"`
	CredentialsKey                    pulumi.StringPtrOutput                `pulumi:"credentialsKey"`
	DaysTrialRemaining                pulumi.IntOutput                      `pulumi:"daysTrialRemaining"`
	DefaultManagementGroupId          pulumi.StringPtrOutput                `pulumi:"defaultManagementGroupId"`
	DisplayName                       pulumi.StringPtrOutput                `pulumi:"displayName"`
	ExternalBillingAccountId          pulumi.StringOutput                   `pulumi:"externalBillingAccountId"`
	Kind                              pulumi.StringPtrOutput                `pulumi:"kind"`
	ModifiedOn                        pulumi.StringOutput                   `pulumi:"modifiedOn"`
	Name                              pulumi.StringOutput                   `pulumi:"name"`
	ProviderBillingAccountDisplayName pulumi.StringOutput                   `pulumi:"providerBillingAccountDisplayName"`
	ProviderBillingAccountId          pulumi.StringOutput                   `pulumi:"providerBillingAccountId"`
	ReportId                          pulumi.StringPtrOutput                `pulumi:"reportId"`
	Status                            pulumi.StringOutput                   `pulumi:"status"`
	SubscriptionId                    pulumi.StringPtrOutput                `pulumi:"subscriptionId"`
	Type                              pulumi.StringOutput                   `pulumi:"type"`
}


func NewCloudConnector(ctx *pulumi.Context,
	name string, args *CloudConnectorArgs, opts ...pulumi.ResourceOption) (*CloudConnector, error) {
	if args == nil {
		args = &CloudConnectorArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement/v20180801preview:CloudConnector"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20190301preview:CloudConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource CloudConnector
	err := ctx.RegisterResource("azure-native:costmanagement:CloudConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCloudConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudConnectorState, opts ...pulumi.ResourceOption) (*CloudConnector, error) {
	var resource CloudConnector
	err := ctx.ReadResource("azure-native:costmanagement:CloudConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type cloudConnectorState struct {
}

type CloudConnectorState struct {
}

func (CloudConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudConnectorState)(nil)).Elem()
}

type cloudConnectorArgs struct {
	BillingModel             *string `pulumi:"billingModel"`
	ConnectorName            *string `pulumi:"connectorName"`
	CredentialsKey           *string `pulumi:"credentialsKey"`
	CredentialsSecret        *string `pulumi:"credentialsSecret"`
	DefaultManagementGroupId *string `pulumi:"defaultManagementGroupId"`
	DisplayName              *string `pulumi:"displayName"`
	Kind                     *string `pulumi:"kind"`
	ReportId                 *string `pulumi:"reportId"`
	SubscriptionId           *string `pulumi:"subscriptionId"`
}


type CloudConnectorArgs struct {
	BillingModel             pulumi.StringPtrInput
	ConnectorName            pulumi.StringPtrInput
	CredentialsKey           pulumi.StringPtrInput
	CredentialsSecret        pulumi.StringPtrInput
	DefaultManagementGroupId pulumi.StringPtrInput
	DisplayName              pulumi.StringPtrInput
	Kind                     pulumi.StringPtrInput
	ReportId                 pulumi.StringPtrInput
	SubscriptionId           pulumi.StringPtrInput
}

func (CloudConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudConnectorArgs)(nil)).Elem()
}

type CloudConnectorInput interface {
	pulumi.Input

	ToCloudConnectorOutput() CloudConnectorOutput
	ToCloudConnectorOutputWithContext(ctx context.Context) CloudConnectorOutput
}

func (*CloudConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudConnector)(nil)).Elem()
}

func (i *CloudConnector) ToCloudConnectorOutput() CloudConnectorOutput {
	return i.ToCloudConnectorOutputWithContext(context.Background())
}

func (i *CloudConnector) ToCloudConnectorOutputWithContext(ctx context.Context) CloudConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudConnectorOutput)
}

type CloudConnectorOutput struct{ *pulumi.OutputState }

func (CloudConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudConnector)(nil)).Elem()
}

func (o CloudConnectorOutput) ToCloudConnectorOutput() CloudConnectorOutput {
	return o
}

func (o CloudConnectorOutput) ToCloudConnectorOutputWithContext(ctx context.Context) CloudConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CloudConnectorOutput{})
}
