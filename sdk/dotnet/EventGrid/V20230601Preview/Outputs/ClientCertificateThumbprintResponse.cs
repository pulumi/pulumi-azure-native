// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.V20230601Preview.Outputs
{

    /// <summary>
    /// Thumbprints are used by the service to validate the device permission when authentication is done using self signed certificate.
    /// </summary>
    [OutputType]
    public sealed class ClientCertificateThumbprintResponse
    {
        /// <summary>
        /// The primary thumbprint used for validation.
        /// </summary>
        public readonly string? Primary;
        /// <summary>
        /// The secondary thumbprint used for validation.
        /// </summary>
        public readonly string? Secondary;

        [OutputConstructor]
        private ClientCertificateThumbprintResponse(
            string? primary,

            string? secondary)
        {
            Primary = primary;
            Secondary = secondary;
        }
    }
}
