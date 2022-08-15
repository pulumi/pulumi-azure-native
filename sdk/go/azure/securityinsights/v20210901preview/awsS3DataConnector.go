


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AwsS3DataConnector struct {
	pulumi.CustomResourceState

	DataTypes        AwsS3DataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	DestinationTable pulumi.StringOutput                       `pulumi:"destinationTable"`
	Etag             pulumi.StringPtrOutput                    `pulumi:"etag"`
	Kind             pulumi.StringOutput                       `pulumi:"kind"`
	Name             pulumi.StringOutput                       `pulumi:"name"`
	RoleArn          pulumi.StringOutput                       `pulumi:"roleArn"`
	SqsUrls          pulumi.StringArrayOutput                  `pulumi:"sqsUrls"`
	SystemData       SystemDataResponseOutput                  `pulumi:"systemData"`
	Type             pulumi.StringOutput                       `pulumi:"type"`
}


func NewAwsS3DataConnector(ctx *pulumi.Context,
	name string, args *AwsS3DataConnectorArgs, opts ...pulumi.ResourceOption) (*AwsS3DataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataTypes == nil {
		return nil, errors.New("invalid value for required argument 'DataTypes'")
	}
	if args.DestinationTable == nil {
		return nil, errors.New("invalid value for required argument 'DestinationTable'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	if args.SqsUrls == nil {
		return nil, errors.New("invalid value for required argument 'SqsUrls'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("AmazonWebServicesS3")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:AwsS3DataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource AwsS3DataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20210901preview:AwsS3DataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAwsS3DataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AwsS3DataConnectorState, opts ...pulumi.ResourceOption) (*AwsS3DataConnector, error) {
	var resource AwsS3DataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20210901preview:AwsS3DataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type awsS3DataConnectorState struct {
}

type AwsS3DataConnectorState struct {
}

func (AwsS3DataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*awsS3DataConnectorState)(nil)).Elem()
}

type awsS3DataConnectorArgs struct {
	DataConnectorId   *string                     `pulumi:"dataConnectorId"`
	DataTypes         AwsS3DataConnectorDataTypes `pulumi:"dataTypes"`
	DestinationTable  string                      `pulumi:"destinationTable"`
	Kind              string                      `pulumi:"kind"`
	ResourceGroupName string                      `pulumi:"resourceGroupName"`
	RoleArn           string                      `pulumi:"roleArn"`
	SqsUrls           []string                    `pulumi:"sqsUrls"`
	WorkspaceName     string                      `pulumi:"workspaceName"`
}


type AwsS3DataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         AwsS3DataConnectorDataTypesInput
	DestinationTable  pulumi.StringInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	RoleArn           pulumi.StringInput
	SqsUrls           pulumi.StringArrayInput
	WorkspaceName     pulumi.StringInput
}

func (AwsS3DataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awsS3DataConnectorArgs)(nil)).Elem()
}

type AwsS3DataConnectorInput interface {
	pulumi.Input

	ToAwsS3DataConnectorOutput() AwsS3DataConnectorOutput
	ToAwsS3DataConnectorOutputWithContext(ctx context.Context) AwsS3DataConnectorOutput
}

func (*AwsS3DataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsS3DataConnector)(nil)).Elem()
}

func (i *AwsS3DataConnector) ToAwsS3DataConnectorOutput() AwsS3DataConnectorOutput {
	return i.ToAwsS3DataConnectorOutputWithContext(context.Background())
}

func (i *AwsS3DataConnector) ToAwsS3DataConnectorOutputWithContext(ctx context.Context) AwsS3DataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsS3DataConnectorOutput)
}

type AwsS3DataConnectorOutput struct{ *pulumi.OutputState }

func (AwsS3DataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsS3DataConnector)(nil)).Elem()
}

func (o AwsS3DataConnectorOutput) ToAwsS3DataConnectorOutput() AwsS3DataConnectorOutput {
	return o
}

func (o AwsS3DataConnectorOutput) ToAwsS3DataConnectorOutputWithContext(ctx context.Context) AwsS3DataConnectorOutput {
	return o
}

func (o AwsS3DataConnectorOutput) DataTypes() AwsS3DataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) AwsS3DataConnectorDataTypesResponseOutput { return v.DataTypes }).(AwsS3DataConnectorDataTypesResponseOutput)
}

func (o AwsS3DataConnectorOutput) DestinationTable() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) pulumi.StringOutput { return v.DestinationTable }).(pulumi.StringOutput)
}

func (o AwsS3DataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o AwsS3DataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o AwsS3DataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AwsS3DataConnectorOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

func (o AwsS3DataConnectorOutput) SqsUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) pulumi.StringArrayOutput { return v.SqsUrls }).(pulumi.StringArrayOutput)
}

func (o AwsS3DataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AwsS3DataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AwsS3DataConnectorOutput{})
}
