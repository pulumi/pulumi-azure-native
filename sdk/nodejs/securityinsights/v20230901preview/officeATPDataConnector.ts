// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Represents OfficeATP (Office 365 Advanced Threat Protection) data connector.
 */
export class OfficeATPDataConnector extends pulumi.CustomResource {
    /**
     * Get an existing OfficeATPDataConnector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OfficeATPDataConnector {
        return new OfficeATPDataConnector(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:securityinsights/v20230901preview:OfficeATPDataConnector';

    /**
     * Returns true if the given object is an instance of OfficeATPDataConnector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OfficeATPDataConnector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OfficeATPDataConnector.__pulumiType;
    }

    /**
     * The available data types for the connector.
     */
    public readonly dataTypes!: pulumi.Output<outputs.securityinsights.v20230901preview.AlertsDataTypeOfDataConnectorResponse | undefined>;
    /**
     * Etag of the azure resource
     */
    public /*out*/ readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The kind of the data connector
     * Expected value is 'OfficeATP'.
     */
    public readonly kind!: pulumi.Output<"OfficeATP">;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.securityinsights.v20230901preview.SystemDataResponse>;
    /**
     * The tenant id to connect to, and get the data from.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a OfficeATPDataConnector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OfficeATPDataConnectorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
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
            resourceInputs["dataConnectorId"] = args ? args.dataConnectorId : undefined;
            resourceInputs["dataTypes"] = args ? args.dataTypes : undefined;
            resourceInputs["kind"] = "OfficeATP";
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["workspaceName"] = args ? args.workspaceName : undefined;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["dataTypes"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tenantId"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:securityinsights/v20190101preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20200101:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20210301preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20210901preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20211001:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20211001preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20220101preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20220401preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20220501preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20220601preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20220701preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20220801:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20220801preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20220901preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20221001preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20221101:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20221101preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20221201preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20230201:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20230201preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20230301preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20230401preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20230501preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20230601preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20230701preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20230801preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20231001preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20231101:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20231201preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20240101preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20240301:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20240401preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20240901:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20241001preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20250101preview:OfficeATPDataConnector" }, { type: "azure-native:securityinsights/v20250301:OfficeATPDataConnector" }, { type: "azure-native:securityinsights:OfficeATPDataConnector" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(OfficeATPDataConnector.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a OfficeATPDataConnector resource.
 */
export interface OfficeATPDataConnectorArgs {
    /**
     * Connector ID
     */
    dataConnectorId?: pulumi.Input<string>;
    /**
     * The available data types for the connector.
     */
    dataTypes?: pulumi.Input<inputs.securityinsights.v20230901preview.AlertsDataTypeOfDataConnectorArgs>;
    /**
     * The kind of the data connector
     * Expected value is 'OfficeATP'.
     */
    kind: pulumi.Input<"OfficeATP">;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The tenant id to connect to, and get the data from.
     */
    tenantId: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    workspaceName: pulumi.Input<string>;
}
