


package v20220101

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
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220301:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220501:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220901:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221001:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221201:Cluster"),
		},
	})
	opts = append(opts, aliases)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:azurestackhci/v20220101:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:azurestackhci/v20220101:Cluster", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func (o ClusterOutput) AadClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.AadClientId }).(pulumi.StringOutput)
}

func (o ClusterOutput) AadTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.AadTenantId }).(pulumi.StringOutput)
}

func (o ClusterOutput) BillingModel() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.BillingModel }).(pulumi.StringOutput)
}

func (o ClusterOutput) CloudId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.CloudId }).(pulumi.StringOutput)
}

func (o ClusterOutput) CloudManagementEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.CloudManagementEndpoint }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) DesiredProperties() ClusterDesiredPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterDesiredPropertiesResponsePtrOutput { return v.DesiredProperties }).(ClusterDesiredPropertiesResponsePtrOutput)
}

func (o ClusterOutput) LastBillingTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.LastBillingTimestamp }).(pulumi.StringOutput)
}

func (o ClusterOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) LastSyncTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.LastSyncTimestamp }).(pulumi.StringOutput)
}

func (o ClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ClusterOutput) RegistrationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.RegistrationTimestamp }).(pulumi.StringOutput)
}

func (o ClusterOutput) ReportedProperties() ClusterReportedPropertiesResponseOutput {
	return o.ApplyT(func(v *Cluster) ClusterReportedPropertiesResponseOutput { return v.ReportedProperties }).(ClusterReportedPropertiesResponseOutput)
}

func (o ClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ClusterOutput) TrialDaysRemaining() pulumi.Float64Output {
	return o.ApplyT(func(v *Cluster) pulumi.Float64Output { return v.TrialDaysRemaining }).(pulumi.Float64Output)
}

func (o ClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
