// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MobileNetwork.Outputs
{

    /// <summary>
    /// Data flow template
    /// </summary>
    [OutputType]
    public sealed class ServiceDataFlowTemplateResponse
    {
        /// <summary>
        /// The direction of this flow.
        /// </summary>
        public readonly string Direction;
        /// <summary>
        /// The port(s) to which UEs will connect for this flow. You can specify zero or more ports or port ranges. If you specify one or more ports or port ranges then you must specify a value other than `ip` in the `protocol` field. This is an optional setting. If you do not specify it then connections will be allowed on all ports. Port ranges must be specified as &lt;FirstPort&gt;-&lt;LastPort&gt;. For example: [`8080`, `8082-8085`].
        /// </summary>
        public readonly ImmutableArray<string> Ports;
        /// <summary>
        /// A list of the allowed protocol(s) for this flow. If you want this flow to be able to use any protocol within the internet protocol suite, use the value `ip`. If you only want to allow a selection of protocols, you must use the corresponding IANA Assigned Internet Protocol Number for each protocol, as described in https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml. For example, for UDP, you must use 17. If you use the value `ip` then you must leave the field `port` unspecified.
        /// </summary>
        public readonly ImmutableArray<string> Protocol;
        /// <summary>
        /// The remote IP address(es) to which UEs will connect for this flow. If you want to allow connections on any IP address, use the value `any`. Otherwise, you must provide each of the remote IP addresses to which the packet core instance will connect for this flow. You must provide each IP address in CIDR notation, including the netmask (for example, 192.0.2.54/24).
        /// </summary>
        public readonly ImmutableArray<string> RemoteIpList;
        /// <summary>
        /// The name of the data flow template. This must be unique within the parent data flow policy rule. You must not use any of the following reserved strings - `default`, `requested` or `service`.
        /// </summary>
        public readonly string TemplateName;

        [OutputConstructor]
        private ServiceDataFlowTemplateResponse(
            string direction,

            ImmutableArray<string> ports,

            ImmutableArray<string> protocol,

            ImmutableArray<string> remoteIpList,

            string templateName)
        {
            Direction = direction;
            Ports = ports;
            Protocol = protocol;
            RemoteIpList = remoteIpList;
            TemplateName = templateName;
        }
    }
}
