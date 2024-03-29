// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { APICollectionByAzureApiManagementServiceArgs } from "./apicollectionByAzureApiManagementService";
export type APICollectionByAzureApiManagementService = import("./apicollectionByAzureApiManagementService").APICollectionByAzureApiManagementService;
export const APICollectionByAzureApiManagementService: typeof import("./apicollectionByAzureApiManagementService").APICollectionByAzureApiManagementService = null as any;
utilities.lazyLoad(exports, ["APICollectionByAzureApiManagementService"], () => require("./apicollectionByAzureApiManagementService"));

export { GetAPICollectionByAzureApiManagementServiceArgs, GetAPICollectionByAzureApiManagementServiceResult, GetAPICollectionByAzureApiManagementServiceOutputArgs } from "./getAPICollectionByAzureApiManagementService";
export const getAPICollectionByAzureApiManagementService: typeof import("./getAPICollectionByAzureApiManagementService").getAPICollectionByAzureApiManagementService = null as any;
export const getAPICollectionByAzureApiManagementServiceOutput: typeof import("./getAPICollectionByAzureApiManagementService").getAPICollectionByAzureApiManagementServiceOutput = null as any;
utilities.lazyLoad(exports, ["getAPICollectionByAzureApiManagementService","getAPICollectionByAzureApiManagementServiceOutput"], () => require("./getAPICollectionByAzureApiManagementService"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:security/v20231115:APICollectionByAzureApiManagementService":
                return new APICollectionByAzureApiManagementService(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "security/v20231115", _module)
