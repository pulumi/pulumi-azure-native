// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VirtualMachineImages.Outputs
{

    /// <summary>
    /// Purchase plan configuration for platform image.
    /// </summary>
    [OutputType]
    public sealed class PlatformImagePurchasePlanResponse
    {
        /// <summary>
        /// Name of the purchase plan.
        /// </summary>
        public readonly string PlanName;
        /// <summary>
        /// Product of the purchase plan.
        /// </summary>
        public readonly string PlanProduct;
        /// <summary>
        /// Publisher of the purchase plan.
        /// </summary>
        public readonly string PlanPublisher;

        [OutputConstructor]
        private PlatformImagePurchasePlanResponse(
            string planName,

            string planProduct,

            string planPublisher)
        {
            PlanName = planName;
            PlanProduct = planProduct;
            PlanPublisher = planPublisher;
        }
    }
}
