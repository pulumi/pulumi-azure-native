// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20230601Preview.Inputs
{

    public sealed class HdfsDatastoreArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Required] Account credentials.
        /// </summary>
        [Input("credentials", required: true)]
        public object Credentials { get; set; } = null!;

        /// <summary>
        /// Enum to determine the datastore contents type.
        /// Expected value is 'Hdfs'.
        /// </summary>
        [Input("datastoreType", required: true)]
        public Input<string> DatastoreType { get; set; } = null!;

        /// <summary>
        /// The asset description text.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The TLS cert of the HDFS server. Needs to be a base64 encoded string. Required if "Https" protocol is selected.
        /// </summary>
        [Input("hdfsServerCertificate")]
        public Input<string>? HdfsServerCertificate { get; set; }

        /// <summary>
        /// Intellectual Property details.
        /// </summary>
        [Input("intellectualProperty")]
        public Input<Inputs.IntellectualPropertyArgs>? IntellectualProperty { get; set; }

        /// <summary>
        /// [Required] IP Address or DNS HostName.
        /// </summary>
        [Input("nameNodeAddress", required: true)]
        public Input<string> NameNodeAddress { get; set; } = null!;

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// The asset property dictionary.
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        /// <summary>
        /// Protocol used to communicate with the storage account (Https/Http).
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tag dictionary. Tags can be added, removed, and updated.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public HdfsDatastoreArgs()
        {
            Protocol = "http";
        }
        public static new HdfsDatastoreArgs Empty => new HdfsDatastoreArgs();
    }
}
