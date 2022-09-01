


package v20150831preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupUserAssignedIdentity(ctx *pulumi.Context, args *LookupUserAssignedIdentityArgs, opts ...pulumi.InvokeOption) (*LookupUserAssignedIdentityResult, error) {
	var rv LookupUserAssignedIdentityResult
	err := ctx.Invoke("azure-native:managedidentity/v20150831preview:getUserAssignedIdentity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserAssignedIdentityArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupUserAssignedIdentityResult struct {
	ClientId        string            `pulumi:"clientId"`
	ClientSecretUrl string            `pulumi:"clientSecretUrl"`
	Id              string            `pulumi:"id"`
	Location        string            `pulumi:"location"`
	Name            string            `pulumi:"name"`
	PrincipalId     string            `pulumi:"principalId"`
	Tags            map[string]string `pulumi:"tags"`
	TenantId        string            `pulumi:"tenantId"`
	Type            string            `pulumi:"type"`
}

func LookupUserAssignedIdentityOutput(ctx *pulumi.Context, args LookupUserAssignedIdentityOutputArgs, opts ...pulumi.InvokeOption) LookupUserAssignedIdentityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserAssignedIdentityResult, error) {
			args := v.(LookupUserAssignedIdentityArgs)
			r, err := LookupUserAssignedIdentity(ctx, &args, opts...)
			var s LookupUserAssignedIdentityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserAssignedIdentityResultOutput)
}

type LookupUserAssignedIdentityOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupUserAssignedIdentityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserAssignedIdentityArgs)(nil)).Elem()
}


type LookupUserAssignedIdentityResultOutput struct{ *pulumi.OutputState }

func (LookupUserAssignedIdentityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserAssignedIdentityResult)(nil)).Elem()
}

func (o LookupUserAssignedIdentityResultOutput) ToLookupUserAssignedIdentityResultOutput() LookupUserAssignedIdentityResultOutput {
	return o
}

func (o LookupUserAssignedIdentityResultOutput) ToLookupUserAssignedIdentityResultOutputWithContext(ctx context.Context) LookupUserAssignedIdentityResultOutput {
	return o
}

func (o LookupUserAssignedIdentityResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o LookupUserAssignedIdentityResultOutput) ClientSecretUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) string { return v.ClientSecretUrl }).(pulumi.StringOutput)
}

func (o LookupUserAssignedIdentityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupUserAssignedIdentityResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupUserAssignedIdentityResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupUserAssignedIdentityResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o LookupUserAssignedIdentityResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupUserAssignedIdentityResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupUserAssignedIdentityResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserAssignedIdentityResultOutput{})
}
