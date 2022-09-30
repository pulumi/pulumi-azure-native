


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NspProfile struct {
	pulumi.CustomResourceState

	AccessRulesVersion        pulumi.StringOutput    `pulumi:"accessRulesVersion"`
	DiagnosticSettingsVersion pulumi.StringOutput    `pulumi:"diagnosticSettingsVersion"`
	Location                  pulumi.StringPtrOutput `pulumi:"location"`
	Name                      pulumi.StringOutput    `pulumi:"name"`
	Tags                      pulumi.StringMapOutput `pulumi:"tags"`
	Type                      pulumi.StringOutput    `pulumi:"type"`
}


func NewNspProfile(ctx *pulumi.Context,
	name string, args *NspProfileArgs, opts ...pulumi.ResourceOption) (*NspProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkSecurityPerimeterName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSecurityPerimeterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NspProfile"),
		},
	})
	opts = append(opts, aliases)
	var resource NspProfile
	err := ctx.RegisterResource("azure-native:network/v20210201preview:NspProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNspProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NspProfileState, opts ...pulumi.ResourceOption) (*NspProfile, error) {
	var resource NspProfile
	err := ctx.ReadResource("azure-native:network/v20210201preview:NspProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type nspProfileState struct {
}

type NspProfileState struct {
}

func (NspProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*nspProfileState)(nil)).Elem()
}

type nspProfileArgs struct {
	Id                           *string           `pulumi:"id"`
	Location                     *string           `pulumi:"location"`
	Name                         *string           `pulumi:"name"`
	NetworkSecurityPerimeterName string            `pulumi:"networkSecurityPerimeterName"`
	ProfileName                  *string           `pulumi:"profileName"`
	ResourceGroupName            string            `pulumi:"resourceGroupName"`
	Tags                         map[string]string `pulumi:"tags"`
}


type NspProfileArgs struct {
	Id                           pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	Name                         pulumi.StringPtrInput
	NetworkSecurityPerimeterName pulumi.StringInput
	ProfileName                  pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	Tags                         pulumi.StringMapInput
}

func (NspProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nspProfileArgs)(nil)).Elem()
}

type NspProfileInput interface {
	pulumi.Input

	ToNspProfileOutput() NspProfileOutput
	ToNspProfileOutputWithContext(ctx context.Context) NspProfileOutput
}

func (*NspProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**NspProfile)(nil)).Elem()
}

func (i *NspProfile) ToNspProfileOutput() NspProfileOutput {
	return i.ToNspProfileOutputWithContext(context.Background())
}

func (i *NspProfile) ToNspProfileOutputWithContext(ctx context.Context) NspProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NspProfileOutput)
}

type NspProfileOutput struct{ *pulumi.OutputState }

func (NspProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NspProfile)(nil)).Elem()
}

func (o NspProfileOutput) ToNspProfileOutput() NspProfileOutput {
	return o
}

func (o NspProfileOutput) ToNspProfileOutputWithContext(ctx context.Context) NspProfileOutput {
	return o
}

func (o NspProfileOutput) AccessRulesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *NspProfile) pulumi.StringOutput { return v.AccessRulesVersion }).(pulumi.StringOutput)
}

func (o NspProfileOutput) DiagnosticSettingsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *NspProfile) pulumi.StringOutput { return v.DiagnosticSettingsVersion }).(pulumi.StringOutput)
}

func (o NspProfileOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NspProfile) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NspProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NspProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NspProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NspProfile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NspProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NspProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NspProfileOutput{})
}
