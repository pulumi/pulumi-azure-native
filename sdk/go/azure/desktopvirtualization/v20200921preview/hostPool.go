


package v20200921preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HostPool struct {
	pulumi.CustomResourceState

	ApplicationGroupReferences    pulumi.StringArrayOutput          `pulumi:"applicationGroupReferences"`
	CustomRdpProperty             pulumi.StringPtrOutput            `pulumi:"customRdpProperty"`
	Description                   pulumi.StringPtrOutput            `pulumi:"description"`
	FriendlyName                  pulumi.StringPtrOutput            `pulumi:"friendlyName"`
	HostPoolType                  pulumi.StringOutput               `pulumi:"hostPoolType"`
	LoadBalancerType              pulumi.StringOutput               `pulumi:"loadBalancerType"`
	Location                      pulumi.StringOutput               `pulumi:"location"`
	MaxSessionLimit               pulumi.IntPtrOutput               `pulumi:"maxSessionLimit"`
	Name                          pulumi.StringOutput               `pulumi:"name"`
	PersonalDesktopAssignmentType pulumi.StringPtrOutput            `pulumi:"personalDesktopAssignmentType"`
	PreferredAppGroupType         pulumi.StringOutput               `pulumi:"preferredAppGroupType"`
	RegistrationInfo              RegistrationInfoResponsePtrOutput `pulumi:"registrationInfo"`
	Ring                          pulumi.IntPtrOutput               `pulumi:"ring"`
	SsoContext                    pulumi.StringPtrOutput            `pulumi:"ssoContext"`
	Tags                          pulumi.StringMapOutput            `pulumi:"tags"`
	Type                          pulumi.StringOutput               `pulumi:"type"`
	ValidationEnvironment         pulumi.BoolPtrOutput              `pulumi:"validationEnvironment"`
	VmTemplate                    pulumi.StringPtrOutput            `pulumi:"vmTemplate"`
}


func NewHostPool(ctx *pulumi.Context,
	name string, args *HostPoolArgs, opts ...pulumi.ResourceOption) (*HostPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostPoolType == nil {
		return nil, errors.New("invalid value for required argument 'HostPoolType'")
	}
	if args.LoadBalancerType == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerType'")
	}
	if args.PreferredAppGroupType == nil {
		return nil, errors.New("invalid value for required argument 'PreferredAppGroupType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:desktopvirtualization:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190123preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190924preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20191210preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201019preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201102preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201110preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210114preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210201preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210309preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210401preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210712:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:HostPool"),
		},
	})
	opts = append(opts, aliases)
	var resource HostPool
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20200921preview:HostPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHostPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostPoolState, opts ...pulumi.ResourceOption) (*HostPool, error) {
	var resource HostPool
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20200921preview:HostPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hostPoolState struct {
}

type HostPoolState struct {
}

func (HostPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostPoolState)(nil)).Elem()
}

type hostPoolArgs struct {
	CustomRdpProperty             *string           `pulumi:"customRdpProperty"`
	Description                   *string           `pulumi:"description"`
	FriendlyName                  *string           `pulumi:"friendlyName"`
	HostPoolName                  *string           `pulumi:"hostPoolName"`
	HostPoolType                  string            `pulumi:"hostPoolType"`
	LoadBalancerType              string            `pulumi:"loadBalancerType"`
	Location                      *string           `pulumi:"location"`
	MaxSessionLimit               *int              `pulumi:"maxSessionLimit"`
	PersonalDesktopAssignmentType *string           `pulumi:"personalDesktopAssignmentType"`
	PreferredAppGroupType         string            `pulumi:"preferredAppGroupType"`
	RegistrationInfo              *RegistrationInfo `pulumi:"registrationInfo"`
	ResourceGroupName             string            `pulumi:"resourceGroupName"`
	Ring                          *int              `pulumi:"ring"`
	SsoContext                    *string           `pulumi:"ssoContext"`
	Tags                          map[string]string `pulumi:"tags"`
	ValidationEnvironment         *bool             `pulumi:"validationEnvironment"`
	VmTemplate                    *string           `pulumi:"vmTemplate"`
}


type HostPoolArgs struct {
	CustomRdpProperty             pulumi.StringPtrInput
	Description                   pulumi.StringPtrInput
	FriendlyName                  pulumi.StringPtrInput
	HostPoolName                  pulumi.StringPtrInput
	HostPoolType                  pulumi.StringInput
	LoadBalancerType              pulumi.StringInput
	Location                      pulumi.StringPtrInput
	MaxSessionLimit               pulumi.IntPtrInput
	PersonalDesktopAssignmentType pulumi.StringPtrInput
	PreferredAppGroupType         pulumi.StringInput
	RegistrationInfo              RegistrationInfoPtrInput
	ResourceGroupName             pulumi.StringInput
	Ring                          pulumi.IntPtrInput
	SsoContext                    pulumi.StringPtrInput
	Tags                          pulumi.StringMapInput
	ValidationEnvironment         pulumi.BoolPtrInput
	VmTemplate                    pulumi.StringPtrInput
}

func (HostPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostPoolArgs)(nil)).Elem()
}

type HostPoolInput interface {
	pulumi.Input

	ToHostPoolOutput() HostPoolOutput
	ToHostPoolOutputWithContext(ctx context.Context) HostPoolOutput
}

func (*HostPool) ElementType() reflect.Type {
	return reflect.TypeOf((**HostPool)(nil)).Elem()
}

func (i *HostPool) ToHostPoolOutput() HostPoolOutput {
	return i.ToHostPoolOutputWithContext(context.Background())
}

func (i *HostPool) ToHostPoolOutputWithContext(ctx context.Context) HostPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostPoolOutput)
}

type HostPoolOutput struct{ *pulumi.OutputState }

func (HostPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostPool)(nil)).Elem()
}

func (o HostPoolOutput) ToHostPoolOutput() HostPoolOutput {
	return o
}

func (o HostPoolOutput) ToHostPoolOutputWithContext(ctx context.Context) HostPoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HostPoolOutput{})
}
