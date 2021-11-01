


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TiTaxiiDataConnector struct {
	pulumi.CustomResourceState

	CollectionId        pulumi.StringPtrOutput                      `pulumi:"collectionId"`
	DataTypes           TiTaxiiDataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	Etag                pulumi.StringPtrOutput                      `pulumi:"etag"`
	FriendlyName        pulumi.StringPtrOutput                      `pulumi:"friendlyName"`
	Kind                pulumi.StringOutput                         `pulumi:"kind"`
	Name                pulumi.StringOutput                         `pulumi:"name"`
	Password            pulumi.StringPtrOutput                      `pulumi:"password"`
	PollingFrequency    pulumi.StringOutput                         `pulumi:"pollingFrequency"`
	SystemData          SystemDataResponseOutput                    `pulumi:"systemData"`
	TaxiiLookbackPeriod pulumi.StringPtrOutput                      `pulumi:"taxiiLookbackPeriod"`
	TaxiiServer         pulumi.StringPtrOutput                      `pulumi:"taxiiServer"`
	TenantId            pulumi.StringOutput                         `pulumi:"tenantId"`
	Type                pulumi.StringOutput                         `pulumi:"type"`
	UserName            pulumi.StringPtrOutput                      `pulumi:"userName"`
	WorkspaceId         pulumi.StringPtrOutput                      `pulumi:"workspaceId"`
}


func NewTiTaxiiDataConnector(ctx *pulumi.Context,
	name string, args *TiTaxiiDataConnectorArgs, opts ...pulumi.ResourceOption) (*TiTaxiiDataConnector, error) {
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
	if args.PollingFrequency == nil {
		return nil, errors.New("invalid value for required argument 'PollingFrequency'")
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
	args.Kind = pulumi.String("ThreatIntelligenceTaxii")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20210301preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20190101preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20200101:TiTaxiiDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource TiTaxiiDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20210301preview:TiTaxiiDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTiTaxiiDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TiTaxiiDataConnectorState, opts ...pulumi.ResourceOption) (*TiTaxiiDataConnector, error) {
	var resource TiTaxiiDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20210301preview:TiTaxiiDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type tiTaxiiDataConnectorState struct {
}

type TiTaxiiDataConnectorState struct {
}

func (TiTaxiiDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*tiTaxiiDataConnectorState)(nil)).Elem()
}

type tiTaxiiDataConnectorArgs struct {
	CollectionId                        *string                       `pulumi:"collectionId"`
	DataConnectorId                     *string                       `pulumi:"dataConnectorId"`
	DataTypes                           TiTaxiiDataConnectorDataTypes `pulumi:"dataTypes"`
	Etag                                *string                       `pulumi:"etag"`
	FriendlyName                        *string                       `pulumi:"friendlyName"`
	Kind                                string                        `pulumi:"kind"`
	OperationalInsightsResourceProvider string                        `pulumi:"operationalInsightsResourceProvider"`
	Password                            *string                       `pulumi:"password"`
	PollingFrequency                    string                        `pulumi:"pollingFrequency"`
	ResourceGroupName                   string                        `pulumi:"resourceGroupName"`
	TaxiiLookbackPeriod                 *string                       `pulumi:"taxiiLookbackPeriod"`
	TaxiiServer                         *string                       `pulumi:"taxiiServer"`
	TenantId                            string                        `pulumi:"tenantId"`
	UserName                            *string                       `pulumi:"userName"`
	WorkspaceId                         *string                       `pulumi:"workspaceId"`
	WorkspaceName                       string                        `pulumi:"workspaceName"`
}


type TiTaxiiDataConnectorArgs struct {
	CollectionId                        pulumi.StringPtrInput
	DataConnectorId                     pulumi.StringPtrInput
	DataTypes                           TiTaxiiDataConnectorDataTypesInput
	Etag                                pulumi.StringPtrInput
	FriendlyName                        pulumi.StringPtrInput
	Kind                                pulumi.StringInput
	OperationalInsightsResourceProvider pulumi.StringInput
	Password                            pulumi.StringPtrInput
	PollingFrequency                    pulumi.StringInput
	ResourceGroupName                   pulumi.StringInput
	TaxiiLookbackPeriod                 pulumi.StringPtrInput
	TaxiiServer                         pulumi.StringPtrInput
	TenantId                            pulumi.StringInput
	UserName                            pulumi.StringPtrInput
	WorkspaceId                         pulumi.StringPtrInput
	WorkspaceName                       pulumi.StringInput
}

func (TiTaxiiDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tiTaxiiDataConnectorArgs)(nil)).Elem()
}

type TiTaxiiDataConnectorInput interface {
	pulumi.Input

	ToTiTaxiiDataConnectorOutput() TiTaxiiDataConnectorOutput
	ToTiTaxiiDataConnectorOutputWithContext(ctx context.Context) TiTaxiiDataConnectorOutput
}

func (*TiTaxiiDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((*TiTaxiiDataConnector)(nil))
}

func (i *TiTaxiiDataConnector) ToTiTaxiiDataConnectorOutput() TiTaxiiDataConnectorOutput {
	return i.ToTiTaxiiDataConnectorOutputWithContext(context.Background())
}

func (i *TiTaxiiDataConnector) ToTiTaxiiDataConnectorOutputWithContext(ctx context.Context) TiTaxiiDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiTaxiiDataConnectorOutput)
}

type TiTaxiiDataConnectorOutput struct{ *pulumi.OutputState }

func (TiTaxiiDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TiTaxiiDataConnector)(nil))
}

func (o TiTaxiiDataConnectorOutput) ToTiTaxiiDataConnectorOutput() TiTaxiiDataConnectorOutput {
	return o
}

func (o TiTaxiiDataConnectorOutput) ToTiTaxiiDataConnectorOutputWithContext(ctx context.Context) TiTaxiiDataConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TiTaxiiDataConnectorOutput{})
}
