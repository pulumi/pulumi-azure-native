// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Inputs
{

    public sealed class ResourceTypeRegistrationPropertiesTemplateDeploymentOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("preflightOptions")]
        private InputList<Union<string, Pulumi.AzureNative.ProviderHub.PreflightOption>>? _preflightOptions;
        public InputList<Union<string, Pulumi.AzureNative.ProviderHub.PreflightOption>> PreflightOptions
        {
            get => _preflightOptions ?? (_preflightOptions = new InputList<Union<string, Pulumi.AzureNative.ProviderHub.PreflightOption>>());
            set => _preflightOptions = value;
        }

        [Input("preflightSupported")]
        public Input<bool>? PreflightSupported { get; set; }

        public ResourceTypeRegistrationPropertiesTemplateDeploymentOptionsArgs()
        {
        }
        public static new ResourceTypeRegistrationPropertiesTemplateDeploymentOptionsArgs Empty => new ResourceTypeRegistrationPropertiesTemplateDeploymentOptionsArgs();
    }
}
