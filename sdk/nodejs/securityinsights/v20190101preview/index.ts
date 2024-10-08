// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { AutomationRuleArgs } from "./automationRule";
export type AutomationRule = import("./automationRule").AutomationRule;
export const AutomationRule: typeof import("./automationRule").AutomationRule = null as any;
utilities.lazyLoad(exports, ["AutomationRule"], () => require("./automationRule"));

export { BookmarkArgs } from "./bookmark";
export type Bookmark = import("./bookmark").Bookmark;
export const Bookmark: typeof import("./bookmark").Bookmark = null as any;
utilities.lazyLoad(exports, ["Bookmark"], () => require("./bookmark"));

export { BookmarkRelationArgs } from "./bookmarkRelation";
export type BookmarkRelation = import("./bookmarkRelation").BookmarkRelation;
export const BookmarkRelation: typeof import("./bookmarkRelation").BookmarkRelation = null as any;
utilities.lazyLoad(exports, ["BookmarkRelation"], () => require("./bookmarkRelation"));

export { GetAutomationRuleArgs, GetAutomationRuleResult, GetAutomationRuleOutputArgs } from "./getAutomationRule";
export const getAutomationRule: typeof import("./getAutomationRule").getAutomationRule = null as any;
export const getAutomationRuleOutput: typeof import("./getAutomationRule").getAutomationRuleOutput = null as any;
utilities.lazyLoad(exports, ["getAutomationRule","getAutomationRuleOutput"], () => require("./getAutomationRule"));

export { GetBookmarkArgs, GetBookmarkResult, GetBookmarkOutputArgs } from "./getBookmark";
export const getBookmark: typeof import("./getBookmark").getBookmark = null as any;
export const getBookmarkOutput: typeof import("./getBookmark").getBookmarkOutput = null as any;
utilities.lazyLoad(exports, ["getBookmark","getBookmarkOutput"], () => require("./getBookmark"));

export { GetBookmarkRelationArgs, GetBookmarkRelationResult, GetBookmarkRelationOutputArgs } from "./getBookmarkRelation";
export const getBookmarkRelation: typeof import("./getBookmarkRelation").getBookmarkRelation = null as any;
export const getBookmarkRelationOutput: typeof import("./getBookmarkRelation").getBookmarkRelationOutput = null as any;
utilities.lazyLoad(exports, ["getBookmarkRelation","getBookmarkRelationOutput"], () => require("./getBookmarkRelation"));

export { GetEntitiesGetTimelineArgs, GetEntitiesGetTimelineResult, GetEntitiesGetTimelineOutputArgs } from "./getEntitiesGetTimeline";
export const getEntitiesGetTimeline: typeof import("./getEntitiesGetTimeline").getEntitiesGetTimeline = null as any;
export const getEntitiesGetTimelineOutput: typeof import("./getEntitiesGetTimeline").getEntitiesGetTimelineOutput = null as any;
utilities.lazyLoad(exports, ["getEntitiesGetTimeline","getEntitiesGetTimelineOutput"], () => require("./getEntitiesGetTimeline"));

export { GetEntityInsightsArgs, GetEntityInsightsResult, GetEntityInsightsOutputArgs } from "./getEntityInsights";
export const getEntityInsights: typeof import("./getEntityInsights").getEntityInsights = null as any;
export const getEntityInsightsOutput: typeof import("./getEntityInsights").getEntityInsightsOutput = null as any;
utilities.lazyLoad(exports, ["getEntityInsights","getEntityInsightsOutput"], () => require("./getEntityInsights"));

export { GetIPSyncerArgs, GetIPSyncerResult, GetIPSyncerOutputArgs } from "./getIPSyncer";
export const getIPSyncer: typeof import("./getIPSyncer").getIPSyncer = null as any;
export const getIPSyncerOutput: typeof import("./getIPSyncer").getIPSyncerOutput = null as any;
utilities.lazyLoad(exports, ["getIPSyncer","getIPSyncerOutput"], () => require("./getIPSyncer"));

export { GetWatchlistArgs, GetWatchlistResult, GetWatchlistOutputArgs } from "./getWatchlist";
export const getWatchlist: typeof import("./getWatchlist").getWatchlist = null as any;
export const getWatchlistOutput: typeof import("./getWatchlist").getWatchlistOutput = null as any;
utilities.lazyLoad(exports, ["getWatchlist","getWatchlistOutput"], () => require("./getWatchlist"));

export { IPSyncerArgs } from "./ipsyncer";
export type IPSyncer = import("./ipsyncer").IPSyncer;
export const IPSyncer: typeof import("./ipsyncer").IPSyncer = null as any;
utilities.lazyLoad(exports, ["IPSyncer"], () => require("./ipsyncer"));

export { WatchlistArgs } from "./watchlist";
export type Watchlist = import("./watchlist").Watchlist;
export const Watchlist: typeof import("./watchlist").Watchlist = null as any;
utilities.lazyLoad(exports, ["Watchlist"], () => require("./watchlist"));


// Export enums:
export * from "../../types/enums/securityinsights/v20190101preview";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:securityinsights/v20190101preview:AutomationRule":
                return new AutomationRule(name, <any>undefined, { urn })
            case "azure-native:securityinsights/v20190101preview:Bookmark":
                return new Bookmark(name, <any>undefined, { urn })
            case "azure-native:securityinsights/v20190101preview:BookmarkRelation":
                return new BookmarkRelation(name, <any>undefined, { urn })
            case "azure-native:securityinsights/v20190101preview:IPSyncer":
                return new IPSyncer(name, <any>undefined, { urn })
            case "azure-native:securityinsights/v20190101preview:Watchlist":
                return new Watchlist(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "securityinsights/v20190101preview", _module)
