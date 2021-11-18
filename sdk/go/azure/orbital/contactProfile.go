


package orbital

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContactProfile struct {
	pulumi.CustomResourceState

	AutoTrackingConfiguration    pulumi.StringPtrOutput                `pulumi:"autoTrackingConfiguration"`
	Etag                         pulumi.StringOutput                   `pulumi:"etag"`
	Links                        ContactProfileLinkResponseArrayOutput `pulumi:"links"`
	Location                     pulumi.StringOutput                   `pulumi:"location"`
	MinimumElevationDegrees      pulumi.Float64PtrOutput               `pulumi:"minimumElevationDegrees"`
	MinimumViableContactDuration pulumi.StringPtrOutput                `pulumi:"minimumViableContactDuration"`
	Name                         pulumi.StringOutput                   `pulumi:"name"`
	SystemData                   SystemDataResponseOutput              `pulumi:"systemData"`
	Tags                         pulumi.StringMapOutput                `pulumi:"tags"`
	Type                         pulumi.StringOutput                   `pulumi:"type"`
}


func NewContactProfile(ctx *pulumi.Context,
	name string, args *ContactProfileArgs, opts ...pulumi.ResourceOption) (*ContactProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Links == nil {
		return nil, errors.New("invalid value for required argument 'Links'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:orbital/v20210404preview:ContactProfile"),
		},
	})
	opts = append(opts, aliases)
	var resource ContactProfile
	err := ctx.RegisterResource("azure-native:orbital:ContactProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetContactProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContactProfileState, opts ...pulumi.ResourceOption) (*ContactProfile, error) {
	var resource ContactProfile
	err := ctx.ReadResource("azure-native:orbital:ContactProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type contactProfileState struct {
}

type ContactProfileState struct {
}

func (ContactProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*contactProfileState)(nil)).Elem()
}

type contactProfileArgs struct {
	AutoTrackingConfiguration    *AutoTrackingConfiguration `pulumi:"autoTrackingConfiguration"`
	ContactProfileName           *string                    `pulumi:"contactProfileName"`
	Links                        []ContactProfileLink       `pulumi:"links"`
	Location                     *string                    `pulumi:"location"`
	MinimumElevationDegrees      *float64                   `pulumi:"minimumElevationDegrees"`
	MinimumViableContactDuration *string                    `pulumi:"minimumViableContactDuration"`
	ResourceGroupName            string                     `pulumi:"resourceGroupName"`
	Tags                         map[string]string          `pulumi:"tags"`
}


type ContactProfileArgs struct {
	AutoTrackingConfiguration    AutoTrackingConfigurationPtrInput
	ContactProfileName           pulumi.StringPtrInput
	Links                        ContactProfileLinkArrayInput
	Location                     pulumi.StringPtrInput
	MinimumElevationDegrees      pulumi.Float64PtrInput
	MinimumViableContactDuration pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	Tags                         pulumi.StringMapInput
}

func (ContactProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contactProfileArgs)(nil)).Elem()
}

type ContactProfileInput interface {
	pulumi.Input

	ToContactProfileOutput() ContactProfileOutput
	ToContactProfileOutputWithContext(ctx context.Context) ContactProfileOutput
}

func (*ContactProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactProfile)(nil))
}

func (i *ContactProfile) ToContactProfileOutput() ContactProfileOutput {
	return i.ToContactProfileOutputWithContext(context.Background())
}

func (i *ContactProfile) ToContactProfileOutputWithContext(ctx context.Context) ContactProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactProfileOutput)
}

type ContactProfileOutput struct{ *pulumi.OutputState }

func (ContactProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactProfile)(nil))
}

func (o ContactProfileOutput) ToContactProfileOutput() ContactProfileOutput {
	return o
}

func (o ContactProfileOutput) ToContactProfileOutputWithContext(ctx context.Context) ContactProfileOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ContactProfileOutput{})
}
