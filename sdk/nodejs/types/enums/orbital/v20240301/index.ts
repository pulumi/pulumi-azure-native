// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const Capability = {
    EarthObservation: "EarthObservation",
    Communication: "Communication",
} as const;

/**
 * Capability of the Ground Station.
 */
export type Capability = (typeof Capability)[keyof typeof Capability];
