// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetRemoteRenderingAccountArgs, GetRemoteRenderingAccountResult, GetRemoteRenderingAccountOutputArgs } from "./getRemoteRenderingAccount";
export const getRemoteRenderingAccount: typeof import("./getRemoteRenderingAccount").getRemoteRenderingAccount = null as any;
export const getRemoteRenderingAccountOutput: typeof import("./getRemoteRenderingAccount").getRemoteRenderingAccountOutput = null as any;
utilities.lazyLoad(exports, ["getRemoteRenderingAccount","getRemoteRenderingAccountOutput"], () => require("./getRemoteRenderingAccount"));

export { ListRemoteRenderingAccountKeysArgs, ListRemoteRenderingAccountKeysResult, ListRemoteRenderingAccountKeysOutputArgs } from "./listRemoteRenderingAccountKeys";
export const listRemoteRenderingAccountKeys: typeof import("./listRemoteRenderingAccountKeys").listRemoteRenderingAccountKeys = null as any;
export const listRemoteRenderingAccountKeysOutput: typeof import("./listRemoteRenderingAccountKeys").listRemoteRenderingAccountKeysOutput = null as any;
utilities.lazyLoad(exports, ["listRemoteRenderingAccountKeys","listRemoteRenderingAccountKeysOutput"], () => require("./listRemoteRenderingAccountKeys"));

export { RemoteRenderingAccountArgs } from "./remoteRenderingAccount";
export type RemoteRenderingAccount = import("./remoteRenderingAccount").RemoteRenderingAccount;
export const RemoteRenderingAccount: typeof import("./remoteRenderingAccount").RemoteRenderingAccount = null as any;
utilities.lazyLoad(exports, ["RemoteRenderingAccount"], () => require("./remoteRenderingAccount"));


// Export enums:
export * from "../../types/enums/mixedreality/v20250101";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:mixedreality/v20250101:RemoteRenderingAccount":
                return new RemoteRenderingAccount(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "mixedreality/v20250101", _module)
