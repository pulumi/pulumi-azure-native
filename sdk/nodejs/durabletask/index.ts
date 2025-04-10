// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetSchedulerArgs, GetSchedulerResult, GetSchedulerOutputArgs } from "./getScheduler";
export const getScheduler: typeof import("./getScheduler").getScheduler = null as any;
export const getSchedulerOutput: typeof import("./getScheduler").getSchedulerOutput = null as any;
utilities.lazyLoad(exports, ["getScheduler","getSchedulerOutput"], () => require("./getScheduler"));

export { GetTaskHubArgs, GetTaskHubResult, GetTaskHubOutputArgs } from "./getTaskHub";
export const getTaskHub: typeof import("./getTaskHub").getTaskHub = null as any;
export const getTaskHubOutput: typeof import("./getTaskHub").getTaskHubOutput = null as any;
utilities.lazyLoad(exports, ["getTaskHub","getTaskHubOutput"], () => require("./getTaskHub"));

export { SchedulerArgs } from "./scheduler";
export type Scheduler = import("./scheduler").Scheduler;
export const Scheduler: typeof import("./scheduler").Scheduler = null as any;
utilities.lazyLoad(exports, ["Scheduler"], () => require("./scheduler"));

export { TaskHubArgs } from "./taskHub";
export type TaskHub = import("./taskHub").TaskHub;
export const TaskHub: typeof import("./taskHub").TaskHub = null as any;
utilities.lazyLoad(exports, ["TaskHub"], () => require("./taskHub"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:durabletask:Scheduler":
                return new Scheduler(name, <any>undefined, { urn })
            case "azure-native:durabletask:TaskHub":
                return new TaskHub(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "durabletask", _module)
