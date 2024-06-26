// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridCompute.V20240331Preview.Outputs
{

    /// <summary>
    /// License Profile Instance View in Machine Properties.
    /// </summary>
    [OutputType]
    public sealed class LicenseProfileMachineInstanceViewResponse
    {
        /// <summary>
        /// The timestamp in UTC when the billing starts.
        /// </summary>
        public readonly string BillingStartDate;
        /// <summary>
        /// The timestamp in UTC when the user disenrolled the feature.
        /// </summary>
        public readonly string DisenrollmentDate;
        /// <summary>
        /// The timestamp in UTC when the user enrolls the feature.
        /// </summary>
        public readonly string EnrollmentDate;
        /// <summary>
        /// Properties for the Machine ESU profile.
        /// </summary>
        public readonly Outputs.LicenseProfileMachineInstanceViewEsuPropertiesResponse? EsuProfile;
        /// <summary>
        /// Indicates the license channel.
        /// </summary>
        public readonly string LicenseChannel;
        /// <summary>
        /// Indicates the license status of the OS.
        /// </summary>
        public readonly string LicenseStatus;
        /// <summary>
        /// The list of product features.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProductFeatureResponse> ProductFeatures;
        /// <summary>
        /// Indicates the product type of the license.
        /// </summary>
        public readonly string? ProductType;
        /// <summary>
        /// Specifies if this machine is licensed as part of a Software Assurance agreement.
        /// </summary>
        public readonly bool? SoftwareAssuranceCustomer;
        /// <summary>
        /// Indicates the subscription status of the product.
        /// </summary>
        public readonly string? SubscriptionStatus;

        [OutputConstructor]
        private LicenseProfileMachineInstanceViewResponse(
            string billingStartDate,

            string disenrollmentDate,

            string enrollmentDate,

            Outputs.LicenseProfileMachineInstanceViewEsuPropertiesResponse? esuProfile,

            string licenseChannel,

            string licenseStatus,

            ImmutableArray<Outputs.ProductFeatureResponse> productFeatures,

            string? productType,

            bool? softwareAssuranceCustomer,

            string? subscriptionStatus)
        {
            BillingStartDate = billingStartDate;
            DisenrollmentDate = disenrollmentDate;
            EnrollmentDate = enrollmentDate;
            EsuProfile = esuProfile;
            LicenseChannel = licenseChannel;
            LicenseStatus = licenseStatus;
            ProductFeatures = productFeatures;
            ProductType = productType;
            SoftwareAssuranceCustomer = softwareAssuranceCustomer;
            SubscriptionStatus = subscriptionStatus;
        }
    }
}
