// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const IntervalInMins = {
    ThreeMins: "ThreeMins",
    FiveMins: "FiveMins",
    ThirtyMins: "ThirtyMins",
    SixtyMins: "SixtyMins",
} as const;

/**
 * Interval value in minutes used to create LogAnalytics call rate logs.
 */
export type IntervalInMins = (typeof IntervalInMins)[keyof typeof IntervalInMins];

export const StatusLevelTypes = {
    Info: "Info",
    Warning: "Warning",
    Error: "Error",
} as const;

/**
 * The level code.
 */
export type StatusLevelTypes = (typeof StatusLevelTypes)[keyof typeof StatusLevelTypes];
