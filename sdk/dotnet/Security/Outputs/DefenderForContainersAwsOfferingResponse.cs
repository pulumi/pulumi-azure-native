// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Outputs
{

    /// <summary>
    /// The Defender for Containers AWS offering
    /// </summary>
    [OutputType]
    public sealed class DefenderForContainersAwsOfferingResponse
    {
        /// <summary>
        /// The cloudwatch to kinesis connection configuration
        /// </summary>
        public readonly Outputs.DefenderForContainersAwsOfferingResponseCloudWatchToKinesis? CloudWatchToKinesis;
        /// <summary>
        /// The externalId used by the data reader to prevent the confused deputy attack
        /// </summary>
        public readonly string? DataCollectionExternalId;
        /// <summary>
        /// The offering description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Is audit logs data collection enabled
        /// </summary>
        public readonly bool? EnableAuditLogsAutoProvisioning;
        /// <summary>
        /// Is Microsoft Defender for Cloud Kubernetes agent auto provisioning enabled
        /// </summary>
        public readonly bool? EnableDefenderAgentAutoProvisioning;
        /// <summary>
        /// Is Policy Kubernetes agent auto provisioning enabled
        /// </summary>
        public readonly bool? EnablePolicyAgentAutoProvisioning;
        /// <summary>
        /// The kinesis to s3 connection configuration
        /// </summary>
        public readonly Outputs.DefenderForContainersAwsOfferingResponseKinesisToS3? KinesisToS3;
        /// <summary>
        /// The retention time in days of kube audit logs set on the CloudWatch log group
        /// </summary>
        public readonly double? KubeAuditRetentionTime;
        /// <summary>
        /// The kubernetes data collection connection configuration
        /// </summary>
        public readonly Outputs.DefenderForContainersAwsOfferingResponseKubernetesDataCollection? KubernetesDataCollection;
        /// <summary>
        /// The kubernetes service connection configuration
        /// </summary>
        public readonly Outputs.DefenderForContainersAwsOfferingResponseKubernetesService? KubernetesService;
        /// <summary>
        /// The Microsoft Defender container agentless discovery K8s configuration
        /// </summary>
        public readonly Outputs.DefenderForContainersAwsOfferingResponseMdcContainersAgentlessDiscoveryK8s? MdcContainersAgentlessDiscoveryK8s;
        /// <summary>
        /// The Microsoft Defender container image assessment configuration
        /// </summary>
        public readonly Outputs.DefenderForContainersAwsOfferingResponseMdcContainersImageAssessment? MdcContainersImageAssessment;
        /// <summary>
        /// The type of the security offering.
        /// Expected value is 'DefenderForContainersAws'.
        /// </summary>
        public readonly string OfferingType;
        /// <summary>
        /// The Microsoft Defender for Container K8s VM host scanning configuration
        /// </summary>
        public readonly Outputs.DefenderForContainersAwsOfferingResponseVmScanners? VmScanners;

        [OutputConstructor]
        private DefenderForContainersAwsOfferingResponse(
            Outputs.DefenderForContainersAwsOfferingResponseCloudWatchToKinesis? cloudWatchToKinesis,

            string? dataCollectionExternalId,

            string description,

            bool? enableAuditLogsAutoProvisioning,

            bool? enableDefenderAgentAutoProvisioning,

            bool? enablePolicyAgentAutoProvisioning,

            Outputs.DefenderForContainersAwsOfferingResponseKinesisToS3? kinesisToS3,

            double? kubeAuditRetentionTime,

            Outputs.DefenderForContainersAwsOfferingResponseKubernetesDataCollection? kubernetesDataCollection,

            Outputs.DefenderForContainersAwsOfferingResponseKubernetesService? kubernetesService,

            Outputs.DefenderForContainersAwsOfferingResponseMdcContainersAgentlessDiscoveryK8s? mdcContainersAgentlessDiscoveryK8s,

            Outputs.DefenderForContainersAwsOfferingResponseMdcContainersImageAssessment? mdcContainersImageAssessment,

            string offeringType,

            Outputs.DefenderForContainersAwsOfferingResponseVmScanners? vmScanners)
        {
            CloudWatchToKinesis = cloudWatchToKinesis;
            DataCollectionExternalId = dataCollectionExternalId;
            Description = description;
            EnableAuditLogsAutoProvisioning = enableAuditLogsAutoProvisioning;
            EnableDefenderAgentAutoProvisioning = enableDefenderAgentAutoProvisioning;
            EnablePolicyAgentAutoProvisioning = enablePolicyAgentAutoProvisioning;
            KinesisToS3 = kinesisToS3;
            KubeAuditRetentionTime = kubeAuditRetentionTime;
            KubernetesDataCollection = kubernetesDataCollection;
            KubernetesService = kubernetesService;
            MdcContainersAgentlessDiscoveryK8s = mdcContainersAgentlessDiscoveryK8s;
            MdcContainersImageAssessment = mdcContainersImageAssessment;
            OfferingType = offeringType;
            VmScanners = vmScanners;
        }
    }
}
