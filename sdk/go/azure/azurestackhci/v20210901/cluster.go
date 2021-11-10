


package v20210901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Cluster struct {
	pulumi.CustomResourceState

	AadClientId             pulumi.StringOutput                       `pulumi:"aadClientId"`
	AadTenantId             pulumi.StringOutput                       `pulumi:"aadTenantId"`
	BillingModel            pulumi.StringOutput                       `pulumi:"billingModel"`
	CloudId                 pulumi.StringOutput                       `pulumi:"cloudId"`
	CloudManagementEndpoint pulumi.StringPtrOutput                    `pulumi:"cloudManagementEndpoint"`
	CreatedAt               pulumi.StringPtrOutput                    `pulumi:"createdAt"`
	CreatedBy               pulumi.StringPtrOutput                    `pulumi:"createdBy"`
	CreatedByType           pulumi.StringPtrOutput                    `pulumi:"createdByType"`
	DesiredProperties       ClusterDesiredPropertiesResponsePtrOutput `pulumi:"desiredProperties"`
	LastBillingTimestamp    pulumi.StringOutput                       `pulumi:"lastBillingTimestamp"`
	LastModifiedAt          pulumi.StringPtrOutput                    `pulumi:"lastModifiedAt"`
	LastModifiedBy          pulumi.StringPtrOutput                    `pulumi:"lastModifiedBy"`
	LastModifiedByType      pulumi.StringPtrOutput                    `pulumi:"lastModifiedByType"`
	LastSyncTimestamp       pulumi.StringOutput                       `pulumi:"lastSyncTimestamp"`
	Location                pulumi.StringOutput                       `pulumi:"location"`
	Name                    pulumi.StringOutput                       `pulumi:"name"`
	ProvisioningState       pulumi.StringOutput                       `pulumi:"provisioningState"`
	RegistrationTimestamp   pulumi.StringOutput                       `pulumi:"registrationTimestamp"`
	ReportedProperties      ClusterReportedPropertiesResponseOutput   `pulumi:"reportedProperties"`
	Status                  pulumi.StringOutput                       `pulumi:"status"`
	Tags                    pulumi.StringMapOutput                    `pulumi:"tags"`
	TrialDaysRemaining      pulumi.Float64Output                      `pulumi:"trialDaysRemaining"`
	Type                    pulumi.StringOutput                       `pulumi:"type"`
}


func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AadClientId == nil {
		return nil, errors.New("invalid value for required argument 'AadClientId'")
	}
	if args.AadTenantId == nil {
		return nil, errors.New("invalid value for required argument 'AadTenantId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20200301preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20201001:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210101preview:Cluster"),
		},
	})
	opts = append(opts, aliases)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:azurestackhci/v20210901:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:azurestackhci/v20210901:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type clusterState struct {
}

type ClusterState struct {
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	AadClientId             string                    `pulumi:"aadClientId"`
	AadTenantId             string                    `pulumi:"aadTenantId"`
	CloudManagementEndpoint *string                   `pulumi:"cloudManagementEndpoint"`
	ClusterName             *string                   `pulumi:"clusterName"`
	CreatedAt               *string                   `pulumi:"createdAt"`
	CreatedBy               *string                   `pulumi:"createdBy"`
	CreatedByType           *string                   `pulumi:"createdByType"`
	DesiredProperties       *ClusterDesiredProperties `pulumi:"desiredProperties"`
	LastModifiedAt          *string                   `pulumi:"lastModifiedAt"`
	LastModifiedBy          *string                   `pulumi:"lastModifiedBy"`
	LastModifiedByType      *string                   `pulumi:"lastModifiedByType"`
	Location                *string                   `pulumi:"location"`
	ResourceGroupName       string                    `pulumi:"resourceGroupName"`
	Tags                    map[string]string         `pulumi:"tags"`
}


type ClusterArgs struct {
	AadClientId             pulumi.StringInput
	AadTenantId             pulumi.StringInput
	CloudManagementEndpoint pulumi.StringPtrInput
	ClusterName             pulumi.StringPtrInput
	CreatedAt               pulumi.StringPtrInput
	CreatedBy               pulumi.StringPtrInput
	CreatedByType           pulumi.StringPtrInput
	DesiredProperties       ClusterDesiredPropertiesPtrInput
	LastModifiedAt          pulumi.StringPtrInput
	LastModifiedBy          pulumi.StringPtrInput
	LastModifiedByType      pulumi.StringPtrInput
	Location                pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	Tags                    pulumi.StringMapInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil))
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil))
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
