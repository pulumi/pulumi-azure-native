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
    /// Parameters for Kubernetes Cluster Backup Datasource
    /// </summary>
    public sealed class KubernetesClusterBackupDatasourceParametersArgs : global::Pulumi.ResourceArgs
    {
        [Input("backupHookReferences")]
        private InputList<Inputs.NamespacedNameResourceArgs>? _backupHookReferences;

        /// <summary>
        /// Gets or sets the backup hook references. This property sets the hook reference to be executed during backup.
        /// </summary>
        public InputList<Inputs.NamespacedNameResourceArgs> BackupHookReferences
        {
            get => _backupHookReferences ?? (_backupHookReferences = new InputList<Inputs.NamespacedNameResourceArgs>());
            set => _backupHookReferences = value;
        }

        [Input("excludedNamespaces")]
        private InputList<string>? _excludedNamespaces;

        /// <summary>
        /// Gets or sets the exclude namespaces property. This property sets the namespaces to be excluded during backup.
        /// </summary>
        public InputList<string> ExcludedNamespaces
        {
            get => _excludedNamespaces ?? (_excludedNamespaces = new InputList<string>());
            set => _excludedNamespaces = value;
        }

        [Input("excludedResourceTypes")]
        private InputList<string>? _excludedResourceTypes;

        /// <summary>
        /// Gets or sets the exclude resource types property. This property sets the resource types to be excluded during backup.
        /// </summary>
        public InputList<string> ExcludedResourceTypes
        {
            get => _excludedResourceTypes ?? (_excludedResourceTypes = new InputList<string>());
            set => _excludedResourceTypes = value;
        }

        /// <summary>
        /// Gets or sets the include cluster resources property. This property if enabled will include cluster scope resources during backup.
        /// </summary>
        [Input("includeClusterScopeResources", required: true)]
        public Input<bool> IncludeClusterScopeResources { get; set; } = null!;

        [Input("includedNamespaces")]
        private InputList<string>? _includedNamespaces;

        /// <summary>
        /// Gets or sets the include namespaces property. This property sets the namespaces to be included during backup.
        /// </summary>
        public InputList<string> IncludedNamespaces
        {
            get => _includedNamespaces ?? (_includedNamespaces = new InputList<string>());
            set => _includedNamespaces = value;
        }

        [Input("includedResourceTypes")]
        private InputList<string>? _includedResourceTypes;

        /// <summary>
        /// Gets or sets the include resource types property. This property sets the resource types to be included during backup.
        /// </summary>
        public InputList<string> IncludedResourceTypes
        {
            get => _includedResourceTypes ?? (_includedResourceTypes = new InputList<string>());
            set => _includedResourceTypes = value;
        }

        [Input("includedVolumeTypes")]
        private InputList<Union<string, Pulumi.AzureNative.DataProtection.AKSVolumeTypes>>? _includedVolumeTypes;

        /// <summary>
        /// Gets or sets the include volume types property. This property sets the volume types to be included during backup.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.DataProtection.AKSVolumeTypes>> IncludedVolumeTypes
        {
            get => _includedVolumeTypes ?? (_includedVolumeTypes = new InputList<Union<string, Pulumi.AzureNative.DataProtection.AKSVolumeTypes>>());
            set => _includedVolumeTypes = value;
        }

        [Input("labelSelectors")]
        private InputList<string>? _labelSelectors;

        /// <summary>
        /// Gets or sets the LabelSelectors property. This property sets the resource with such label selectors to be included during backup.
        /// </summary>
        public InputList<string> LabelSelectors
        {
            get => _labelSelectors ?? (_labelSelectors = new InputList<string>());
            set => _labelSelectors = value;
        }

        /// <summary>
        /// Type of the specific object - used for deserializing
        /// Expected value is 'KubernetesClusterBackupDatasourceParameters'.
        /// </summary>
        [Input("objectType", required: true)]
        public Input<string> ObjectType { get; set; } = null!;

        /// <summary>
        /// Gets or sets the volume snapshot property. This property if enabled will take volume snapshots during backup.
        /// </summary>
        [Input("snapshotVolumes", required: true)]
        public Input<bool> SnapshotVolumes { get; set; } = null!;

        public KubernetesClusterBackupDatasourceParametersArgs()
        {
        }
        public static new KubernetesClusterBackupDatasourceParametersArgs Empty => new KubernetesClusterBackupDatasourceParametersArgs();
    }
}
