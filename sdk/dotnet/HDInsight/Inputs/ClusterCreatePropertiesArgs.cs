// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Inputs
{

    /// <summary>
    /// The cluster create parameters.
    /// </summary>
    public sealed class ClusterCreatePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cluster definition.
        /// </summary>
        [Input("clusterDefinition")]
        public Input<Inputs.ClusterDefinitionArgs>? ClusterDefinition { get; set; }

        /// <summary>
        /// The version of the cluster.
        /// </summary>
        [Input("clusterVersion")]
        public Input<string>? ClusterVersion { get; set; }

        /// <summary>
        /// The compute isolation properties.
        /// </summary>
        [Input("computeIsolationProperties")]
        public Input<Inputs.ComputeIsolationPropertiesArgs>? ComputeIsolationProperties { get; set; }

        /// <summary>
        /// The compute profile.
        /// </summary>
        [Input("computeProfile")]
        public Input<Inputs.ComputeProfileArgs>? ComputeProfile { get; set; }

        /// <summary>
        /// The disk encryption properties.
        /// </summary>
        [Input("diskEncryptionProperties")]
        public Input<Inputs.DiskEncryptionPropertiesArgs>? DiskEncryptionProperties { get; set; }

        /// <summary>
        /// The encryption-in-transit properties.
        /// </summary>
        [Input("encryptionInTransitProperties")]
        public Input<Inputs.EncryptionInTransitPropertiesArgs>? EncryptionInTransitProperties { get; set; }

        /// <summary>
        /// The cluster kafka rest proxy configuration.
        /// </summary>
        [Input("kafkaRestProperties")]
        public Input<Inputs.KafkaRestPropertiesArgs>? KafkaRestProperties { get; set; }

        /// <summary>
        /// The minimal supported tls version.
        /// </summary>
        [Input("minSupportedTlsVersion")]
        public Input<string>? MinSupportedTlsVersion { get; set; }

        /// <summary>
        /// The network properties.
        /// </summary>
        [Input("networkProperties")]
        public Input<Inputs.NetworkPropertiesArgs>? NetworkProperties { get; set; }

        /// <summary>
        /// The type of operating system.
        /// </summary>
        [Input("osType")]
        public InputUnion<string, Pulumi.AzureNative.HDInsight.OSType>? OsType { get; set; }

        [Input("privateLinkConfigurations")]
        private InputList<Inputs.PrivateLinkConfigurationArgs>? _privateLinkConfigurations;

        /// <summary>
        /// The private link configurations.
        /// </summary>
        public InputList<Inputs.PrivateLinkConfigurationArgs> PrivateLinkConfigurations
        {
            get => _privateLinkConfigurations ?? (_privateLinkConfigurations = new InputList<Inputs.PrivateLinkConfigurationArgs>());
            set => _privateLinkConfigurations = value;
        }

        /// <summary>
        /// The security profile.
        /// </summary>
        [Input("securityProfile")]
        public Input<Inputs.SecurityProfileArgs>? SecurityProfile { get; set; }

        /// <summary>
        /// The storage profile.
        /// </summary>
        [Input("storageProfile")]
        public Input<Inputs.StorageProfileArgs>? StorageProfile { get; set; }

        /// <summary>
        /// The cluster tier.
        /// </summary>
        [Input("tier")]
        public InputUnion<string, Pulumi.AzureNative.HDInsight.Tier>? Tier { get; set; }

        public ClusterCreatePropertiesArgs()
        {
            Tier = "Standard";
        }
        public static new ClusterCreatePropertiesArgs Empty => new ClusterCreatePropertiesArgs();
    }
}
