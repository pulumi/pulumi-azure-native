// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetLandingZoneAccountOperationArgs, GetLandingZoneAccountOperationResult, GetLandingZoneAccountOperationOutputArgs } from "./getLandingZoneAccountOperation";
export const getLandingZoneAccountOperation: typeof import("./getLandingZoneAccountOperation").getLandingZoneAccountOperation = null as any;
export const getLandingZoneAccountOperationOutput: typeof import("./getLandingZoneAccountOperation").getLandingZoneAccountOperationOutput = null as any;
utilities.lazyLoad(exports, ["getLandingZoneAccountOperation","getLandingZoneAccountOperationOutput"], () => require("./getLandingZoneAccountOperation"));

export { GetLandingZoneConfigurationOperationArgs, GetLandingZoneConfigurationOperationResult, GetLandingZoneConfigurationOperationOutputArgs } from "./getLandingZoneConfigurationOperation";
export const getLandingZoneConfigurationOperation: typeof import("./getLandingZoneConfigurationOperation").getLandingZoneConfigurationOperation = null as any;
export const getLandingZoneConfigurationOperationOutput: typeof import("./getLandingZoneConfigurationOperation").getLandingZoneConfigurationOperationOutput = null as any;
utilities.lazyLoad(exports, ["getLandingZoneConfigurationOperation","getLandingZoneConfigurationOperationOutput"], () => require("./getLandingZoneConfigurationOperation"));

export { GetLandingZoneRegistrationOperationArgs, GetLandingZoneRegistrationOperationResult, GetLandingZoneRegistrationOperationOutputArgs } from "./getLandingZoneRegistrationOperation";
export const getLandingZoneRegistrationOperation: typeof import("./getLandingZoneRegistrationOperation").getLandingZoneRegistrationOperation = null as any;
export const getLandingZoneRegistrationOperationOutput: typeof import("./getLandingZoneRegistrationOperation").getLandingZoneRegistrationOperationOutput = null as any;
utilities.lazyLoad(exports, ["getLandingZoneRegistrationOperation","getLandingZoneRegistrationOperationOutput"], () => require("./getLandingZoneRegistrationOperation"));

export { LandingZoneAccountOperationArgs } from "./landingZoneAccountOperation";
export type LandingZoneAccountOperation = import("./landingZoneAccountOperation").LandingZoneAccountOperation;
export const LandingZoneAccountOperation: typeof import("./landingZoneAccountOperation").LandingZoneAccountOperation = null as any;
utilities.lazyLoad(exports, ["LandingZoneAccountOperation"], () => require("./landingZoneAccountOperation"));

export { LandingZoneConfigurationOperationArgs } from "./landingZoneConfigurationOperation";
export type LandingZoneConfigurationOperation = import("./landingZoneConfigurationOperation").LandingZoneConfigurationOperation;
export const LandingZoneConfigurationOperation: typeof import("./landingZoneConfigurationOperation").LandingZoneConfigurationOperation = null as any;
utilities.lazyLoad(exports, ["LandingZoneConfigurationOperation"], () => require("./landingZoneConfigurationOperation"));

export { LandingZoneRegistrationOperationArgs } from "./landingZoneRegistrationOperation";
export type LandingZoneRegistrationOperation = import("./landingZoneRegistrationOperation").LandingZoneRegistrationOperation;
export const LandingZoneRegistrationOperation: typeof import("./landingZoneRegistrationOperation").LandingZoneRegistrationOperation = null as any;
utilities.lazyLoad(exports, ["LandingZoneRegistrationOperation"], () => require("./landingZoneRegistrationOperation"));


// Export enums:
export * from "../types/enums/sovereign";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:sovereign:LandingZoneAccountOperation":
                return new LandingZoneAccountOperation(name, <any>undefined, { urn })
            case "azure-native:sovereign:LandingZoneConfigurationOperation":
                return new LandingZoneConfigurationOperation(name, <any>undefined, { urn })
            case "azure-native:sovereign:LandingZoneRegistrationOperation":
                return new LandingZoneRegistrationOperation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "sovereign", _module)
