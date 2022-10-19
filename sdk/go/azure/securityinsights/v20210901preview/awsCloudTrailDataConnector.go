


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AwsCloudTrailDataConnector struct {
	pulumi.CustomResourceState

	AwsRoleArn pulumi.StringPtrOutput                            `pulumi:"awsRoleArn"`
	DataTypes  AwsCloudTrailDataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	Etag       pulumi.StringPtrOutput                            `pulumi:"etag"`
	Kind       pulumi.StringOutput                               `pulumi:"kind"`
	Name       pulumi.StringOutput                               `pulumi:"name"`
	SystemData SystemDataResponseOutput                          `pulumi:"systemData"`
	Type       pulumi.StringOutput                               `pulumi:"type"`
}


func NewAwsCloudTrailDataConnector(ctx *pulumi.Context,
	name string, args *AwsCloudTrailDataConnectorArgs, opts ...pulumi.ResourceOption) (*AwsCloudTrailDataConnector, error) {
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
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("AmazonWebServicesCloudTrail")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:AwsCloudTrailDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource AwsCloudTrailDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20210901preview:AwsCloudTrailDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAwsCloudTrailDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AwsCloudTrailDataConnectorState, opts ...pulumi.ResourceOption) (*AwsCloudTrailDataConnector, error) {
	var resource AwsCloudTrailDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20210901preview:AwsCloudTrailDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type awsCloudTrailDataConnectorState struct {
}

type AwsCloudTrailDataConnectorState struct {
}

func (AwsCloudTrailDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*awsCloudTrailDataConnectorState)(nil)).Elem()
}

type awsCloudTrailDataConnectorArgs struct {
	AwsRoleArn        *string                             `pulumi:"awsRoleArn"`
	DataConnectorId   *string                             `pulumi:"dataConnectorId"`
	DataTypes         AwsCloudTrailDataConnectorDataTypes `pulumi:"dataTypes"`
	Kind              string                              `pulumi:"kind"`
	ResourceGroupName string                              `pulumi:"resourceGroupName"`
	WorkspaceName     string                              `pulumi:"workspaceName"`
}


type AwsCloudTrailDataConnectorArgs struct {
	AwsRoleArn        pulumi.StringPtrInput
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         AwsCloudTrailDataConnectorDataTypesInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (AwsCloudTrailDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awsCloudTrailDataConnectorArgs)(nil)).Elem()
}

type AwsCloudTrailDataConnectorInput interface {
	pulumi.Input

	ToAwsCloudTrailDataConnectorOutput() AwsCloudTrailDataConnectorOutput
	ToAwsCloudTrailDataConnectorOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorOutput
}

func (*AwsCloudTrailDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsCloudTrailDataConnector)(nil)).Elem()
}

func (i *AwsCloudTrailDataConnector) ToAwsCloudTrailDataConnectorOutput() AwsCloudTrailDataConnectorOutput {
	return i.ToAwsCloudTrailDataConnectorOutputWithContext(context.Background())
}

func (i *AwsCloudTrailDataConnector) ToAwsCloudTrailDataConnectorOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCloudTrailDataConnectorOutput)
}

type AwsCloudTrailDataConnectorOutput struct{ *pulumi.OutputState }

func (AwsCloudTrailDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsCloudTrailDataConnector)(nil)).Elem()
}

func (o AwsCloudTrailDataConnectorOutput) ToAwsCloudTrailDataConnectorOutput() AwsCloudTrailDataConnectorOutput {
	return o
}

func (o AwsCloudTrailDataConnectorOutput) ToAwsCloudTrailDataConnectorOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorOutput {
	return o
}

func (o AwsCloudTrailDataConnectorOutput) AwsRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnector) pulumi.StringPtrOutput { return v.AwsRoleArn }).(pulumi.StringPtrOutput)
}

func (o AwsCloudTrailDataConnectorOutput) DataTypes() AwsCloudTrailDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnector) AwsCloudTrailDataConnectorDataTypesResponseOutput {
		return v.DataTypes
	}).(AwsCloudTrailDataConnectorDataTypesResponseOutput)
}

func (o AwsCloudTrailDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o AwsCloudTrailDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o AwsCloudTrailDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AwsCloudTrailDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AwsCloudTrailDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorOutput{})
}
