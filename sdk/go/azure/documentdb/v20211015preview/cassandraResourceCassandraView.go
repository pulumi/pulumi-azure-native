


package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CassandraResourceCassandraView struct {
	pulumi.CustomResourceState

	Identity ManagedServiceIdentityResponsePtrOutput             `pulumi:"identity"`
	Location pulumi.StringPtrOutput                              `pulumi:"location"`
	Name     pulumi.StringOutput                                 `pulumi:"name"`
	Options  CassandraViewGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource CassandraViewGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	Tags     pulumi.StringMapOutput                              `pulumi:"tags"`
	Type     pulumi.StringOutput                                 `pulumi:"type"`
}


func NewCassandraResourceCassandraView(ctx *pulumi.Context,
	name string, args *CassandraResourceCassandraViewArgs, opts ...pulumi.ResourceOption) (*CassandraResourceCassandraView, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.KeyspaceName == nil {
		return nil, errors.New("invalid value for required argument 'KeyspaceName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:CassandraResourceCassandraView"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:CassandraResourceCassandraView"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:CassandraResourceCassandraView"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:CassandraResourceCassandraView"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:CassandraResourceCassandraView"),
		},
	})
	opts = append(opts, aliases)
	var resource CassandraResourceCassandraView
	err := ctx.RegisterResource("azure-native:documentdb/v20211015preview:CassandraResourceCassandraView", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCassandraResourceCassandraView(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CassandraResourceCassandraViewState, opts ...pulumi.ResourceOption) (*CassandraResourceCassandraView, error) {
	var resource CassandraResourceCassandraView
	err := ctx.ReadResource("azure-native:documentdb/v20211015preview:CassandraResourceCassandraView", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type cassandraResourceCassandraViewState struct {
}

type CassandraResourceCassandraViewState struct {
}

func (CassandraResourceCassandraViewState) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraResourceCassandraViewState)(nil)).Elem()
}

type cassandraResourceCassandraViewArgs struct {
	AccountName       string                  `pulumi:"accountName"`
	Identity          *ManagedServiceIdentity `pulumi:"identity"`
	KeyspaceName      string                  `pulumi:"keyspaceName"`
	Location          *string                 `pulumi:"location"`
	Options           *CreateUpdateOptions    `pulumi:"options"`
	Resource          CassandraViewResource   `pulumi:"resource"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Tags              map[string]string       `pulumi:"tags"`
	ViewName          *string                 `pulumi:"viewName"`
}


type CassandraResourceCassandraViewArgs struct {
	AccountName       pulumi.StringInput
	Identity          ManagedServiceIdentityPtrInput
	KeyspaceName      pulumi.StringInput
	Location          pulumi.StringPtrInput
	Options           CreateUpdateOptionsPtrInput
	Resource          CassandraViewResourceInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	ViewName          pulumi.StringPtrInput
}

func (CassandraResourceCassandraViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraResourceCassandraViewArgs)(nil)).Elem()
}

type CassandraResourceCassandraViewInput interface {
	pulumi.Input

	ToCassandraResourceCassandraViewOutput() CassandraResourceCassandraViewOutput
	ToCassandraResourceCassandraViewOutputWithContext(ctx context.Context) CassandraResourceCassandraViewOutput
}

func (*CassandraResourceCassandraView) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraResourceCassandraView)(nil)).Elem()
}

func (i *CassandraResourceCassandraView) ToCassandraResourceCassandraViewOutput() CassandraResourceCassandraViewOutput {
	return i.ToCassandraResourceCassandraViewOutputWithContext(context.Background())
}

func (i *CassandraResourceCassandraView) ToCassandraResourceCassandraViewOutputWithContext(ctx context.Context) CassandraResourceCassandraViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraResourceCassandraViewOutput)
}

type CassandraResourceCassandraViewOutput struct{ *pulumi.OutputState }

func (CassandraResourceCassandraViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraResourceCassandraView)(nil)).Elem()
}

func (o CassandraResourceCassandraViewOutput) ToCassandraResourceCassandraViewOutput() CassandraResourceCassandraViewOutput {
	return o
}

func (o CassandraResourceCassandraViewOutput) ToCassandraResourceCassandraViewOutputWithContext(ctx context.Context) CassandraResourceCassandraViewOutput {
	return o
}

func (o CassandraResourceCassandraViewOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraView) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o CassandraResourceCassandraViewOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraView) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o CassandraResourceCassandraViewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraView) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CassandraResourceCassandraViewOutput) Options() CassandraViewGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraView) CassandraViewGetPropertiesResponseOptionsPtrOutput {
		return v.Options
	}).(CassandraViewGetPropertiesResponseOptionsPtrOutput)
}

func (o CassandraResourceCassandraViewOutput) Resource() CassandraViewGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraView) CassandraViewGetPropertiesResponseResourcePtrOutput {
		return v.Resource
	}).(CassandraViewGetPropertiesResponseResourcePtrOutput)
}

func (o CassandraResourceCassandraViewOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraView) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CassandraResourceCassandraViewOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraView) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CassandraResourceCassandraViewOutput{})
}
