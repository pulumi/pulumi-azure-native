// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { DnsForwardingRulesetArgs } from "./dnsForwardingRuleset";
export type DnsForwardingRuleset = import("./dnsForwardingRuleset").DnsForwardingRuleset;
export const DnsForwardingRuleset: typeof import("./dnsForwardingRuleset").DnsForwardingRuleset = null as any;
utilities.lazyLoad(exports, ["DnsForwardingRuleset"], () => require("./dnsForwardingRuleset"));

export { GetDnsForwardingRulesetArgs, GetDnsForwardingRulesetResult, GetDnsForwardingRulesetOutputArgs } from "./getDnsForwardingRuleset";
export const getDnsForwardingRuleset: typeof import("./getDnsForwardingRuleset").getDnsForwardingRuleset = null as any;
export const getDnsForwardingRulesetOutput: typeof import("./getDnsForwardingRuleset").getDnsForwardingRulesetOutput = null as any;
utilities.lazyLoad(exports, ["getDnsForwardingRuleset","getDnsForwardingRulesetOutput"], () => require("./getDnsForwardingRuleset"));

export { GetInboundEndpointArgs, GetInboundEndpointResult, GetInboundEndpointOutputArgs } from "./getInboundEndpoint";
export const getInboundEndpoint: typeof import("./getInboundEndpoint").getInboundEndpoint = null as any;
export const getInboundEndpointOutput: typeof import("./getInboundEndpoint").getInboundEndpointOutput = null as any;
utilities.lazyLoad(exports, ["getInboundEndpoint","getInboundEndpointOutput"], () => require("./getInboundEndpoint"));

export { GetOutboundEndpointArgs, GetOutboundEndpointResult, GetOutboundEndpointOutputArgs } from "./getOutboundEndpoint";
export const getOutboundEndpoint: typeof import("./getOutboundEndpoint").getOutboundEndpoint = null as any;
export const getOutboundEndpointOutput: typeof import("./getOutboundEndpoint").getOutboundEndpointOutput = null as any;
utilities.lazyLoad(exports, ["getOutboundEndpoint","getOutboundEndpointOutput"], () => require("./getOutboundEndpoint"));

export { GetPrivateResolverVirtualNetworkLinkArgs, GetPrivateResolverVirtualNetworkLinkResult, GetPrivateResolverVirtualNetworkLinkOutputArgs } from "./getPrivateResolverVirtualNetworkLink";
export const getPrivateResolverVirtualNetworkLink: typeof import("./getPrivateResolverVirtualNetworkLink").getPrivateResolverVirtualNetworkLink = null as any;
export const getPrivateResolverVirtualNetworkLinkOutput: typeof import("./getPrivateResolverVirtualNetworkLink").getPrivateResolverVirtualNetworkLinkOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateResolverVirtualNetworkLink","getPrivateResolverVirtualNetworkLinkOutput"], () => require("./getPrivateResolverVirtualNetworkLink"));

export { InboundEndpointArgs } from "./inboundEndpoint";
export type InboundEndpoint = import("./inboundEndpoint").InboundEndpoint;
export const InboundEndpoint: typeof import("./inboundEndpoint").InboundEndpoint = null as any;
utilities.lazyLoad(exports, ["InboundEndpoint"], () => require("./inboundEndpoint"));

export { ListDnsForwardingRulesetByVirtualNetworkArgs, ListDnsForwardingRulesetByVirtualNetworkResult, ListDnsForwardingRulesetByVirtualNetworkOutputArgs } from "./listDnsForwardingRulesetByVirtualNetwork";
export const listDnsForwardingRulesetByVirtualNetwork: typeof import("./listDnsForwardingRulesetByVirtualNetwork").listDnsForwardingRulesetByVirtualNetwork = null as any;
export const listDnsForwardingRulesetByVirtualNetworkOutput: typeof import("./listDnsForwardingRulesetByVirtualNetwork").listDnsForwardingRulesetByVirtualNetworkOutput = null as any;
utilities.lazyLoad(exports, ["listDnsForwardingRulesetByVirtualNetwork","listDnsForwardingRulesetByVirtualNetworkOutput"], () => require("./listDnsForwardingRulesetByVirtualNetwork"));

export { ListDnsResolverByVirtualNetworkArgs, ListDnsResolverByVirtualNetworkResult, ListDnsResolverByVirtualNetworkOutputArgs } from "./listDnsResolverByVirtualNetwork";
export const listDnsResolverByVirtualNetwork: typeof import("./listDnsResolverByVirtualNetwork").listDnsResolverByVirtualNetwork = null as any;
export const listDnsResolverByVirtualNetworkOutput: typeof import("./listDnsResolverByVirtualNetwork").listDnsResolverByVirtualNetworkOutput = null as any;
utilities.lazyLoad(exports, ["listDnsResolverByVirtualNetwork","listDnsResolverByVirtualNetworkOutput"], () => require("./listDnsResolverByVirtualNetwork"));

export { OutboundEndpointArgs } from "./outboundEndpoint";
export type OutboundEndpoint = import("./outboundEndpoint").OutboundEndpoint;
export const OutboundEndpoint: typeof import("./outboundEndpoint").OutboundEndpoint = null as any;
utilities.lazyLoad(exports, ["OutboundEndpoint"], () => require("./outboundEndpoint"));

export { PrivateResolverVirtualNetworkLinkArgs } from "./privateResolverVirtualNetworkLink";
export type PrivateResolverVirtualNetworkLink = import("./privateResolverVirtualNetworkLink").PrivateResolverVirtualNetworkLink;
export const PrivateResolverVirtualNetworkLink: typeof import("./privateResolverVirtualNetworkLink").PrivateResolverVirtualNetworkLink = null as any;
utilities.lazyLoad(exports, ["PrivateResolverVirtualNetworkLink"], () => require("./privateResolverVirtualNetworkLink"));


// Export enums:
export * from "../../types/enums/network/v20200401preview";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:network/v20200401preview:DnsForwardingRuleset":
                return new DnsForwardingRuleset(name, <any>undefined, { urn })
            case "azure-native:network/v20200401preview:InboundEndpoint":
                return new InboundEndpoint(name, <any>undefined, { urn })
            case "azure-native:network/v20200401preview:OutboundEndpoint":
                return new OutboundEndpoint(name, <any>undefined, { urn })
            case "azure-native:network/v20200401preview:PrivateResolverVirtualNetworkLink":
                return new PrivateResolverVirtualNetworkLink(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "network/v20200401preview", _module)
