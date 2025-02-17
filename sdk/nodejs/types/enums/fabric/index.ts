// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// Export sub-modules:
import * as v20231101 from "./v20231101";
import * as v20250115preview from "./v20250115preview";

export {
    v20231101,
    v20250115preview,
};

export const RpSkuTier = {
    /**
     * Fabric tier
     */
    Fabric: "Fabric",
} as const;

/**
 * The name of the Azure pricing tier to which the SKU applies.
 */
export type RpSkuTier = (typeof RpSkuTier)[keyof typeof RpSkuTier];
