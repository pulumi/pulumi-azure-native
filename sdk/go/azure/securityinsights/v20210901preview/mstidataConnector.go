


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MSTIDataConnector struct {
	pulumi.CustomResourceState

	DataTypes  MSTIDataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	Etag       pulumi.StringPtrOutput                   `pulumi:"etag"`
	Kind       pulumi.StringOutput                      `pulumi:"kind"`
	Name       pulumi.StringOutput                      `pulumi:"name"`
	SystemData SystemDataResponseOutput                 `pulumi:"systemData"`
	TenantId   pulumi.StringOutput                      `pulumi:"tenantId"`
	Type       pulumi.StringOutput                      `pulumi:"type"`
}


func NewMSTIDataConnector(ctx *pulumi.Context,
	name string, args *MSTIDataConnectorArgs, opts ...pulumi.ResourceOption) (*MSTIDataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataTypes == nil {
		return nil, errors.New("invalid value for required argument 'DataTypes'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
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
	args.Kind = pulumi.String("MicrosoftThreatIntelligence")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:MSTIDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource MSTIDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20210901preview:MSTIDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMSTIDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MSTIDataConnectorState, opts ...pulumi.ResourceOption) (*MSTIDataConnector, error) {
	var resource MSTIDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20210901preview:MSTIDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mstidataConnectorState struct {
}

type MSTIDataConnectorState struct {
}

func (MSTIDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*mstidataConnectorState)(nil)).Elem()
}

type mstidataConnectorArgs struct {
	DataConnectorId   *string                    `pulumi:"dataConnectorId"`
	DataTypes         MSTIDataConnectorDataTypes `pulumi:"dataTypes"`
	Etag              *string                    `pulumi:"etag"`
	Kind              string                     `pulumi:"kind"`
	ResourceGroupName string                     `pulumi:"resourceGroupName"`
	TenantId          string                     `pulumi:"tenantId"`
	WorkspaceName     string                     `pulumi:"workspaceName"`
}


type MSTIDataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         MSTIDataConnectorDataTypesInput
	Etag              pulumi.StringPtrInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TenantId          pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (MSTIDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mstidataConnectorArgs)(nil)).Elem()
}

type MSTIDataConnectorInput interface {
	pulumi.Input

	ToMSTIDataConnectorOutput() MSTIDataConnectorOutput
	ToMSTIDataConnectorOutputWithContext(ctx context.Context) MSTIDataConnectorOutput
}

func (*MSTIDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnector)(nil)).Elem()
}

func (i *MSTIDataConnector) ToMSTIDataConnectorOutput() MSTIDataConnectorOutput {
	return i.ToMSTIDataConnectorOutputWithContext(context.Background())
}

func (i *MSTIDataConnector) ToMSTIDataConnectorOutputWithContext(ctx context.Context) MSTIDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorOutput)
}

type MSTIDataConnectorOutput struct{ *pulumi.OutputState }

func (MSTIDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnector)(nil)).Elem()
}

func (o MSTIDataConnectorOutput) ToMSTIDataConnectorOutput() MSTIDataConnectorOutput {
	return o
}

func (o MSTIDataConnectorOutput) ToMSTIDataConnectorOutputWithContext(ctx context.Context) MSTIDataConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MSTIDataConnectorOutput{})
}
