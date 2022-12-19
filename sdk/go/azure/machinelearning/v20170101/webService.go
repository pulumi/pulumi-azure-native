


package v20170101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebService struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                        `pulumi:"location"`
	Name       pulumi.StringOutput                        `pulumi:"name"`
	Properties WebServicePropertiesForGraphResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                     `pulumi:"tags"`
	Type       pulumi.StringOutput                        `pulumi:"type"`
}


func NewWebService(ctx *pulumi.Context,
	name string, args *WebServiceArgs, opts ...pulumi.ResourceOption) (*WebService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Properties = args.Properties.ToWebServicePropertiesForGraphOutput().ApplyT(func(v WebServicePropertiesForGraph) WebServicePropertiesForGraph { return *v.Defaults() }).(WebServicePropertiesForGraphOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearning:WebService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearning/v20160501preview:WebService"),
		},
	})
	opts = append(opts, aliases)
	var resource WebService
	err := ctx.RegisterResource("azure-native:machinelearning/v20170101:WebService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebServiceState, opts ...pulumi.ResourceOption) (*WebService, error) {
	var resource WebService
	err := ctx.ReadResource("azure-native:machinelearning/v20170101:WebService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webServiceState struct {
}

type WebServiceState struct {
}

func (WebServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*webServiceState)(nil)).Elem()
}

type webServiceArgs struct {
	Location          *string                      `pulumi:"location"`
	Properties        WebServicePropertiesForGraph `pulumi:"properties"`
	ResourceGroupName string                       `pulumi:"resourceGroupName"`
	Tags              map[string]string            `pulumi:"tags"`
	WebServiceName    *string                      `pulumi:"webServiceName"`
}


type WebServiceArgs struct {
	Location          pulumi.StringPtrInput
	Properties        WebServicePropertiesForGraphInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	WebServiceName    pulumi.StringPtrInput
}

func (WebServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webServiceArgs)(nil)).Elem()
}

type WebServiceInput interface {
	pulumi.Input

	ToWebServiceOutput() WebServiceOutput
	ToWebServiceOutputWithContext(ctx context.Context) WebServiceOutput
}

func (*WebService) ElementType() reflect.Type {
	return reflect.TypeOf((**WebService)(nil)).Elem()
}

func (i *WebService) ToWebServiceOutput() WebServiceOutput {
	return i.ToWebServiceOutputWithContext(context.Background())
}

func (i *WebService) ToWebServiceOutputWithContext(ctx context.Context) WebServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServiceOutput)
}

type WebServiceOutput struct{ *pulumi.OutputState }

func (WebServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebService)(nil)).Elem()
}

func (o WebServiceOutput) ToWebServiceOutput() WebServiceOutput {
	return o
}

func (o WebServiceOutput) ToWebServiceOutputWithContext(ctx context.Context) WebServiceOutput {
	return o
}

func (o WebServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o WebServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebServiceOutput) Properties() WebServicePropertiesForGraphResponseOutput {
	return o.ApplyT(func(v *WebService) WebServicePropertiesForGraphResponseOutput { return v.Properties }).(WebServicePropertiesForGraphResponseOutput)
}

func (o WebServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WebServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebServiceOutput{})
}
