


package portal

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type UserSettings struct {
	pulumi.CustomResourceState

	Properties UserPropertiesResponseOutput `pulumi:"properties"`
}


func NewUserSettings(ctx *pulumi.Context,
	name string, args *UserSettingsArgs, opts ...pulumi.ResourceOption) (*UserSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:portal:UserSettings"),
		},
		{
			Type: pulumi.String("azure-native:portal/v20181001:UserSettings"),
		},
		{
			Type: pulumi.String("azure-nextgen:portal/v20181001:UserSettings"),
		},
	})
	opts = append(opts, aliases)
	var resource UserSettings
	err := ctx.RegisterResource("azure-native:portal:UserSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetUserSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserSettingsState, opts ...pulumi.ResourceOption) (*UserSettings, error) {
	var resource UserSettings
	err := ctx.ReadResource("azure-native:portal:UserSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type userSettingsState struct {
}

type UserSettingsState struct {
}

func (UserSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*userSettingsState)(nil)).Elem()
}

type userSettingsArgs struct {
	Properties       UserProperties `pulumi:"properties"`
	UserSettingsName *string        `pulumi:"userSettingsName"`
}


type UserSettingsArgs struct {
	Properties       UserPropertiesInput
	UserSettingsName pulumi.StringPtrInput
}

func (UserSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userSettingsArgs)(nil)).Elem()
}

type UserSettingsInput interface {
	pulumi.Input

	ToUserSettingsOutput() UserSettingsOutput
	ToUserSettingsOutputWithContext(ctx context.Context) UserSettingsOutput
}

func (*UserSettings) ElementType() reflect.Type {
	return reflect.TypeOf((*UserSettings)(nil))
}

func (i *UserSettings) ToUserSettingsOutput() UserSettingsOutput {
	return i.ToUserSettingsOutputWithContext(context.Background())
}

func (i *UserSettings) ToUserSettingsOutputWithContext(ctx context.Context) UserSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSettingsOutput)
}

type UserSettingsOutput struct{ *pulumi.OutputState }

func (UserSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserSettings)(nil))
}

func (o UserSettingsOutput) ToUserSettingsOutput() UserSettingsOutput {
	return o
}

func (o UserSettingsOutput) ToUserSettingsOutputWithContext(ctx context.Context) UserSettingsOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(UserSettingsOutput{})
}
