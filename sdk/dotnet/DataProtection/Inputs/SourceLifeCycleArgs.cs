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
    /// Source LifeCycle
    /// </summary>
    public sealed class SourceLifeCycleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Delete Option
        /// </summary>
        [Input("deleteAfter", required: true)]
        public Input<Inputs.AbsoluteDeleteOptionArgs> DeleteAfter { get; set; } = null!;

        /// <summary>
        /// DataStoreInfo base
        /// </summary>
        [Input("sourceDataStore", required: true)]
        public Input<Inputs.DataStoreInfoBaseArgs> SourceDataStore { get; set; } = null!;

        [Input("targetDataStoreCopySettings")]
        private InputList<Inputs.TargetCopySettingArgs>? _targetDataStoreCopySettings;
        public InputList<Inputs.TargetCopySettingArgs> TargetDataStoreCopySettings
        {
            get => _targetDataStoreCopySettings ?? (_targetDataStoreCopySettings = new InputList<Inputs.TargetCopySettingArgs>());
            set => _targetDataStoreCopySettings = value;
        }

        public SourceLifeCycleArgs()
        {
        }
        public static new SourceLifeCycleArgs Empty => new SourceLifeCycleArgs();
    }
}
