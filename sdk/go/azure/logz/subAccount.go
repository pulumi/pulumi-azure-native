


package logz

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SubAccount struct {
	pulumi.CustomResourceState

	Identity   IdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringOutput                 `pulumi:"location"`
	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties MonitorPropertiesResponseOutput     `pulumi:"properties"`
	SystemData SystemDataResponseOutput            `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput              `pulumi:"tags"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewSubAccount(ctx *pulumi.Context,
	name string, args *SubAccountArgs, opts ...pulumi.ResourceOption) (*SubAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MonitorName == nil {
		return nil, errors.New("invalid value for required argument 'MonitorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logz/v20201001:SubAccount"),
		},
		{
			Type: pulumi.String("azure-native:logz/v20201001preview:SubAccount"),
		},
		{
			Type: pulumi.String("azure-native:logz/v20220101preview:SubAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource SubAccount
	err := ctx.RegisterResource("azure-native:logz:SubAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSubAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubAccountState, opts ...pulumi.ResourceOption) (*SubAccount, error) {
	var resource SubAccount
	err := ctx.ReadResource("azure-native:logz:SubAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type subAccountState struct {
}

type SubAccountState struct {
}

func (SubAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*subAccountState)(nil)).Elem()
}

type subAccountArgs struct {
	Identity          *IdentityProperties `pulumi:"identity"`
	Location          *string             `pulumi:"location"`
	MonitorName       string              `pulumi:"monitorName"`
	Properties        *MonitorProperties  `pulumi:"properties"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	SubAccountName    *string             `pulumi:"subAccountName"`
	Tags              map[string]string   `pulumi:"tags"`
}


type SubAccountArgs struct {
	Identity          IdentityPropertiesPtrInput
	Location          pulumi.StringPtrInput
	MonitorName       pulumi.StringInput
	Properties        MonitorPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	SubAccountName    pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (SubAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subAccountArgs)(nil)).Elem()
}

type SubAccountInput interface {
	pulumi.Input

	ToSubAccountOutput() SubAccountOutput
	ToSubAccountOutputWithContext(ctx context.Context) SubAccountOutput
}

func (*SubAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**SubAccount)(nil)).Elem()
}

func (i *SubAccount) ToSubAccountOutput() SubAccountOutput {
	return i.ToSubAccountOutputWithContext(context.Background())
}

func (i *SubAccount) ToSubAccountOutputWithContext(ctx context.Context) SubAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubAccountOutput)
}

type SubAccountOutput struct{ *pulumi.OutputState }

func (SubAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubAccount)(nil)).Elem()
}

func (o SubAccountOutput) ToSubAccountOutput() SubAccountOutput {
	return o
}

func (o SubAccountOutput) ToSubAccountOutputWithContext(ctx context.Context) SubAccountOutput {
	return o
}

func (o SubAccountOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SubAccount) IdentityPropertiesResponsePtrOutput { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o SubAccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SubAccount) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SubAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SubAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SubAccountOutput) Properties() MonitorPropertiesResponseOutput {
	return o.ApplyT(func(v *SubAccount) MonitorPropertiesResponseOutput { return v.Properties }).(MonitorPropertiesResponseOutput)
}

func (o SubAccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SubAccount) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SubAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SubAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SubAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SubAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SubAccountOutput{})
}
