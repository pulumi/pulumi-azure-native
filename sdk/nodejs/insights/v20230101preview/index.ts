// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { ActivityLogAlertArgs } from "./activityLogAlert";
export type ActivityLogAlert = import("./activityLogAlert").ActivityLogAlert;
export const ActivityLogAlert: typeof import("./activityLogAlert").ActivityLogAlert = null as any;
utilities.lazyLoad(exports, ["ActivityLogAlert"], () => require("./activityLogAlert"));

export { GetActivityLogAlertArgs, GetActivityLogAlertResult, GetActivityLogAlertOutputArgs } from "./getActivityLogAlert";
export const getActivityLogAlert: typeof import("./getActivityLogAlert").getActivityLogAlert = null as any;
export const getActivityLogAlertOutput: typeof import("./getActivityLogAlert").getActivityLogAlertOutput = null as any;
utilities.lazyLoad(exports, ["getActivityLogAlert","getActivityLogAlertOutput"], () => require("./getActivityLogAlert"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:insights/v20230101preview:ActivityLogAlert":
                return new ActivityLogAlert(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "insights/v20230101preview", _module)
