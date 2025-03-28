// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Represents PurviewAudit data connector.
 */
export class PurviewAuditDataConnector extends pulumi.CustomResource {
    /**
     * Get an existing PurviewAuditDataConnector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PurviewAuditDataConnector {
        return new PurviewAuditDataConnector(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:securityinsights/v20241001preview:PurviewAuditDataConnector';

    /**
     * Returns true if the given object is an instance of PurviewAuditDataConnector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PurviewAuditDataConnector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PurviewAuditDataConnector.__pulumiType;
    }

    /**
     * The connector definition name (the dataConnectorDefinition resource id).
     */
    public readonly connectorDefinitionName!: pulumi.Output<string | undefined>;
    /**
     * The available data types for the connector.
     */
    public readonly dataTypes!: pulumi.Output<outputs.securityinsights.v20241001preview.PurviewAuditConnectorDataTypesResponse>;
    /**
     * The DCR related properties.
     */
    public readonly dcrConfig!: pulumi.Output<outputs.securityinsights.v20241001preview.DCRConfigurationResponse | undefined>;
    /**
     * Etag of the azure resource
     */
    public /*out*/ readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The kind of the data connector
     * Expected value is 'PurviewAudit'.
     */
    public readonly kind!: pulumi.Output<"PurviewAudit">;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The source type indicates which kind of data is relevant for this connector.
     */
    public readonly sourceType!: pulumi.Output<string | undefined>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.securityinsights.v20241001preview.SystemDataResponse>;
    /**
     * The tenant id to connect to, and get the data from.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a PurviewAuditDataConnector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PurviewAuditDataConnectorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dataTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataTypes'");
            }
            if ((!args || args.kind === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kind'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            if ((!args || args.workspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceName'");
            }
            resourceInputs["connectorDefinitionName"] = args ? args.connectorDefinitionName : undefined;
            resourceInputs["dataConnectorId"] = args ? args.dataConnectorId : undefined;
            resourceInputs["dataTypes"] = args ? args.dataTypes : undefined;
            resourceInputs["dcrConfig"] = args ? args.dcrConfig : undefined;
            resourceInputs["kind"] = "PurviewAudit";
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["sourceType"] = args ? args.sourceType : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["workspaceName"] = args ? args.workspaceName : undefined;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["connectorDefinitionName"] = undefined /*out*/;
            resourceInputs["dataTypes"] = undefined /*out*/;
            resourceInputs["dcrConfig"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["sourceType"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tenantId"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:securityinsights/v20190101preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20200101:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20210301preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20210901preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20211001:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20211001preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20220101preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20220401preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20220501preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20220601preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20220701preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20220801:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20220801preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20220901preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20221001preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20221101:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20221101preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20221201preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20230201:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20230201preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20230301preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20230401preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20230501preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20230601preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20230701preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20230801preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20230901preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20231001preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20231101:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20231201preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20240101preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20240301:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20240401preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20240901:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20250101preview:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights/v20250301:PurviewAuditDataConnector" }, { type: "azure-native:securityinsights:PurviewAuditDataConnector" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(PurviewAuditDataConnector.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PurviewAuditDataConnector resource.
 */
export interface PurviewAuditDataConnectorArgs {
    /**
     * The connector definition name (the dataConnectorDefinition resource id).
     */
    connectorDefinitionName?: pulumi.Input<string>;
    /**
     * Connector ID
     */
    dataConnectorId?: pulumi.Input<string>;
    /**
     * The available data types for the connector.
     */
    dataTypes: pulumi.Input<inputs.securityinsights.v20241001preview.PurviewAuditConnectorDataTypesArgs>;
    /**
     * The DCR related properties.
     */
    dcrConfig?: pulumi.Input<inputs.securityinsights.v20241001preview.DCRConfigurationArgs>;
    /**
     * The kind of the data connector
     * Expected value is 'PurviewAudit'.
     */
    kind: pulumi.Input<"PurviewAudit">;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The source type indicates which kind of data is relevant for this connector.
     */
    sourceType?: pulumi.Input<string>;
    /**
     * The tenant id to connect to, and get the data from.
     */
    tenantId: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    workspaceName: pulumi.Input<string>;
}
