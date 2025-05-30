// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureFleet.Outputs
{

    /// <summary>
    /// Specifies the Security profile settings for the virtual machine or virtual
    /// machine scale set.
    /// </summary>
    [OutputType]
    public sealed class SecurityProfileResponse
    {
        /// <summary>
        /// This property can be used by user in the request to enable or disable the Host
        /// Encryption for the virtual machine or virtual machine scale set. This will
        /// enable the encryption for all the disks including Resource/Temp disk at host
        /// itself. The default behavior is: The Encryption at host will be disabled unless
        /// this property is set to true for the resource.
        /// </summary>
        public readonly bool? EncryptionAtHost;
        /// <summary>
        /// Specifies the Managed Identity used by ADE to get access token for keyvault
        /// operations.
        /// </summary>
        public readonly Outputs.EncryptionIdentityResponse? EncryptionIdentity;
        /// <summary>
        /// Specifies ProxyAgent settings while creating the virtual machine. Minimum
        /// api-version: 2023-09-01.
        /// </summary>
        public readonly Outputs.ProxyAgentSettingsResponse? ProxyAgentSettings;
        /// <summary>
        /// Specifies the SecurityType of the virtual machine. It has to be set to any
        /// specified value to enable UefiSettings. The default behavior is: UefiSettings
        /// will not be enabled unless this property is set.
        /// </summary>
        public readonly string? SecurityType;
        /// <summary>
        /// Specifies the security settings like secure boot and vTPM used while creating
        /// the virtual machine. Minimum api-version: 2020-12-01.
        /// </summary>
        public readonly Outputs.UefiSettingsResponse? UefiSettings;

        [OutputConstructor]
        private SecurityProfileResponse(
            bool? encryptionAtHost,

            Outputs.EncryptionIdentityResponse? encryptionIdentity,

            Outputs.ProxyAgentSettingsResponse? proxyAgentSettings,

            string? securityType,

            Outputs.UefiSettingsResponse? uefiSettings)
        {
            EncryptionAtHost = encryptionAtHost;
            EncryptionIdentity = encryptionIdentity;
            ProxyAgentSettings = proxyAgentSettings;
            SecurityType = securityType;
            UefiSettings = uefiSettings;
        }
    }
}
