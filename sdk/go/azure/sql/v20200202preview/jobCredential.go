


package v20200202preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type JobCredential struct {
	pulumi.CustomResourceState

	Name     pulumi.StringOutput `pulumi:"name"`
	Type     pulumi.StringOutput `pulumi:"type"`
	Username pulumi.StringOutput `pulumi:"username"`
}


func NewJobCredential(ctx *pulumi.Context,
	name string, args *JobCredentialArgs, opts ...pulumi.ResourceOption) (*JobCredential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobAgentName == nil {
		return nil, errors.New("invalid value for required argument 'JobAgentName'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:JobCredential"),
		},
	})
	opts = append(opts, aliases)
	var resource JobCredential
	err := ctx.RegisterResource("azure-native:sql/v20200202preview:JobCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetJobCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobCredentialState, opts ...pulumi.ResourceOption) (*JobCredential, error) {
	var resource JobCredential
	err := ctx.ReadResource("azure-native:sql/v20200202preview:JobCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type jobCredentialState struct {
}

type JobCredentialState struct {
}

func (JobCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobCredentialState)(nil)).Elem()
}

type jobCredentialArgs struct {
	CredentialName    *string `pulumi:"credentialName"`
	JobAgentName      string  `pulumi:"jobAgentName"`
	Password          string  `pulumi:"password"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServerName        string  `pulumi:"serverName"`
	Username          string  `pulumi:"username"`
}


type JobCredentialArgs struct {
	CredentialName    pulumi.StringPtrInput
	JobAgentName      pulumi.StringInput
	Password          pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	Username          pulumi.StringInput
}

func (JobCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobCredentialArgs)(nil)).Elem()
}

type JobCredentialInput interface {
	pulumi.Input

	ToJobCredentialOutput() JobCredentialOutput
	ToJobCredentialOutputWithContext(ctx context.Context) JobCredentialOutput
}

func (*JobCredential) ElementType() reflect.Type {
	return reflect.TypeOf((**JobCredential)(nil)).Elem()
}

func (i *JobCredential) ToJobCredentialOutput() JobCredentialOutput {
	return i.ToJobCredentialOutputWithContext(context.Background())
}

func (i *JobCredential) ToJobCredentialOutputWithContext(ctx context.Context) JobCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobCredentialOutput)
}

type JobCredentialOutput struct{ *pulumi.OutputState }

func (JobCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobCredential)(nil)).Elem()
}

func (o JobCredentialOutput) ToJobCredentialOutput() JobCredentialOutput {
	return o
}

func (o JobCredentialOutput) ToJobCredentialOutputWithContext(ctx context.Context) JobCredentialOutput {
	return o
}

func (o JobCredentialOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JobCredential) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o JobCredentialOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JobCredential) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o JobCredentialOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *JobCredential) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobCredentialOutput{})
}
