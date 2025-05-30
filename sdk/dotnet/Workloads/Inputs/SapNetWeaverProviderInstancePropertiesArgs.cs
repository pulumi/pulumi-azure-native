// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.Inputs
{

    /// <summary>
    /// Gets or sets the provider properties.
    /// </summary>
    public sealed class SapNetWeaverProviderInstancePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The provider type. For example, the value can be SapHana.
        /// Expected value is 'SapNetWeaver'.
        /// </summary>
        [Input("providerType", required: true)]
        public Input<string> ProviderType { get; set; } = null!;

        /// <summary>
        /// Gets or sets the SAP Client ID.
        /// </summary>
        [Input("sapClientId")]
        public Input<string>? SapClientId { get; set; }

        [Input("sapHostFileEntries")]
        private InputList<string>? _sapHostFileEntries;

        /// <summary>
        /// Gets or sets the list of HostFile Entries
        /// </summary>
        public InputList<string> SapHostFileEntries
        {
            get => _sapHostFileEntries ?? (_sapHostFileEntries = new InputList<string>());
            set => _sapHostFileEntries = value;
        }

        /// <summary>
        /// Gets or sets the target virtual machine IP Address/FQDN.
        /// </summary>
        [Input("sapHostname")]
        public Input<string>? SapHostname { get; set; }

        /// <summary>
        /// Gets or sets the instance number of SAP NetWeaver.
        /// </summary>
        [Input("sapInstanceNr")]
        public Input<string>? SapInstanceNr { get; set; }

        /// <summary>
        /// Sets the SAP password.
        /// </summary>
        [Input("sapPassword")]
        public Input<string>? SapPassword { get; set; }

        /// <summary>
        /// Gets or sets the key vault URI to secret with the SAP password.
        /// </summary>
        [Input("sapPasswordUri")]
        public Input<string>? SapPasswordUri { get; set; }

        /// <summary>
        /// Gets or sets the SAP HTTP port number.
        /// </summary>
        [Input("sapPortNumber")]
        public Input<string>? SapPortNumber { get; set; }

        /// <summary>
        /// Gets or sets the SAP System Identifier
        /// </summary>
        [Input("sapSid")]
        public Input<string>? SapSid { get; set; }

        /// <summary>
        /// Gets or sets the SAP user name.
        /// </summary>
        [Input("sapUsername")]
        public Input<string>? SapUsername { get; set; }

        /// <summary>
        /// Gets or sets the blob URI to SSL certificate for the SAP system.
        /// </summary>
        [Input("sslCertificateUri")]
        public Input<string>? SslCertificateUri { get; set; }

        /// <summary>
        /// Gets or sets certificate preference if secure communication is enabled.
        /// </summary>
        [Input("sslPreference")]
        public InputUnion<string, Pulumi.AzureNative.Workloads.SslPreference>? SslPreference { get; set; }

        public SapNetWeaverProviderInstancePropertiesArgs()
        {
        }
        public static new SapNetWeaverProviderInstancePropertiesArgs Empty => new SapNetWeaverProviderInstancePropertiesArgs();
    }
}
