


package v20200301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkSecurityGroup struct {
	pulumi.CustomResourceState

	DefaultSecurityRules SecurityRuleResponseArrayOutput     `pulumi:"defaultSecurityRules"`
	Etag                 pulumi.StringOutput                 `pulumi:"etag"`
	FlowLogs             FlowLogResponseArrayOutput          `pulumi:"flowLogs"`
	Location             pulumi.StringPtrOutput              `pulumi:"location"`
	Name                 pulumi.StringOutput                 `pulumi:"name"`
	NetworkInterfaces    NetworkInterfaceResponseArrayOutput `pulumi:"networkInterfaces"`
	ProvisioningState    pulumi.StringOutput                 `pulumi:"provisioningState"`
	ResourceGuid         pulumi.StringOutput                 `pulumi:"resourceGuid"`
	SecurityRules        SecurityRuleResponseArrayOutput     `pulumi:"securityRules"`
	Subnets              SubnetResponseArrayOutput           `pulumi:"subnets"`
	Tags                 pulumi.StringMapOutput              `pulumi:"tags"`
	Type                 pulumi.StringOutput                 `pulumi:"type"`
}


func NewNetworkSecurityGroup(ctx *pulumi.Context,
	name string, args *NetworkSecurityGroupArgs, opts ...pulumi.ResourceOption) (*NetworkSecurityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:NetworkSecurityGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkSecurityGroup
	err := ctx.RegisterResource("azure-native:network/v20200301:NetworkSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSecurityGroupState, opts ...pulumi.ResourceOption) (*NetworkSecurityGroup, error) {
	var resource NetworkSecurityGroup
	err := ctx.ReadResource("azure-native:network/v20200301:NetworkSecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type networkSecurityGroupState struct {
}

type NetworkSecurityGroupState struct {
}

func (NetworkSecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityGroupState)(nil)).Elem()
}

type networkSecurityGroupArgs struct {
	Id                       *string            `pulumi:"id"`
	Location                 *string            `pulumi:"location"`
	NetworkSecurityGroupName *string            `pulumi:"networkSecurityGroupName"`
	ResourceGroupName        string             `pulumi:"resourceGroupName"`
	SecurityRules            []SecurityRuleType `pulumi:"securityRules"`
	Tags                     map[string]string  `pulumi:"tags"`
}


type NetworkSecurityGroupArgs struct {
	Id                       pulumi.StringPtrInput
	Location                 pulumi.StringPtrInput
	NetworkSecurityGroupName pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	SecurityRules            SecurityRuleTypeArrayInput
	Tags                     pulumi.StringMapInput
}

func (NetworkSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityGroupArgs)(nil)).Elem()
}

type NetworkSecurityGroupInput interface {
	pulumi.Input

	ToNetworkSecurityGroupOutput() NetworkSecurityGroupOutput
	ToNetworkSecurityGroupOutputWithContext(ctx context.Context) NetworkSecurityGroupOutput
}

func (*NetworkSecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroup)(nil))
}

func (i *NetworkSecurityGroup) ToNetworkSecurityGroupOutput() NetworkSecurityGroupOutput {
	return i.ToNetworkSecurityGroupOutputWithContext(context.Background())
}

func (i *NetworkSecurityGroup) ToNetworkSecurityGroupOutputWithContext(ctx context.Context) NetworkSecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupOutput)
}

type NetworkSecurityGroupOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroup)(nil))
}

func (o NetworkSecurityGroupOutput) ToNetworkSecurityGroupOutput() NetworkSecurityGroupOutput {
	return o
}

func (o NetworkSecurityGroupOutput) ToNetworkSecurityGroupOutputWithContext(ctx context.Context) NetworkSecurityGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkSecurityGroupOutput{})
}
