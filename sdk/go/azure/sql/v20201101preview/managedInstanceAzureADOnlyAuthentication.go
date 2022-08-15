


package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedInstanceAzureADOnlyAuthentication struct {
	pulumi.CustomResourceState

	AzureADOnlyAuthentication pulumi.BoolOutput   `pulumi:"azureADOnlyAuthentication"`
	Name                      pulumi.StringOutput `pulumi:"name"`
	Type                      pulumi.StringOutput `pulumi:"type"`
}


func NewManagedInstanceAzureADOnlyAuthentication(ctx *pulumi.Context,
	name string, args *ManagedInstanceAzureADOnlyAuthenticationArgs, opts ...pulumi.ResourceOption) (*ManagedInstanceAzureADOnlyAuthentication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AzureADOnlyAuthentication == nil {
		return nil, errors.New("invalid value for required argument 'AzureADOnlyAuthentication'")
	}
	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedInstanceAzureADOnlyAuthentication
	err := ctx.RegisterResource("azure-native:sql/v20201101preview:ManagedInstanceAzureADOnlyAuthentication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedInstanceAzureADOnlyAuthentication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedInstanceAzureADOnlyAuthenticationState, opts ...pulumi.ResourceOption) (*ManagedInstanceAzureADOnlyAuthentication, error) {
	var resource ManagedInstanceAzureADOnlyAuthentication
	err := ctx.ReadResource("azure-native:sql/v20201101preview:ManagedInstanceAzureADOnlyAuthentication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedInstanceAzureADOnlyAuthenticationState struct {
}

type ManagedInstanceAzureADOnlyAuthenticationState struct {
}

func (ManagedInstanceAzureADOnlyAuthenticationState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceAzureADOnlyAuthenticationState)(nil)).Elem()
}

type managedInstanceAzureADOnlyAuthenticationArgs struct {
	AuthenticationName        *string `pulumi:"authenticationName"`
	AzureADOnlyAuthentication bool    `pulumi:"azureADOnlyAuthentication"`
	ManagedInstanceName       string  `pulumi:"managedInstanceName"`
	ResourceGroupName         string  `pulumi:"resourceGroupName"`
}


type ManagedInstanceAzureADOnlyAuthenticationArgs struct {
	AuthenticationName        pulumi.StringPtrInput
	AzureADOnlyAuthentication pulumi.BoolInput
	ManagedInstanceName       pulumi.StringInput
	ResourceGroupName         pulumi.StringInput
}

func (ManagedInstanceAzureADOnlyAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceAzureADOnlyAuthenticationArgs)(nil)).Elem()
}

type ManagedInstanceAzureADOnlyAuthenticationInput interface {
	pulumi.Input

	ToManagedInstanceAzureADOnlyAuthenticationOutput() ManagedInstanceAzureADOnlyAuthenticationOutput
	ToManagedInstanceAzureADOnlyAuthenticationOutputWithContext(ctx context.Context) ManagedInstanceAzureADOnlyAuthenticationOutput
}

func (*ManagedInstanceAzureADOnlyAuthentication) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceAzureADOnlyAuthentication)(nil)).Elem()
}

func (i *ManagedInstanceAzureADOnlyAuthentication) ToManagedInstanceAzureADOnlyAuthenticationOutput() ManagedInstanceAzureADOnlyAuthenticationOutput {
	return i.ToManagedInstanceAzureADOnlyAuthenticationOutputWithContext(context.Background())
}

func (i *ManagedInstanceAzureADOnlyAuthentication) ToManagedInstanceAzureADOnlyAuthenticationOutputWithContext(ctx context.Context) ManagedInstanceAzureADOnlyAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceAzureADOnlyAuthenticationOutput)
}

type ManagedInstanceAzureADOnlyAuthenticationOutput struct{ *pulumi.OutputState }

func (ManagedInstanceAzureADOnlyAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceAzureADOnlyAuthentication)(nil)).Elem()
}

func (o ManagedInstanceAzureADOnlyAuthenticationOutput) ToManagedInstanceAzureADOnlyAuthenticationOutput() ManagedInstanceAzureADOnlyAuthenticationOutput {
	return o
}

func (o ManagedInstanceAzureADOnlyAuthenticationOutput) ToManagedInstanceAzureADOnlyAuthenticationOutputWithContext(ctx context.Context) ManagedInstanceAzureADOnlyAuthenticationOutput {
	return o
}

func (o ManagedInstanceAzureADOnlyAuthenticationOutput) AzureADOnlyAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v *ManagedInstanceAzureADOnlyAuthentication) pulumi.BoolOutput {
		return v.AzureADOnlyAuthentication
	}).(pulumi.BoolOutput)
}

func (o ManagedInstanceAzureADOnlyAuthenticationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceAzureADOnlyAuthentication) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedInstanceAzureADOnlyAuthenticationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceAzureADOnlyAuthentication) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedInstanceAzureADOnlyAuthenticationOutput{})
}
