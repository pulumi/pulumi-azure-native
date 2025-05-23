// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MobileNetwork.Inputs
{

    /// <summary>
    /// Public land mobile network (PLMN) ID. This is made up of the mobile country code and mobile network code, as defined in https://www.itu.int/rec/T-REC-E.212. The values 001-01 and 001-001 can be used for testing and the values 999-99 and 999-999 can be used on internal private networks.
    /// </summary>
    public sealed class PlmnIdArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mobile country code (MCC).
        /// </summary>
        [Input("mcc", required: true)]
        public Input<string> Mcc { get; set; } = null!;

        /// <summary>
        /// Mobile network code (MNC).
        /// </summary>
        [Input("mnc", required: true)]
        public Input<string> Mnc { get; set; } = null!;

        public PlmnIdArgs()
        {
        }
        public static new PlmnIdArgs Empty => new PlmnIdArgs();
    }
}
