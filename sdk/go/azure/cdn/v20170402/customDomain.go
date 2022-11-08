


package v20170402

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type CustomDomain struct {
	pulumi.CustomResourceState

	CustomHttpsProvisioningState    pulumi.StringOutput    `pulumi:"customHttpsProvisioningState"`
	CustomHttpsProvisioningSubstate pulumi.StringOutput    `pulumi:"customHttpsProvisioningSubstate"`
	HostName                        pulumi.StringOutput    `pulumi:"hostName"`
	Name                            pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState               pulumi.StringOutput    `pulumi:"provisioningState"`
	ResourceState                   pulumi.StringOutput    `pulumi:"resourceState"`
	Type                            pulumi.StringOutput    `pulumi:"type"`
	ValidationData                  pulumi.StringPtrOutput `pulumi:"validationData"`
}


func NewCustomDomain(ctx *pulumi.Context,
	name string, args *CustomDomainArgs, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20150601:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20160402:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20161002:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20171012:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190415:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20191231:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200331:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200415:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20221101preview:CustomDomain"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomDomain
	err := ctx.RegisterResource("azure-native:cdn/v20170402:CustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomDomainState, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	var resource CustomDomain
	err := ctx.ReadResource("azure-native:cdn/v20170402:CustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type customDomainState struct {
}

type CustomDomainState struct {
}

func (CustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*customDomainState)(nil)).Elem()
}

type customDomainArgs struct {
	CustomDomainName  *string `pulumi:"customDomainName"`
	EndpointName      string  `pulumi:"endpointName"`
	HostName          string  `pulumi:"hostName"`
	ProfileName       string  `pulumi:"profileName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type CustomDomainArgs struct {
	CustomDomainName  pulumi.StringPtrInput
	EndpointName      pulumi.StringInput
	HostName          pulumi.StringInput
	ProfileName       pulumi.StringInput
	ResourceGroupName pulumi.StringInput
}

func (CustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customDomainArgs)(nil)).Elem()
}

type CustomDomainInput interface {
	pulumi.Input

	ToCustomDomainOutput() CustomDomainOutput
	ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput
}

func (*CustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (i *CustomDomain) ToCustomDomainOutput() CustomDomainOutput {
	return i.ToCustomDomainOutputWithContext(context.Background())
}

func (i *CustomDomain) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput)
}

type CustomDomainOutput struct{ *pulumi.OutputState }

func (CustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (o CustomDomainOutput) ToCustomDomainOutput() CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) CustomHttpsProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.CustomHttpsProvisioningState }).(pulumi.StringOutput)
}

func (o CustomDomainOutput) CustomHttpsProvisioningSubstate() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.CustomHttpsProvisioningSubstate }).(pulumi.StringOutput)
}

func (o CustomDomainOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

func (o CustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CustomDomainOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CustomDomainOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

func (o CustomDomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o CustomDomainOutput) ValidationData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringPtrOutput { return v.ValidationData }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomDomainOutput{})
}
