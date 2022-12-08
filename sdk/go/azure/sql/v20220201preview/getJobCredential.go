


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJobCredential(ctx *pulumi.Context, args *LookupJobCredentialArgs, opts ...pulumi.InvokeOption) (*LookupJobCredentialResult, error) {
	var rv LookupJobCredentialResult
	err := ctx.Invoke("azure-native:sql/v20220201preview:getJobCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobCredentialArgs struct {
	CredentialName    string `pulumi:"credentialName"`
	JobAgentName      string `pulumi:"jobAgentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupJobCredentialResult struct {
	Id       string `pulumi:"id"`
	Name     string `pulumi:"name"`
	Type     string `pulumi:"type"`
	Username string `pulumi:"username"`
}

func LookupJobCredentialOutput(ctx *pulumi.Context, args LookupJobCredentialOutputArgs, opts ...pulumi.InvokeOption) LookupJobCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobCredentialResult, error) {
			args := v.(LookupJobCredentialArgs)
			r, err := LookupJobCredential(ctx, &args, opts...)
			var s LookupJobCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobCredentialResultOutput)
}

type LookupJobCredentialOutputArgs struct {
	CredentialName    pulumi.StringInput `pulumi:"credentialName"`
	JobAgentName      pulumi.StringInput `pulumi:"jobAgentName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupJobCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobCredentialArgs)(nil)).Elem()
}


type LookupJobCredentialResultOutput struct{ *pulumi.OutputState }

func (LookupJobCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobCredentialResult)(nil)).Elem()
}

func (o LookupJobCredentialResultOutput) ToLookupJobCredentialResultOutput() LookupJobCredentialResultOutput {
	return o
}

func (o LookupJobCredentialResultOutput) ToLookupJobCredentialResultOutputWithContext(ctx context.Context) LookupJobCredentialResultOutput {
	return o
}

func (o LookupJobCredentialResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobCredentialResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupJobCredentialResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobCredentialResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupJobCredentialResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobCredentialResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupJobCredentialResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobCredentialResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobCredentialResultOutput{})
}
