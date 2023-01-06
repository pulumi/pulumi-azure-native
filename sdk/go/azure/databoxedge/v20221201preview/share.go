


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Share struct {
	pulumi.CustomResourceState

	AccessProtocol     pulumi.StringOutput                  `pulumi:"accessProtocol"`
	AzureContainerInfo AzureContainerInfoResponsePtrOutput  `pulumi:"azureContainerInfo"`
	ClientAccessRights ClientAccessRightResponseArrayOutput `pulumi:"clientAccessRights"`
	DataPolicy         pulumi.StringPtrOutput               `pulumi:"dataPolicy"`
	Description        pulumi.StringPtrOutput               `pulumi:"description"`
	MonitoringStatus   pulumi.StringOutput                  `pulumi:"monitoringStatus"`
	Name               pulumi.StringOutput                  `pulumi:"name"`
	RefreshDetails     RefreshDetailsResponsePtrOutput      `pulumi:"refreshDetails"`
	ShareMappings      MountPointMapResponseArrayOutput     `pulumi:"shareMappings"`
	ShareStatus        pulumi.StringOutput                  `pulumi:"shareStatus"`
	SystemData         SystemDataResponseOutput             `pulumi:"systemData"`
	Type               pulumi.StringOutput                  `pulumi:"type"`
	UserAccessRights   UserAccessRightResponseArrayOutput   `pulumi:"userAccessRights"`
}


func NewShare(ctx *pulumi.Context,
	name string, args *ShareArgs, opts ...pulumi.ResourceOption) (*Share, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessProtocol == nil {
		return nil, errors.New("invalid value for required argument 'AccessProtocol'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.MonitoringStatus == nil {
		return nil, errors.New("invalid value for required argument 'MonitoringStatus'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareStatus == nil {
		return nil, errors.New("invalid value for required argument 'ShareStatus'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:Share"),
		},
	})
	opts = append(opts, aliases)
	var resource Share
	err := ctx.RegisterResource("azure-native:databoxedge/v20221201preview:Share", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShareState, opts ...pulumi.ResourceOption) (*Share, error) {
	var resource Share
	err := ctx.ReadResource("azure-native:databoxedge/v20221201preview:Share", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type shareState struct {
}

type ShareState struct {
}

func (ShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*shareState)(nil)).Elem()
}

type shareArgs struct {
	AccessProtocol     string              `pulumi:"accessProtocol"`
	AzureContainerInfo *AzureContainerInfo `pulumi:"azureContainerInfo"`
	ClientAccessRights []ClientAccessRight `pulumi:"clientAccessRights"`
	DataPolicy         *string             `pulumi:"dataPolicy"`
	Description        *string             `pulumi:"description"`
	DeviceName         string              `pulumi:"deviceName"`
	MonitoringStatus   string              `pulumi:"monitoringStatus"`
	Name               *string             `pulumi:"name"`
	RefreshDetails     *RefreshDetails     `pulumi:"refreshDetails"`
	ResourceGroupName  string              `pulumi:"resourceGroupName"`
	ShareStatus        string              `pulumi:"shareStatus"`
	UserAccessRights   []UserAccessRight   `pulumi:"userAccessRights"`
}


type ShareArgs struct {
	AccessProtocol     pulumi.StringInput
	AzureContainerInfo AzureContainerInfoPtrInput
	ClientAccessRights ClientAccessRightArrayInput
	DataPolicy         pulumi.StringPtrInput
	Description        pulumi.StringPtrInput
	DeviceName         pulumi.StringInput
	MonitoringStatus   pulumi.StringInput
	Name               pulumi.StringPtrInput
	RefreshDetails     RefreshDetailsPtrInput
	ResourceGroupName  pulumi.StringInput
	ShareStatus        pulumi.StringInput
	UserAccessRights   UserAccessRightArrayInput
}

func (ShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shareArgs)(nil)).Elem()
}

type ShareInput interface {
	pulumi.Input

	ToShareOutput() ShareOutput
	ToShareOutputWithContext(ctx context.Context) ShareOutput
}

func (*Share) ElementType() reflect.Type {
	return reflect.TypeOf((**Share)(nil)).Elem()
}

func (i *Share) ToShareOutput() ShareOutput {
	return i.ToShareOutputWithContext(context.Background())
}

func (i *Share) ToShareOutputWithContext(ctx context.Context) ShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareOutput)
}

type ShareOutput struct{ *pulumi.OutputState }

func (ShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Share)(nil)).Elem()
}

func (o ShareOutput) ToShareOutput() ShareOutput {
	return o
}

func (o ShareOutput) ToShareOutputWithContext(ctx context.Context) ShareOutput {
	return o
}

func (o ShareOutput) AccessProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.AccessProtocol }).(pulumi.StringOutput)
}

func (o ShareOutput) AzureContainerInfo() AzureContainerInfoResponsePtrOutput {
	return o.ApplyT(func(v *Share) AzureContainerInfoResponsePtrOutput { return v.AzureContainerInfo }).(AzureContainerInfoResponsePtrOutput)
}

func (o ShareOutput) ClientAccessRights() ClientAccessRightResponseArrayOutput {
	return o.ApplyT(func(v *Share) ClientAccessRightResponseArrayOutput { return v.ClientAccessRights }).(ClientAccessRightResponseArrayOutput)
}

func (o ShareOutput) DataPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Share) pulumi.StringPtrOutput { return v.DataPolicy }).(pulumi.StringPtrOutput)
}

func (o ShareOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Share) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ShareOutput) MonitoringStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.MonitoringStatus }).(pulumi.StringOutput)
}

func (o ShareOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ShareOutput) RefreshDetails() RefreshDetailsResponsePtrOutput {
	return o.ApplyT(func(v *Share) RefreshDetailsResponsePtrOutput { return v.RefreshDetails }).(RefreshDetailsResponsePtrOutput)
}

func (o ShareOutput) ShareMappings() MountPointMapResponseArrayOutput {
	return o.ApplyT(func(v *Share) MountPointMapResponseArrayOutput { return v.ShareMappings }).(MountPointMapResponseArrayOutput)
}

func (o ShareOutput) ShareStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.ShareStatus }).(pulumi.StringOutput)
}

func (o ShareOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Share) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ShareOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ShareOutput) UserAccessRights() UserAccessRightResponseArrayOutput {
	return o.ApplyT(func(v *Share) UserAccessRightResponseArrayOutput { return v.UserAccessRights }).(UserAccessRightResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ShareOutput{})
}
