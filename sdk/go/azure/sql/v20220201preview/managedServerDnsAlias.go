


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedServerDnsAlias struct {
	pulumi.CustomResourceState

	AzureDnsRecord       pulumi.StringOutput `pulumi:"azureDnsRecord"`
	Name                 pulumi.StringOutput `pulumi:"name"`
	PublicAzureDnsRecord pulumi.StringOutput `pulumi:"publicAzureDnsRecord"`
	Type                 pulumi.StringOutput `pulumi:"type"`
}


func NewManagedServerDnsAlias(ctx *pulumi.Context,
	name string, args *ManagedServerDnsAliasArgs, opts ...pulumi.ResourceOption) (*ManagedServerDnsAlias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.CreateDnsRecord) {
		args.CreateDnsRecord = pulumi.BoolPtr(true)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ManagedServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ManagedServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ManagedServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ManagedServerDnsAlias"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedServerDnsAlias
	err := ctx.RegisterResource("azure-native:sql/v20220201preview:ManagedServerDnsAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedServerDnsAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedServerDnsAliasState, opts ...pulumi.ResourceOption) (*ManagedServerDnsAlias, error) {
	var resource ManagedServerDnsAlias
	err := ctx.ReadResource("azure-native:sql/v20220201preview:ManagedServerDnsAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedServerDnsAliasState struct {
}

type ManagedServerDnsAliasState struct {
}

func (ManagedServerDnsAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedServerDnsAliasState)(nil)).Elem()
}

type managedServerDnsAliasArgs struct {
	CreateDnsRecord     *bool   `pulumi:"createDnsRecord"`
	DnsAliasName        *string `pulumi:"dnsAliasName"`
	ManagedInstanceName string  `pulumi:"managedInstanceName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
}


type ManagedServerDnsAliasArgs struct {
	CreateDnsRecord     pulumi.BoolPtrInput
	DnsAliasName        pulumi.StringPtrInput
	ManagedInstanceName pulumi.StringInput
	ResourceGroupName   pulumi.StringInput
}

func (ManagedServerDnsAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedServerDnsAliasArgs)(nil)).Elem()
}

type ManagedServerDnsAliasInput interface {
	pulumi.Input

	ToManagedServerDnsAliasOutput() ManagedServerDnsAliasOutput
	ToManagedServerDnsAliasOutputWithContext(ctx context.Context) ManagedServerDnsAliasOutput
}

func (*ManagedServerDnsAlias) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServerDnsAlias)(nil)).Elem()
}

func (i *ManagedServerDnsAlias) ToManagedServerDnsAliasOutput() ManagedServerDnsAliasOutput {
	return i.ToManagedServerDnsAliasOutputWithContext(context.Background())
}

func (i *ManagedServerDnsAlias) ToManagedServerDnsAliasOutputWithContext(ctx context.Context) ManagedServerDnsAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServerDnsAliasOutput)
}

type ManagedServerDnsAliasOutput struct{ *pulumi.OutputState }

func (ManagedServerDnsAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServerDnsAlias)(nil)).Elem()
}

func (o ManagedServerDnsAliasOutput) ToManagedServerDnsAliasOutput() ManagedServerDnsAliasOutput {
	return o
}

func (o ManagedServerDnsAliasOutput) ToManagedServerDnsAliasOutputWithContext(ctx context.Context) ManagedServerDnsAliasOutput {
	return o
}

func (o ManagedServerDnsAliasOutput) AzureDnsRecord() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedServerDnsAlias) pulumi.StringOutput { return v.AzureDnsRecord }).(pulumi.StringOutput)
}

func (o ManagedServerDnsAliasOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedServerDnsAlias) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedServerDnsAliasOutput) PublicAzureDnsRecord() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedServerDnsAlias) pulumi.StringOutput { return v.PublicAzureDnsRecord }).(pulumi.StringOutput)
}

func (o ManagedServerDnsAliasOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedServerDnsAlias) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedServerDnsAliasOutput{})
}
