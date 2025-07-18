// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents an attached NetworkConnection.
 *
 * Uses Azure REST API version 2024-02-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01.
 *
 * Other available API versions: 2023-04-01, 2023-08-01-preview, 2023-10-01-preview, 2024-05-01-preview, 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-02-01, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class AttachedNetworkByDevCenter extends pulumi.CustomResource {
    /**
     * Get an existing AttachedNetworkByDevCenter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AttachedNetworkByDevCenter {
        return new AttachedNetworkByDevCenter(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:devcenter:AttachedNetworkByDevCenter';

    /**
     * Returns true if the given object is an instance of AttachedNetworkByDevCenter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AttachedNetworkByDevCenter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AttachedNetworkByDevCenter.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * AAD Join type of the network. This is populated based on the referenced Network Connection.
     */
    public /*out*/ readonly domainJoinType!: pulumi.Output<string>;
    /**
     * Health check status values
     */
    public /*out*/ readonly healthCheckStatus!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The resource ID of the NetworkConnection you want to attach.
     */
    public readonly networkConnectionId!: pulumi.Output<string>;
    /**
     * The geo-location where the NetworkConnection resource specified in 'networkConnectionResourceId' property lives.
     */
    public /*out*/ readonly networkConnectionLocation!: pulumi.Output<string>;
    /**
     * The provisioning state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.devcenter.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AttachedNetworkByDevCenter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AttachedNetworkByDevCenterArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.devCenterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'devCenterName'");
            }
            if ((!args || args.networkConnectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkConnectionId'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["attachedNetworkConnectionName"] = args ? args.attachedNetworkConnectionName : undefined;
            resourceInputs["devCenterName"] = args ? args.devCenterName : undefined;
            resourceInputs["networkConnectionId"] = args ? args.networkConnectionId : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["domainJoinType"] = undefined /*out*/;
            resourceInputs["healthCheckStatus"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["networkConnectionLocation"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["domainJoinType"] = undefined /*out*/;
            resourceInputs["healthCheckStatus"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["networkConnectionId"] = undefined /*out*/;
            resourceInputs["networkConnectionLocation"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:devcenter/v20220801preview:AttachedNetworkByDevCenter" }, { type: "azure-native:devcenter/v20220901preview:AttachedNetworkByDevCenter" }, { type: "azure-native:devcenter/v20221012preview:AttachedNetworkByDevCenter" }, { type: "azure-native:devcenter/v20221111preview:AttachedNetworkByDevCenter" }, { type: "azure-native:devcenter/v20230101preview:AttachedNetworkByDevCenter" }, { type: "azure-native:devcenter/v20230401:AttachedNetworkByDevCenter" }, { type: "azure-native:devcenter/v20230801preview:AttachedNetworkByDevCenter" }, { type: "azure-native:devcenter/v20231001preview:AttachedNetworkByDevCenter" }, { type: "azure-native:devcenter/v20240201:AttachedNetworkByDevCenter" }, { type: "azure-native:devcenter/v20240501preview:AttachedNetworkByDevCenter" }, { type: "azure-native:devcenter/v20240601preview:AttachedNetworkByDevCenter" }, { type: "azure-native:devcenter/v20240701preview:AttachedNetworkByDevCenter" }, { type: "azure-native:devcenter/v20240801preview:AttachedNetworkByDevCenter" }, { type: "azure-native:devcenter/v20241001preview:AttachedNetworkByDevCenter" }, { type: "azure-native:devcenter/v20250201:AttachedNetworkByDevCenter" }, { type: "azure-native:devcenter/v20250401preview:AttachedNetworkByDevCenter" }, { type: "azure-native:devcenter/v20250701preview:AttachedNetworkByDevCenter" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(AttachedNetworkByDevCenter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AttachedNetworkByDevCenter resource.
 */
export interface AttachedNetworkByDevCenterArgs {
    /**
     * The name of the attached NetworkConnection.
     */
    attachedNetworkConnectionName?: pulumi.Input<string>;
    /**
     * The name of the devcenter.
     */
    devCenterName: pulumi.Input<string>;
    /**
     * The resource ID of the NetworkConnection you want to attach.
     */
    networkConnectionId: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
