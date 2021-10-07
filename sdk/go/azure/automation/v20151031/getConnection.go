


package v20151031

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnection(ctx *pulumi.Context, args *LookupConnectionArgs, opts ...pulumi.InvokeOption) (*LookupConnectionResult, error) {
	var rv LookupConnectionResult
	err := ctx.Invoke("azure-native:automation/v20151031:getConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ConnectionName        string `pulumi:"connectionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupConnectionResult struct {
	ConnectionType        *ConnectionTypeAssociationPropertyResponse `pulumi:"connectionType"`
	CreationTime          string                                     `pulumi:"creationTime"`
	Description           *string                                    `pulumi:"description"`
	FieldDefinitionValues map[string]string                          `pulumi:"fieldDefinitionValues"`
	Id                    string                                     `pulumi:"id"`
	LastModifiedTime      string                                     `pulumi:"lastModifiedTime"`
	Name                  string                                     `pulumi:"name"`
	Type                  string                                     `pulumi:"type"`
}
