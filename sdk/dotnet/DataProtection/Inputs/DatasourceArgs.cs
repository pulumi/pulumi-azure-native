// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataProtection.Inputs
{

    /// <summary>
    /// Datasource to be backed up
    /// </summary>
    public sealed class DatasourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DatasourceType of the resource.
        /// </summary>
        [Input("datasourceType")]
        public Input<string>? DatasourceType { get; set; }

        /// <summary>
        /// Type of Datasource object, used to initialize the right inherited type
        /// </summary>
        [Input("objectType")]
        public Input<string>? ObjectType { get; set; }

        /// <summary>
        /// Full ARM ID of the resource. For azure resources, this is ARM ID. For non azure resources, this will be the ID created by backup service via Fabric/Vault.
        /// </summary>
        [Input("resourceID", required: true)]
        public Input<string> ResourceID { get; set; } = null!;

        /// <summary>
        /// Location of datasource.
        /// </summary>
        [Input("resourceLocation")]
        public Input<string>? ResourceLocation { get; set; }

        /// <summary>
        /// Unique identifier of the resource in the context of parent.
        /// </summary>
        [Input("resourceName")]
        public Input<string>? ResourceName { get; set; }

        /// <summary>
        /// Properties specific to data source
        /// </summary>
        [Input("resourceProperties")]
        public Input<Inputs.DefaultResourcePropertiesArgs>? ResourceProperties { get; set; }

        /// <summary>
        /// Resource Type of Datasource.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        /// <summary>
        /// Uri of the resource.
        /// </summary>
        [Input("resourceUri")]
        public Input<string>? ResourceUri { get; set; }

        public DatasourceArgs()
        {
        }
        public static new DatasourceArgs Empty => new DatasourceArgs();
    }
}
