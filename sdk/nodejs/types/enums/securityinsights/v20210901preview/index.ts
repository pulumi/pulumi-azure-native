// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const EntityTimelineKind = {
    /**
     * activity
     */
    Activity: "Activity",
    /**
     * bookmarks
     */
    Bookmark: "Bookmark",
    /**
     * security alerts
     */
    SecurityAlert: "SecurityAlert",
} as const;

/**
 * The entity query kind
 */
export type EntityTimelineKind = (typeof EntityTimelineKind)[keyof typeof EntityTimelineKind];

export const ThreatIntelligenceResourceKindEnum = {
    /**
     * Entity represents threat intelligence indicator in the system.
     */
    Indicator: "indicator",
} as const;

/**
 * The kind of the entity.
 */
export type ThreatIntelligenceResourceKindEnum = (typeof ThreatIntelligenceResourceKindEnum)[keyof typeof ThreatIntelligenceResourceKindEnum];
