// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Packet capture session resource.
 *
 * Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2023-06-01.
 *
 * Other available API versions: 2023-06-01, 2023-09-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mobilenetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class PacketCapture extends pulumi.CustomResource {
    /**
     * Get an existing PacketCapture resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PacketCapture {
        return new PacketCapture(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:mobilenetwork:PacketCapture';

    /**
     * Returns true if the given object is an instance of PacketCapture.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PacketCapture {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PacketCapture.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Number of bytes captured per packet, the remaining bytes are truncated. The default "0" means the entire packet is captured.
     */
    public readonly bytesToCapturePerPacket!: pulumi.Output<number | undefined>;
    /**
     * The start time of the packet capture session.
     */
    public /*out*/ readonly captureStartTime!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * List of network interfaces to capture on.
     */
    public readonly networkInterfaces!: pulumi.Output<string[] | undefined>;
    /**
     * The list of output files of a packet capture session.
     */
    public /*out*/ readonly outputFiles!: pulumi.Output<string[]>;
    /**
     * The provisioning state of the packet capture session resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The reason the current packet capture session state.
     */
    public /*out*/ readonly reason!: pulumi.Output<string>;
    /**
     * The status of the packet capture session.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.mobilenetwork.SystemDataResponse>;
    /**
     * Maximum duration of the capture session in seconds.
     */
    public readonly timeLimitInSeconds!: pulumi.Output<number | undefined>;
    /**
     * Maximum size of the capture output.
     */
    public readonly totalBytesPerSession!: pulumi.Output<number | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a PacketCapture resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PacketCaptureArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.packetCoreControlPlaneName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'packetCoreControlPlaneName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["bytesToCapturePerPacket"] = (args ? args.bytesToCapturePerPacket : undefined) ?? 0;
            resourceInputs["networkInterfaces"] = args ? args.networkInterfaces : undefined;
            resourceInputs["packetCaptureName"] = args ? args.packetCaptureName : undefined;
            resourceInputs["packetCoreControlPlaneName"] = args ? args.packetCoreControlPlaneName : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["timeLimitInSeconds"] = (args ? args.timeLimitInSeconds : undefined) ?? 18000;
            resourceInputs["totalBytesPerSession"] = (args ? args.totalBytesPerSession : undefined) ?? 67108864;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["captureStartTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["outputFiles"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["reason"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["bytesToCapturePerPacket"] = undefined /*out*/;
            resourceInputs["captureStartTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["networkInterfaces"] = undefined /*out*/;
            resourceInputs["outputFiles"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["reason"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["timeLimitInSeconds"] = undefined /*out*/;
            resourceInputs["totalBytesPerSession"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:mobilenetwork/v20230601:PacketCapture" }, { type: "azure-native:mobilenetwork/v20230901:PacketCapture" }, { type: "azure-native:mobilenetwork/v20240201:PacketCapture" }, { type: "azure-native:mobilenetwork/v20240401:PacketCapture" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(PacketCapture.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PacketCapture resource.
 */
export interface PacketCaptureArgs {
    /**
     * Number of bytes captured per packet, the remaining bytes are truncated. The default "0" means the entire packet is captured.
     */
    bytesToCapturePerPacket?: pulumi.Input<number>;
    /**
     * List of network interfaces to capture on.
     */
    networkInterfaces?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the packet capture session.
     */
    packetCaptureName?: pulumi.Input<string>;
    /**
     * The name of the packet core control plane.
     */
    packetCoreControlPlaneName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Maximum duration of the capture session in seconds.
     */
    timeLimitInSeconds?: pulumi.Input<number>;
    /**
     * Maximum size of the capture output.
     */
    totalBytesPerSession?: pulumi.Input<number>;
}
