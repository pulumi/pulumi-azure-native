


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AnalyticsConnector struct {
	pulumi.CustomResourceState

	DataDestinationConfiguration AnalyticsConnectorDataLakeDataDestinationResponseOutput `pulumi:"dataDestinationConfiguration"`
	DataMappingConfiguration     AnalyticsConnectorFhirToParquetMappingResponseOutput    `pulumi:"dataMappingConfiguration"`
	DataSourceConfiguration      AnalyticsConnectorFhirServiceDataSourceResponseOutput   `pulumi:"dataSourceConfiguration"`
	Etag                         pulumi.StringPtrOutput                                  `pulumi:"etag"`
	Identity                     ServiceManagedIdentityResponseIdentityPtrOutput         `pulumi:"identity"`
	Location                     pulumi.StringPtrOutput                                  `pulumi:"location"`
	Name                         pulumi.StringOutput                                     `pulumi:"name"`
	ProvisioningState            pulumi.StringOutput                                     `pulumi:"provisioningState"`
	SystemData                   SystemDataResponseOutput                                `pulumi:"systemData"`
	Tags                         pulumi.StringMapOutput                                  `pulumi:"tags"`
	Type                         pulumi.StringOutput                                     `pulumi:"type"`
}


func NewAnalyticsConnector(ctx *pulumi.Context,
	name string, args *AnalyticsConnectorArgs, opts ...pulumi.ResourceOption) (*AnalyticsConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataDestinationConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'DataDestinationConfiguration'")
	}
	if args.DataMappingConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'DataMappingConfiguration'")
	}
	if args.DataSourceConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'DataSourceConfiguration'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	var resource AnalyticsConnector
	err := ctx.RegisterResource("azure-native:healthcareapis/v20221001preview:AnalyticsConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAnalyticsConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalyticsConnectorState, opts ...pulumi.ResourceOption) (*AnalyticsConnector, error) {
	var resource AnalyticsConnector
	err := ctx.ReadResource("azure-native:healthcareapis/v20221001preview:AnalyticsConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type analyticsConnectorState struct {
}

type AnalyticsConnectorState struct {
}

func (AnalyticsConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsConnectorState)(nil)).Elem()
}

type analyticsConnectorArgs struct {
	AnalyticsConnectorName       *string                                   `pulumi:"analyticsConnectorName"`
	DataDestinationConfiguration AnalyticsConnectorDataLakeDataDestination `pulumi:"dataDestinationConfiguration"`
	DataMappingConfiguration     AnalyticsConnectorFhirToParquetMapping    `pulumi:"dataMappingConfiguration"`
	DataSourceConfiguration      AnalyticsConnectorFhirServiceDataSource   `pulumi:"dataSourceConfiguration"`
	Identity                     *ServiceManagedIdentityIdentity           `pulumi:"identity"`
	Location                     *string                                   `pulumi:"location"`
	ResourceGroupName            string                                    `pulumi:"resourceGroupName"`
	Tags                         map[string]string                         `pulumi:"tags"`
	WorkspaceName                string                                    `pulumi:"workspaceName"`
}


type AnalyticsConnectorArgs struct {
	AnalyticsConnectorName       pulumi.StringPtrInput
	DataDestinationConfiguration AnalyticsConnectorDataLakeDataDestinationInput
	DataMappingConfiguration     AnalyticsConnectorFhirToParquetMappingInput
	DataSourceConfiguration      AnalyticsConnectorFhirServiceDataSourceInput
	Identity                     ServiceManagedIdentityIdentityPtrInput
	Location                     pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	Tags                         pulumi.StringMapInput
	WorkspaceName                pulumi.StringInput
}

func (AnalyticsConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsConnectorArgs)(nil)).Elem()
}

type AnalyticsConnectorInput interface {
	pulumi.Input

	ToAnalyticsConnectorOutput() AnalyticsConnectorOutput
	ToAnalyticsConnectorOutputWithContext(ctx context.Context) AnalyticsConnectorOutput
}

func (*AnalyticsConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticsConnector)(nil)).Elem()
}

func (i *AnalyticsConnector) ToAnalyticsConnectorOutput() AnalyticsConnectorOutput {
	return i.ToAnalyticsConnectorOutputWithContext(context.Background())
}

func (i *AnalyticsConnector) ToAnalyticsConnectorOutputWithContext(ctx context.Context) AnalyticsConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsConnectorOutput)
}

type AnalyticsConnectorOutput struct{ *pulumi.OutputState }

func (AnalyticsConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticsConnector)(nil)).Elem()
}

func (o AnalyticsConnectorOutput) ToAnalyticsConnectorOutput() AnalyticsConnectorOutput {
	return o
}

func (o AnalyticsConnectorOutput) ToAnalyticsConnectorOutputWithContext(ctx context.Context) AnalyticsConnectorOutput {
	return o
}

func (o AnalyticsConnectorOutput) DataDestinationConfiguration() AnalyticsConnectorDataLakeDataDestinationResponseOutput {
	return o.ApplyT(func(v *AnalyticsConnector) AnalyticsConnectorDataLakeDataDestinationResponseOutput {
		return v.DataDestinationConfiguration
	}).(AnalyticsConnectorDataLakeDataDestinationResponseOutput)
}

func (o AnalyticsConnectorOutput) DataMappingConfiguration() AnalyticsConnectorFhirToParquetMappingResponseOutput {
	return o.ApplyT(func(v *AnalyticsConnector) AnalyticsConnectorFhirToParquetMappingResponseOutput {
		return v.DataMappingConfiguration
	}).(AnalyticsConnectorFhirToParquetMappingResponseOutput)
}

func (o AnalyticsConnectorOutput) DataSourceConfiguration() AnalyticsConnectorFhirServiceDataSourceResponseOutput {
	return o.ApplyT(func(v *AnalyticsConnector) AnalyticsConnectorFhirServiceDataSourceResponseOutput {
		return v.DataSourceConfiguration
	}).(AnalyticsConnectorFhirServiceDataSourceResponseOutput)
}

func (o AnalyticsConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnalyticsConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o AnalyticsConnectorOutput) Identity() ServiceManagedIdentityResponseIdentityPtrOutput {
	return o.ApplyT(func(v *AnalyticsConnector) ServiceManagedIdentityResponseIdentityPtrOutput { return v.Identity }).(ServiceManagedIdentityResponseIdentityPtrOutput)
}

func (o AnalyticsConnectorOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnalyticsConnector) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o AnalyticsConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalyticsConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AnalyticsConnectorOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalyticsConnector) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AnalyticsConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AnalyticsConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AnalyticsConnectorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AnalyticsConnector) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AnalyticsConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalyticsConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AnalyticsConnectorOutput{})
}
