


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OfficeDataConnector struct {
	pulumi.CustomResourceState

	DataTypes  OfficeDataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	Etag       pulumi.StringPtrOutput                     `pulumi:"etag"`
	Kind       pulumi.StringOutput                        `pulumi:"kind"`
	Name       pulumi.StringOutput                        `pulumi:"name"`
	SystemData SystemDataResponseOutput                   `pulumi:"systemData"`
	TenantId   pulumi.StringOutput                        `pulumi:"tenantId"`
	Type       pulumi.StringOutput                        `pulumi:"type"`
}


func NewOfficeDataConnector(ctx *pulumi.Context,
	name string, args *OfficeDataConnectorArgs, opts ...pulumi.ResourceOption) (*OfficeDataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataTypes == nil {
		return nil, errors.New("invalid value for required argument 'DataTypes'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("Office365")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20210301preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20190101preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20200101:OfficeDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource OfficeDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20210301preview:OfficeDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOfficeDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OfficeDataConnectorState, opts ...pulumi.ResourceOption) (*OfficeDataConnector, error) {
	var resource OfficeDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20210301preview:OfficeDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type officeDataConnectorState struct {
}

type OfficeDataConnectorState struct {
}

func (OfficeDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*officeDataConnectorState)(nil)).Elem()
}

type officeDataConnectorArgs struct {
	DataConnectorId                     *string                      `pulumi:"dataConnectorId"`
	DataTypes                           OfficeDataConnectorDataTypes `pulumi:"dataTypes"`
	Etag                                *string                      `pulumi:"etag"`
	Kind                                string                       `pulumi:"kind"`
	OperationalInsightsResourceProvider string                       `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string                       `pulumi:"resourceGroupName"`
	TenantId                            string                       `pulumi:"tenantId"`
	WorkspaceName                       string                       `pulumi:"workspaceName"`
}


type OfficeDataConnectorArgs struct {
	DataConnectorId                     pulumi.StringPtrInput
	DataTypes                           OfficeDataConnectorDataTypesInput
	Etag                                pulumi.StringPtrInput
	Kind                                pulumi.StringInput
	OperationalInsightsResourceProvider pulumi.StringInput
	ResourceGroupName                   pulumi.StringInput
	TenantId                            pulumi.StringInput
	WorkspaceName                       pulumi.StringInput
}

func (OfficeDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*officeDataConnectorArgs)(nil)).Elem()
}

type OfficeDataConnectorInput interface {
	pulumi.Input

	ToOfficeDataConnectorOutput() OfficeDataConnectorOutput
	ToOfficeDataConnectorOutputWithContext(ctx context.Context) OfficeDataConnectorOutput
}

func (*OfficeDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnector)(nil))
}

func (i *OfficeDataConnector) ToOfficeDataConnectorOutput() OfficeDataConnectorOutput {
	return i.ToOfficeDataConnectorOutputWithContext(context.Background())
}

func (i *OfficeDataConnector) ToOfficeDataConnectorOutputWithContext(ctx context.Context) OfficeDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorOutput)
}

type OfficeDataConnectorOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnector)(nil))
}

func (o OfficeDataConnectorOutput) ToOfficeDataConnectorOutput() OfficeDataConnectorOutput {
	return o
}

func (o OfficeDataConnectorOutput) ToOfficeDataConnectorOutputWithContext(ctx context.Context) OfficeDataConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OfficeDataConnectorOutput{})
}
