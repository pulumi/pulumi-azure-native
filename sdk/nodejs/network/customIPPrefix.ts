// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Custom IP prefix resource.
 *
 * Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
 *
 * Other available API versions: 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class CustomIPPrefix extends pulumi.CustomResource {
    /**
     * Get an existing CustomIPPrefix resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomIPPrefix {
        return new CustomIPPrefix(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:network:CustomIPPrefix';

    /**
     * Returns true if the given object is an instance of CustomIPPrefix.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomIPPrefix {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomIPPrefix.__pulumiType;
    }

    /**
     * The ASN for CIDR advertising. Should be an integer as string.
     */
    public readonly asn!: pulumi.Output<string | undefined>;
    /**
     * Authorization message for WAN validation.
     */
    public readonly authorizationMessage!: pulumi.Output<string | undefined>;
    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The list of all Children for IPv6 /48 CustomIpPrefix.
     */
    public /*out*/ readonly childCustomIpPrefixes!: pulumi.Output<outputs.network.SubResourceResponse[]>;
    /**
     * The prefix range in CIDR notation. Should include the start address and the prefix length.
     */
    public readonly cidr!: pulumi.Output<string | undefined>;
    /**
     * The commissioned state of the Custom IP Prefix.
     */
    public readonly commissionedState!: pulumi.Output<string | undefined>;
    /**
     * The Parent CustomIpPrefix for IPv6 /64 CustomIpPrefix.
     */
    public readonly customIpPrefixParent!: pulumi.Output<outputs.network.SubResourceResponse | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Whether to do express route advertise.
     */
    public readonly expressRouteAdvertise!: pulumi.Output<boolean | undefined>;
    /**
     * The extended location of the custom IP prefix.
     */
    public readonly extendedLocation!: pulumi.Output<outputs.network.ExtendedLocationResponse | undefined>;
    /**
     * The reason why resource is in failed state.
     */
    public /*out*/ readonly failedReason!: pulumi.Output<string>;
    /**
     * The Geo for CIDR advertising. Should be an Geo code.
     */
    public readonly geo!: pulumi.Output<string | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Whether to Advertise the range to Internet.
     */
    public readonly noInternetAdvertise!: pulumi.Output<boolean | undefined>;
    /**
     * Type of custom IP prefix. Should be Singular, Parent, or Child.
     */
    public readonly prefixType!: pulumi.Output<string | undefined>;
    /**
     * The provisioning state of the custom IP prefix resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The list of all referenced PublicIpPrefixes.
     */
    public /*out*/ readonly publicIpPrefixes!: pulumi.Output<outputs.network.SubResourceResponse[]>;
    /**
     * The resource GUID property of the custom IP prefix resource.
     */
    public /*out*/ readonly resourceGuid!: pulumi.Output<string>;
    /**
     * Signed message for WAN validation.
     */
    public readonly signedMessage!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * A list of availability zones denoting the IP allocated for the resource needs to come from.
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a CustomIPPrefix resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomIPPrefixArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["asn"] = args ? args.asn : undefined;
            resourceInputs["authorizationMessage"] = args ? args.authorizationMessage : undefined;
            resourceInputs["cidr"] = args ? args.cidr : undefined;
            resourceInputs["commissionedState"] = args ? args.commissionedState : undefined;
            resourceInputs["customIpPrefixName"] = args ? args.customIpPrefixName : undefined;
            resourceInputs["customIpPrefixParent"] = args ? args.customIpPrefixParent : undefined;
            resourceInputs["expressRouteAdvertise"] = args ? args.expressRouteAdvertise : undefined;
            resourceInputs["extendedLocation"] = args ? args.extendedLocation : undefined;
            resourceInputs["geo"] = args ? args.geo : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["noInternetAdvertise"] = args ? args.noInternetAdvertise : undefined;
            resourceInputs["prefixType"] = args ? args.prefixType : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["signedMessage"] = args ? args.signedMessage : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["zones"] = args ? args.zones : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["childCustomIpPrefixes"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["failedReason"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["publicIpPrefixes"] = undefined /*out*/;
            resourceInputs["resourceGuid"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["asn"] = undefined /*out*/;
            resourceInputs["authorizationMessage"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["childCustomIpPrefixes"] = undefined /*out*/;
            resourceInputs["cidr"] = undefined /*out*/;
            resourceInputs["commissionedState"] = undefined /*out*/;
            resourceInputs["customIpPrefixParent"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["expressRouteAdvertise"] = undefined /*out*/;
            resourceInputs["extendedLocation"] = undefined /*out*/;
            resourceInputs["failedReason"] = undefined /*out*/;
            resourceInputs["geo"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["noInternetAdvertise"] = undefined /*out*/;
            resourceInputs["prefixType"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["publicIpPrefixes"] = undefined /*out*/;
            resourceInputs["resourceGuid"] = undefined /*out*/;
            resourceInputs["signedMessage"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["zones"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:network/v20200601:CustomIPPrefix" }, { type: "azure-native:network/v20200701:CustomIPPrefix" }, { type: "azure-native:network/v20200801:CustomIPPrefix" }, { type: "azure-native:network/v20201101:CustomIPPrefix" }, { type: "azure-native:network/v20210201:CustomIPPrefix" }, { type: "azure-native:network/v20210301:CustomIPPrefix" }, { type: "azure-native:network/v20210501:CustomIPPrefix" }, { type: "azure-native:network/v20210801:CustomIPPrefix" }, { type: "azure-native:network/v20220101:CustomIPPrefix" }, { type: "azure-native:network/v20220501:CustomIPPrefix" }, { type: "azure-native:network/v20220701:CustomIPPrefix" }, { type: "azure-native:network/v20220901:CustomIPPrefix" }, { type: "azure-native:network/v20221101:CustomIPPrefix" }, { type: "azure-native:network/v20230201:CustomIPPrefix" }, { type: "azure-native:network/v20230401:CustomIPPrefix" }, { type: "azure-native:network/v20230501:CustomIPPrefix" }, { type: "azure-native:network/v20230601:CustomIPPrefix" }, { type: "azure-native:network/v20230901:CustomIPPrefix" }, { type: "azure-native:network/v20231101:CustomIPPrefix" }, { type: "azure-native:network/v20240101:CustomIPPrefix" }, { type: "azure-native:network/v20240301:CustomIPPrefix" }, { type: "azure-native:network/v20240501:CustomIPPrefix" }, { type: "azure-native:network/v20240701:CustomIPPrefix" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(CustomIPPrefix.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CustomIPPrefix resource.
 */
export interface CustomIPPrefixArgs {
    /**
     * The ASN for CIDR advertising. Should be an integer as string.
     */
    asn?: pulumi.Input<string>;
    /**
     * Authorization message for WAN validation.
     */
    authorizationMessage?: pulumi.Input<string>;
    /**
     * The prefix range in CIDR notation. Should include the start address and the prefix length.
     */
    cidr?: pulumi.Input<string>;
    /**
     * The commissioned state of the Custom IP Prefix.
     */
    commissionedState?: pulumi.Input<string | enums.network.CommissionedState>;
    /**
     * The name of the custom IP prefix.
     */
    customIpPrefixName?: pulumi.Input<string>;
    /**
     * The Parent CustomIpPrefix for IPv6 /64 CustomIpPrefix.
     */
    customIpPrefixParent?: pulumi.Input<inputs.network.SubResourceArgs>;
    /**
     * Whether to do express route advertise.
     */
    expressRouteAdvertise?: pulumi.Input<boolean>;
    /**
     * The extended location of the custom IP prefix.
     */
    extendedLocation?: pulumi.Input<inputs.network.ExtendedLocationArgs>;
    /**
     * The Geo for CIDR advertising. Should be an Geo code.
     */
    geo?: pulumi.Input<string | enums.network.Geo>;
    /**
     * Resource ID.
     */
    id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    location?: pulumi.Input<string>;
    /**
     * Whether to Advertise the range to Internet.
     */
    noInternetAdvertise?: pulumi.Input<boolean>;
    /**
     * Type of custom IP prefix. Should be Singular, Parent, or Child.
     */
    prefixType?: pulumi.Input<string | enums.network.CustomIpPrefixType>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Signed message for WAN validation.
     */
    signedMessage?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of availability zones denoting the IP allocated for the resource needs to come from.
     */
    zones?: pulumi.Input<pulumi.Input<string>[]>;
}
