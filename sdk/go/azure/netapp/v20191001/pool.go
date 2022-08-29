


package v20191001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Pool struct {
	pulumi.CustomResourceState

	Location          pulumi.StringOutput    `pulumi:"location"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	PoolId            pulumi.StringOutput    `pulumi:"poolId"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	ServiceLevel      pulumi.StringOutput    `pulumi:"serviceLevel"`
	Size              pulumi.Float64Output   `pulumi:"size"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	if isZero(args.ServiceLevel) {
		args.ServiceLevel = pulumi.String("Premium")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:netapp:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20170815:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190501:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190601:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190701:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190801:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191101:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200201:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200301:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200501:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200701:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201101:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210201:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210601:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210801:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20211001:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:Pool"),
		},
	})
	opts = append(opts, aliases)
	var resource Pool
	err := ctx.RegisterResource("azure-native:netapp/v20191001:Pool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolState, opts ...pulumi.ResourceOption) (*Pool, error) {
	var resource Pool
	err := ctx.ReadResource("azure-native:netapp/v20191001:Pool", name, id, state, &resource, opts...)
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
	AccountName       string            `pulumi:"accountName"`
	Location          *string           `pulumi:"location"`
	PoolName          *string           `pulumi:"poolName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ServiceLevel      string            `pulumi:"serviceLevel"`
	Size              float64           `pulumi:"size"`
	Tags              map[string]string `pulumi:"tags"`
}


type PoolArgs struct {
	AccountName       pulumi.StringInput
	Location          pulumi.StringPtrInput
	PoolName          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceLevel      pulumi.StringInput
	Size              pulumi.Float64Input
	Tags              pulumi.StringMapInput
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

func (o PoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PoolOutput) PoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.PoolId }).(pulumi.StringOutput)
}

func (o PoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PoolOutput) ServiceLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.ServiceLevel }).(pulumi.StringOutput)
}

func (o PoolOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v *Pool) pulumi.Float64Output { return v.Size }).(pulumi.Float64Output)
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
