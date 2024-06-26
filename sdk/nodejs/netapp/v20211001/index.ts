// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetVolumeArgs, GetVolumeResult, GetVolumeOutputArgs } from "./getVolume";
export const getVolume: typeof import("./getVolume").getVolume = null as any;
export const getVolumeOutput: typeof import("./getVolume").getVolumeOutput = null as any;
utilities.lazyLoad(exports, ["getVolume","getVolumeOutput"], () => require("./getVolume"));

export { GetVolumeGroupArgs, GetVolumeGroupResult, GetVolumeGroupOutputArgs } from "./getVolumeGroup";
export const getVolumeGroup: typeof import("./getVolumeGroup").getVolumeGroup = null as any;
export const getVolumeGroupOutput: typeof import("./getVolumeGroup").getVolumeGroupOutput = null as any;
utilities.lazyLoad(exports, ["getVolumeGroup","getVolumeGroupOutput"], () => require("./getVolumeGroup"));

export { VolumeArgs } from "./volume";
export type Volume = import("./volume").Volume;
export const Volume: typeof import("./volume").Volume = null as any;
utilities.lazyLoad(exports, ["Volume"], () => require("./volume"));

export { VolumeGroupArgs } from "./volumeGroup";
export type VolumeGroup = import("./volumeGroup").VolumeGroup;
export const VolumeGroup: typeof import("./volumeGroup").VolumeGroup = null as any;
utilities.lazyLoad(exports, ["VolumeGroup"], () => require("./volumeGroup"));


// Export enums:
export * from "../../types/enums/netapp/v20211001";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:netapp/v20211001:Volume":
                return new Volume(name, <any>undefined, { urn })
            case "azure-native:netapp/v20211001:VolumeGroup":
                return new VolumeGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "netapp/v20211001", _module)
