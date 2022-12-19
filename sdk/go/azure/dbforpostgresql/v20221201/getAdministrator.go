


package v20221201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAdministrator(ctx *pulumi.Context, args *LookupAdministratorArgs, opts ...pulumi.InvokeOption) (*LookupAdministratorResult, error) {
	var rv LookupAdministratorResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20221201:getAdministrator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAdministratorArgs struct {
	ObjectId          string `pulumi:"objectId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupAdministratorResult struct {
	Id            string             `pulumi:"id"`
	Name          string             `pulumi:"name"`
	ObjectId      *string            `pulumi:"objectId"`
	PrincipalName *string            `pulumi:"principalName"`
	PrincipalType *string            `pulumi:"principalType"`
	SystemData    SystemDataResponse `pulumi:"systemData"`
	TenantId      *string            `pulumi:"tenantId"`
	Type          string             `pulumi:"type"`
}

func LookupAdministratorOutput(ctx *pulumi.Context, args LookupAdministratorOutputArgs, opts ...pulumi.InvokeOption) LookupAdministratorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAdministratorResult, error) {
			args := v.(LookupAdministratorArgs)
			r, err := LookupAdministrator(ctx, &args, opts...)
			var s LookupAdministratorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAdministratorResultOutput)
}

type LookupAdministratorOutputArgs struct {
	ObjectId          pulumi.StringInput `pulumi:"objectId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupAdministratorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdministratorArgs)(nil)).Elem()
}


type LookupAdministratorResultOutput struct{ *pulumi.OutputState }

func (LookupAdministratorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdministratorResult)(nil)).Elem()
}

func (o LookupAdministratorResultOutput) ToLookupAdministratorResultOutput() LookupAdministratorResultOutput {
	return o
}

func (o LookupAdministratorResultOutput) ToLookupAdministratorResultOutputWithContext(ctx context.Context) LookupAdministratorResultOutput {
	return o
}

func (o LookupAdministratorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdministratorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAdministratorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdministratorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAdministratorResultOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAdministratorResult) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o LookupAdministratorResultOutput) PrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAdministratorResult) *string { return v.PrincipalName }).(pulumi.StringPtrOutput)
}

func (o LookupAdministratorResultOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAdministratorResult) *string { return v.PrincipalType }).(pulumi.StringPtrOutput)
}

func (o LookupAdministratorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAdministratorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAdministratorResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAdministratorResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupAdministratorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdministratorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAdministratorResultOutput{})
}
