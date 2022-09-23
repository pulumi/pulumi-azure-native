


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Instance struct {
	pulumi.CustomResourceState

	AccountName                 pulumi.StringOutput                          `pulumi:"accountName"`
	DiagnosticStorageProperties DiagnosticStoragePropertiesResponsePtrOutput `pulumi:"diagnosticStorageProperties"`
	EnableDiagnostics           pulumi.BoolPtrOutput                         `pulumi:"enableDiagnostics"`
	IotHubs                     IotHubSettingsResponseArrayOutput            `pulumi:"iotHubs"`
	Location                    pulumi.StringOutput                          `pulumi:"location"`
	Name                        pulumi.StringOutput                          `pulumi:"name"`
	ProvisioningState           pulumi.StringOutput                          `pulumi:"provisioningState"`
	SystemData                  SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                        pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                        pulumi.StringOutput                          `pulumi:"type"`
}


func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:deviceupdate:Instance"),
		},
		{
			Type: pulumi.String("azure-native:deviceupdate/v20200301preview:Instance"),
		},
		{
			Type: pulumi.String("azure-native:deviceupdate/v20221001:Instance"),
		},
	})
	opts = append(opts, aliases)
	var resource Instance
	err := ctx.RegisterResource("azure-native:deviceupdate/v20220401preview:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("azure-native:deviceupdate/v20220401preview:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type instanceState struct {
}

type InstanceState struct {
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	AccountName                 string                       `pulumi:"accountName"`
	DiagnosticStorageProperties *DiagnosticStorageProperties `pulumi:"diagnosticStorageProperties"`
	EnableDiagnostics           *bool                        `pulumi:"enableDiagnostics"`
	InstanceName                *string                      `pulumi:"instanceName"`
	IotHubs                     []IotHubSettings             `pulumi:"iotHubs"`
	Location                    *string                      `pulumi:"location"`
	ResourceGroupName           string                       `pulumi:"resourceGroupName"`
	Tags                        map[string]string            `pulumi:"tags"`
}


type InstanceArgs struct {
	AccountName                 pulumi.StringInput
	DiagnosticStorageProperties DiagnosticStoragePropertiesPtrInput
	EnableDiagnostics           pulumi.BoolPtrInput
	InstanceName                pulumi.StringPtrInput
	IotHubs                     IotHubSettingsArrayInput
	Location                    pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	Tags                        pulumi.StringMapInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func (o InstanceOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

func (o InstanceOutput) DiagnosticStorageProperties() DiagnosticStoragePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Instance) DiagnosticStoragePropertiesResponsePtrOutput { return v.DiagnosticStorageProperties }).(DiagnosticStoragePropertiesResponsePtrOutput)
}

func (o InstanceOutput) EnableDiagnostics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.EnableDiagnostics }).(pulumi.BoolPtrOutput)
}

func (o InstanceOutput) IotHubs() IotHubSettingsResponseArrayOutput {
	return o.ApplyT(func(v *Instance) IotHubSettingsResponseArrayOutput { return v.IotHubs }).(IotHubSettingsResponseArrayOutput)
}

func (o InstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o InstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o InstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Instance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o InstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o InstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceOutput{})
}
