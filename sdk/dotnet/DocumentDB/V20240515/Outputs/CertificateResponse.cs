// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DocumentDB.V20240515.Outputs
{

    [OutputType]
    public sealed class CertificateResponse
    {
        /// <summary>
        /// PEM formatted public key.
        /// </summary>
        public readonly string? Pem;

        [OutputConstructor]
        private CertificateResponse(string? pem)
        {
            Pem = pem;
        }
    }
}
