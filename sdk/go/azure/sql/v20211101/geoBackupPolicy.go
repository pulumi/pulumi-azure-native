


package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GeoBackupPolicy struct {
	pulumi.CustomResourceState

	Kind        pulumi.StringOutput `pulumi:"kind"`
	Location    pulumi.StringOutput `pulumi:"location"`
	Name        pulumi.StringOutput `pulumi:"name"`
	State       pulumi.StringOutput `pulumi:"state"`
	StorageType pulumi.StringOutput `pulumi:"storageType"`
	Type        pulumi.StringOutput `pulumi:"type"`
}


func NewGeoBackupPolicy(ctx *pulumi.Context,
	name string, args *GeoBackupPolicyArgs, opts ...pulumi.ResourceOption) (*GeoBackupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:GeoBackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:GeoBackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:GeoBackupPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource GeoBackupPolicy
	err := ctx.RegisterResource("azure-native:sql/v20211101:GeoBackupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGeoBackupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GeoBackupPolicyState, opts ...pulumi.ResourceOption) (*GeoBackupPolicy, error) {
	var resource GeoBackupPolicy
	err := ctx.ReadResource("azure-native:sql/v20211101:GeoBackupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type geoBackupPolicyState struct {
}

type GeoBackupPolicyState struct {
}

func (GeoBackupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*geoBackupPolicyState)(nil)).Elem()
}

type geoBackupPolicyArgs struct {
	DatabaseName        string                   `pulumi:"databaseName"`
	GeoBackupPolicyName *string                  `pulumi:"geoBackupPolicyName"`
	ResourceGroupName   string                   `pulumi:"resourceGroupName"`
	ServerName          string                   `pulumi:"serverName"`
	State               GeoBackupPolicyStateEnum `pulumi:"state"`
}


type GeoBackupPolicyArgs struct {
	DatabaseName        pulumi.StringInput
	GeoBackupPolicyName pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	ServerName          pulumi.StringInput
	State               GeoBackupPolicyStateEnumInput
}

func (GeoBackupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*geoBackupPolicyArgs)(nil)).Elem()
}

type GeoBackupPolicyInput interface {
	pulumi.Input

	ToGeoBackupPolicyOutput() GeoBackupPolicyOutput
	ToGeoBackupPolicyOutputWithContext(ctx context.Context) GeoBackupPolicyOutput
}

func (*GeoBackupPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**GeoBackupPolicy)(nil)).Elem()
}

func (i *GeoBackupPolicy) ToGeoBackupPolicyOutput() GeoBackupPolicyOutput {
	return i.ToGeoBackupPolicyOutputWithContext(context.Background())
}

func (i *GeoBackupPolicy) ToGeoBackupPolicyOutputWithContext(ctx context.Context) GeoBackupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoBackupPolicyOutput)
}

type GeoBackupPolicyOutput struct{ *pulumi.OutputState }

func (GeoBackupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GeoBackupPolicy)(nil)).Elem()
}

func (o GeoBackupPolicyOutput) ToGeoBackupPolicyOutput() GeoBackupPolicyOutput {
	return o
}

func (o GeoBackupPolicyOutput) ToGeoBackupPolicyOutputWithContext(ctx context.Context) GeoBackupPolicyOutput {
	return o
}

func (o GeoBackupPolicyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *GeoBackupPolicy) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o GeoBackupPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GeoBackupPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o GeoBackupPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GeoBackupPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GeoBackupPolicyOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *GeoBackupPolicy) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o GeoBackupPolicyOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v *GeoBackupPolicy) pulumi.StringOutput { return v.StorageType }).(pulumi.StringOutput)
}

func (o GeoBackupPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GeoBackupPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GeoBackupPolicyOutput{})
}
