// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DocumentDB.V20240901Preview.Inputs
{

    /// <summary>
    /// Configuration values for periodic mode backup
    /// </summary>
    public sealed class ContinuousModePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enum to indicate type of Continuos backup mode
        /// </summary>
        [Input("tier")]
        public InputUnion<string, Pulumi.AzureNative.DocumentDB.V20240901Preview.ContinuousTier>? Tier { get; set; }

        public ContinuousModePropertiesArgs()
        {
        }
        public static new ContinuousModePropertiesArgs Empty => new ContinuousModePropertiesArgs();
    }
}
