// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DocumentDB.V20241201Preview.Inputs
{

    /// <summary>
    /// Materialized View definition for the container.
    /// </summary>
    public sealed class MaterializedViewDefinitionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The definition should be an SQL query which would be used to fetch data from the source container to populate into the Materialized View container.
        /// </summary>
        [Input("definition", required: true)]
        public Input<string> Definition { get; set; } = null!;

        /// <summary>
        /// The name of the source container on which the Materialized View will be created.
        /// </summary>
        [Input("sourceCollectionId", required: true)]
        public Input<string> SourceCollectionId { get; set; } = null!;

        public MaterializedViewDefinitionArgs()
        {
        }
        public static new MaterializedViewDefinitionArgs Empty => new MaterializedViewDefinitionArgs();
    }
}
