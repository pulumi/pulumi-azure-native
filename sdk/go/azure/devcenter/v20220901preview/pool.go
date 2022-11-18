


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Pool struct {
	pulumi.CustomResourceState

	DevBoxDefinitionName  pulumi.StringOutput      `pulumi:"devBoxDefinitionName"`
	LicenseType           pulumi.StringOutput      `pulumi:"licenseType"`
	LocalAdministrator    pulumi.StringOutput      `pulumi:"localAdministrator"`
	Location              pulumi.StringOutput      `pulumi:"location"`
	Name                  pulumi.StringOutput      `pulumi:"name"`
	NetworkConnectionName pulumi.StringOutput      `pulumi:"networkConnectionName"`
	ProvisioningState     pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData            SystemDataResponseOutput `pulumi:"systemData"`
	Tags                  pulumi.StringMapOutput   `pulumi:"tags"`
	Type                  pulumi.StringOutput      `pulumi:"type"`
}


func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DevBoxDefinitionName == nil {
		return nil, errors.New("invalid value for required argument 'DevBoxDefinitionName'")
	}
	if args.LicenseType == nil {
		return nil, errors.New("invalid value for required argument 'LicenseType'")
	}
	if args.LocalAdministrator == nil {
		return nil, errors.New("invalid value for required argument 'LocalAdministrator'")
	}
	if args.NetworkConnectionName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkConnectionName'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221012preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221111preview:Pool"),
		},
	})
	opts = append(opts, aliases)
	var resource Pool
	err := ctx.RegisterResource("azure-native:devcenter/v20220901preview:Pool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolState, opts ...pulumi.ResourceOption) (*Pool, error) {
	var resource Pool
	err := ctx.ReadResource("azure-native:devcenter/v20220901preview:Pool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type poolState struct {
}

type PoolState struct {
}

func (PoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolState)(nil)).Elem()
}

type poolArgs struct {
	DevBoxDefinitionName  string            `pulumi:"devBoxDefinitionName"`
	LicenseType           string            `pulumi:"licenseType"`
	LocalAdministrator    string            `pulumi:"localAdministrator"`
	Location              *string           `pulumi:"location"`
	NetworkConnectionName string            `pulumi:"networkConnectionName"`
	PoolName              *string           `pulumi:"poolName"`
	ProjectName           string            `pulumi:"projectName"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	Tags                  map[string]string `pulumi:"tags"`
}


type PoolArgs struct {
	DevBoxDefinitionName  pulumi.StringInput
	LicenseType           pulumi.StringInput
	LocalAdministrator    pulumi.StringInput
	Location              pulumi.StringPtrInput
	NetworkConnectionName pulumi.StringInput
	PoolName              pulumi.StringPtrInput
	ProjectName           pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	Tags                  pulumi.StringMapInput
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}

type PoolInput interface {
	pulumi.Input

	ToPoolOutput() PoolOutput
	ToPoolOutputWithContext(ctx context.Context) PoolOutput
}

func (*Pool) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (i *Pool) ToPoolOutput() PoolOutput {
	return i.ToPoolOutputWithContext(context.Background())
}

func (i *Pool) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolOutput)
}

type PoolOutput struct{ *pulumi.OutputState }

func (PoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (o PoolOutput) ToPoolOutput() PoolOutput {
	return o
}

func (o PoolOutput) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return o
}

func (o PoolOutput) DevBoxDefinitionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.DevBoxDefinitionName }).(pulumi.StringOutput)
}

func (o PoolOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.LicenseType }).(pulumi.StringOutput)
}

func (o PoolOutput) LocalAdministrator() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.LocalAdministrator }).(pulumi.StringOutput)
}

func (o PoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PoolOutput) NetworkConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.NetworkConnectionName }).(pulumi.StringOutput)
}

func (o PoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Pool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PoolOutput{})
}
