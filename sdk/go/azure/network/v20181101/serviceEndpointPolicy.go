


package v20181101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServiceEndpointPolicy struct {
	pulumi.CustomResourceState

	Etag                             pulumi.StringPtrOutput                             `pulumi:"etag"`
	Location                         pulumi.StringPtrOutput                             `pulumi:"location"`
	Name                             pulumi.StringOutput                                `pulumi:"name"`
	ProvisioningState                pulumi.StringOutput                                `pulumi:"provisioningState"`
	ResourceGuid                     pulumi.StringOutput                                `pulumi:"resourceGuid"`
	ServiceEndpointPolicyDefinitions ServiceEndpointPolicyDefinitionResponseArrayOutput `pulumi:"serviceEndpointPolicyDefinitions"`
	Subnets                          SubnetResponseArrayOutput                          `pulumi:"subnets"`
	Tags                             pulumi.StringMapOutput                             `pulumi:"tags"`
	Type                             pulumi.StringOutput                                `pulumi:"type"`
}


func NewServiceEndpointPolicy(ctx *pulumi.Context,
	name string, args *ServiceEndpointPolicyArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ServiceEndpointPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceEndpointPolicy
	err := ctx.RegisterResource("azure-native:network/v20181101:ServiceEndpointPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServiceEndpointPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointPolicyState, opts ...pulumi.ResourceOption) (*ServiceEndpointPolicy, error) {
	var resource ServiceEndpointPolicy
	err := ctx.ReadResource("azure-native:network/v20181101:ServiceEndpointPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serviceEndpointPolicyState struct {
}

type ServiceEndpointPolicyState struct {
}

func (ServiceEndpointPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointPolicyState)(nil)).Elem()
}

type serviceEndpointPolicyArgs struct {
	Id                               *string                               `pulumi:"id"`
	Location                         *string                               `pulumi:"location"`
	ResourceGroupName                string                                `pulumi:"resourceGroupName"`
	ServiceEndpointPolicyDefinitions []ServiceEndpointPolicyDefinitionType `pulumi:"serviceEndpointPolicyDefinitions"`
	ServiceEndpointPolicyName        *string                               `pulumi:"serviceEndpointPolicyName"`
	Tags                             map[string]string                     `pulumi:"tags"`
}


type ServiceEndpointPolicyArgs struct {
	Id                               pulumi.StringPtrInput
	Location                         pulumi.StringPtrInput
	ResourceGroupName                pulumi.StringInput
	ServiceEndpointPolicyDefinitions ServiceEndpointPolicyDefinitionTypeArrayInput
	ServiceEndpointPolicyName        pulumi.StringPtrInput
	Tags                             pulumi.StringMapInput
}

func (ServiceEndpointPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointPolicyArgs)(nil)).Elem()
}

type ServiceEndpointPolicyInput interface {
	pulumi.Input

	ToServiceEndpointPolicyOutput() ServiceEndpointPolicyOutput
	ToServiceEndpointPolicyOutputWithContext(ctx context.Context) ServiceEndpointPolicyOutput
}

func (*ServiceEndpointPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointPolicy)(nil)).Elem()
}

func (i *ServiceEndpointPolicy) ToServiceEndpointPolicyOutput() ServiceEndpointPolicyOutput {
	return i.ToServiceEndpointPolicyOutputWithContext(context.Background())
}

func (i *ServiceEndpointPolicy) ToServiceEndpointPolicyOutputWithContext(ctx context.Context) ServiceEndpointPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPolicyOutput)
}

type ServiceEndpointPolicyOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointPolicy)(nil)).Elem()
}

func (o ServiceEndpointPolicyOutput) ToServiceEndpointPolicyOutput() ServiceEndpointPolicyOutput {
	return o
}

func (o ServiceEndpointPolicyOutput) ToServiceEndpointPolicyOutputWithContext(ctx context.Context) ServiceEndpointPolicyOutput {
	return o
}

func (o ServiceEndpointPolicyOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceEndpointPolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServiceEndpointPolicyOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o ServiceEndpointPolicyOutput) ServiceEndpointPolicyDefinitions() ServiceEndpointPolicyDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) ServiceEndpointPolicyDefinitionResponseArrayOutput {
		return v.ServiceEndpointPolicyDefinitions
	}).(ServiceEndpointPolicyDefinitionResponseArrayOutput)
}

func (o ServiceEndpointPolicyOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) SubnetResponseArrayOutput { return v.Subnets }).(SubnetResponseArrayOutput)
}

func (o ServiceEndpointPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ServiceEndpointPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceEndpointPolicyOutput{})
}
