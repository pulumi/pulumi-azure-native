// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.KubernetesConfiguration.Inputs
{

    /// <summary>
    /// Parameters to authenticate using Service Principal.
    /// </summary>
    public sealed class ServicePrincipalDefinitionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Base64-encoded certificate used to authenticate a Service Principal 
        /// </summary>
        [Input("clientCertificate")]
        public Input<string>? ClientCertificate { get; set; }

        /// <summary>
        /// The password for the certificate used to authenticate a Service Principal 
        /// </summary>
        [Input("clientCertificatePassword")]
        public Input<string>? ClientCertificatePassword { get; set; }

        /// <summary>
        /// Specifies whether to include x5c header in client claims when acquiring a token to enable subject name / issuer based authentication for the Client Certificate
        /// </summary>
        [Input("clientCertificateSendChain")]
        public Input<bool>? ClientCertificateSendChain { get; set; }

        /// <summary>
        /// The client Id for authenticating a Service Principal.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The client secret for authenticating a Service Principal
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// The tenant Id for authenticating a Service Principal
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public ServicePrincipalDefinitionArgs()
        {
            ClientCertificateSendChain = false;
        }
        public static new ServicePrincipalDefinitionArgs Empty => new ServicePrincipalDefinitionArgs();
    }
}
