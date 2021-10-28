


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomDomain struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                  `pulumi:"name"`
	Properties CustomDomainPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                  `pulumi:"type"`
}


func NewCustomDomain(ctx *pulumi.Context,
	name string, args *CustomDomainArgs, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppName == nil {
		return nil, errors.New("invalid value for required argument 'AppName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:appplatform/v20210601preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-nextgen:appplatform:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20190501preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-nextgen:appplatform/v20190501preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20200701:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-nextgen:appplatform/v20200701:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20201101preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-nextgen:appplatform/v20201101preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-nextgen:appplatform/v20210901preview:CustomDomain"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomDomain
	err := ctx.RegisterResource("azure-native:appplatform/v20210601preview:CustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomDomainState, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	var resource CustomDomain
	err := ctx.ReadResource("azure-native:appplatform/v20210601preview:CustomDomain", name, id, state, &resource, opts...)
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
	AppName           string                  `pulumi:"appName"`
	DomainName        *string                 `pulumi:"domainName"`
	Properties        *CustomDomainProperties `pulumi:"properties"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	ServiceName       string                  `pulumi:"serviceName"`
}


type CustomDomainArgs struct {
	AppName           pulumi.StringInput
	DomainName        pulumi.StringPtrInput
	Properties        CustomDomainPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
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
	return reflect.TypeOf((*CustomDomain)(nil))
}

func (i *CustomDomain) ToCustomDomainOutput() CustomDomainOutput {
	return i.ToCustomDomainOutputWithContext(context.Background())
}

func (i *CustomDomain) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput)
}

type CustomDomainOutput struct{ *pulumi.OutputState }

func (CustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomain)(nil))
}

func (o CustomDomainOutput) ToCustomDomainOutput() CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDomainInput)(nil)).Elem(), &CustomDomain{})
	pulumi.RegisterOutputType(CustomDomainOutput{})
}
