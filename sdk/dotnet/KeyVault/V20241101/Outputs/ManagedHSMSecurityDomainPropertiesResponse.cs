// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.KeyVault.V20241101.Outputs
{

    /// <summary>
    /// The security domain properties of the managed hsm.
    /// </summary>
    [OutputType]
    public sealed class ManagedHSMSecurityDomainPropertiesResponse
    {
        /// <summary>
        /// Activation Status
        /// </summary>
        public readonly string ActivationStatus;
        /// <summary>
        /// Activation Status Message.
        /// </summary>
        public readonly string ActivationStatusMessage;

        [OutputConstructor]
        private ManagedHSMSecurityDomainPropertiesResponse(
            string activationStatus,

            string activationStatusMessage)
        {
            ActivationStatus = activationStatus;
            ActivationStatusMessage = activationStatusMessage;
        }
    }
}
