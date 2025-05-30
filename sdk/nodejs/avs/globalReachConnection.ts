// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A global reach connection resource
 *
 * Uses Azure REST API version 2023-09-01. In version 2.x of the Azure Native provider, it used API version 2022-05-01.
 *
 * Other available API versions: 2022-05-01, 2023-03-01, 2024-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native avs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class GlobalReachConnection extends pulumi.CustomResource {
    /**
     * Get an existing GlobalReachConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GlobalReachConnection {
        return new GlobalReachConnection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:avs:GlobalReachConnection';

    /**
     * Returns true if the given object is an instance of GlobalReachConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GlobalReachConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GlobalReachConnection.__pulumiType;
    }

    /**
     * The network used for global reach carved out from the original network block
     * provided for the private cloud
     */
    public /*out*/ readonly addressPrefix!: pulumi.Output<string>;
    /**
     * Authorization key from the peer express route used for the global reach
     * connection
     */
    public readonly authorizationKey!: pulumi.Output<string | undefined>;
    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The connection status of the global reach connection
     */
    public /*out*/ readonly circuitConnectionStatus!: pulumi.Output<string>;
    /**
     * The ID of the Private Cloud's ExpressRoute Circuit that is participating in the
     * global reach connection
     */
    public readonly expressRouteId!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Identifier of the ExpressRoute Circuit to peer with in the global reach
     * connection
     */
    public readonly peerExpressRouteCircuit!: pulumi.Output<string | undefined>;
    /**
     * The state of the  ExpressRoute Circuit Authorization provisioning
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.avs.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a GlobalReachConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GlobalReachConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.privateCloudName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateCloudName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["authorizationKey"] = args ? args.authorizationKey : undefined;
            resourceInputs["expressRouteId"] = args ? args.expressRouteId : undefined;
            resourceInputs["globalReachConnectionName"] = args ? args.globalReachConnectionName : undefined;
            resourceInputs["peerExpressRouteCircuit"] = args ? args.peerExpressRouteCircuit : undefined;
            resourceInputs["privateCloudName"] = args ? args.privateCloudName : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["addressPrefix"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["circuitConnectionStatus"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["addressPrefix"] = undefined /*out*/;
            resourceInputs["authorizationKey"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["circuitConnectionStatus"] = undefined /*out*/;
            resourceInputs["expressRouteId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["peerExpressRouteCircuit"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:avs/v20200717preview:GlobalReachConnection" }, { type: "azure-native:avs/v20210101preview:GlobalReachConnection" }, { type: "azure-native:avs/v20210601:GlobalReachConnection" }, { type: "azure-native:avs/v20211201:GlobalReachConnection" }, { type: "azure-native:avs/v20220501:GlobalReachConnection" }, { type: "azure-native:avs/v20230301:GlobalReachConnection" }, { type: "azure-native:avs/v20230901:GlobalReachConnection" }, { type: "azure-native:avs/v20240901:GlobalReachConnection" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(GlobalReachConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a GlobalReachConnection resource.
 */
export interface GlobalReachConnectionArgs {
    /**
     * Authorization key from the peer express route used for the global reach
     * connection
     */
    authorizationKey?: pulumi.Input<string>;
    /**
     * The ID of the Private Cloud's ExpressRoute Circuit that is participating in the
     * global reach connection
     */
    expressRouteId?: pulumi.Input<string>;
    /**
     * Name of the global reach connection
     */
    globalReachConnectionName?: pulumi.Input<string>;
    /**
     * Identifier of the ExpressRoute Circuit to peer with in the global reach
     * connection
     */
    peerExpressRouteCircuit?: pulumi.Input<string>;
    /**
     * Name of the private cloud
     */
    privateCloudName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
