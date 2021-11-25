


package orbital

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Spacecraft struct {
	pulumi.CustomResourceState

	AuthorizationStatus         pulumi.StringOutput               `pulumi:"authorizationStatus"`
	AuthorizationStatusExtended pulumi.StringOutput               `pulumi:"authorizationStatusExtended"`
	Etag                        pulumi.StringOutput               `pulumi:"etag"`
	Links                       SpacecraftLinkResponseArrayOutput `pulumi:"links"`
	Location                    pulumi.StringOutput               `pulumi:"location"`
	Name                        pulumi.StringOutput               `pulumi:"name"`
	NoradId                     pulumi.StringOutput               `pulumi:"noradId"`
	SystemData                  SystemDataResponseOutput          `pulumi:"systemData"`
	Tags                        pulumi.StringMapOutput            `pulumi:"tags"`
	TitleLine                   pulumi.StringPtrOutput            `pulumi:"titleLine"`
	TleLine1                    pulumi.StringPtrOutput            `pulumi:"tleLine1"`
	TleLine2                    pulumi.StringPtrOutput            `pulumi:"tleLine2"`
	Type                        pulumi.StringOutput               `pulumi:"type"`
}


func NewSpacecraft(ctx *pulumi.Context,
	name string, args *SpacecraftArgs, opts ...pulumi.ResourceOption) (*Spacecraft, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NoradId == nil {
		return nil, errors.New("invalid value for required argument 'NoradId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:orbital/v20210404preview:Spacecraft"),
		},
	})
	opts = append(opts, aliases)
	var resource Spacecraft
	err := ctx.RegisterResource("azure-native:orbital:Spacecraft", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSpacecraft(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpacecraftState, opts ...pulumi.ResourceOption) (*Spacecraft, error) {
	var resource Spacecraft
	err := ctx.ReadResource("azure-native:orbital:Spacecraft", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type spacecraftState struct {
}

type SpacecraftState struct {
}

func (SpacecraftState) ElementType() reflect.Type {
	return reflect.TypeOf((*spacecraftState)(nil)).Elem()
}

type spacecraftArgs struct {
	Links             []SpacecraftLink  `pulumi:"links"`
	Location          *string           `pulumi:"location"`
	NoradId           string            `pulumi:"noradId"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	SpacecraftName    *string           `pulumi:"spacecraftName"`
	Tags              map[string]string `pulumi:"tags"`
	TitleLine         *string           `pulumi:"titleLine"`
	TleLine1          *string           `pulumi:"tleLine1"`
	TleLine2          *string           `pulumi:"tleLine2"`
}


type SpacecraftArgs struct {
	Links             SpacecraftLinkArrayInput
	Location          pulumi.StringPtrInput
	NoradId           pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SpacecraftName    pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	TitleLine         pulumi.StringPtrInput
	TleLine1          pulumi.StringPtrInput
	TleLine2          pulumi.StringPtrInput
}

func (SpacecraftArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spacecraftArgs)(nil)).Elem()
}

type SpacecraftInput interface {
	pulumi.Input

	ToSpacecraftOutput() SpacecraftOutput
	ToSpacecraftOutputWithContext(ctx context.Context) SpacecraftOutput
}

func (*Spacecraft) ElementType() reflect.Type {
	return reflect.TypeOf((*Spacecraft)(nil))
}

func (i *Spacecraft) ToSpacecraftOutput() SpacecraftOutput {
	return i.ToSpacecraftOutputWithContext(context.Background())
}

func (i *Spacecraft) ToSpacecraftOutputWithContext(ctx context.Context) SpacecraftOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpacecraftOutput)
}

type SpacecraftOutput struct{ *pulumi.OutputState }

func (SpacecraftOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Spacecraft)(nil))
}

func (o SpacecraftOutput) ToSpacecraftOutput() SpacecraftOutput {
	return o
}

func (o SpacecraftOutput) ToSpacecraftOutputWithContext(ctx context.Context) SpacecraftOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SpacecraftOutput{})
}
