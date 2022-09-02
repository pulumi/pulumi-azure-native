


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupActivityCustomEntityQuery(ctx *pulumi.Context, args *LookupActivityCustomEntityQueryArgs, opts ...pulumi.InvokeOption) (*LookupActivityCustomEntityQueryResult, error) {
	var rv LookupActivityCustomEntityQueryResult
	err := ctx.Invoke("azure-native:securityinsights/v20220501preview:getActivityCustomEntityQuery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupActivityCustomEntityQueryArgs struct {
	EntityQueryId     string `pulumi:"entityQueryId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupActivityCustomEntityQueryResult struct {
	Content                 *string                                                  `pulumi:"content"`
	CreatedTimeUtc          string                                                   `pulumi:"createdTimeUtc"`
	Description             *string                                                  `pulumi:"description"`
	Enabled                 *bool                                                    `pulumi:"enabled"`
	EntitiesFilter          map[string][]string                                      `pulumi:"entitiesFilter"`
	Etag                    *string                                                  `pulumi:"etag"`
	Id                      string                                                   `pulumi:"id"`
	InputEntityType         *string                                                  `pulumi:"inputEntityType"`
	Kind                    string                                                   `pulumi:"kind"`
	LastModifiedTimeUtc     string                                                   `pulumi:"lastModifiedTimeUtc"`
	Name                    string                                                   `pulumi:"name"`
	QueryDefinitions        *ActivityEntityQueriesPropertiesResponseQueryDefinitions `pulumi:"queryDefinitions"`
	RequiredInputFieldsSets [][]string                                               `pulumi:"requiredInputFieldsSets"`
	SystemData              SystemDataResponse                                       `pulumi:"systemData"`
	TemplateName            *string                                                  `pulumi:"templateName"`
	Title                   *string                                                  `pulumi:"title"`
	Type                    string                                                   `pulumi:"type"`
}

func LookupActivityCustomEntityQueryOutput(ctx *pulumi.Context, args LookupActivityCustomEntityQueryOutputArgs, opts ...pulumi.InvokeOption) LookupActivityCustomEntityQueryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupActivityCustomEntityQueryResult, error) {
			args := v.(LookupActivityCustomEntityQueryArgs)
			r, err := LookupActivityCustomEntityQuery(ctx, &args, opts...)
			var s LookupActivityCustomEntityQueryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupActivityCustomEntityQueryResultOutput)
}

type LookupActivityCustomEntityQueryOutputArgs struct {
	EntityQueryId     pulumi.StringInput `pulumi:"entityQueryId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupActivityCustomEntityQueryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActivityCustomEntityQueryArgs)(nil)).Elem()
}


type LookupActivityCustomEntityQueryResultOutput struct{ *pulumi.OutputState }

func (LookupActivityCustomEntityQueryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActivityCustomEntityQueryResult)(nil)).Elem()
}

func (o LookupActivityCustomEntityQueryResultOutput) ToLookupActivityCustomEntityQueryResultOutput() LookupActivityCustomEntityQueryResultOutput {
	return o
}

func (o LookupActivityCustomEntityQueryResultOutput) ToLookupActivityCustomEntityQueryResultOutputWithContext(ctx context.Context) LookupActivityCustomEntityQueryResultOutput {
	return o
}

func (o LookupActivityCustomEntityQueryResultOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o LookupActivityCustomEntityQueryResultOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) string { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

func (o LookupActivityCustomEntityQueryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupActivityCustomEntityQueryResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupActivityCustomEntityQueryResultOutput) EntitiesFilter() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) map[string][]string { return v.EntitiesFilter }).(pulumi.StringArrayMapOutput)
}

func (o LookupActivityCustomEntityQueryResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupActivityCustomEntityQueryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupActivityCustomEntityQueryResultOutput) InputEntityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) *string { return v.InputEntityType }).(pulumi.StringPtrOutput)
}

func (o LookupActivityCustomEntityQueryResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupActivityCustomEntityQueryResultOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) string { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

func (o LookupActivityCustomEntityQueryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupActivityCustomEntityQueryResultOutput) QueryDefinitions() ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) *ActivityEntityQueriesPropertiesResponseQueryDefinitions {
		return v.QueryDefinitions
	}).(ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput)
}

func (o LookupActivityCustomEntityQueryResultOutput) RequiredInputFieldsSets() pulumi.StringArrayArrayOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) [][]string { return v.RequiredInputFieldsSets }).(pulumi.StringArrayArrayOutput)
}

func (o LookupActivityCustomEntityQueryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupActivityCustomEntityQueryResultOutput) TemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) *string { return v.TemplateName }).(pulumi.StringPtrOutput)
}

func (o LookupActivityCustomEntityQueryResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o LookupActivityCustomEntityQueryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupActivityCustomEntityQueryResultOutput{})
}
