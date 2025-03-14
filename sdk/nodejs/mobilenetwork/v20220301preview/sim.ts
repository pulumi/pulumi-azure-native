// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Sim resource.
 */
export class Sim extends pulumi.CustomResource {
    /**
     * Get an existing Sim resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Sim {
        return new Sim(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:mobilenetwork/v20220301preview:Sim';

    /**
     * Returns true if the given object is an instance of Sim.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Sim {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Sim.__pulumiType;
    }

    /**
     * The timestamp of resource creation (UTC).
     */
    public readonly createdAt!: pulumi.Output<string | undefined>;
    /**
     * The identity that created the resource.
     */
    public readonly createdBy!: pulumi.Output<string | undefined>;
    /**
     * The type of identity that created the resource.
     */
    public readonly createdByType!: pulumi.Output<string | undefined>;
    /**
     * An optional free-form text field that can be used to record the device type this sim is associated with, for example 'Video camera'. The Azure portal allows Sims to be grouped and filtered based on this value.
     */
    public readonly deviceType!: pulumi.Output<string | undefined>;
    /**
     * The Integrated Circuit Card ID (ICC Id) for the sim.
     */
    public readonly integratedCircuitCardIdentifier!: pulumi.Output<string | undefined>;
    /**
     * The International Mobile Subscriber Identity (IMSI) for the sim.
     */
    public readonly internationalMobileSubscriberIdentity!: pulumi.Output<string>;
    /**
     * The timestamp of resource last modification (UTC)
     */
    public readonly lastModifiedAt!: pulumi.Output<string | undefined>;
    /**
     * The identity that last modified the resource.
     */
    public readonly lastModifiedBy!: pulumi.Output<string | undefined>;
    /**
     * The type of identity that last modified the resource.
     */
    public readonly lastModifiedByType!: pulumi.Output<string | undefined>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Mobile network that this sim belongs to
     */
    public readonly mobileNetwork!: pulumi.Output<outputs.mobilenetwork.v20220301preview.MobileNetworkResourceIdResponse | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the sim resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The simPolicy used by this sim.
     */
    public readonly simPolicy!: pulumi.Output<outputs.mobilenetwork.v20220301preview.SimPolicyResourceIdResponse | undefined>;
    /**
     * The state of the sim resource.
     */
    public /*out*/ readonly simState!: pulumi.Output<string>;
    /**
     * A list of static IP addresses assigned to this sim. Each address is assigned at a defined network scope, made up of {attached data network, slice}.
     */
    public readonly staticIpConfiguration!: pulumi.Output<outputs.mobilenetwork.v20220301preview.SimStaticIpPropertiesResponse[] | undefined>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.mobilenetwork.v20220301preview.SystemDataResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Sim resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SimArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.internationalMobileSubscriberIdentity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'internationalMobileSubscriberIdentity'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["authenticationKey"] = args ? args.authenticationKey : undefined;
            resourceInputs["createdAt"] = args ? args.createdAt : undefined;
            resourceInputs["createdBy"] = args ? args.createdBy : undefined;
            resourceInputs["createdByType"] = args ? args.createdByType : undefined;
            resourceInputs["deviceType"] = args ? args.deviceType : undefined;
            resourceInputs["integratedCircuitCardIdentifier"] = args ? args.integratedCircuitCardIdentifier : undefined;
            resourceInputs["internationalMobileSubscriberIdentity"] = args ? args.internationalMobileSubscriberIdentity : undefined;
            resourceInputs["lastModifiedAt"] = args ? args.lastModifiedAt : undefined;
            resourceInputs["lastModifiedBy"] = args ? args.lastModifiedBy : undefined;
            resourceInputs["lastModifiedByType"] = args ? args.lastModifiedByType : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["mobileNetwork"] = args ? args.mobileNetwork : undefined;
            resourceInputs["operatorKeyCode"] = args ? args.operatorKeyCode : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["simName"] = args ? args.simName : undefined;
            resourceInputs["simPolicy"] = args ? args.simPolicy : undefined;
            resourceInputs["staticIpConfiguration"] = args ? args.staticIpConfiguration : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["simState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["createdBy"] = undefined /*out*/;
            resourceInputs["createdByType"] = undefined /*out*/;
            resourceInputs["deviceType"] = undefined /*out*/;
            resourceInputs["integratedCircuitCardIdentifier"] = undefined /*out*/;
            resourceInputs["internationalMobileSubscriberIdentity"] = undefined /*out*/;
            resourceInputs["lastModifiedAt"] = undefined /*out*/;
            resourceInputs["lastModifiedBy"] = undefined /*out*/;
            resourceInputs["lastModifiedByType"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["mobileNetwork"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["simPolicy"] = undefined /*out*/;
            resourceInputs["simState"] = undefined /*out*/;
            resourceInputs["staticIpConfiguration"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Sim.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Sim resource.
 */
export interface SimArgs {
    /**
     * The ki value for the sim.
     */
    authenticationKey?: pulumi.Input<string>;
    /**
     * The timestamp of resource creation (UTC).
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The identity that created the resource.
     */
    createdBy?: pulumi.Input<string>;
    /**
     * The type of identity that created the resource.
     */
    createdByType?: pulumi.Input<string | enums.mobilenetwork.v20220301preview.CreatedByType>;
    /**
     * An optional free-form text field that can be used to record the device type this sim is associated with, for example 'Video camera'. The Azure portal allows Sims to be grouped and filtered based on this value.
     */
    deviceType?: pulumi.Input<string>;
    /**
     * The Integrated Circuit Card ID (ICC Id) for the sim.
     */
    integratedCircuitCardIdentifier?: pulumi.Input<string>;
    /**
     * The International Mobile Subscriber Identity (IMSI) for the sim.
     */
    internationalMobileSubscriberIdentity: pulumi.Input<string>;
    /**
     * The timestamp of resource last modification (UTC)
     */
    lastModifiedAt?: pulumi.Input<string>;
    /**
     * The identity that last modified the resource.
     */
    lastModifiedBy?: pulumi.Input<string>;
    /**
     * The type of identity that last modified the resource.
     */
    lastModifiedByType?: pulumi.Input<string | enums.mobilenetwork.v20220301preview.CreatedByType>;
    /**
     * The geo-location where the resource lives
     */
    location?: pulumi.Input<string>;
    /**
     * Mobile network that this sim belongs to
     */
    mobileNetwork?: pulumi.Input<inputs.mobilenetwork.v20220301preview.MobileNetworkResourceIdArgs>;
    /**
     * The Opc value for the sim.
     */
    operatorKeyCode?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the SIM.
     */
    simName?: pulumi.Input<string>;
    /**
     * The simPolicy used by this sim.
     */
    simPolicy?: pulumi.Input<inputs.mobilenetwork.v20220301preview.SimPolicyResourceIdArgs>;
    /**
     * A list of static IP addresses assigned to this sim. Each address is assigned at a defined network scope, made up of {attached data network, slice}.
     */
    staticIpConfiguration?: pulumi.Input<pulumi.Input<inputs.mobilenetwork.v20220301preview.SimStaticIpPropertiesArgs>[]>;
    /**
     * Resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
