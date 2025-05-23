// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Inputs
{

    /// <summary>
    /// Istio service mesh configuration.
    /// </summary>
    public sealed class IstioServiceMeshArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Istio Service Mesh Certificate Authority (CA) configuration. For now, we only support plugin certificates as described here https://aka.ms/asm-plugin-ca
        /// </summary>
        [Input("certificateAuthority")]
        public Input<Inputs.IstioCertificateAuthorityArgs>? CertificateAuthority { get; set; }

        /// <summary>
        /// Istio components configuration.
        /// </summary>
        [Input("components")]
        public Input<Inputs.IstioComponentsArgs>? Components { get; set; }

        [Input("revisions")]
        private InputList<string>? _revisions;

        /// <summary>
        /// The list of revisions of the Istio control plane. When an upgrade is not in progress, this holds one value. When canary upgrade is in progress, this can only hold two consecutive values. For more information, see: https://learn.microsoft.com/en-us/azure/aks/istio-upgrade
        /// </summary>
        public InputList<string> Revisions
        {
            get => _revisions ?? (_revisions = new InputList<string>());
            set => _revisions = value;
        }

        public IstioServiceMeshArgs()
        {
        }
        public static new IstioServiceMeshArgs Empty => new IstioServiceMeshArgs();
    }
}
