// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of awsLightsailInstance
    /// </summary>
    [OutputType]
    public sealed class AwsLightsailInstancePropertiesResponse
    {
        /// <summary>
        /// An array of objects representing the add-ons to enable for the new instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.AddOnResponse> AddOns;
        /// <summary>
        /// The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.
        /// </summary>
        public readonly string? AvailabilityZone;
        /// <summary>
        /// The ID for a virtual private server image (e.g., app_wordpress_4_4 or app_lamp_7_0 ). Use the get blueprints operation to return a list of available images (or blueprints ).
        /// </summary>
        public readonly string? BlueprintId;
        /// <summary>
        /// The bundle of specification information for your virtual private server (or instance ), including the pricing plan (e.g., micro_1_0 ).
        /// </summary>
        public readonly string? BundleId;
        /// <summary>
        /// Hardware of the Instance.
        /// </summary>
        public readonly Outputs.HardwareResponse? Hardware;
        /// <summary>
        /// Property instanceArn
        /// </summary>
        public readonly string? InstanceArn;
        /// <summary>
        /// The names to use for your new Lightsail instance.
        /// </summary>
        public readonly string? InstanceName;
        /// <summary>
        /// Is the IP Address of the Instance is the static IP
        /// </summary>
        public readonly bool? IsStaticIp;
        /// <summary>
        /// The name of your key pair.
        /// </summary>
        public readonly string? KeyPairName;
        /// <summary>
        /// Location of a resource.
        /// </summary>
        public readonly Outputs.LocationResponse? Location;
        /// <summary>
        /// Networking of the Instance.
        /// </summary>
        public readonly Outputs.NetworkingResponse? Networking;
        /// <summary>
        /// Private IP Address of the Instance
        /// </summary>
        public readonly string? PrivateIpAddress;
        /// <summary>
        /// Public IP Address of the Instance
        /// </summary>
        public readonly string? PublicIpAddress;
        /// <summary>
        /// Resource type of Lightsail instance.
        /// </summary>
        public readonly string? ResourceType;
        /// <summary>
        /// SSH Key Name of the  Lightsail instance.
        /// </summary>
        public readonly string? SshKeyName;
        /// <summary>
        /// Current State of the Instance.
        /// </summary>
        public readonly Outputs.StateResponse? State;
        /// <summary>
        /// Support code to help identify any issues
        /// </summary>
        public readonly string? SupportCode;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.TagResponse> Tags;
        /// <summary>
        /// A launch script you can create that configures a server with additional user data. For example, you might want to run apt-get -y update.
        /// </summary>
        public readonly string? UserData;
        /// <summary>
        /// Username of the  Lightsail instance.
        /// </summary>
        public readonly string? UserName;

        [OutputConstructor]
        private AwsLightsailInstancePropertiesResponse(
            ImmutableArray<Outputs.AddOnResponse> addOns,

            string? availabilityZone,

            string? blueprintId,

            string? bundleId,

            Outputs.HardwareResponse? hardware,

            string? instanceArn,

            string? instanceName,

            bool? isStaticIp,

            string? keyPairName,

            Outputs.LocationResponse? location,

            Outputs.NetworkingResponse? networking,

            string? privateIpAddress,

            string? publicIpAddress,

            string? resourceType,

            string? sshKeyName,

            Outputs.StateResponse? state,

            string? supportCode,

            ImmutableArray<Outputs.TagResponse> tags,

            string? userData,

            string? userName)
        {
            AddOns = addOns;
            AvailabilityZone = availabilityZone;
            BlueprintId = blueprintId;
            BundleId = bundleId;
            Hardware = hardware;
            InstanceArn = instanceArn;
            InstanceName = instanceName;
            IsStaticIp = isStaticIp;
            KeyPairName = keyPairName;
            Location = location;
            Networking = networking;
            PrivateIpAddress = privateIpAddress;
            PublicIpAddress = publicIpAddress;
            ResourceType = resourceType;
            SshKeyName = sshKeyName;
            State = state;
            SupportCode = supportCode;
            Tags = tags;
            UserData = userData;
            UserName = userName;
        }
    }
}
