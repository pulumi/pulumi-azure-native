


package v20211115preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMongoDBResourceMongoRoleDefinition(ctx *pulumi.Context, args *LookupMongoDBResourceMongoRoleDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupMongoDBResourceMongoRoleDefinitionResult, error) {
	var rv LookupMongoDBResourceMongoRoleDefinitionResult
	err := ctx.Invoke("azure-native:documentdb/v20211115preview:getMongoDBResourceMongoRoleDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMongoDBResourceMongoRoleDefinitionArgs struct {
	AccountName           string `pulumi:"accountName"`
	MongoRoleDefinitionId string `pulumi:"mongoRoleDefinitionId"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupMongoDBResourceMongoRoleDefinitionResult struct {
	DatabaseName *string             `pulumi:"databaseName"`
	Id           string              `pulumi:"id"`
	Name         string              `pulumi:"name"`
	Privileges   []PrivilegeResponse `pulumi:"privileges"`
	RoleName     *string             `pulumi:"roleName"`
	Roles        []RoleResponse      `pulumi:"roles"`
	Type         string              `pulumi:"type"`
}

func LookupMongoDBResourceMongoRoleDefinitionOutput(ctx *pulumi.Context, args LookupMongoDBResourceMongoRoleDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupMongoDBResourceMongoRoleDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMongoDBResourceMongoRoleDefinitionResult, error) {
			args := v.(LookupMongoDBResourceMongoRoleDefinitionArgs)
			r, err := LookupMongoDBResourceMongoRoleDefinition(ctx, &args, opts...)
			var s LookupMongoDBResourceMongoRoleDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMongoDBResourceMongoRoleDefinitionResultOutput)
}

type LookupMongoDBResourceMongoRoleDefinitionOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	MongoRoleDefinitionId pulumi.StringInput `pulumi:"mongoRoleDefinitionId"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMongoDBResourceMongoRoleDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDBResourceMongoRoleDefinitionArgs)(nil)).Elem()
}


type LookupMongoDBResourceMongoRoleDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupMongoDBResourceMongoRoleDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDBResourceMongoRoleDefinitionResult)(nil)).Elem()
}

func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) ToLookupMongoDBResourceMongoRoleDefinitionResultOutput() LookupMongoDBResourceMongoRoleDefinitionResultOutput {
	return o
}

func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) ToLookupMongoDBResourceMongoRoleDefinitionResultOutputWithContext(ctx context.Context) LookupMongoDBResourceMongoRoleDefinitionResultOutput {
	return o
}

func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoRoleDefinitionResult) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoRoleDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoRoleDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) Privileges() PrivilegeResponseArrayOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoRoleDefinitionResult) []PrivilegeResponse { return v.Privileges }).(PrivilegeResponseArrayOutput)
}

func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoRoleDefinitionResult) *string { return v.RoleName }).(pulumi.StringPtrOutput)
}

func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) Roles() RoleResponseArrayOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoRoleDefinitionResult) []RoleResponse { return v.Roles }).(RoleResponseArrayOutput)
}

func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoRoleDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMongoDBResourceMongoRoleDefinitionResultOutput{})
}
