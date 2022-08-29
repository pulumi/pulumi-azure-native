


package v20200605preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VmmServer struct {
	pulumi.CustomResourceState

	ConnectionStatus  pulumi.StringOutput                             `pulumi:"connectionStatus"`
	Credentials       VMMServerPropertiesResponseCredentialsPtrOutput `pulumi:"credentials"`
	ErrorMessage      pulumi.StringOutput                             `pulumi:"errorMessage"`
	ExtendedLocation  ExtendedLocationResponseOutput                  `pulumi:"extendedLocation"`
	Fqdn              pulumi.StringOutput                             `pulumi:"fqdn"`
	Location          pulumi.StringOutput                             `pulumi:"location"`
	Name              pulumi.StringOutput                             `pulumi:"name"`
	Port              pulumi.IntPtrOutput                             `pulumi:"port"`
	ProvisioningState pulumi.StringOutput                             `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput                        `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                          `pulumi:"tags"`
	Type              pulumi.StringOutput                             `pulumi:"type"`
	Uuid              pulumi.StringOutput                             `pulumi:"uuid"`
	Version           pulumi.StringOutput                             `pulumi:"version"`
}


func NewVmmServer(ctx *pulumi.Context,
	name string, args *VmmServerArgs, opts ...pulumi.ResourceOption) (*VmmServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.Fqdn == nil {
		return nil, errors.New("invalid value for required argument 'Fqdn'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:scvmm:VmmServer"),
		},
	})
	opts = append(opts, aliases)
	var resource VmmServer
	err := ctx.RegisterResource("azure-native:scvmm/v20200605preview:VmmServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVmmServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VmmServerState, opts ...pulumi.ResourceOption) (*VmmServer, error) {
	var resource VmmServer
	err := ctx.ReadResource("azure-native:scvmm/v20200605preview:VmmServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type vmmServerState struct {
}

type VmmServerState struct {
}

func (VmmServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmmServerState)(nil)).Elem()
}

type vmmServerArgs struct {
	Credentials       *VMMServerPropertiesCredentials `pulumi:"credentials"`
	ExtendedLocation  ExtendedLocation                `pulumi:"extendedLocation"`
	Fqdn              string                          `pulumi:"fqdn"`
	Location          *string                         `pulumi:"location"`
	Port              *int                            `pulumi:"port"`
	ResourceGroupName string                          `pulumi:"resourceGroupName"`
	Tags              map[string]string               `pulumi:"tags"`
	VmmServerName     *string                         `pulumi:"vmmServerName"`
}


type VmmServerArgs struct {
	Credentials       VMMServerPropertiesCredentialsPtrInput
	ExtendedLocation  ExtendedLocationInput
	Fqdn              pulumi.StringInput
	Location          pulumi.StringPtrInput
	Port              pulumi.IntPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VmmServerName     pulumi.StringPtrInput
}

func (VmmServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmmServerArgs)(nil)).Elem()
}

type VmmServerInput interface {
	pulumi.Input

	ToVmmServerOutput() VmmServerOutput
	ToVmmServerOutputWithContext(ctx context.Context) VmmServerOutput
}

func (*VmmServer) ElementType() reflect.Type {
	return reflect.TypeOf((**VmmServer)(nil)).Elem()
}

func (i *VmmServer) ToVmmServerOutput() VmmServerOutput {
	return i.ToVmmServerOutputWithContext(context.Background())
}

func (i *VmmServer) ToVmmServerOutputWithContext(ctx context.Context) VmmServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmmServerOutput)
}

type VmmServerOutput struct{ *pulumi.OutputState }

func (VmmServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmmServer)(nil)).Elem()
}

func (o VmmServerOutput) ToVmmServerOutput() VmmServerOutput {
	return o
}

func (o VmmServerOutput) ToVmmServerOutputWithContext(ctx context.Context) VmmServerOutput {
	return o
}

func (o VmmServerOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o VmmServerOutput) Credentials() VMMServerPropertiesResponseCredentialsPtrOutput {
	return o.ApplyT(func(v *VmmServer) VMMServerPropertiesResponseCredentialsPtrOutput { return v.Credentials }).(VMMServerPropertiesResponseCredentialsPtrOutput)
}

func (o VmmServerOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o VmmServerOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *VmmServer) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

func (o VmmServerOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

func (o VmmServerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VmmServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VmmServerOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o VmmServerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VmmServerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VmmServer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VmmServerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VmmServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VmmServerOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func (o VmmServerOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VmmServerOutput{})
}
