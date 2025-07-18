// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Action for alert rule.
 *
 * Uses Azure REST API version 2024-09-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
 *
 * Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-11-01, 2023-12-01-preview, 2024-01-01-preview, 2024-03-01, 2024-04-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-03-01, 2025-04-01-preview, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class Action extends pulumi.CustomResource {
    /**
     * Get an existing Action resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Action {
        return new Action(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:securityinsights:Action';

    /**
     * Returns true if the given object is an instance of Action.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Action {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Action.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Etag of the action.
     */
    public /*out*/ readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Logic App Resource Id, /subscriptions/{my-subscription}/resourceGroups/{my-resource-group}/providers/Microsoft.Logic/workflows/{my-workflow-id}.
     */
    public readonly logicAppResourceId!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.securityinsights.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The name of the logic app's workflow.
     */
    public /*out*/ readonly workflowId!: pulumi.Output<string | undefined>;

    /**
     * Create a Action resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.logicAppResourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logicAppResourceId'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.ruleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleId'");
            }
            if ((!args || args.triggerUri === undefined) && !opts.urn) {
                throw new Error("Missing required property 'triggerUri'");
            }
            if ((!args || args.workspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceName'");
            }
            resourceInputs["actionId"] = args ? args.actionId : undefined;
            resourceInputs["logicAppResourceId"] = args ? args.logicAppResourceId : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["ruleId"] = args ? args.ruleId : undefined;
            resourceInputs["triggerUri"] = args ? args.triggerUri : undefined;
            resourceInputs["workspaceName"] = args ? args.workspaceName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["workflowId"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["logicAppResourceId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["workflowId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:securityinsights/v20190101preview:Action" }, { type: "azure-native:securityinsights/v20200101:Action" }, { type: "azure-native:securityinsights/v20210301preview:Action" }, { type: "azure-native:securityinsights/v20210901preview:Action" }, { type: "azure-native:securityinsights/v20211001:Action" }, { type: "azure-native:securityinsights/v20211001preview:Action" }, { type: "azure-native:securityinsights/v20220101preview:Action" }, { type: "azure-native:securityinsights/v20220401preview:Action" }, { type: "azure-native:securityinsights/v20220501preview:Action" }, { type: "azure-native:securityinsights/v20220601preview:Action" }, { type: "azure-native:securityinsights/v20220701preview:Action" }, { type: "azure-native:securityinsights/v20220801:Action" }, { type: "azure-native:securityinsights/v20220801preview:Action" }, { type: "azure-native:securityinsights/v20220901preview:Action" }, { type: "azure-native:securityinsights/v20221001preview:Action" }, { type: "azure-native:securityinsights/v20221101:Action" }, { type: "azure-native:securityinsights/v20221101preview:Action" }, { type: "azure-native:securityinsights/v20221201preview:Action" }, { type: "azure-native:securityinsights/v20230201:Action" }, { type: "azure-native:securityinsights/v20230201preview:Action" }, { type: "azure-native:securityinsights/v20230301preview:Action" }, { type: "azure-native:securityinsights/v20230401preview:Action" }, { type: "azure-native:securityinsights/v20230501preview:Action" }, { type: "azure-native:securityinsights/v20230601preview:Action" }, { type: "azure-native:securityinsights/v20230701preview:Action" }, { type: "azure-native:securityinsights/v20230801preview:Action" }, { type: "azure-native:securityinsights/v20230901preview:Action" }, { type: "azure-native:securityinsights/v20231001preview:Action" }, { type: "azure-native:securityinsights/v20231101:Action" }, { type: "azure-native:securityinsights/v20231201preview:Action" }, { type: "azure-native:securityinsights/v20240101preview:Action" }, { type: "azure-native:securityinsights/v20240301:Action" }, { type: "azure-native:securityinsights/v20240401preview:Action" }, { type: "azure-native:securityinsights/v20240901:Action" }, { type: "azure-native:securityinsights/v20241001preview:Action" }, { type: "azure-native:securityinsights/v20250101preview:Action" }, { type: "azure-native:securityinsights/v20250301:Action" }, { type: "azure-native:securityinsights/v20250401preview:Action" }, { type: "azure-native:securityinsights/v20250601:Action" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Action.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Action resource.
 */
export interface ActionArgs {
    /**
     * Action ID
     */
    actionId?: pulumi.Input<string>;
    /**
     * Logic App Resource Id, /subscriptions/{my-subscription}/resourceGroups/{my-resource-group}/providers/Microsoft.Logic/workflows/{my-workflow-id}.
     */
    logicAppResourceId: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Alert rule ID
     */
    ruleId: pulumi.Input<string>;
    /**
     * Logic App Callback URL for this specific workflow.
     */
    triggerUri: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    workspaceName: pulumi.Input<string>;
}
