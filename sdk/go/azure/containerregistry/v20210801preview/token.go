


package v20210801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Token struct {
	pulumi.CustomResourceState

	CreationDate      pulumi.StringOutput                         `pulumi:"creationDate"`
	Credentials       TokenCredentialsPropertiesResponsePtrOutput `pulumi:"credentials"`
	Name              pulumi.StringOutput                         `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                         `pulumi:"provisioningState"`
	ScopeMapId        pulumi.StringPtrOutput                      `pulumi:"scopeMapId"`
	Status            pulumi.StringPtrOutput                      `pulumi:"status"`
	SystemData        SystemDataResponseOutput                    `pulumi:"systemData"`
	Type              pulumi.StringOutput                         `pulumi:"type"`
}


func NewToken(ctx *pulumi.Context,
	name string, args *TokenArgs, opts ...pulumi.ResourceOption) (*Token, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20210801preview:Token"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry:Token"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry:Token"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20190501preview:Token"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20190501preview:Token"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20201101preview:Token"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20201101preview:Token"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210601preview:Token"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20210601preview:Token"),
		},
	})
	opts = append(opts, aliases)
	var resource Token
	err := ctx.RegisterResource("azure-native:containerregistry/v20210801preview:Token", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TokenState, opts ...pulumi.ResourceOption) (*Token, error) {
	var resource Token
	err := ctx.ReadResource("azure-native:containerregistry/v20210801preview:Token", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type tokenState struct {
}

type TokenState struct {
}

func (TokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenState)(nil)).Elem()
}

type tokenArgs struct {
	Credentials       *TokenCredentialsProperties `pulumi:"credentials"`
	RegistryName      string                      `pulumi:"registryName"`
	ResourceGroupName string                      `pulumi:"resourceGroupName"`
	ScopeMapId        *string                     `pulumi:"scopeMapId"`
	Status            *string                     `pulumi:"status"`
	TokenName         *string                     `pulumi:"tokenName"`
}


type TokenArgs struct {
	Credentials       TokenCredentialsPropertiesPtrInput
	RegistryName      pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ScopeMapId        pulumi.StringPtrInput
	Status            pulumi.StringPtrInput
	TokenName         pulumi.StringPtrInput
}

func (TokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenArgs)(nil)).Elem()
}

type TokenInput interface {
	pulumi.Input

	ToTokenOutput() TokenOutput
	ToTokenOutputWithContext(ctx context.Context) TokenOutput
}

func (*Token) ElementType() reflect.Type {
	return reflect.TypeOf((*Token)(nil))
}

func (i *Token) ToTokenOutput() TokenOutput {
	return i.ToTokenOutputWithContext(context.Background())
}

func (i *Token) ToTokenOutputWithContext(ctx context.Context) TokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenOutput)
}

type TokenOutput struct{ *pulumi.OutputState }

func (TokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Token)(nil))
}

func (o TokenOutput) ToTokenOutput() TokenOutput {
	return o
}

func (o TokenOutput) ToTokenOutputWithContext(ctx context.Context) TokenOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TokenInput)(nil)).Elem(), &Token{})
	pulumi.RegisterOutputType(TokenOutput{})
}
