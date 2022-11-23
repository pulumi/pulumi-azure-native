


package v20220701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecurityPartnerProvider struct {
	pulumi.CustomResourceState

	ConnectionStatus     pulumi.StringOutput          `pulumi:"connectionStatus"`
	Etag                 pulumi.StringOutput          `pulumi:"etag"`
	Location             pulumi.StringPtrOutput       `pulumi:"location"`
	Name                 pulumi.StringOutput          `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput          `pulumi:"provisioningState"`
	SecurityProviderName pulumi.StringPtrOutput       `pulumi:"securityProviderName"`
	Tags                 pulumi.StringMapOutput       `pulumi:"tags"`
	Type                 pulumi.StringOutput          `pulumi:"type"`
	VirtualHub           SubResourceResponsePtrOutput `pulumi:"virtualHub"`
}


func NewSecurityPartnerProvider(ctx *pulumi.Context,
	name string, args *SecurityPartnerProviderArgs, opts ...pulumi.ResourceOption) (*SecurityPartnerProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:SecurityPartnerProvider"),
		},
	})
	opts = append(opts, aliases)
	var resource SecurityPartnerProvider
	err := ctx.RegisterResource("azure-native:network/v20220701:SecurityPartnerProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSecurityPartnerProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityPartnerProviderState, opts ...pulumi.ResourceOption) (*SecurityPartnerProvider, error) {
	var resource SecurityPartnerProvider
	err := ctx.ReadResource("azure-native:network/v20220701:SecurityPartnerProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type securityPartnerProviderState struct {
}

type SecurityPartnerProviderState struct {
}

func (SecurityPartnerProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPartnerProviderState)(nil)).Elem()
}

type securityPartnerProviderArgs struct {
	Id                          *string           `pulumi:"id"`
	Location                    *string           `pulumi:"location"`
	ResourceGroupName           string            `pulumi:"resourceGroupName"`
	SecurityPartnerProviderName *string           `pulumi:"securityPartnerProviderName"`
	SecurityProviderName        *string           `pulumi:"securityProviderName"`
	Tags                        map[string]string `pulumi:"tags"`
	VirtualHub                  *SubResource      `pulumi:"virtualHub"`
}


type SecurityPartnerProviderArgs struct {
	Id                          pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	SecurityPartnerProviderName pulumi.StringPtrInput
	SecurityProviderName        pulumi.StringPtrInput
	Tags                        pulumi.StringMapInput
	VirtualHub                  SubResourcePtrInput
}

func (SecurityPartnerProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPartnerProviderArgs)(nil)).Elem()
}

type SecurityPartnerProviderInput interface {
	pulumi.Input

	ToSecurityPartnerProviderOutput() SecurityPartnerProviderOutput
	ToSecurityPartnerProviderOutputWithContext(ctx context.Context) SecurityPartnerProviderOutput
}

func (*SecurityPartnerProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPartnerProvider)(nil)).Elem()
}

func (i *SecurityPartnerProvider) ToSecurityPartnerProviderOutput() SecurityPartnerProviderOutput {
	return i.ToSecurityPartnerProviderOutputWithContext(context.Background())
}

func (i *SecurityPartnerProvider) ToSecurityPartnerProviderOutputWithContext(ctx context.Context) SecurityPartnerProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPartnerProviderOutput)
}

type SecurityPartnerProviderOutput struct{ *pulumi.OutputState }

func (SecurityPartnerProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPartnerProvider)(nil)).Elem()
}

func (o SecurityPartnerProviderOutput) ToSecurityPartnerProviderOutput() SecurityPartnerProviderOutput {
	return o
}

func (o SecurityPartnerProviderOutput) ToSecurityPartnerProviderOutputWithContext(ctx context.Context) SecurityPartnerProviderOutput {
	return o
}

func (o SecurityPartnerProviderOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) pulumi.StringOutput { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o SecurityPartnerProviderOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o SecurityPartnerProviderOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o SecurityPartnerProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SecurityPartnerProviderOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SecurityPartnerProviderOutput) SecurityProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) pulumi.StringPtrOutput { return v.SecurityProviderName }).(pulumi.StringPtrOutput)
}

func (o SecurityPartnerProviderOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SecurityPartnerProviderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SecurityPartnerProviderOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) SubResourceResponsePtrOutput { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityPartnerProviderOutput{})
}
