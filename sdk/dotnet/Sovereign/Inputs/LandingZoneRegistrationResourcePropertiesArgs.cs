// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sovereign.Inputs
{

    /// <summary>
    /// The properties of landing zone registration resource type.
    /// </summary>
    public sealed class LandingZoneRegistrationResourcePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource id of the associated landing zone configuration.
        /// </summary>
        [Input("existingLandingZoneConfigurationId", required: true)]
        public Input<string> ExistingLandingZoneConfigurationId { get; set; } = null!;

        /// <summary>
        /// The resource id of the top level management group
        /// </summary>
        [Input("existingTopLevelMgId", required: true)]
        public Input<string> ExistingTopLevelMgId { get; set; } = null!;

        /// <summary>
        /// The managed identity to be assigned to this landing zone registration.
        /// </summary>
        [Input("managedIdentity")]
        public Input<Inputs.ManagedIdentityPropertiesArgs>? ManagedIdentity { get; set; }

        public LandingZoneRegistrationResourcePropertiesArgs()
        {
        }
        public static new LandingZoneRegistrationResourcePropertiesArgs Empty => new LandingZoneRegistrationResourcePropertiesArgs();
    }
}
