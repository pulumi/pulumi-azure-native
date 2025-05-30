// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * VirtualHubRouteTableV2 Resource.
 *
 * Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
 *
 * Other available API versions: 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class VirtualHubRouteTableV2 extends pulumi.CustomResource {
    /**
     * Get an existing VirtualHubRouteTableV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualHubRouteTableV2 {
        return new VirtualHubRouteTableV2(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:network:VirtualHubRouteTableV2';

    /**
     * Returns true if the given object is an instance of VirtualHubRouteTableV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualHubRouteTableV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualHubRouteTableV2.__pulumiType;
    }

    /**
     * List of all connections attached to this route table v2.
     */
    public readonly attachedConnections!: pulumi.Output<string[] | undefined>;
    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The provisioning state of the virtual hub route table v2 resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * List of all routes.
     */
    public readonly routes!: pulumi.Output<outputs.network.VirtualHubRouteV2Response[] | undefined>;

    /**
     * Create a VirtualHubRouteTableV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualHubRouteTableV2Args, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.virtualHubName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualHubName'");
            }
            resourceInputs["attachedConnections"] = args ? args.attachedConnections : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["routeTableName"] = args ? args.routeTableName : undefined;
            resourceInputs["routes"] = args ? args.routes : undefined;
            resourceInputs["virtualHubName"] = args ? args.virtualHubName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
        } else {
            resourceInputs["attachedConnections"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["routes"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:network/v20190901:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20191101:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20191201:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20200301:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20200401:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20200501:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20200601:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20200701:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20200801:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20201101:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20210201:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20210301:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20210501:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20210801:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20220101:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20220501:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20220701:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20220901:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20221101:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20230201:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20230401:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20230501:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20230601:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20230901:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20231101:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20240101:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20240301:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20240501:VirtualHubRouteTableV2" }, { type: "azure-native:network/v20240701:VirtualHubRouteTableV2" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(VirtualHubRouteTableV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualHubRouteTableV2 resource.
 */
export interface VirtualHubRouteTableV2Args {
    /**
     * List of all connections attached to this route table v2.
     */
    attachedConnections?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Resource ID.
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The resource group name of the VirtualHub.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the VirtualHubRouteTableV2.
     */
    routeTableName?: pulumi.Input<string>;
    /**
     * List of all routes.
     */
    routes?: pulumi.Input<pulumi.Input<inputs.network.VirtualHubRouteV2Args>[]>;
    /**
     * The name of the VirtualHub.
     */
    virtualHubName: pulumi.Input<string>;
}
