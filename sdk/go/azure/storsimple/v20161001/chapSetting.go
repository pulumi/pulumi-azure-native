


package v20161001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ChapSetting struct {
	pulumi.CustomResourceState

	Name     pulumi.StringOutput                     `pulumi:"name"`
	Password AsymmetricEncryptedSecretResponseOutput `pulumi:"password"`
	Type     pulumi.StringOutput                     `pulumi:"type"`
}


func NewChapSetting(ctx *pulumi.Context,
	name string, args *ChapSettingArgs, opts ...pulumi.ResourceOption) (*ChapSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ChapSetting
	err := ctx.RegisterResource("azure-native:storsimple/v20161001:ChapSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetChapSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChapSettingState, opts ...pulumi.ResourceOption) (*ChapSetting, error) {
	var resource ChapSetting
	err := ctx.ReadResource("azure-native:storsimple/v20161001:ChapSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type chapSettingState struct {
}

type ChapSettingState struct {
}

func (ChapSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*chapSettingState)(nil)).Elem()
}

type chapSettingArgs struct {
	ChapUserName      *string                   `pulumi:"chapUserName"`
	DeviceName        string                    `pulumi:"deviceName"`
	ManagerName       string                    `pulumi:"managerName"`
	Password          AsymmetricEncryptedSecret `pulumi:"password"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
}


type ChapSettingArgs struct {
	ChapUserName      pulumi.StringPtrInput
	DeviceName        pulumi.StringInput
	ManagerName       pulumi.StringInput
	Password          AsymmetricEncryptedSecretInput
	ResourceGroupName pulumi.StringInput
}

func (ChapSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*chapSettingArgs)(nil)).Elem()
}

type ChapSettingInput interface {
	pulumi.Input

	ToChapSettingOutput() ChapSettingOutput
	ToChapSettingOutputWithContext(ctx context.Context) ChapSettingOutput
}

func (*ChapSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**ChapSetting)(nil)).Elem()
}

func (i *ChapSetting) ToChapSettingOutput() ChapSettingOutput {
	return i.ToChapSettingOutputWithContext(context.Background())
}

func (i *ChapSetting) ToChapSettingOutputWithContext(ctx context.Context) ChapSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChapSettingOutput)
}

type ChapSettingOutput struct{ *pulumi.OutputState }

func (ChapSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChapSetting)(nil)).Elem()
}

func (o ChapSettingOutput) ToChapSettingOutput() ChapSettingOutput {
	return o
}

func (o ChapSettingOutput) ToChapSettingOutputWithContext(ctx context.Context) ChapSettingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ChapSettingOutput{})
}
