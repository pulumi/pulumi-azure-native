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
    /// Azure backup rule
    /// </summary>
    public sealed class AzureBackupRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// BackupParameters base
        /// </summary>
        [Input("backupParameters")]
        public Input<Inputs.AzureBackupParamsArgs>? BackupParameters { get; set; }

        /// <summary>
        /// DataStoreInfo base
        /// </summary>
        [Input("dataStore", required: true)]
        public Input<Inputs.DataStoreInfoBaseArgs> DataStore { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// 
        /// Expected value is 'AzureBackupRule'.
        /// </summary>
        [Input("objectType", required: true)]
        public Input<string> ObjectType { get; set; } = null!;

        /// <summary>
        /// Trigger context
        /// </summary>
        [Input("trigger", required: true)]
        public InputUnion<Inputs.AdhocBasedTriggerContextArgs, Inputs.ScheduleBasedTriggerContextArgs> Trigger { get; set; } = null!;

        public AzureBackupRuleArgs()
        {
        }
        public static new AzureBackupRuleArgs Empty => new AzureBackupRuleArgs();
    }
}
