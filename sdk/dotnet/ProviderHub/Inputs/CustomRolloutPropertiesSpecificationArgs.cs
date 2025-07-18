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
    /// The specification.
    /// </summary>
    public sealed class CustomRolloutPropertiesSpecificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The auto provisioning configuration.
        /// </summary>
        [Input("autoProvisionConfig")]
        public Input<Inputs.CustomRolloutSpecificationAutoProvisionConfigArgs>? AutoProvisionConfig { get; set; }

        /// <summary>
        /// The canary region configuration.
        /// </summary>
        [Input("canary")]
        public Input<Inputs.CustomRolloutSpecificationCanaryArgs>? Canary { get; set; }

        /// <summary>
        /// The provider registration.
        /// </summary>
        [Input("providerRegistration")]
        public Input<Inputs.CustomRolloutSpecificationProviderRegistrationArgs>? ProviderRegistration { get; set; }

        /// <summary>
        /// Whether refreshing subscription registration is enabled or disabled.
        /// </summary>
        [Input("refreshSubscriptionRegistration")]
        public Input<bool>? RefreshSubscriptionRegistration { get; set; }

        [Input("releaseScopes")]
        private InputList<string>? _releaseScopes;

        /// <summary>
        /// The list of ARM regions scoped for the release.
        /// </summary>
        public InputList<string> ReleaseScopes
        {
            get => _releaseScopes ?? (_releaseScopes = new InputList<string>());
            set => _releaseScopes = value;
        }

        [Input("resourceTypeRegistrations")]
        private InputList<Inputs.ResourceTypeRegistrationArgs>? _resourceTypeRegistrations;

        /// <summary>
        /// The resource type registrations.
        /// </summary>
        public InputList<Inputs.ResourceTypeRegistrationArgs> ResourceTypeRegistrations
        {
            get => _resourceTypeRegistrations ?? (_resourceTypeRegistrations = new InputList<Inputs.ResourceTypeRegistrationArgs>());
            set => _resourceTypeRegistrations = value;
        }

        /// <summary>
        /// Whether release scope validation should be skipped.
        /// </summary>
        [Input("skipReleaseScopeValidation")]
        public Input<bool>? SkipReleaseScopeValidation { get; set; }

        public CustomRolloutPropertiesSpecificationArgs()
        {
        }
        public static new CustomRolloutPropertiesSpecificationArgs Empty => new CustomRolloutPropertiesSpecificationArgs();
    }
}
