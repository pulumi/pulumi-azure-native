// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.LabServices.V20181015
{
    /// <summary>
    /// Enum indicating if user is adding or removing a favorite lab
    /// </summary>
    [EnumType]
    public readonly struct AddRemove : IEquatable<AddRemove>
    {
        private readonly string _value;

        private AddRemove(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Indicates that a user is adding a favorite lab
        /// </summary>
        public static AddRemove Add { get; } = new AddRemove("Add");
        /// <summary>
        /// Indicates that a user is removing a favorite lab
        /// </summary>
        public static AddRemove Remove { get; } = new AddRemove("Remove");

        public static bool operator ==(AddRemove left, AddRemove right) => left.Equals(right);
        public static bool operator !=(AddRemove left, AddRemove right) => !left.Equals(right);

        public static explicit operator string(AddRemove value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AddRemove other && Equals(other);
        public bool Equals(AddRemove other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Describes the user's progress in configuring their environment setting
    /// </summary>
    [EnumType]
    public readonly struct ConfigurationState : IEquatable<ConfigurationState>
    {
        private readonly string _value;

        private ConfigurationState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// User either hasn't started configuring their template
        /// or they haven't started the configuration process.
        /// </summary>
        public static ConfigurationState NotApplicable { get; } = new ConfigurationState("NotApplicable");
        /// <summary>
        /// User is finished modifying the template.
        /// </summary>
        public static ConfigurationState Completed { get; } = new ConfigurationState("Completed");

        public static bool operator ==(ConfigurationState left, ConfigurationState right) => left.Equals(right);
        public static bool operator !=(ConfigurationState left, ConfigurationState right) => !left.Equals(right);

        public static explicit operator string(ConfigurationState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConfigurationState other && Equals(other);
        public bool Equals(ConfigurationState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Lab user access mode (open to all vs. restricted to those listed on the lab).
    /// </summary>
    [EnumType]
    public readonly struct LabUserAccessMode : IEquatable<LabUserAccessMode>
    {
        private readonly string _value;

        private LabUserAccessMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Only users registered with the lab can access VMs.
        /// </summary>
        public static LabUserAccessMode Restricted { get; } = new LabUserAccessMode("Restricted");
        /// <summary>
        /// Any user can register with the lab and access its VMs.
        /// </summary>
        public static LabUserAccessMode Open { get; } = new LabUserAccessMode("Open");

        public static bool operator ==(LabUserAccessMode left, LabUserAccessMode right) => left.Equals(right);
        public static bool operator !=(LabUserAccessMode left, LabUserAccessMode right) => !left.Equals(right);

        public static explicit operator string(LabUserAccessMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LabUserAccessMode other && Equals(other);
        public bool Equals(LabUserAccessMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The size of the virtual machine
    /// </summary>
    [EnumType]
    public readonly struct ManagedLabVmSize : IEquatable<ManagedLabVmSize>
    {
        private readonly string _value;

        private ManagedLabVmSize(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The base VM size
        /// </summary>
        public static ManagedLabVmSize Basic { get; } = new ManagedLabVmSize("Basic");
        /// <summary>
        /// The standard or default VM size
        /// </summary>
        public static ManagedLabVmSize Standard { get; } = new ManagedLabVmSize("Standard");
        /// <summary>
        /// The most performant VM size
        /// </summary>
        public static ManagedLabVmSize Performance { get; } = new ManagedLabVmSize("Performance");

        public static bool operator ==(ManagedLabVmSize left, ManagedLabVmSize right) => left.Equals(right);
        public static bool operator !=(ManagedLabVmSize left, ManagedLabVmSize right) => !left.Equals(right);

        public static explicit operator string(ManagedLabVmSize value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ManagedLabVmSize other && Equals(other);
        public bool Equals(ManagedLabVmSize other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
