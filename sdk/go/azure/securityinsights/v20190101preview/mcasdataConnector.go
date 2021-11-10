


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MCASDataConnector struct {
	pulumi.CustomResourceState

	DataTypes MCASDataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	Etag      pulumi.StringPtrOutput                   `pulumi:"etag"`
	Kind      pulumi.StringOutput                      `pulumi:"kind"`
	Name      pulumi.StringOutput                      `pulumi:"name"`
	TenantId  pulumi.StringOutput                      `pulumi:"tenantId"`
	Type      pulumi.StringOutput                      `pulumi:"type"`
}


func NewMCASDataConnector(ctx *pulumi.Context,
	name string, args *MCASDataConnectorArgs, opts ...pulumi.ResourceOption) (*MCASDataConnector, error) {
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
	args.Kind = pulumi.String("MicrosoftCloudAppSecurity")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:MCASDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource MCASDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20190101preview:MCASDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMCASDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MCASDataConnectorState, opts ...pulumi.ResourceOption) (*MCASDataConnector, error) {
	var resource MCASDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20190101preview:MCASDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mcasdataConnectorState struct {
}

type MCASDataConnectorState struct {
}

func (MCASDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*mcasdataConnectorState)(nil)).Elem()
}

type mcasdataConnectorArgs struct {
	DataConnectorId                     *string                    `pulumi:"dataConnectorId"`
	DataTypes                           MCASDataConnectorDataTypes `pulumi:"dataTypes"`
	Etag                                *string                    `pulumi:"etag"`
	Kind                                string                     `pulumi:"kind"`
	OperationalInsightsResourceProvider string                     `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string                     `pulumi:"resourceGroupName"`
	TenantId                            string                     `pulumi:"tenantId"`
	WorkspaceName                       string                     `pulumi:"workspaceName"`
}


type MCASDataConnectorArgs struct {
	DataConnectorId                     pulumi.StringPtrInput
	DataTypes                           MCASDataConnectorDataTypesInput
	Etag                                pulumi.StringPtrInput
	Kind                                pulumi.StringInput
	OperationalInsightsResourceProvider pulumi.StringInput
	ResourceGroupName                   pulumi.StringInput
	TenantId                            pulumi.StringInput
	WorkspaceName                       pulumi.StringInput
}

func (MCASDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mcasdataConnectorArgs)(nil)).Elem()
}

type MCASDataConnectorInput interface {
	pulumi.Input

	ToMCASDataConnectorOutput() MCASDataConnectorOutput
	ToMCASDataConnectorOutputWithContext(ctx context.Context) MCASDataConnectorOutput
}

func (*MCASDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((*MCASDataConnector)(nil))
}

func (i *MCASDataConnector) ToMCASDataConnectorOutput() MCASDataConnectorOutput {
	return i.ToMCASDataConnectorOutputWithContext(context.Background())
}

func (i *MCASDataConnector) ToMCASDataConnectorOutputWithContext(ctx context.Context) MCASDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MCASDataConnectorOutput)
}

type MCASDataConnectorOutput struct{ *pulumi.OutputState }

func (MCASDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MCASDataConnector)(nil))
}

func (o MCASDataConnectorOutput) ToMCASDataConnectorOutput() MCASDataConnectorOutput {
	return o
}

func (o MCASDataConnectorOutput) ToMCASDataConnectorOutputWithContext(ctx context.Context) MCASDataConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MCASDataConnectorOutput{})
}
