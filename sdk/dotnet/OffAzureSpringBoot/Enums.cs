// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.OffAzureSpringBoot
{
    /// <summary>
    /// The resource provisioning state.
    /// </summary>
    [EnumType]
    public readonly struct ProvisioningState : IEquatable<ProvisioningState>
    {
        private readonly string _value;

        private ProvisioningState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ProvisioningState Unknown { get; } = new ProvisioningState("Unknown");
        public static ProvisioningState Succeeded { get; } = new ProvisioningState("Succeeded");
        public static ProvisioningState Failed { get; } = new ProvisioningState("Failed");
        public static ProvisioningState Canceled { get; } = new ProvisioningState("Canceled");
        public static ProvisioningState Accepted { get; } = new ProvisioningState("Accepted");
        public static ProvisioningState Provisioning { get; } = new ProvisioningState("Provisioning");
        public static ProvisioningState Deleting { get; } = new ProvisioningState("Deleting");

        public static bool operator ==(ProvisioningState left, ProvisioningState right) => left.Equals(right);
        public static bool operator !=(ProvisioningState left, ProvisioningState right) => !left.Equals(right);

        public static explicit operator string(ProvisioningState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProvisioningState other && Equals(other);
        public bool Equals(ProvisioningState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
