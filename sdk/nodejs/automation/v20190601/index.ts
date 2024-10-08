// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetSoftwareUpdateConfigurationByNameArgs, GetSoftwareUpdateConfigurationByNameResult, GetSoftwareUpdateConfigurationByNameOutputArgs } from "./getSoftwareUpdateConfigurationByName";
export const getSoftwareUpdateConfigurationByName: typeof import("./getSoftwareUpdateConfigurationByName").getSoftwareUpdateConfigurationByName = null as any;
export const getSoftwareUpdateConfigurationByNameOutput: typeof import("./getSoftwareUpdateConfigurationByName").getSoftwareUpdateConfigurationByNameOutput = null as any;
utilities.lazyLoad(exports, ["getSoftwareUpdateConfigurationByName","getSoftwareUpdateConfigurationByNameOutput"], () => require("./getSoftwareUpdateConfigurationByName"));

export { SoftwareUpdateConfigurationByNameArgs } from "./softwareUpdateConfigurationByName";
export type SoftwareUpdateConfigurationByName = import("./softwareUpdateConfigurationByName").SoftwareUpdateConfigurationByName;
export const SoftwareUpdateConfigurationByName: typeof import("./softwareUpdateConfigurationByName").SoftwareUpdateConfigurationByName = null as any;
utilities.lazyLoad(exports, ["SoftwareUpdateConfigurationByName"], () => require("./softwareUpdateConfigurationByName"));


// Export enums:
export * from "../../types/enums/automation/v20190601";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:automation/v20190601:SoftwareUpdateConfigurationByName":
                return new SoftwareUpdateConfigurationByName(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "automation/v20190601", _module)
