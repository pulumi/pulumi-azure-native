


package v20191202preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RemoteRenderingAccount struct {
	pulumi.CustomResourceState

	AccountDomain pulumi.StringOutput       `pulumi:"accountDomain"`
	AccountId     pulumi.StringOutput       `pulumi:"accountId"`
	Identity      IdentityResponsePtrOutput `pulumi:"identity"`
	Location      pulumi.StringOutput       `pulumi:"location"`
	Name          pulumi.StringOutput       `pulumi:"name"`
	Tags          pulumi.StringMapOutput    `pulumi:"tags"`
	Type          pulumi.StringOutput       `pulumi:"type"`
}


func NewRemoteRenderingAccount(ctx *pulumi.Context,
	name string, args *RemoteRenderingAccountArgs, opts ...pulumi.ResourceOption) (*RemoteRenderingAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:mixedreality/v20191202preview:RemoteRenderingAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality:RemoteRenderingAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:mixedreality:RemoteRenderingAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20200406preview:RemoteRenderingAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:mixedreality/v20200406preview:RemoteRenderingAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20210101:RemoteRenderingAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:mixedreality/v20210101:RemoteRenderingAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20210301preview:RemoteRenderingAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:mixedreality/v20210301preview:RemoteRenderingAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource RemoteRenderingAccount
	err := ctx.RegisterResource("azure-native:mixedreality/v20191202preview:RemoteRenderingAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRemoteRenderingAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemoteRenderingAccountState, opts ...pulumi.ResourceOption) (*RemoteRenderingAccount, error) {
	var resource RemoteRenderingAccount
	err := ctx.ReadResource("azure-native:mixedreality/v20191202preview:RemoteRenderingAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type remoteRenderingAccountState struct {
}

type RemoteRenderingAccountState struct {
}

func (RemoteRenderingAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteRenderingAccountState)(nil)).Elem()
}

type remoteRenderingAccountArgs struct {
	AccountName       *string           `pulumi:"accountName"`
	Identity          *Identity         `pulumi:"identity"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type RemoteRenderingAccountArgs struct {
	AccountName       pulumi.StringPtrInput
	Identity          IdentityPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (RemoteRenderingAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteRenderingAccountArgs)(nil)).Elem()
}

type RemoteRenderingAccountInput interface {
	pulumi.Input

	ToRemoteRenderingAccountOutput() RemoteRenderingAccountOutput
	ToRemoteRenderingAccountOutputWithContext(ctx context.Context) RemoteRenderingAccountOutput
}

func (*RemoteRenderingAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteRenderingAccount)(nil))
}

func (i *RemoteRenderingAccount) ToRemoteRenderingAccountOutput() RemoteRenderingAccountOutput {
	return i.ToRemoteRenderingAccountOutputWithContext(context.Background())
}

func (i *RemoteRenderingAccount) ToRemoteRenderingAccountOutputWithContext(ctx context.Context) RemoteRenderingAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteRenderingAccountOutput)
}

type RemoteRenderingAccountOutput struct{ *pulumi.OutputState }

func (RemoteRenderingAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteRenderingAccount)(nil))
}

func (o RemoteRenderingAccountOutput) ToRemoteRenderingAccountOutput() RemoteRenderingAccountOutput {
	return o
}

func (o RemoteRenderingAccountOutput) ToRemoteRenderingAccountOutputWithContext(ctx context.Context) RemoteRenderingAccountOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RemoteRenderingAccountOutput{})
}
