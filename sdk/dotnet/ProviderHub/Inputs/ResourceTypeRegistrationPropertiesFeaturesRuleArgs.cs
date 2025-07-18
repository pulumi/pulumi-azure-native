// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Inputs
{

    /// <summary>
    /// The features rule.
    /// </summary>
    public sealed class ResourceTypeRegistrationPropertiesFeaturesRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The required feature policy.
        /// </summary>
        [Input("requiredFeaturesPolicy", required: true)]
        public InputUnion<string, Pulumi.AzureNative.ProviderHub.FeaturesPolicy> RequiredFeaturesPolicy { get; set; } = null!;

        public ResourceTypeRegistrationPropertiesFeaturesRuleArgs()
        {
        }
        public static new ResourceTypeRegistrationPropertiesFeaturesRuleArgs Empty => new ResourceTypeRegistrationPropertiesFeaturesRuleArgs();
    }
}
