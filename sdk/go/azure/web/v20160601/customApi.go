


package v20160601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomApi struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput                      `pulumi:"etag"`
	Location   pulumi.StringPtrOutput                      `pulumi:"location"`
	Name       pulumi.StringOutput                         `pulumi:"name"`
	Properties CustomApiPropertiesDefinitionResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                      `pulumi:"tags"`
	Type       pulumi.StringOutput                         `pulumi:"type"`
}


func NewCustomApi(ctx *pulumi.Context,
	name string, args *CustomApiArgs, opts ...pulumi.ResourceOption) (*CustomApi, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:CustomApi"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomApi
	err := ctx.RegisterResource("azure-native:web/v20160601:CustomApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCustomApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomApiState, opts ...pulumi.ResourceOption) (*CustomApi, error) {
	var resource CustomApi
	err := ctx.ReadResource("azure-native:web/v20160601:CustomApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type customApiState struct {
}

type CustomApiState struct {
}

func (CustomApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*customApiState)(nil)).Elem()
}

type customApiArgs struct {
	ApiName           *string                        `pulumi:"apiName"`
	Etag              *string                        `pulumi:"etag"`
	Location          *string                        `pulumi:"location"`
	Properties        *CustomApiPropertiesDefinition `pulumi:"properties"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	SubscriptionId    *string                        `pulumi:"subscriptionId"`
	Tags              map[string]string              `pulumi:"tags"`
}


type CustomApiArgs struct {
	ApiName           pulumi.StringPtrInput
	Etag              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        CustomApiPropertiesDefinitionPtrInput
	ResourceGroupName pulumi.StringInput
	SubscriptionId    pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (CustomApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customApiArgs)(nil)).Elem()
}

type CustomApiInput interface {
	pulumi.Input

	ToCustomApiOutput() CustomApiOutput
	ToCustomApiOutputWithContext(ctx context.Context) CustomApiOutput
}

func (*CustomApi) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomApi)(nil)).Elem()
}

func (i *CustomApi) ToCustomApiOutput() CustomApiOutput {
	return i.ToCustomApiOutputWithContext(context.Background())
}

func (i *CustomApi) ToCustomApiOutputWithContext(ctx context.Context) CustomApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomApiOutput)
}

type CustomApiOutput struct{ *pulumi.OutputState }

func (CustomApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomApi)(nil)).Elem()
}

func (o CustomApiOutput) ToCustomApiOutput() CustomApiOutput {
	return o
}

func (o CustomApiOutput) ToCustomApiOutputWithContext(ctx context.Context) CustomApiOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CustomApiOutput{})
}
