


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MTPDataConnector struct {
	pulumi.CustomResourceState

	DataTypes  MTPDataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	Etag       pulumi.StringPtrOutput                  `pulumi:"etag"`
	Kind       pulumi.StringOutput                     `pulumi:"kind"`
	Name       pulumi.StringOutput                     `pulumi:"name"`
	SystemData SystemDataResponseOutput                `pulumi:"systemData"`
	TenantId   pulumi.StringOutput                     `pulumi:"tenantId"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewMTPDataConnector(ctx *pulumi.Context,
	name string, args *MTPDataConnectorArgs, opts ...pulumi.ResourceOption) (*MTPDataConnector, error) {
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
	args.Kind = pulumi.String("MicrosoftThreatProtection")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:MTPDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource MTPDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20210301preview:MTPDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMTPDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MTPDataConnectorState, opts ...pulumi.ResourceOption) (*MTPDataConnector, error) {
	var resource MTPDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20210301preview:MTPDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mtpdataConnectorState struct {
}

type MTPDataConnectorState struct {
}

func (MTPDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*mtpdataConnectorState)(nil)).Elem()
}

type mtpdataConnectorArgs struct {
	DataConnectorId                     *string                   `pulumi:"dataConnectorId"`
	DataTypes                           MTPDataConnectorDataTypes `pulumi:"dataTypes"`
	Etag                                *string                   `pulumi:"etag"`
	Kind                                string                    `pulumi:"kind"`
	OperationalInsightsResourceProvider string                    `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string                    `pulumi:"resourceGroupName"`
	TenantId                            string                    `pulumi:"tenantId"`
	WorkspaceName                       string                    `pulumi:"workspaceName"`
}


type MTPDataConnectorArgs struct {
	DataConnectorId                     pulumi.StringPtrInput
	DataTypes                           MTPDataConnectorDataTypesInput
	Etag                                pulumi.StringPtrInput
	Kind                                pulumi.StringInput
	OperationalInsightsResourceProvider pulumi.StringInput
	ResourceGroupName                   pulumi.StringInput
	TenantId                            pulumi.StringInput
	WorkspaceName                       pulumi.StringInput
}

func (MTPDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mtpdataConnectorArgs)(nil)).Elem()
}

type MTPDataConnectorInput interface {
	pulumi.Input

	ToMTPDataConnectorOutput() MTPDataConnectorOutput
	ToMTPDataConnectorOutputWithContext(ctx context.Context) MTPDataConnectorOutput
}

func (*MTPDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**MTPDataConnector)(nil)).Elem()
}

func (i *MTPDataConnector) ToMTPDataConnectorOutput() MTPDataConnectorOutput {
	return i.ToMTPDataConnectorOutputWithContext(context.Background())
}

func (i *MTPDataConnector) ToMTPDataConnectorOutputWithContext(ctx context.Context) MTPDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MTPDataConnectorOutput)
}

type MTPDataConnectorOutput struct{ *pulumi.OutputState }

func (MTPDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MTPDataConnector)(nil)).Elem()
}

func (o MTPDataConnectorOutput) ToMTPDataConnectorOutput() MTPDataConnectorOutput {
	return o
}

func (o MTPDataConnectorOutput) ToMTPDataConnectorOutputWithContext(ctx context.Context) MTPDataConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MTPDataConnectorOutput{})
}
