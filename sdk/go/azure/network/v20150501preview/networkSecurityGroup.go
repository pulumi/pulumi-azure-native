


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type NetworkSecurityGroup struct {
	pulumi.CustomResourceState

	DefaultSecurityRules SecurityRuleResponseArrayOutput `pulumi:"defaultSecurityRules"`
	Etag                 pulumi.StringPtrOutput          `pulumi:"etag"`
	Location             pulumi.StringOutput             `pulumi:"location"`
	Name                 pulumi.StringOutput             `pulumi:"name"`
	NetworkInterfaces    SubResourceResponseArrayOutput  `pulumi:"networkInterfaces"`
	ProvisioningState    pulumi.StringPtrOutput          `pulumi:"provisioningState"`
	ResourceGuid         pulumi.StringPtrOutput          `pulumi:"resourceGuid"`
	SecurityRules        SecurityRuleResponseArrayOutput `pulumi:"securityRules"`
	Subnets              SubResourceResponseArrayOutput  `pulumi:"subnets"`
	Tags                 pulumi.StringMapOutput          `pulumi:"tags"`
	Type                 pulumi.StringOutput             `pulumi:"type"`
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
			Type: pulumi.String("azure-native:network/v20200301:NetworkSecurityGroup"),
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
		{
			Type: pulumi.String("azure-native:network/v20210801:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:NetworkSecurityGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkSecurityGroup
	err := ctx.RegisterResource("azure-native:network/v20150501preview:NetworkSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSecurityGroupState, opts ...pulumi.ResourceOption) (*NetworkSecurityGroup, error) {
	var resource NetworkSecurityGroup
	err := ctx.ReadResource("azure-native:network/v20150501preview:NetworkSecurityGroup", name, id, state, &resource, opts...)
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
	DefaultSecurityRules     []SecurityRuleType `pulumi:"defaultSecurityRules"`
	Location                 *string            `pulumi:"location"`
	NetworkInterfaces        []SubResource      `pulumi:"networkInterfaces"`
	NetworkSecurityGroupName *string            `pulumi:"networkSecurityGroupName"`
	ProvisioningState        *string            `pulumi:"provisioningState"`
	ResourceGroupName        string             `pulumi:"resourceGroupName"`
	ResourceGuid             *string            `pulumi:"resourceGuid"`
	SecurityRules            []SecurityRuleType `pulumi:"securityRules"`
	Subnets                  []SubResource      `pulumi:"subnets"`
	Tags                     map[string]string  `pulumi:"tags"`
}


type NetworkSecurityGroupArgs struct {
	DefaultSecurityRules     SecurityRuleTypeArrayInput
	Location                 pulumi.StringPtrInput
	NetworkInterfaces        SubResourceArrayInput
	NetworkSecurityGroupName pulumi.StringPtrInput
	ProvisioningState        pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	ResourceGuid             pulumi.StringPtrInput
	SecurityRules            SecurityRuleTypeArrayInput
	Subnets                  SubResourceArrayInput
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
	return reflect.TypeOf((**NetworkSecurityGroup)(nil)).Elem()
}

func (i *NetworkSecurityGroup) ToNetworkSecurityGroupOutput() NetworkSecurityGroupOutput {
	return i.ToNetworkSecurityGroupOutputWithContext(context.Background())
}

func (i *NetworkSecurityGroup) ToNetworkSecurityGroupOutputWithContext(ctx context.Context) NetworkSecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupOutput)
}

type NetworkSecurityGroupOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityGroup)(nil)).Elem()
}

func (o NetworkSecurityGroupOutput) ToNetworkSecurityGroupOutput() NetworkSecurityGroupOutput {
	return o
}

func (o NetworkSecurityGroupOutput) ToNetworkSecurityGroupOutputWithContext(ctx context.Context) NetworkSecurityGroupOutput {
	return o
}

func (o NetworkSecurityGroupOutput) DefaultSecurityRules() SecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) SecurityRuleResponseArrayOutput { return v.DefaultSecurityRules }).(SecurityRuleResponseArrayOutput)
}

func (o NetworkSecurityGroupOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o NetworkSecurityGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkSecurityGroupOutput) NetworkInterfaces() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) SubResourceResponseArrayOutput { return v.NetworkInterfaces }).(SubResourceResponseArrayOutput)
}

func (o NetworkSecurityGroupOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) pulumi.StringPtrOutput { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupOutput) SecurityRules() SecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) SecurityRuleResponseArrayOutput { return v.SecurityRules }).(SecurityRuleResponseArrayOutput)
}

func (o NetworkSecurityGroupOutput) Subnets() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) SubResourceResponseArrayOutput { return v.Subnets }).(SubResourceResponseArrayOutput)
}

func (o NetworkSecurityGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkSecurityGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkSecurityGroupOutput{})
}
