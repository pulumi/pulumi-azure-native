// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetMigrateProjectArgs, GetMigrateProjectResult, GetMigrateProjectOutputArgs } from "./getMigrateProject";
export const getMigrateProject: typeof import("./getMigrateProject").getMigrateProject = null as any;
export const getMigrateProjectOutput: typeof import("./getMigrateProject").getMigrateProjectOutput = null as any;
utilities.lazyLoad(exports, ["getMigrateProject","getMigrateProjectOutput"], () => require("./getMigrateProject"));

export { GetSolutionArgs, GetSolutionResult, GetSolutionOutputArgs } from "./getSolution";
export const getSolution: typeof import("./getSolution").getSolution = null as any;
export const getSolutionOutput: typeof import("./getSolution").getSolutionOutput = null as any;
utilities.lazyLoad(exports, ["getSolution","getSolutionOutput"], () => require("./getSolution"));

export { GetSolutionConfigArgs, GetSolutionConfigResult, GetSolutionConfigOutputArgs } from "./getSolutionConfig";
export const getSolutionConfig: typeof import("./getSolutionConfig").getSolutionConfig = null as any;
export const getSolutionConfigOutput: typeof import("./getSolutionConfig").getSolutionConfigOutput = null as any;
utilities.lazyLoad(exports, ["getSolutionConfig","getSolutionConfigOutput"], () => require("./getSolutionConfig"));

export { MigrateProjectArgs } from "./migrateProject";
export type MigrateProject = import("./migrateProject").MigrateProject;
export const MigrateProject: typeof import("./migrateProject").MigrateProject = null as any;
utilities.lazyLoad(exports, ["MigrateProject"], () => require("./migrateProject"));

export { SolutionArgs } from "./solution";
export type Solution = import("./solution").Solution;
export const Solution: typeof import("./solution").Solution = null as any;
utilities.lazyLoad(exports, ["Solution"], () => require("./solution"));


// Export enums:
export * from "../../types/enums/migrate/v20180901preview";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:migrate/v20180901preview:MigrateProject":
                return new MigrateProject(name, <any>undefined, { urn })
            case "azure-native:migrate/v20180901preview:Solution":
                return new Solution(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "migrate/v20180901preview", _module)
