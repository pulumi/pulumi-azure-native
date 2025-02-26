// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevCenter.V20241001Preview.Inputs
{

    /// <summary>
    /// Settings to be used when associating a project with a catalog.
    /// </summary>
    public sealed class ProjectCatalogSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("catalogItemSyncTypes")]
        private InputList<Union<string, Pulumi.AzureNative.DevCenter.V20241001Preview.CatalogItemType>>? _catalogItemSyncTypes;

        /// <summary>
        /// Indicates catalog item types that can be synced.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.DevCenter.V20241001Preview.CatalogItemType>> CatalogItemSyncTypes
        {
            get => _catalogItemSyncTypes ?? (_catalogItemSyncTypes = new InputList<Union<string, Pulumi.AzureNative.DevCenter.V20241001Preview.CatalogItemType>>());
            set => _catalogItemSyncTypes = value;
        }

        public ProjectCatalogSettingsArgs()
        {
        }
        public static new ProjectCatalogSettingsArgs Empty => new ProjectCatalogSettingsArgs();
    }
}
