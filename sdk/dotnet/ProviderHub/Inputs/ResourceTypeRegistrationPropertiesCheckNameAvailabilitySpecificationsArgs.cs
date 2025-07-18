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
    /// The check name availability specifications.
    /// </summary>
    public sealed class ResourceTypeRegistrationPropertiesCheckNameAvailabilitySpecificationsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether default validation is enabled.
        /// </summary>
        [Input("enableDefaultValidation")]
        public Input<bool>? EnableDefaultValidation { get; set; }

        [Input("resourceTypesWithCustomValidation")]
        private InputList<string>? _resourceTypesWithCustomValidation;

        /// <summary>
        /// The resource types with custom validation.
        /// </summary>
        public InputList<string> ResourceTypesWithCustomValidation
        {
            get => _resourceTypesWithCustomValidation ?? (_resourceTypesWithCustomValidation = new InputList<string>());
            set => _resourceTypesWithCustomValidation = value;
        }

        public ResourceTypeRegistrationPropertiesCheckNameAvailabilitySpecificationsArgs()
        {
        }
        public static new ResourceTypeRegistrationPropertiesCheckNameAvailabilitySpecificationsArgs Empty => new ResourceTypeRegistrationPropertiesCheckNameAvailabilitySpecificationsArgs();
    }
}
