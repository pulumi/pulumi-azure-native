// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridCompute.Inputs
{

    /// <summary>
    /// License Profile Instance View in Machine Properties.
    /// </summary>
    public sealed class LicenseProfileMachineInstanceViewArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Properties for the Machine ESU profile.
        /// </summary>
        [Input("esuProfile")]
        public Input<Inputs.LicenseProfileMachineInstanceViewEsuPropertiesArgs>? EsuProfile { get; set; }

        [Input("productFeatures")]
        private InputList<Inputs.ProductFeatureArgs>? _productFeatures;

        /// <summary>
        /// The list of product features.
        /// </summary>
        public InputList<Inputs.ProductFeatureArgs> ProductFeatures
        {
            get => _productFeatures ?? (_productFeatures = new InputList<Inputs.ProductFeatureArgs>());
            set => _productFeatures = value;
        }

        /// <summary>
        /// Indicates the product type of the license.
        /// </summary>
        [Input("productType")]
        public InputUnion<string, Pulumi.AzureNative.HybridCompute.LicenseProfileProductType>? ProductType { get; set; }

        /// <summary>
        /// Specifies if this machine is licensed as part of a Software Assurance agreement.
        /// </summary>
        [Input("softwareAssuranceCustomer")]
        public Input<bool>? SoftwareAssuranceCustomer { get; set; }

        /// <summary>
        /// Indicates the subscription status of the product.
        /// </summary>
        [Input("subscriptionStatus")]
        public InputUnion<string, Pulumi.AzureNative.HybridCompute.LicenseProfileSubscriptionStatus>? SubscriptionStatus { get; set; }

        public LicenseProfileMachineInstanceViewArgs()
        {
        }
        public static new LicenseProfileMachineInstanceViewArgs Empty => new LicenseProfileMachineInstanceViewArgs();
    }
}
