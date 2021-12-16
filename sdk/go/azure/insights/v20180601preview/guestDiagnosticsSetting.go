


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GuestDiagnosticsSetting struct {
	pulumi.CustomResourceState

	DataSources  DataSourceResponseArrayOutput `pulumi:"dataSources"`
	Location     pulumi.StringOutput           `pulumi:"location"`
	Name         pulumi.StringOutput           `pulumi:"name"`
	OsType       pulumi.StringPtrOutput        `pulumi:"osType"`
	ProxySetting pulumi.StringPtrOutput        `pulumi:"proxySetting"`
	Tags         pulumi.StringMapOutput        `pulumi:"tags"`
	Type         pulumi.StringOutput           `pulumi:"type"`
}


func NewGuestDiagnosticsSetting(ctx *pulumi.Context,
	name string, args *GuestDiagnosticsSettingArgs, opts ...pulumi.ResourceOption) (*GuestDiagnosticsSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:guestDiagnosticsSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource GuestDiagnosticsSetting
	err := ctx.RegisterResource("azure-native:insights/v20180601preview:guestDiagnosticsSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGuestDiagnosticsSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestDiagnosticsSettingState, opts ...pulumi.ResourceOption) (*GuestDiagnosticsSetting, error) {
	var resource GuestDiagnosticsSetting
	err := ctx.ReadResource("azure-native:insights/v20180601preview:guestDiagnosticsSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type guestDiagnosticsSettingState struct {
}

type GuestDiagnosticsSettingState struct {
}

func (GuestDiagnosticsSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*guestDiagnosticsSettingState)(nil)).Elem()
}

type guestDiagnosticsSettingArgs struct {
	DataSources            []DataSource      `pulumi:"dataSources"`
	DiagnosticSettingsName *string           `pulumi:"diagnosticSettingsName"`
	Location               *string           `pulumi:"location"`
	OsType                 *string           `pulumi:"osType"`
	ProxySetting           *string           `pulumi:"proxySetting"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	Tags                   map[string]string `pulumi:"tags"`
}


type GuestDiagnosticsSettingArgs struct {
	DataSources            DataSourceArrayInput
	DiagnosticSettingsName pulumi.StringPtrInput
	Location               pulumi.StringPtrInput
	OsType                 pulumi.StringPtrInput
	ProxySetting           pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (GuestDiagnosticsSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*guestDiagnosticsSettingArgs)(nil)).Elem()
}

type GuestDiagnosticsSettingInput interface {
	pulumi.Input

	ToGuestDiagnosticsSettingOutput() GuestDiagnosticsSettingOutput
	ToGuestDiagnosticsSettingOutputWithContext(ctx context.Context) GuestDiagnosticsSettingOutput
}

func (*GuestDiagnosticsSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestDiagnosticsSetting)(nil)).Elem()
}

func (i *GuestDiagnosticsSetting) ToGuestDiagnosticsSettingOutput() GuestDiagnosticsSettingOutput {
	return i.ToGuestDiagnosticsSettingOutputWithContext(context.Background())
}

func (i *GuestDiagnosticsSetting) ToGuestDiagnosticsSettingOutputWithContext(ctx context.Context) GuestDiagnosticsSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestDiagnosticsSettingOutput)
}

type GuestDiagnosticsSettingOutput struct{ *pulumi.OutputState }

func (GuestDiagnosticsSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestDiagnosticsSetting)(nil)).Elem()
}

func (o GuestDiagnosticsSettingOutput) ToGuestDiagnosticsSettingOutput() GuestDiagnosticsSettingOutput {
	return o
}

func (o GuestDiagnosticsSettingOutput) ToGuestDiagnosticsSettingOutputWithContext(ctx context.Context) GuestDiagnosticsSettingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GuestDiagnosticsSettingOutput{})
}
