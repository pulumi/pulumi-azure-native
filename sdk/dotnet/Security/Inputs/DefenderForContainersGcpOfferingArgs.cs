// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Inputs
{

    /// <summary>
    /// The containers GCP offering
    /// </summary>
    public sealed class DefenderForContainersGcpOfferingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The native cloud connection configuration
        /// </summary>
        [Input("dataPipelineNativeCloudConnection")]
        public Input<Inputs.DefenderForContainersGcpOfferingDataPipelineNativeCloudConnectionArgs>? DataPipelineNativeCloudConnection { get; set; }

        /// <summary>
        /// Is audit logs data collection enabled
        /// </summary>
        [Input("enableAuditLogsAutoProvisioning")]
        public Input<bool>? EnableAuditLogsAutoProvisioning { get; set; }

        /// <summary>
        /// Is Microsoft Defender for Cloud Kubernetes agent auto provisioning enabled
        /// </summary>
        [Input("enableDefenderAgentAutoProvisioning")]
        public Input<bool>? EnableDefenderAgentAutoProvisioning { get; set; }

        /// <summary>
        /// Is Policy Kubernetes agent auto provisioning enabled
        /// </summary>
        [Input("enablePolicyAgentAutoProvisioning")]
        public Input<bool>? EnablePolicyAgentAutoProvisioning { get; set; }

        /// <summary>
        /// The Microsoft Defender Container agentless discovery configuration
        /// </summary>
        [Input("mdcContainersAgentlessDiscoveryK8s")]
        public Input<Inputs.DefenderForContainersGcpOfferingMdcContainersAgentlessDiscoveryK8sArgs>? MdcContainersAgentlessDiscoveryK8s { get; set; }

        /// <summary>
        /// The Microsoft Defender Container image assessment configuration
        /// </summary>
        [Input("mdcContainersImageAssessment")]
        public Input<Inputs.DefenderForContainersGcpOfferingMdcContainersImageAssessmentArgs>? MdcContainersImageAssessment { get; set; }

        /// <summary>
        /// The native cloud connection configuration
        /// </summary>
        [Input("nativeCloudConnection")]
        public Input<Inputs.DefenderForContainersGcpOfferingNativeCloudConnectionArgs>? NativeCloudConnection { get; set; }

        /// <summary>
        /// The type of the security offering.
        /// Expected value is 'DefenderForContainersGcp'.
        /// </summary>
        [Input("offeringType", required: true)]
        public Input<string> OfferingType { get; set; } = null!;

        /// <summary>
        /// The Microsoft Defender for Container K8s VM host scanning configuration
        /// </summary>
        [Input("vmScanners")]
        public Input<Inputs.DefenderForContainersGcpOfferingVmScannersArgs>? VmScanners { get; set; }

        public DefenderForContainersGcpOfferingArgs()
        {
        }
        public static new DefenderForContainersGcpOfferingArgs Empty => new DefenderForContainersGcpOfferingArgs();
    }
}
