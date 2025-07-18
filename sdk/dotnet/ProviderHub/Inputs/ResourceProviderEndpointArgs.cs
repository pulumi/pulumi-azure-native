// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Inputs
{

    public sealed class ResourceProviderEndpointArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiVersions")]
        private InputList<string>? _apiVersions;

        /// <summary>
        /// The api versions.
        /// </summary>
        public InputList<string> ApiVersions
        {
            get => _apiVersions ?? (_apiVersions = new InputList<string>());
            set => _apiVersions = value;
        }

        /// <summary>
        /// Whether the endpoint is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The endpoint type.
        /// </summary>
        [Input("endpointType")]
        public InputUnion<string, Pulumi.AzureNative.ProviderHub.EndpointType>? EndpointType { get; set; }

        /// <summary>
        /// The endpoint uri.
        /// </summary>
        [Input("endpointUri")]
        public Input<string>? EndpointUri { get; set; }

        /// <summary>
        /// The feature rules.
        /// </summary>
        [Input("featuresRule")]
        public Input<Inputs.ResourceProviderEndpointFeaturesRuleArgs>? FeaturesRule { get; set; }

        [Input("locations")]
        private InputList<string>? _locations;

        /// <summary>
        /// The locations.
        /// </summary>
        public InputList<string> Locations
        {
            get => _locations ?? (_locations = new InputList<string>());
            set => _locations = value;
        }

        [Input("requiredFeatures")]
        private InputList<string>? _requiredFeatures;

        /// <summary>
        /// The required features.
        /// </summary>
        public InputList<string> RequiredFeatures
        {
            get => _requiredFeatures ?? (_requiredFeatures = new InputList<string>());
            set => _requiredFeatures = value;
        }

        /// <summary>
        /// The sku link.
        /// </summary>
        [Input("skuLink")]
        public Input<string>? SkuLink { get; set; }

        /// <summary>
        /// The timeout.
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        public ResourceProviderEndpointArgs()
        {
        }
        public static new ResourceProviderEndpointArgs Empty => new ResourceProviderEndpointArgs();
    }
}
