// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Inputs
{

    /// <summary>
    /// The properties of an add provider request.
    /// </summary>
    public sealed class AddRecoveryServicesProviderInputPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identity provider input for DRA authentication.
        /// </summary>
        [Input("authenticationIdentityInput", required: true)]
        public Input<Inputs.IdentityProviderInputArgs> AuthenticationIdentityInput { get; set; } = null!;

        /// <summary>
        /// The Bios Id of the machine.
        /// </summary>
        [Input("biosId")]
        public Input<string>? BiosId { get; set; }

        /// <summary>
        /// The identity provider input for data plane authentication.
        /// </summary>
        [Input("dataPlaneAuthenticationIdentityInput")]
        public Input<Inputs.IdentityProviderInputArgs>? DataPlaneAuthenticationIdentityInput { get; set; }

        /// <summary>
        /// The Id of the machine where the provider is getting added.
        /// </summary>
        [Input("machineId")]
        public Input<string>? MachineId { get; set; }

        /// <summary>
        /// The name of the machine where the provider is getting added.
        /// </summary>
        [Input("machineName", required: true)]
        public Input<string> MachineName { get; set; } = null!;

        /// <summary>
        /// The identity provider input for resource access.
        /// </summary>
        [Input("resourceAccessIdentityInput", required: true)]
        public Input<Inputs.IdentityProviderInputArgs> ResourceAccessIdentityInput { get; set; } = null!;

        public AddRecoveryServicesProviderInputPropertiesArgs()
        {
        }
        public static new AddRecoveryServicesProviderInputPropertiesArgs Empty => new AddRecoveryServicesProviderInputPropertiesArgs();
    }
}
