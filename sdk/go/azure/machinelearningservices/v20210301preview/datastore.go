


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Datastore struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput               `pulumi:"name"`
	Properties DatastorePropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput          `pulumi:"systemData"`
	Type       pulumi.StringOutput               `pulumi:"type"`
}


func NewDatastore(ctx *pulumi.Context,
	name string, args *DatastoreArgs, opts ...pulumi.ResourceOption) (*Datastore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200501preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:Datastore"),
		},
	})
	opts = append(opts, aliases)
	var resource Datastore
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20210301preview:Datastore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatastore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatastoreState, opts ...pulumi.ResourceOption) (*Datastore, error) {
	var resource Datastore
	err := ctx.ReadResource("azure-native:machinelearningservices/v20210301preview:Datastore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type datastoreState struct {
}

type DatastoreState struct {
}

func (DatastoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreState)(nil)).Elem()
}

type datastoreArgs struct {
	Name              *string             `pulumi:"name"`
	Properties        DatastoreProperties `pulumi:"properties"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	SkipValidation    *bool               `pulumi:"skipValidation"`
	WorkspaceName     string              `pulumi:"workspaceName"`
}


type DatastoreArgs struct {
	Name              pulumi.StringPtrInput
	Properties        DatastorePropertiesInput
	ResourceGroupName pulumi.StringInput
	SkipValidation    pulumi.BoolPtrInput
	WorkspaceName     pulumi.StringInput
}

func (DatastoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreArgs)(nil)).Elem()
}

type DatastoreInput interface {
	pulumi.Input

	ToDatastoreOutput() DatastoreOutput
	ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput
}

func (*Datastore) ElementType() reflect.Type {
	return reflect.TypeOf((**Datastore)(nil)).Elem()
}

func (i *Datastore) ToDatastoreOutput() DatastoreOutput {
	return i.ToDatastoreOutputWithContext(context.Background())
}

func (i *Datastore) ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatastoreOutput)
}

type DatastoreOutput struct{ *pulumi.OutputState }

func (DatastoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Datastore)(nil)).Elem()
}

func (o DatastoreOutput) ToDatastoreOutput() DatastoreOutput {
	return o
}

func (o DatastoreOutput) ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput {
	return o
}

func (o DatastoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatastoreOutput) Properties() DatastorePropertiesResponseOutput {
	return o.ApplyT(func(v *Datastore) DatastorePropertiesResponseOutput { return v.Properties }).(DatastorePropertiesResponseOutput)
}

func (o DatastoreOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Datastore) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DatastoreOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatastoreOutput{})
}
