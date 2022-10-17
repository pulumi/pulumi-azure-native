


package v20200601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SqlResourceSqlTrigger struct {
	pulumi.CustomResourceState

	Identity ManagedServiceIdentityResponsePtrOutput          `pulumi:"identity"`
	Location pulumi.StringPtrOutput                           `pulumi:"location"`
	Name     pulumi.StringOutput                              `pulumi:"name"`
	Resource SqlTriggerGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	Tags     pulumi.StringMapOutput                           `pulumi:"tags"`
	Type     pulumi.StringOutput                              `pulumi:"type"`
}


func NewSqlResourceSqlTrigger(ctx *pulumi.Context,
	name string, args *SqlResourceSqlTriggerArgs, opts ...pulumi.ResourceOption) (*SqlResourceSqlTrigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Options == nil {
		return nil, errors.New("invalid value for required argument 'Options'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:SqlResourceSqlTrigger"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlResourceSqlTrigger
	err := ctx.RegisterResource("azure-native:documentdb/v20200601preview:SqlResourceSqlTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlResourceSqlTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlResourceSqlTriggerState, opts ...pulumi.ResourceOption) (*SqlResourceSqlTrigger, error) {
	var resource SqlResourceSqlTrigger
	err := ctx.ReadResource("azure-native:documentdb/v20200601preview:SqlResourceSqlTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlResourceSqlTriggerState struct {
}

type SqlResourceSqlTriggerState struct {
}

func (SqlResourceSqlTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlTriggerState)(nil)).Elem()
}

type sqlResourceSqlTriggerArgs struct {
	AccountName       string                  `pulumi:"accountName"`
	ContainerName     string                  `pulumi:"containerName"`
	DatabaseName      string                  `pulumi:"databaseName"`
	Identity          *ManagedServiceIdentity `pulumi:"identity"`
	Location          *string                 `pulumi:"location"`
	Options           CreateUpdateOptions     `pulumi:"options"`
	Resource          SqlTriggerResource      `pulumi:"resource"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Tags              map[string]string       `pulumi:"tags"`
	TriggerName       *string                 `pulumi:"triggerName"`
}


type SqlResourceSqlTriggerArgs struct {
	AccountName       pulumi.StringInput
	ContainerName     pulumi.StringInput
	DatabaseName      pulumi.StringInput
	Identity          ManagedServiceIdentityPtrInput
	Location          pulumi.StringPtrInput
	Options           CreateUpdateOptionsInput
	Resource          SqlTriggerResourceInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	TriggerName       pulumi.StringPtrInput
}

func (SqlResourceSqlTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlTriggerArgs)(nil)).Elem()
}

type SqlResourceSqlTriggerInput interface {
	pulumi.Input

	ToSqlResourceSqlTriggerOutput() SqlResourceSqlTriggerOutput
	ToSqlResourceSqlTriggerOutputWithContext(ctx context.Context) SqlResourceSqlTriggerOutput
}

func (*SqlResourceSqlTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlTrigger)(nil)).Elem()
}

func (i *SqlResourceSqlTrigger) ToSqlResourceSqlTriggerOutput() SqlResourceSqlTriggerOutput {
	return i.ToSqlResourceSqlTriggerOutputWithContext(context.Background())
}

func (i *SqlResourceSqlTrigger) ToSqlResourceSqlTriggerOutputWithContext(ctx context.Context) SqlResourceSqlTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlResourceSqlTriggerOutput)
}

type SqlResourceSqlTriggerOutput struct{ *pulumi.OutputState }

func (SqlResourceSqlTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlTrigger)(nil)).Elem()
}

func (o SqlResourceSqlTriggerOutput) ToSqlResourceSqlTriggerOutput() SqlResourceSqlTriggerOutput {
	return o
}

func (o SqlResourceSqlTriggerOutput) ToSqlResourceSqlTriggerOutputWithContext(ctx context.Context) SqlResourceSqlTriggerOutput {
	return o
}

func (o SqlResourceSqlTriggerOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlTrigger) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o SqlResourceSqlTriggerOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlTrigger) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o SqlResourceSqlTriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlTrigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlResourceSqlTriggerOutput) Resource() SqlTriggerGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlTrigger) SqlTriggerGetPropertiesResponseResourcePtrOutput { return v.Resource }).(SqlTriggerGetPropertiesResponseResourcePtrOutput)
}

func (o SqlResourceSqlTriggerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlResourceSqlTrigger) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlResourceSqlTriggerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlTrigger) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlResourceSqlTriggerOutput{})
}
