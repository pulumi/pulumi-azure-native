


package portal

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type UserSettingsWithLocation struct {
	pulumi.CustomResourceState

	Properties UserPropertiesResponseOutput `pulumi:"properties"`
}


func NewUserSettingsWithLocation(ctx *pulumi.Context,
	name string, args *UserSettingsWithLocationArgs, opts ...pulumi.ResourceOption) (*UserSettingsWithLocation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:portal:UserSettingsWithLocation"),
		},
		{
			Type: pulumi.String("azure-native:portal/v20181001:UserSettingsWithLocation"),
		},
		{
			Type: pulumi.String("azure-nextgen:portal/v20181001:UserSettingsWithLocation"),
		},
	})
	opts = append(opts, aliases)
	var resource UserSettingsWithLocation
	err := ctx.RegisterResource("azure-native:portal:UserSettingsWithLocation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetUserSettingsWithLocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserSettingsWithLocationState, opts ...pulumi.ResourceOption) (*UserSettingsWithLocation, error) {
	var resource UserSettingsWithLocation
	err := ctx.ReadResource("azure-native:portal:UserSettingsWithLocation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type userSettingsWithLocationState struct {
}

type UserSettingsWithLocationState struct {
}

func (UserSettingsWithLocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*userSettingsWithLocationState)(nil)).Elem()
}

type userSettingsWithLocationArgs struct {
	Location         string         `pulumi:"location"`
	Properties       UserProperties `pulumi:"properties"`
	UserSettingsName *string        `pulumi:"userSettingsName"`
}


type UserSettingsWithLocationArgs struct {
	Location         pulumi.StringInput
	Properties       UserPropertiesInput
	UserSettingsName pulumi.StringPtrInput
}

func (UserSettingsWithLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userSettingsWithLocationArgs)(nil)).Elem()
}

type UserSettingsWithLocationInput interface {
	pulumi.Input

	ToUserSettingsWithLocationOutput() UserSettingsWithLocationOutput
	ToUserSettingsWithLocationOutputWithContext(ctx context.Context) UserSettingsWithLocationOutput
}

func (*UserSettingsWithLocation) ElementType() reflect.Type {
	return reflect.TypeOf((*UserSettingsWithLocation)(nil))
}

func (i *UserSettingsWithLocation) ToUserSettingsWithLocationOutput() UserSettingsWithLocationOutput {
	return i.ToUserSettingsWithLocationOutputWithContext(context.Background())
}

func (i *UserSettingsWithLocation) ToUserSettingsWithLocationOutputWithContext(ctx context.Context) UserSettingsWithLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSettingsWithLocationOutput)
}

type UserSettingsWithLocationOutput struct{ *pulumi.OutputState }

func (UserSettingsWithLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserSettingsWithLocation)(nil))
}

func (o UserSettingsWithLocationOutput) ToUserSettingsWithLocationOutput() UserSettingsWithLocationOutput {
	return o
}

func (o UserSettingsWithLocationOutput) ToUserSettingsWithLocationOutputWithContext(ctx context.Context) UserSettingsWithLocationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserSettingsWithLocationInput)(nil)).Elem(), &UserSettingsWithLocation{})
	pulumi.RegisterOutputType(UserSettingsWithLocationOutput{})
}
