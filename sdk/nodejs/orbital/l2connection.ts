// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Connects an edge site to an orbital gateway and describes what layer 2 traffic to forward between them.
 *
 * Uses Azure REST API version 2024-03-01-preview. In version 2.x of the Azure Native provider, it used API version 2024-03-01-preview.
 *
 * Other available API versions: 2024-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native orbital [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class L2Connection extends pulumi.CustomResource {
    /**
     * Get an existing L2Connection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): L2Connection {
        return new L2Connection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:orbital:L2Connection';

    /**
     * Returns true if the given object is an instance of L2Connection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is L2Connection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === L2Connection.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Globally-unique identifier for this connection that is to be used as a circuit ID.
     */
    public /*out*/ readonly circuitId!: pulumi.Output<string>;
    /**
     * A reference to an Microsoft.Orbital/edgeSites resource to route traffic for.
     */
    public readonly edgeSite!: pulumi.Output<outputs.orbital.L2ConnectionsPropertiesResponseEdgeSite>;
    /**
     * A reference to an Microsoft.Orbital/groundStations resource to route traffic for.
     */
    public readonly groundStation!: pulumi.Output<outputs.orbital.L2ConnectionsPropertiesResponseGroundStation>;
    /**
     * The name of the partner router to establish a connection to within the ground station.
     */
    public readonly groundStationPartnerRouter!: pulumi.Output<outputs.orbital.L2ConnectionsPropertiesResponseGroundStationPartnerRouter>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.orbital.SystemDataResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The VLAN ID for the L2 connection.
     */
    public readonly vlanId!: pulumi.Output<number>;

    /**
     * Create a L2Connection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: L2ConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.edgeSite === undefined) && !opts.urn) {
                throw new Error("Missing required property 'edgeSite'");
            }
            if ((!args || args.groundStation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groundStation'");
            }
            if ((!args || args.groundStationPartnerRouter === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groundStationPartnerRouter'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.vlanId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vlanId'");
            }
            resourceInputs["edgeSite"] = args ? args.edgeSite : undefined;
            resourceInputs["groundStation"] = args ? args.groundStation : undefined;
            resourceInputs["groundStationPartnerRouter"] = args ? args.groundStationPartnerRouter : undefined;
            resourceInputs["l2ConnectionName"] = args ? args.l2ConnectionName : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vlanId"] = args ? args.vlanId : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["circuitId"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["circuitId"] = undefined /*out*/;
            resourceInputs["edgeSite"] = undefined /*out*/;
            resourceInputs["groundStation"] = undefined /*out*/;
            resourceInputs["groundStationPartnerRouter"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["vlanId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:orbital/v20240301:L2Connection" }, { type: "azure-native:orbital/v20240301preview:L2Connection" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(L2Connection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a L2Connection resource.
 */
export interface L2ConnectionArgs {
    /**
     * A reference to an Microsoft.Orbital/edgeSites resource to route traffic for.
     */
    edgeSite: pulumi.Input<inputs.orbital.L2ConnectionsPropertiesEdgeSiteArgs>;
    /**
     * A reference to an Microsoft.Orbital/groundStations resource to route traffic for.
     */
    groundStation: pulumi.Input<inputs.orbital.L2ConnectionsPropertiesGroundStationArgs>;
    /**
     * The name of the partner router to establish a connection to within the ground station.
     */
    groundStationPartnerRouter: pulumi.Input<inputs.orbital.L2ConnectionsPropertiesGroundStationPartnerRouterArgs>;
    /**
     * L2 Connection name.
     */
    l2ConnectionName?: pulumi.Input<string>;
    /**
     * The geo-location where the resource lives
     */
    location?: pulumi.Input<string>;
    /**
     * The unique name of the partner router that cross-connects with the Orbital Edge Router at the edge site.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The VLAN ID for the L2 connection.
     */
    vlanId: pulumi.Input<number>;
}
