


package v20201201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Service struct {
	pulumi.CustomResourceState

	CorrelationScheme            ServiceCorrelationDescriptionResponseArrayOutput     `pulumi:"correlationScheme"`
	DefaultMoveCost              pulumi.StringPtrOutput                               `pulumi:"defaultMoveCost"`
	Etag                         pulumi.StringOutput                                  `pulumi:"etag"`
	Location                     pulumi.StringPtrOutput                               `pulumi:"location"`
	Name                         pulumi.StringOutput                                  `pulumi:"name"`
	PartitionDescription         pulumi.AnyOutput                                     `pulumi:"partitionDescription"`
	PlacementConstraints         pulumi.StringPtrOutput                               `pulumi:"placementConstraints"`
	ProvisioningState            pulumi.StringOutput                                  `pulumi:"provisioningState"`
	ServiceDnsName               pulumi.StringPtrOutput                               `pulumi:"serviceDnsName"`
	ServiceKind                  pulumi.StringOutput                                  `pulumi:"serviceKind"`
	ServiceLoadMetrics           ServiceLoadMetricDescriptionResponseArrayOutput      `pulumi:"serviceLoadMetrics"`
	ServicePackageActivationMode pulumi.StringPtrOutput                               `pulumi:"servicePackageActivationMode"`
	ServicePlacementPolicies     ServicePlacementPolicyDescriptionResponseArrayOutput `pulumi:"servicePlacementPolicies"`
	ServiceTypeName              pulumi.StringPtrOutput                               `pulumi:"serviceTypeName"`
	SystemData                   SystemDataResponseOutput                             `pulumi:"systemData"`
	Tags                         pulumi.StringMapOutput                               `pulumi:"tags"`
	Type                         pulumi.StringOutput                                  `pulumi:"type"`
}


func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationName'")
	}
	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceKind == nil {
		return nil, errors.New("invalid value for required argument 'ServiceKind'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric:Service"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20170701preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190301:Service"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190301preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190601preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20191101preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20200301:Service"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210601:Service"),
		},
	})
	opts = append(opts, aliases)
	var resource Service
	err := ctx.RegisterResource("azure-native:servicefabric/v20201201preview:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure-native:servicefabric/v20201201preview:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serviceState struct {
}

type ServiceState struct {
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	ApplicationName              string                              `pulumi:"applicationName"`
	ClusterName                  string                              `pulumi:"clusterName"`
	CorrelationScheme            []ServiceCorrelationDescription     `pulumi:"correlationScheme"`
	DefaultMoveCost              *string                             `pulumi:"defaultMoveCost"`
	Location                     *string                             `pulumi:"location"`
	PartitionDescription         interface{}                         `pulumi:"partitionDescription"`
	PlacementConstraints         *string                             `pulumi:"placementConstraints"`
	ResourceGroupName            string                              `pulumi:"resourceGroupName"`
	ServiceDnsName               *string                             `pulumi:"serviceDnsName"`
	ServiceKind                  string                              `pulumi:"serviceKind"`
	ServiceLoadMetrics           []ServiceLoadMetricDescription      `pulumi:"serviceLoadMetrics"`
	ServiceName                  *string                             `pulumi:"serviceName"`
	ServicePackageActivationMode *string                             `pulumi:"servicePackageActivationMode"`
	ServicePlacementPolicies     []ServicePlacementPolicyDescription `pulumi:"servicePlacementPolicies"`
	ServiceTypeName              *string                             `pulumi:"serviceTypeName"`
	Tags                         map[string]string                   `pulumi:"tags"`
}


type ServiceArgs struct {
	ApplicationName              pulumi.StringInput
	ClusterName                  pulumi.StringInput
	CorrelationScheme            ServiceCorrelationDescriptionArrayInput
	DefaultMoveCost              pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	PartitionDescription         pulumi.Input
	PlacementConstraints         pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	ServiceDnsName               pulumi.StringPtrInput
	ServiceKind                  pulumi.StringInput
	ServiceLoadMetrics           ServiceLoadMetricDescriptionArrayInput
	ServiceName                  pulumi.StringPtrInput
	ServicePackageActivationMode pulumi.StringPtrInput
	ServicePlacementPolicies     ServicePlacementPolicyDescriptionArrayInput
	ServiceTypeName              pulumi.StringPtrInput
	Tags                         pulumi.StringMapInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

func (o ServiceOutput) CorrelationScheme() ServiceCorrelationDescriptionResponseArrayOutput {
	return o.ApplyT(func(v *Service) ServiceCorrelationDescriptionResponseArrayOutput { return v.CorrelationScheme }).(ServiceCorrelationDescriptionResponseArrayOutput)
}

func (o ServiceOutput) DefaultMoveCost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.DefaultMoveCost }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceOutput) PartitionDescription() pulumi.AnyOutput {
	return o.ApplyT(func(v *Service) pulumi.AnyOutput { return v.PartitionDescription }).(pulumi.AnyOutput)
}

func (o ServiceOutput) PlacementConstraints() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.PlacementConstraints }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServiceOutput) ServiceDnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.ServiceDnsName }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) ServiceKind() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ServiceKind }).(pulumi.StringOutput)
}

func (o ServiceOutput) ServiceLoadMetrics() ServiceLoadMetricDescriptionResponseArrayOutput {
	return o.ApplyT(func(v *Service) ServiceLoadMetricDescriptionResponseArrayOutput { return v.ServiceLoadMetrics }).(ServiceLoadMetricDescriptionResponseArrayOutput)
}

func (o ServiceOutput) ServicePackageActivationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.ServicePackageActivationMode }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) ServicePlacementPolicies() ServicePlacementPolicyDescriptionResponseArrayOutput {
	return o.ApplyT(func(v *Service) ServicePlacementPolicyDescriptionResponseArrayOutput {
		return v.ServicePlacementPolicies
	}).(ServicePlacementPolicyDescriptionResponseArrayOutput)
}

func (o ServiceOutput) ServiceTypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.ServiceTypeName }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Service) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Service) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
