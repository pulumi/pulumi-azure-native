// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20210301Preview.Outputs
{

    /// <summary>
    /// Model asset version details.
    /// </summary>
    [OutputType]
    public sealed class ModelVersionResponse
    {
        /// <summary>
        /// ARM resource ID of the datastore where the asset is located.
        /// </summary>
        public readonly string? DatastoreId;
        /// <summary>
        /// The asset description text.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Mapping of model flavors to their properties.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.FlavorDataResponse>? Flavors;
        /// <summary>
        /// If the name version are system generated (anonymous registration).
        /// </summary>
        public readonly bool? IsAnonymous;
        /// <summary>
        /// [Required] The path of the file/directory in the datastore.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// The asset property dictionary.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Properties;
        /// <summary>
        /// Tag dictionary. Tags can be added, removed, and updated.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private ModelVersionResponse(
            string? datastoreId,

            string? description,

            ImmutableDictionary<string, Outputs.FlavorDataResponse>? flavors,

            bool? isAnonymous,

            string path,

            ImmutableDictionary<string, string>? properties,

            ImmutableDictionary<string, string>? tags)
        {
            DatastoreId = datastoreId;
            Description = description;
            Flavors = flavors;
            IsAnonymous = isAnonymous;
            Path = path;
            Properties = properties;
            Tags = tags;
        }
    }
}
