// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of awsEksNodegroup
    /// </summary>
    public sealed class AwsEksNodegroupPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AMI type for your node group.
        /// </summary>
        [Input("amiType")]
        public Input<string>? AmiType { get; set; }

        /// <summary>
        /// Property arn
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The capacity type of your managed node group.
        /// </summary>
        [Input("capacityType")]
        public Input<string>? CapacityType { get; set; }

        /// <summary>
        /// Name of the cluster to create the node group in.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// The root device disk size (in GiB) for your node group instances.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
        /// </summary>
        [Input("forceUpdateEnabled")]
        public Input<bool>? ForceUpdateEnabled { get; set; }

        /// <summary>
        /// Property id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("instanceTypes")]
        private InputList<string>? _instanceTypes;

        /// <summary>
        /// Specify the instance types for a node group.
        /// </summary>
        public InputList<string> InstanceTypes
        {
            get => _instanceTypes ?? (_instanceTypes = new InputList<string>());
            set => _instanceTypes = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The Kubernetes labels to be applied to the nodes in the node group when they are created.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// An object representing a node group's launch template specification. An object representing a launch template specification for AWS EKS Nodegroup.
        /// </summary>
        [Input("launchTemplate")]
        public Input<Inputs.LaunchTemplateSpecificationArgs>? LaunchTemplate { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
        /// </summary>
        [Input("nodeRole")]
        public Input<string>? NodeRole { get; set; }

        /// <summary>
        /// The unique name to give your node group.
        /// </summary>
        [Input("nodegroupName")]
        public Input<string>? NodegroupName { get; set; }

        /// <summary>
        /// The AMI version of the Amazon EKS-optimized AMI to use with your node group.
        /// </summary>
        [Input("releaseVersion")]
        public Input<string>? ReleaseVersion { get; set; }

        /// <summary>
        /// The remote access (SSH) configuration to use with your node group. An object representing a remote access configuration specification for AWS EKS Nodegroup.
        /// </summary>
        [Input("remoteAccess")]
        public Input<Inputs.RemoteAccessArgs>? RemoteAccess { get; set; }

        /// <summary>
        /// The scaling configuration details for the Auto Scaling group that is created for your node group. An object representing a auto scaling group specification for AWS EKS Nodegroup.
        /// </summary>
        [Input("scalingConfig")]
        public Input<Inputs.ScalingConfigArgs>? ScalingConfig { get; set; }

        [Input("subnets")]
        private InputList<string>? _subnets;

        /// <summary>
        /// The subnets to use for the Auto Scaling group that is created for your node group.
        /// </summary>
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("taints")]
        private InputList<Inputs.TaintArgs>? _taints;

        /// <summary>
        /// The Kubernetes taints to be applied to the nodes in the node group when they are created.
        /// </summary>
        public InputList<Inputs.TaintArgs> Taints
        {
            get => _taints ?? (_taints = new InputList<Inputs.TaintArgs>());
            set => _taints = value;
        }

        /// <summary>
        /// The node group update configuration. The node group update configuration.
        /// </summary>
        [Input("updateConfig")]
        public Input<Inputs.UpdateConfigArgs>? UpdateConfig { get; set; }

        /// <summary>
        /// The Kubernetes version to use for your managed nodes.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public AwsEksNodegroupPropertiesArgs()
        {
        }
        public static new AwsEksNodegroupPropertiesArgs Empty => new AwsEksNodegroupPropertiesArgs();
    }
}
