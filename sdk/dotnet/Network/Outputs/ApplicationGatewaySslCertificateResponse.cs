// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Outputs
{

    /// <summary>
    /// SSL certificates of an application gateway.
    /// </summary>
    [OutputType]
    public sealed class ApplicationGatewaySslCertificateResponse
    {
        /// <summary>
        /// Base-64 encoded pfx certificate. Only applicable in PUT Request.
        /// </summary>
        public readonly string? Data;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Secret Id of (base-64 encoded unencrypted pfx) 'Secret' or 'Certificate' object stored in KeyVault.
        /// </summary>
        public readonly string? KeyVaultSecretId;
        /// <summary>
        /// Name of the SSL certificate that is unique within an Application Gateway.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Password for the pfx file specified in data. Only applicable in PUT request.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// The provisioning state of the SSL certificate resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Base-64 encoded Public cert data corresponding to pfx specified in data. Only applicable in GET request.
        /// </summary>
        public readonly string PublicCertData;
        /// <summary>
        /// Type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ApplicationGatewaySslCertificateResponse(
            string? data,

            string etag,

            string? id,

            string? keyVaultSecretId,

            string? name,

            string? password,

            string provisioningState,

            string publicCertData,

            string type)
        {
            Data = data;
            Etag = etag;
            Id = id;
            KeyVaultSecretId = keyVaultSecretId;
            Name = name;
            Password = password;
            ProvisioningState = provisioningState;
            PublicCertData = publicCertData;
            Type = type;
        }
    }
}
