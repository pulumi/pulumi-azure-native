// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Inputs
{

    public sealed class ResourceTypeRegistrationPropertiesCheckNameAvailabilitySpecificationsArgs : global::Pulumi.ResourceArgs
    {
        [Input("enableDefaultValidation")]
        public Input<bool>? EnableDefaultValidation { get; set; }

        [Input("resourceTypesWithCustomValidation")]
        private InputList<string>? _resourceTypesWithCustomValidation;
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
