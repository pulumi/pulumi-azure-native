// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { DiscoverySourceArgs } from "./discoverySource";
export type DiscoverySource = import("./discoverySource").DiscoverySource;
export const DiscoverySource: typeof import("./discoverySource").DiscoverySource = null as any;
utilities.lazyLoad(exports, ["DiscoverySource"], () => require("./discoverySource"));

export { GetDiscoverySourceArgs, GetDiscoverySourceResult, GetDiscoverySourceOutputArgs } from "./getDiscoverySource";
export const getDiscoverySource: typeof import("./getDiscoverySource").getDiscoverySource = null as any;
export const getDiscoverySourceOutput: typeof import("./getDiscoverySource").getDiscoverySourceOutput = null as any;
utilities.lazyLoad(exports, ["getDiscoverySource","getDiscoverySourceOutput"], () => require("./getDiscoverySource"));

export { GetMapArgs, GetMapResult, GetMapOutputArgs } from "./getMap";
export const getMap: typeof import("./getMap").getMap = null as any;
export const getMapOutput: typeof import("./getMap").getMapOutput = null as any;
utilities.lazyLoad(exports, ["getMap","getMapOutput"], () => require("./getMap"));

export { MapArgs } from "./map";
export type Map = import("./map").Map;
export const Map: typeof import("./map").Map = null as any;
utilities.lazyLoad(exports, ["Map"], () => require("./map"));


// Export enums:
export * from "../types/enums/dependencymap";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:dependencymap:DiscoverySource":
                return new DiscoverySource(name, <any>undefined, { urn })
            case "azure-native:dependencymap:Map":
                return new Map(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "dependencymap", _module)
