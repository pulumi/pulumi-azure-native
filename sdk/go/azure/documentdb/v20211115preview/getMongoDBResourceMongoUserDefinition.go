


package v20211115preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMongoDBResourceMongoUserDefinition(ctx *pulumi.Context, args *LookupMongoDBResourceMongoUserDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupMongoDBResourceMongoUserDefinitionResult, error) {
	var rv LookupMongoDBResourceMongoUserDefinitionResult
	err := ctx.Invoke("azure-native:documentdb/v20211115preview:getMongoDBResourceMongoUserDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMongoDBResourceMongoUserDefinitionArgs struct {
	AccountName           string `pulumi:"accountName"`
	MongoUserDefinitionId string `pulumi:"mongoUserDefinitionId"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupMongoDBResourceMongoUserDefinitionResult struct {
	CustomData   *string        `pulumi:"customData"`
	DatabaseName *string        `pulumi:"databaseName"`
	Id           string         `pulumi:"id"`
	Mechanisms   *string        `pulumi:"mechanisms"`
	Name         string         `pulumi:"name"`
	Password     *string        `pulumi:"password"`
	Roles        []RoleResponse `pulumi:"roles"`
	Type         string         `pulumi:"type"`
	UserName     *string        `pulumi:"userName"`
}

func LookupMongoDBResourceMongoUserDefinitionOutput(ctx *pulumi.Context, args LookupMongoDBResourceMongoUserDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupMongoDBResourceMongoUserDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMongoDBResourceMongoUserDefinitionResult, error) {
			args := v.(LookupMongoDBResourceMongoUserDefinitionArgs)
			r, err := LookupMongoDBResourceMongoUserDefinition(ctx, &args, opts...)
			var s LookupMongoDBResourceMongoUserDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMongoDBResourceMongoUserDefinitionResultOutput)
}

type LookupMongoDBResourceMongoUserDefinitionOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	MongoUserDefinitionId pulumi.StringInput `pulumi:"mongoUserDefinitionId"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMongoDBResourceMongoUserDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDBResourceMongoUserDefinitionArgs)(nil)).Elem()
}


type LookupMongoDBResourceMongoUserDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupMongoDBResourceMongoUserDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDBResourceMongoUserDefinitionResult)(nil)).Elem()
}

func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) ToLookupMongoDBResourceMongoUserDefinitionResultOutput() LookupMongoDBResourceMongoUserDefinitionResultOutput {
	return o
}

func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) ToLookupMongoDBResourceMongoUserDefinitionResultOutputWithContext(ctx context.Context) LookupMongoDBResourceMongoUserDefinitionResultOutput {
	return o
}

func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) Mechanisms() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) *string { return v.Mechanisms }).(pulumi.StringPtrOutput)
}

func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) Roles() RoleResponseArrayOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) []RoleResponse { return v.Roles }).(RoleResponseArrayOutput)
}

func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMongoDBResourceMongoUserDefinitionResultOutput{})
}
