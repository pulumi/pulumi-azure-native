// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Definition of hybrid runbook worker.
 */
export class HybridRunbookWorker extends pulumi.CustomResource {
    /**
     * Get an existing HybridRunbookWorker resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): HybridRunbookWorker {
        return new HybridRunbookWorker(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:automation/v20241023:HybridRunbookWorker';

    /**
     * Returns true if the given object is an instance of HybridRunbookWorker.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HybridRunbookWorker {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HybridRunbookWorker.__pulumiType;
    }

    /**
     * Gets or sets the assigned machine IP address.
     */
    public /*out*/ readonly ip!: pulumi.Output<string | undefined>;
    /**
     * Last Heartbeat from the Worker
     */
    public /*out*/ readonly lastSeenDateTime!: pulumi.Output<string | undefined>;
    /**
     * The geo-location where the resource lives
     */
    public /*out*/ readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Gets or sets the registration time of the worker machine.
     */
    public /*out*/ readonly registeredDateTime!: pulumi.Output<string | undefined>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.automation.v20241023.SystemDataResponse>;
    /**
     * Resource tags.
     */
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Azure Resource Manager Id for a virtual machine.
     */
    public readonly vmResourceId!: pulumi.Output<string | undefined>;
    /**
     * Name of the HybridWorker.
     */
    public /*out*/ readonly workerName!: pulumi.Output<string | undefined>;
    /**
     * Type of the HybridWorker.
     */
    public /*out*/ readonly workerType!: pulumi.Output<string | undefined>;

    /**
     * Create a HybridRunbookWorker resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HybridRunbookWorkerArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.automationAccountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if ((!args || args.hybridRunbookWorkerGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hybridRunbookWorkerGroupName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            resourceInputs["hybridRunbookWorkerGroupName"] = args ? args.hybridRunbookWorkerGroupName : undefined;
            resourceInputs["hybridRunbookWorkerId"] = args ? args.hybridRunbookWorkerId : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["vmResourceId"] = args ? args.vmResourceId : undefined;
            resourceInputs["ip"] = undefined /*out*/;
            resourceInputs["lastSeenDateTime"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["registeredDateTime"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["workerName"] = undefined /*out*/;
            resourceInputs["workerType"] = undefined /*out*/;
        } else {
            resourceInputs["ip"] = undefined /*out*/;
            resourceInputs["lastSeenDateTime"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["registeredDateTime"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["vmResourceId"] = undefined /*out*/;
            resourceInputs["workerName"] = undefined /*out*/;
            resourceInputs["workerType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:automation/v20210622:HybridRunbookWorker" }, { type: "azure-native:automation/v20220808:HybridRunbookWorker" }, { type: "azure-native:automation/v20230515preview:HybridRunbookWorker" }, { type: "azure-native:automation/v20231101:HybridRunbookWorker" }, { type: "azure-native:automation:HybridRunbookWorker" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(HybridRunbookWorker.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a HybridRunbookWorker resource.
 */
export interface HybridRunbookWorkerArgs {
    /**
     * The name of the automation account.
     */
    automationAccountName: pulumi.Input<string>;
    /**
     * The hybrid runbook worker group name
     */
    hybridRunbookWorkerGroupName: pulumi.Input<string>;
    /**
     * The hybrid runbook worker id
     */
    hybridRunbookWorkerId?: pulumi.Input<string>;
    /**
     * Name of an Azure Resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Azure Resource Manager Id for a virtual machine.
     */
    vmResourceId?: pulumi.Input<string>;
}
