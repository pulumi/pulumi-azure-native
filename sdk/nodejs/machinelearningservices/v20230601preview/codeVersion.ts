// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Azure Resource Manager resource envelope.
 */
export class CodeVersion extends pulumi.CustomResource {
    /**
     * Get an existing CodeVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CodeVersion {
        return new CodeVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:machinelearningservices/v20230601preview:CodeVersion';

    /**
     * Returns true if the given object is an instance of CodeVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CodeVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CodeVersion.__pulumiType;
    }

    /**
     * [Required] Additional attributes of the entity.
     */
    public readonly codeVersionProperties!: pulumi.Output<outputs.machinelearningservices.v20230601preview.CodeVersionResponse>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.machinelearningservices.v20230601preview.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a CodeVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CodeVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.codeVersionProperties === undefined) && !opts.urn) {
                throw new Error("Missing required property 'codeVersionProperties'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.workspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceName'");
            }
            resourceInputs["codeVersionProperties"] = args ? (args.codeVersionProperties ? pulumi.output(args.codeVersionProperties).apply(inputs.machinelearningservices.v20230601preview.codeVersionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["workspaceName"] = args ? args.workspaceName : undefined;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["codeVersionProperties"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:machinelearningservices/v20210301preview:CodeVersion" }, { type: "azure-native:machinelearningservices/v20220201preview:CodeVersion" }, { type: "azure-native:machinelearningservices/v20220501:CodeVersion" }, { type: "azure-native:machinelearningservices/v20220601preview:CodeVersion" }, { type: "azure-native:machinelearningservices/v20221001:CodeVersion" }, { type: "azure-native:machinelearningservices/v20221001preview:CodeVersion" }, { type: "azure-native:machinelearningservices/v20221201preview:CodeVersion" }, { type: "azure-native:machinelearningservices/v20230201preview:CodeVersion" }, { type: "azure-native:machinelearningservices/v20230401:CodeVersion" }, { type: "azure-native:machinelearningservices/v20230401preview:CodeVersion" }, { type: "azure-native:machinelearningservices/v20230801preview:CodeVersion" }, { type: "azure-native:machinelearningservices/v20231001:CodeVersion" }, { type: "azure-native:machinelearningservices/v20240101preview:CodeVersion" }, { type: "azure-native:machinelearningservices/v20240401:CodeVersion" }, { type: "azure-native:machinelearningservices/v20240401preview:CodeVersion" }, { type: "azure-native:machinelearningservices/v20240701preview:CodeVersion" }, { type: "azure-native:machinelearningservices/v20241001:CodeVersion" }, { type: "azure-native:machinelearningservices/v20241001preview:CodeVersion" }, { type: "azure-native:machinelearningservices/v20250101preview:CodeVersion" }, { type: "azure-native:machinelearningservices:CodeVersion" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(CodeVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CodeVersion resource.
 */
export interface CodeVersionArgs {
    /**
     * [Required] Additional attributes of the entity.
     */
    codeVersionProperties: pulumi.Input<inputs.machinelearningservices.v20230601preview.CodeVersionArgs>;
    /**
     * Container name. This is case-sensitive.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Version identifier. This is case-sensitive.
     */
    version?: pulumi.Input<string>;
    /**
     * Name of Azure Machine Learning workspace.
     */
    workspaceName: pulumi.Input<string>;
}
