// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerInstance.Inputs
{

    /// <summary>
    /// A container instance.
    /// </summary>
    public sealed class ContainerArgs : global::Pulumi.ResourceArgs
    {
        [Input("command")]
        private InputList<string>? _command;

        /// <summary>
        /// The commands to execute within the container instance in exec form.
        /// </summary>
        public InputList<string> Command
        {
            get => _command ?? (_command = new InputList<string>());
            set => _command = value;
        }

        /// <summary>
        /// The config map.
        /// </summary>
        [Input("configMap")]
        public Input<Inputs.ConfigMapArgs>? ConfigMap { get; set; }

        [Input("environmentVariables")]
        private InputList<Inputs.EnvironmentVariableArgs>? _environmentVariables;

        /// <summary>
        /// The environment variables to set in the container instance.
        /// </summary>
        public InputList<Inputs.EnvironmentVariableArgs> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputList<Inputs.EnvironmentVariableArgs>());
            set => _environmentVariables = value;
        }

        /// <summary>
        /// The name of the image used to create the container instance.
        /// </summary>
        [Input("image")]
        public Input<string>? Image { get; set; }

        /// <summary>
        /// The liveness probe.
        /// </summary>
        [Input("livenessProbe")]
        public Input<Inputs.ContainerProbeArgs>? LivenessProbe { get; set; }

        /// <summary>
        /// The user-provided name of the container instance.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("ports")]
        private InputList<Inputs.ContainerPortArgs>? _ports;

        /// <summary>
        /// The exposed ports on the container instance.
        /// </summary>
        public InputList<Inputs.ContainerPortArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.ContainerPortArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// The readiness probe.
        /// </summary>
        [Input("readinessProbe")]
        public Input<Inputs.ContainerProbeArgs>? ReadinessProbe { get; set; }

        /// <summary>
        /// The resource requirements of the container instance.
        /// </summary>
        [Input("resources")]
        public Input<Inputs.ResourceRequirementsArgs>? Resources { get; set; }

        /// <summary>
        /// The container security properties.
        /// </summary>
        [Input("securityContext")]
        public Input<Inputs.SecurityContextDefinitionArgs>? SecurityContext { get; set; }

        [Input("volumeMounts")]
        private InputList<Inputs.VolumeMountArgs>? _volumeMounts;

        /// <summary>
        /// The volume mounts available to the container instance.
        /// </summary>
        public InputList<Inputs.VolumeMountArgs> VolumeMounts
        {
            get => _volumeMounts ?? (_volumeMounts = new InputList<Inputs.VolumeMountArgs>());
            set => _volumeMounts = value;
        }

        public ContainerArgs()
        {
        }
        public static new ContainerArgs Empty => new ContainerArgs();
    }
}
