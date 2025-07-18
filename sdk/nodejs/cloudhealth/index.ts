// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AuthenticationSettingArgs } from "./authenticationSetting";
export type AuthenticationSetting = import("./authenticationSetting").AuthenticationSetting;
export const AuthenticationSetting: typeof import("./authenticationSetting").AuthenticationSetting = null as any;
utilities.lazyLoad(exports, ["AuthenticationSetting"], () => require("./authenticationSetting"));

export { DiscoveryRuleArgs } from "./discoveryRule";
export type DiscoveryRule = import("./discoveryRule").DiscoveryRule;
export const DiscoveryRule: typeof import("./discoveryRule").DiscoveryRule = null as any;
utilities.lazyLoad(exports, ["DiscoveryRule"], () => require("./discoveryRule"));

export { EntityArgs } from "./entity";
export type Entity = import("./entity").Entity;
export const Entity: typeof import("./entity").Entity = null as any;
utilities.lazyLoad(exports, ["Entity"], () => require("./entity"));

export { GetAuthenticationSettingArgs, GetAuthenticationSettingResult, GetAuthenticationSettingOutputArgs } from "./getAuthenticationSetting";
export const getAuthenticationSetting: typeof import("./getAuthenticationSetting").getAuthenticationSetting = null as any;
export const getAuthenticationSettingOutput: typeof import("./getAuthenticationSetting").getAuthenticationSettingOutput = null as any;
utilities.lazyLoad(exports, ["getAuthenticationSetting","getAuthenticationSettingOutput"], () => require("./getAuthenticationSetting"));

export { GetDiscoveryRuleArgs, GetDiscoveryRuleResult, GetDiscoveryRuleOutputArgs } from "./getDiscoveryRule";
export const getDiscoveryRule: typeof import("./getDiscoveryRule").getDiscoveryRule = null as any;
export const getDiscoveryRuleOutput: typeof import("./getDiscoveryRule").getDiscoveryRuleOutput = null as any;
utilities.lazyLoad(exports, ["getDiscoveryRule","getDiscoveryRuleOutput"], () => require("./getDiscoveryRule"));

export { GetEntityArgs, GetEntityResult, GetEntityOutputArgs } from "./getEntity";
export const getEntity: typeof import("./getEntity").getEntity = null as any;
export const getEntityOutput: typeof import("./getEntity").getEntityOutput = null as any;
utilities.lazyLoad(exports, ["getEntity","getEntityOutput"], () => require("./getEntity"));

export { GetHealthModelArgs, GetHealthModelResult, GetHealthModelOutputArgs } from "./getHealthModel";
export const getHealthModel: typeof import("./getHealthModel").getHealthModel = null as any;
export const getHealthModelOutput: typeof import("./getHealthModel").getHealthModelOutput = null as any;
utilities.lazyLoad(exports, ["getHealthModel","getHealthModelOutput"], () => require("./getHealthModel"));

export { GetRelationshipArgs, GetRelationshipResult, GetRelationshipOutputArgs } from "./getRelationship";
export const getRelationship: typeof import("./getRelationship").getRelationship = null as any;
export const getRelationshipOutput: typeof import("./getRelationship").getRelationshipOutput = null as any;
utilities.lazyLoad(exports, ["getRelationship","getRelationshipOutput"], () => require("./getRelationship"));

export { GetSignalDefinitionArgs, GetSignalDefinitionResult, GetSignalDefinitionOutputArgs } from "./getSignalDefinition";
export const getSignalDefinition: typeof import("./getSignalDefinition").getSignalDefinition = null as any;
export const getSignalDefinitionOutput: typeof import("./getSignalDefinition").getSignalDefinitionOutput = null as any;
utilities.lazyLoad(exports, ["getSignalDefinition","getSignalDefinitionOutput"], () => require("./getSignalDefinition"));

export { HealthModelArgs } from "./healthModel";
export type HealthModel = import("./healthModel").HealthModel;
export const HealthModel: typeof import("./healthModel").HealthModel = null as any;
utilities.lazyLoad(exports, ["HealthModel"], () => require("./healthModel"));

export { RelationshipArgs } from "./relationship";
export type Relationship = import("./relationship").Relationship;
export const Relationship: typeof import("./relationship").Relationship = null as any;
utilities.lazyLoad(exports, ["Relationship"], () => require("./relationship"));

export { SignalDefinitionArgs } from "./signalDefinition";
export type SignalDefinition = import("./signalDefinition").SignalDefinition;
export const SignalDefinition: typeof import("./signalDefinition").SignalDefinition = null as any;
utilities.lazyLoad(exports, ["SignalDefinition"], () => require("./signalDefinition"));


// Export enums:
export * from "../types/enums/cloudhealth";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:cloudhealth:AuthenticationSetting":
                return new AuthenticationSetting(name, <any>undefined, { urn })
            case "azure-native:cloudhealth:DiscoveryRule":
                return new DiscoveryRule(name, <any>undefined, { urn })
            case "azure-native:cloudhealth:Entity":
                return new Entity(name, <any>undefined, { urn })
            case "azure-native:cloudhealth:HealthModel":
                return new HealthModel(name, <any>undefined, { urn })
            case "azure-native:cloudhealth:Relationship":
                return new Relationship(name, <any>undefined, { urn })
            case "azure-native:cloudhealth:SignalDefinition":
                return new SignalDefinition(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "cloudhealth", _module)
