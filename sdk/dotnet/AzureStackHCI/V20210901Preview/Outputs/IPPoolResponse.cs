// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.V20210901Preview.Outputs
{

    [OutputType]
    public sealed class IPPoolResponse
    {
        /// <summary>
        /// end of the ip address pool
        /// </summary>
        public readonly string? End;
        public readonly Outputs.IPPoolInfoResponse? Info;
        /// <summary>
        /// ip pool type
        /// </summary>
        public readonly string? IpPoolType;
        /// <summary>
        /// start of the ip address pool
        /// </summary>
        public readonly string? Start;

        [OutputConstructor]
        private IPPoolResponse(
            string? end,

            Outputs.IPPoolInfoResponse? info,

            string? ipPoolType,

            string? start)
        {
            End = end;
            Info = info;
            IpPoolType = ipPoolType;
            Start = start;
        }
    }
}
