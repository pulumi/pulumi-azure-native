// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetPrivateEndpointConnectionArgs, GetPrivateEndpointConnectionResult, GetPrivateEndpointConnectionOutputArgs } from "./getPrivateEndpointConnection";
export const getPrivateEndpointConnection: typeof import("./getPrivateEndpointConnection").getPrivateEndpointConnection = null as any;
export const getPrivateEndpointConnectionOutput: typeof import("./getPrivateEndpointConnection").getPrivateEndpointConnectionOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateEndpointConnection","getPrivateEndpointConnectionOutput"], () => require("./getPrivateEndpointConnection"));

export { GetPrivateLinkScopeArgs, GetPrivateLinkScopeResult, GetPrivateLinkScopeOutputArgs } from "./getPrivateLinkScope";
export const getPrivateLinkScope: typeof import("./getPrivateLinkScope").getPrivateLinkScope = null as any;
export const getPrivateLinkScopeOutput: typeof import("./getPrivateLinkScope").getPrivateLinkScopeOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateLinkScope","getPrivateLinkScopeOutput"], () => require("./getPrivateLinkScope"));

export { GetPrivateLinkScopedResourceArgs, GetPrivateLinkScopedResourceResult, GetPrivateLinkScopedResourceOutputArgs } from "./getPrivateLinkScopedResource";
export const getPrivateLinkScopedResource: typeof import("./getPrivateLinkScopedResource").getPrivateLinkScopedResource = null as any;
export const getPrivateLinkScopedResourceOutput: typeof import("./getPrivateLinkScopedResource").getPrivateLinkScopedResourceOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateLinkScopedResource","getPrivateLinkScopedResourceOutput"], () => require("./getPrivateLinkScopedResource"));

export { PrivateEndpointConnectionArgs } from "./privateEndpointConnection";
export type PrivateEndpointConnection = import("./privateEndpointConnection").PrivateEndpointConnection;
export const PrivateEndpointConnection: typeof import("./privateEndpointConnection").PrivateEndpointConnection = null as any;
utilities.lazyLoad(exports, ["PrivateEndpointConnection"], () => require("./privateEndpointConnection"));

export { PrivateLinkScopeArgs } from "./privateLinkScope";
export type PrivateLinkScope = import("./privateLinkScope").PrivateLinkScope;
export const PrivateLinkScope: typeof import("./privateLinkScope").PrivateLinkScope = null as any;
utilities.lazyLoad(exports, ["PrivateLinkScope"], () => require("./privateLinkScope"));

export { PrivateLinkScopedResourceArgs } from "./privateLinkScopedResource";
export type PrivateLinkScopedResource = import("./privateLinkScopedResource").PrivateLinkScopedResource;
export const PrivateLinkScopedResource: typeof import("./privateLinkScopedResource").PrivateLinkScopedResource = null as any;
utilities.lazyLoad(exports, ["PrivateLinkScopedResource"], () => require("./privateLinkScopedResource"));


// Export enums:
export * from "../../types/enums/insights/v20210901";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:insights/v20210901:PrivateEndpointConnection":
                return new PrivateEndpointConnection(name, <any>undefined, { urn })
            case "azure-native:insights/v20210901:PrivateLinkScope":
                return new PrivateLinkScope(name, <any>undefined, { urn })
            case "azure-native:insights/v20210901:PrivateLinkScopedResource":
                return new PrivateLinkScopedResource(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "insights/v20210901", _module)
