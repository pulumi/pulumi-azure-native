// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DocumentDB.V20240901Preview.Outputs
{

    /// <summary>
    /// Materialized View definition for the container.
    /// </summary>
    [OutputType]
    public sealed class MaterializedViewDefinitionResponse
    {
        /// <summary>
        /// The definition should be an SQL query which would be used to fetch data from the source container to populate into the Materialized View container.
        /// </summary>
        public readonly string Definition;
        /// <summary>
        /// The name of the source container on which the Materialized View will be created.
        /// </summary>
        public readonly string SourceCollectionId;
        /// <summary>
        /// An unique identifier for the source collection. This is a system generated property.
        /// </summary>
        public readonly string SourceCollectionRid;

        [OutputConstructor]
        private MaterializedViewDefinitionResponse(
            string definition,

            string sourceCollectionId,

            string sourceCollectionRid)
        {
            Definition = definition;
            SourceCollectionId = sourceCollectionId;
            SourceCollectionRid = sourceCollectionRid;
        }
    }
}
