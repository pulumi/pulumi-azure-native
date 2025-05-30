// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Uses Azure REST API version 2025-02-01. In version 2.x of the Azure Native provider, it used API version 2023-10-01-preview.
 *
 * Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class BareMetalMachine extends pulumi.CustomResource {
    /**
     * Get an existing BareMetalMachine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BareMetalMachine {
        return new BareMetalMachine(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:networkcloud:BareMetalMachine';

    /**
     * Returns true if the given object is an instance of BareMetalMachine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BareMetalMachine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BareMetalMachine.__pulumiType;
    }

    /**
     * The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
     */
    public /*out*/ readonly associatedResourceIds!: pulumi.Output<string[]>;
    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The connection string for the baseboard management controller including IP address and protocol.
     */
    public readonly bmcConnectionString!: pulumi.Output<string>;
    /**
     * The credentials of the baseboard management controller on this bare metal machine.
     */
    public readonly bmcCredentials!: pulumi.Output<outputs.networkcloud.AdministrativeCredentialsResponse>;
    /**
     * The MAC address of the BMC device.
     */
    public readonly bmcMacAddress!: pulumi.Output<string>;
    /**
     * The MAC address of a NIC connected to the PXE network.
     */
    public readonly bootMacAddress!: pulumi.Output<string>;
    /**
     * The resource ID of the cluster this bare metal machine is associated with.
     */
    public /*out*/ readonly clusterId!: pulumi.Output<string>;
    /**
     * The cordon status of the bare metal machine.
     */
    public /*out*/ readonly cordonStatus!: pulumi.Output<string>;
    /**
     * The more detailed status of the bare metal machine.
     */
    public /*out*/ readonly detailedStatus!: pulumi.Output<string>;
    /**
     * The descriptive message about the current detailed status.
     */
    public /*out*/ readonly detailedStatusMessage!: pulumi.Output<string>;
    /**
     * Resource ETag.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The extended location of the cluster associated with the resource.
     */
    public readonly extendedLocation!: pulumi.Output<outputs.networkcloud.ExtendedLocationResponse>;
    /**
     * The hardware inventory, including information acquired from the model/sku information and from the ironic inspector.
     */
    public /*out*/ readonly hardwareInventory!: pulumi.Output<outputs.networkcloud.HardwareInventoryResponse>;
    /**
     * The details of the latest hardware validation performed for this bare metal machine.
     */
    public /*out*/ readonly hardwareValidationStatus!: pulumi.Output<outputs.networkcloud.HardwareValidationStatusResponse>;
    /**
     * Field Deprecated. These fields will be empty/omitted. The list of the resource IDs for the HybridAksClusters that have nodes hosted on this bare metal machine.
     */
    public /*out*/ readonly hybridAksClustersAssociatedIds!: pulumi.Output<string[]>;
    /**
     * The name of this machine represented by the host object in the Cluster's Kubernetes control plane.
     */
    public /*out*/ readonly kubernetesNodeName!: pulumi.Output<string>;
    /**
     * The version of Kubernetes running on this machine.
     */
    public /*out*/ readonly kubernetesVersion!: pulumi.Output<string>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The cluster version that has been applied to this machine during deployment or a version update.
     */
    public readonly machineClusterVersion!: pulumi.Output<string | undefined>;
    /**
     * The custom details provided by the customer.
     */
    public readonly machineDetails!: pulumi.Output<string>;
    /**
     * The OS-level hostname assigned to this machine.
     */
    public readonly machineName!: pulumi.Output<string>;
    /**
     * The list of roles that are assigned to the cluster node running on this machine.
     */
    public /*out*/ readonly machineRoles!: pulumi.Output<string[]>;
    /**
     * The unique internal identifier of the bare metal machine SKU.
     */
    public readonly machineSkuId!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The IPv4 address that is assigned to the bare metal machine during the cluster deployment.
     */
    public /*out*/ readonly oamIpv4Address!: pulumi.Output<string>;
    /**
     * The IPv6 address that is assigned to the bare metal machine during the cluster deployment.
     */
    public /*out*/ readonly oamIpv6Address!: pulumi.Output<string>;
    /**
     * The image that is currently provisioned to the OS disk.
     */
    public /*out*/ readonly osImage!: pulumi.Output<string>;
    /**
     * The power state derived from the baseboard management controller.
     */
    public /*out*/ readonly powerState!: pulumi.Output<string>;
    /**
     * The provisioning state of the bare metal machine.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The resource ID of the rack where this bare metal machine resides.
     */
    public readonly rackId!: pulumi.Output<string>;
    /**
     * The rack slot in which this bare metal machine is located, ordered from the bottom up i.e. the lowest slot is 1.
     */
    public readonly rackSlot!: pulumi.Output<number>;
    /**
     * The indicator of whether the bare metal machine is ready to receive workloads.
     */
    public /*out*/ readonly readyState!: pulumi.Output<string>;
    /**
     * The runtime protection status of the bare metal machine.
     */
    public /*out*/ readonly runtimeProtectionStatus!: pulumi.Output<outputs.networkcloud.RuntimeProtectionStatusResponse>;
    /**
     * The list of statuses that represent secret rotation activity.
     */
    public /*out*/ readonly secretRotationStatus!: pulumi.Output<outputs.networkcloud.SecretRotationStatusResponse[]>;
    /**
     * The serial number of the bare metal machine.
     */
    public readonly serialNumber!: pulumi.Output<string>;
    /**
     * The discovered value of the machine's service tag.
     */
    public /*out*/ readonly serviceTag!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.networkcloud.SystemDataResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Field Deprecated. These fields will be empty/omitted. The list of the resource IDs for the VirtualMachines that are hosted on this bare metal machine.
     */
    public /*out*/ readonly virtualMachinesAssociatedIds!: pulumi.Output<string[]>;

    /**
     * Create a BareMetalMachine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BareMetalMachineArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.bmcConnectionString === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bmcConnectionString'");
            }
            if ((!args || args.bmcCredentials === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bmcCredentials'");
            }
            if ((!args || args.bmcMacAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bmcMacAddress'");
            }
            if ((!args || args.bootMacAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bootMacAddress'");
            }
            if ((!args || args.extendedLocation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'extendedLocation'");
            }
            if ((!args || args.machineDetails === undefined) && !opts.urn) {
                throw new Error("Missing required property 'machineDetails'");
            }
            if ((!args || args.machineName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'machineName'");
            }
            if ((!args || args.machineSkuId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'machineSkuId'");
            }
            if ((!args || args.rackId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rackId'");
            }
            if ((!args || args.rackSlot === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rackSlot'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serialNumber === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serialNumber'");
            }
            resourceInputs["bareMetalMachineName"] = args ? args.bareMetalMachineName : undefined;
            resourceInputs["bmcConnectionString"] = args ? args.bmcConnectionString : undefined;
            resourceInputs["bmcCredentials"] = args ? args.bmcCredentials : undefined;
            resourceInputs["bmcMacAddress"] = args ? args.bmcMacAddress : undefined;
            resourceInputs["bootMacAddress"] = args ? args.bootMacAddress : undefined;
            resourceInputs["extendedLocation"] = args ? args.extendedLocation : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["machineClusterVersion"] = args ? args.machineClusterVersion : undefined;
            resourceInputs["machineDetails"] = args ? args.machineDetails : undefined;
            resourceInputs["machineName"] = args ? args.machineName : undefined;
            resourceInputs["machineSkuId"] = args ? args.machineSkuId : undefined;
            resourceInputs["rackId"] = args ? args.rackId : undefined;
            resourceInputs["rackSlot"] = args ? args.rackSlot : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["serialNumber"] = args ? args.serialNumber : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["associatedResourceIds"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["clusterId"] = undefined /*out*/;
            resourceInputs["cordonStatus"] = undefined /*out*/;
            resourceInputs["detailedStatus"] = undefined /*out*/;
            resourceInputs["detailedStatusMessage"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["hardwareInventory"] = undefined /*out*/;
            resourceInputs["hardwareValidationStatus"] = undefined /*out*/;
            resourceInputs["hybridAksClustersAssociatedIds"] = undefined /*out*/;
            resourceInputs["kubernetesNodeName"] = undefined /*out*/;
            resourceInputs["kubernetesVersion"] = undefined /*out*/;
            resourceInputs["machineRoles"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["oamIpv4Address"] = undefined /*out*/;
            resourceInputs["oamIpv6Address"] = undefined /*out*/;
            resourceInputs["osImage"] = undefined /*out*/;
            resourceInputs["powerState"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["readyState"] = undefined /*out*/;
            resourceInputs["runtimeProtectionStatus"] = undefined /*out*/;
            resourceInputs["secretRotationStatus"] = undefined /*out*/;
            resourceInputs["serviceTag"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["virtualMachinesAssociatedIds"] = undefined /*out*/;
        } else {
            resourceInputs["associatedResourceIds"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["bmcConnectionString"] = undefined /*out*/;
            resourceInputs["bmcCredentials"] = undefined /*out*/;
            resourceInputs["bmcMacAddress"] = undefined /*out*/;
            resourceInputs["bootMacAddress"] = undefined /*out*/;
            resourceInputs["clusterId"] = undefined /*out*/;
            resourceInputs["cordonStatus"] = undefined /*out*/;
            resourceInputs["detailedStatus"] = undefined /*out*/;
            resourceInputs["detailedStatusMessage"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["extendedLocation"] = undefined /*out*/;
            resourceInputs["hardwareInventory"] = undefined /*out*/;
            resourceInputs["hardwareValidationStatus"] = undefined /*out*/;
            resourceInputs["hybridAksClustersAssociatedIds"] = undefined /*out*/;
            resourceInputs["kubernetesNodeName"] = undefined /*out*/;
            resourceInputs["kubernetesVersion"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["machineClusterVersion"] = undefined /*out*/;
            resourceInputs["machineDetails"] = undefined /*out*/;
            resourceInputs["machineName"] = undefined /*out*/;
            resourceInputs["machineRoles"] = undefined /*out*/;
            resourceInputs["machineSkuId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["oamIpv4Address"] = undefined /*out*/;
            resourceInputs["oamIpv6Address"] = undefined /*out*/;
            resourceInputs["osImage"] = undefined /*out*/;
            resourceInputs["powerState"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["rackId"] = undefined /*out*/;
            resourceInputs["rackSlot"] = undefined /*out*/;
            resourceInputs["readyState"] = undefined /*out*/;
            resourceInputs["runtimeProtectionStatus"] = undefined /*out*/;
            resourceInputs["secretRotationStatus"] = undefined /*out*/;
            resourceInputs["serialNumber"] = undefined /*out*/;
            resourceInputs["serviceTag"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["virtualMachinesAssociatedIds"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:networkcloud/v20230701:BareMetalMachine" }, { type: "azure-native:networkcloud/v20231001preview:BareMetalMachine" }, { type: "azure-native:networkcloud/v20240601preview:BareMetalMachine" }, { type: "azure-native:networkcloud/v20240701:BareMetalMachine" }, { type: "azure-native:networkcloud/v20241001preview:BareMetalMachine" }, { type: "azure-native:networkcloud/v20250201:BareMetalMachine" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(BareMetalMachine.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a BareMetalMachine resource.
 */
export interface BareMetalMachineArgs {
    /**
     * The name of the bare metal machine.
     */
    bareMetalMachineName?: pulumi.Input<string>;
    /**
     * The connection string for the baseboard management controller including IP address and protocol.
     */
    bmcConnectionString: pulumi.Input<string>;
    /**
     * The credentials of the baseboard management controller on this bare metal machine.
     */
    bmcCredentials: pulumi.Input<inputs.networkcloud.AdministrativeCredentialsArgs>;
    /**
     * The MAC address of the BMC device.
     */
    bmcMacAddress: pulumi.Input<string>;
    /**
     * The MAC address of a NIC connected to the PXE network.
     */
    bootMacAddress: pulumi.Input<string>;
    /**
     * The extended location of the cluster associated with the resource.
     */
    extendedLocation: pulumi.Input<inputs.networkcloud.ExtendedLocationArgs>;
    /**
     * The geo-location where the resource lives
     */
    location?: pulumi.Input<string>;
    /**
     * The cluster version that has been applied to this machine during deployment or a version update.
     */
    machineClusterVersion?: pulumi.Input<string>;
    /**
     * The custom details provided by the customer.
     */
    machineDetails: pulumi.Input<string>;
    /**
     * The OS-level hostname assigned to this machine.
     */
    machineName: pulumi.Input<string>;
    /**
     * The unique internal identifier of the bare metal machine SKU.
     */
    machineSkuId: pulumi.Input<string>;
    /**
     * The resource ID of the rack where this bare metal machine resides.
     */
    rackId: pulumi.Input<string>;
    /**
     * The rack slot in which this bare metal machine is located, ordered from the bottom up i.e. the lowest slot is 1.
     */
    rackSlot: pulumi.Input<number>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The serial number of the bare metal machine.
     */
    serialNumber: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
