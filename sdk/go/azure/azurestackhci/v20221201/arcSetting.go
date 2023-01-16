


package v20221201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ArcSetting struct {
	pulumi.CustomResourceState

	AggregateState              pulumi.StringOutput                          `pulumi:"aggregateState"`
	ArcApplicationClientId      pulumi.StringPtrOutput                       `pulumi:"arcApplicationClientId"`
	ArcApplicationObjectId      pulumi.StringPtrOutput                       `pulumi:"arcApplicationObjectId"`
	ArcApplicationTenantId      pulumi.StringPtrOutput                       `pulumi:"arcApplicationTenantId"`
	ArcInstanceResourceGroup    pulumi.StringPtrOutput                       `pulumi:"arcInstanceResourceGroup"`
	ArcServicePrincipalObjectId pulumi.StringPtrOutput                       `pulumi:"arcServicePrincipalObjectId"`
	ConnectivityProperties      ArcConnectivityPropertiesResponseArrayOutput `pulumi:"connectivityProperties"`
	Name                        pulumi.StringOutput                          `pulumi:"name"`
	PerNodeDetails              PerNodeStateResponseArrayOutput              `pulumi:"perNodeDetails"`
	ProvisioningState           pulumi.StringOutput                          `pulumi:"provisioningState"`
	SystemData                  SystemDataResponseOutput                     `pulumi:"systemData"`
	Type                        pulumi.StringOutput                          `pulumi:"type"`
}


func NewArcSetting(ctx *pulumi.Context,
	name string, args *ArcSettingArgs, opts ...pulumi.ResourceOption) (*ArcSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210101preview:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901preview:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220101:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220301:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220501:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220901:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221001:ArcSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource ArcSetting
	err := ctx.RegisterResource("azure-native:azurestackhci/v20221201:ArcSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetArcSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArcSettingState, opts ...pulumi.ResourceOption) (*ArcSetting, error) {
	var resource ArcSetting
	err := ctx.ReadResource("azure-native:azurestackhci/v20221201:ArcSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type arcSettingState struct {
}

type ArcSettingState struct {
}

func (ArcSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*arcSettingState)(nil)).Elem()
}

type arcSettingArgs struct {
	ArcApplicationClientId      *string                     `pulumi:"arcApplicationClientId"`
	ArcApplicationObjectId      *string                     `pulumi:"arcApplicationObjectId"`
	ArcApplicationTenantId      *string                     `pulumi:"arcApplicationTenantId"`
	ArcInstanceResourceGroup    *string                     `pulumi:"arcInstanceResourceGroup"`
	ArcServicePrincipalObjectId *string                     `pulumi:"arcServicePrincipalObjectId"`
	ArcSettingName              *string                     `pulumi:"arcSettingName"`
	ClusterName                 string                      `pulumi:"clusterName"`
	ConnectivityProperties      []ArcConnectivityProperties `pulumi:"connectivityProperties"`
	ResourceGroupName           string                      `pulumi:"resourceGroupName"`
}


type ArcSettingArgs struct {
	ArcApplicationClientId      pulumi.StringPtrInput
	ArcApplicationObjectId      pulumi.StringPtrInput
	ArcApplicationTenantId      pulumi.StringPtrInput
	ArcInstanceResourceGroup    pulumi.StringPtrInput
	ArcServicePrincipalObjectId pulumi.StringPtrInput
	ArcSettingName              pulumi.StringPtrInput
	ClusterName                 pulumi.StringInput
	ConnectivityProperties      ArcConnectivityPropertiesArrayInput
	ResourceGroupName           pulumi.StringInput
}

func (ArcSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*arcSettingArgs)(nil)).Elem()
}

type ArcSettingInput interface {
	pulumi.Input

	ToArcSettingOutput() ArcSettingOutput
	ToArcSettingOutputWithContext(ctx context.Context) ArcSettingOutput
}

func (*ArcSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcSetting)(nil)).Elem()
}

func (i *ArcSetting) ToArcSettingOutput() ArcSettingOutput {
	return i.ToArcSettingOutputWithContext(context.Background())
}

func (i *ArcSetting) ToArcSettingOutputWithContext(ctx context.Context) ArcSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArcSettingOutput)
}

type ArcSettingOutput struct{ *pulumi.OutputState }

func (ArcSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcSetting)(nil)).Elem()
}

func (o ArcSettingOutput) ToArcSettingOutput() ArcSettingOutput {
	return o
}

func (o ArcSettingOutput) ToArcSettingOutputWithContext(ctx context.Context) ArcSettingOutput {
	return o
}

func (o ArcSettingOutput) AggregateState() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringOutput { return v.AggregateState }).(pulumi.StringOutput)
}

func (o ArcSettingOutput) ArcApplicationClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringPtrOutput { return v.ArcApplicationClientId }).(pulumi.StringPtrOutput)
}

func (o ArcSettingOutput) ArcApplicationObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringPtrOutput { return v.ArcApplicationObjectId }).(pulumi.StringPtrOutput)
}

func (o ArcSettingOutput) ArcApplicationTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringPtrOutput { return v.ArcApplicationTenantId }).(pulumi.StringPtrOutput)
}

func (o ArcSettingOutput) ArcInstanceResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringPtrOutput { return v.ArcInstanceResourceGroup }).(pulumi.StringPtrOutput)
}

func (o ArcSettingOutput) ArcServicePrincipalObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringPtrOutput { return v.ArcServicePrincipalObjectId }).(pulumi.StringPtrOutput)
}

func (o ArcSettingOutput) ConnectivityProperties() ArcConnectivityPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *ArcSetting) ArcConnectivityPropertiesResponseArrayOutput { return v.ConnectivityProperties }).(ArcConnectivityPropertiesResponseArrayOutput)
}

func (o ArcSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ArcSettingOutput) PerNodeDetails() PerNodeStateResponseArrayOutput {
	return o.ApplyT(func(v *ArcSetting) PerNodeStateResponseArrayOutput { return v.PerNodeDetails }).(PerNodeStateResponseArrayOutput)
}

func (o ArcSettingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ArcSettingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ArcSetting) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ArcSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ArcSettingOutput{})
}
