// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperations.Outputs
{

    /// <summary>
    /// X509 Certificate Authentication properties.
    /// </summary>
    [OutputType]
    public sealed class X509ManualCertificateResponse
    {
        /// <summary>
        /// Kubernetes secret containing an X.509 client certificate. This is a reference to the secret through an identifying name, not the secret itself.
        /// </summary>
        public readonly string SecretRef;

        [OutputConstructor]
        private X509ManualCertificateResponse(string secretRef)
        {
            SecretRef = secretRef;
        }
    }
}
