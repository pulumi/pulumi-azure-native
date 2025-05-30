// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.Maintenance
{
    /// <summary>
    /// Gets or sets maintenanceScope of the configuration
    /// </summary>
    [EnumType]
    public readonly struct MaintenanceScope : IEquatable<MaintenanceScope>
    {
        private readonly string _value;

        private MaintenanceScope(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// This maintenance scope controls installation of azure platform updates i.e. services on physical nodes hosting customer VMs.
        /// </summary>
        public static MaintenanceScope Host { get; } = new MaintenanceScope("Host");
        /// <summary>
        /// This maintenance scope controls the default update maintenance of the Azure Resource
        /// </summary>
        public static MaintenanceScope Resource { get; } = new MaintenanceScope("Resource");
        /// <summary>
        /// This maintenance scope controls os image installation on VM/VMSS
        /// </summary>
        public static MaintenanceScope OSImage { get; } = new MaintenanceScope("OSImage");
        /// <summary>
        /// This maintenance scope controls extension installation on VM/VMSS
        /// </summary>
        public static MaintenanceScope Extension { get; } = new MaintenanceScope("Extension");
        /// <summary>
        /// This maintenance scope controls installation of windows and linux packages on VM/VMSS
        /// </summary>
        public static MaintenanceScope InGuestPatch { get; } = new MaintenanceScope("InGuestPatch");
        /// <summary>
        /// This maintenance scope controls installation of SQL server platform updates.
        /// </summary>
        public static MaintenanceScope SQLDB { get; } = new MaintenanceScope("SQLDB");
        /// <summary>
        /// This maintenance scope controls installation of SQL managed instance platform update.
        /// </summary>
        public static MaintenanceScope SQLManagedInstance { get; } = new MaintenanceScope("SQLManagedInstance");

        public static bool operator ==(MaintenanceScope left, MaintenanceScope right) => left.Equals(right);
        public static bool operator !=(MaintenanceScope left, MaintenanceScope right) => !left.Equals(right);

        public static explicit operator string(MaintenanceScope value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MaintenanceScope other && Equals(other);
        public bool Equals(MaintenanceScope other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Possible reboot preference as defined by the user based on which it would be decided to reboot the machine or not after the patch operation is completed.
    /// </summary>
    [EnumType]
    public readonly struct RebootOptions : IEquatable<RebootOptions>
    {
        private readonly string _value;

        private RebootOptions(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RebootOptions IfRequired { get; } = new RebootOptions("IfRequired");
        public static RebootOptions Never { get; } = new RebootOptions("Never");
        public static RebootOptions Always { get; } = new RebootOptions("Always");

        public static bool operator ==(RebootOptions left, RebootOptions right) => left.Equals(right);
        public static bool operator !=(RebootOptions left, RebootOptions right) => !left.Equals(right);

        public static explicit operator string(RebootOptions value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RebootOptions other && Equals(other);
        public bool Equals(RebootOptions other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Filter VMs by Any or All specified tags.
    /// </summary>
    [EnumType]
    public readonly struct TagOperators : IEquatable<TagOperators>
    {
        private readonly string _value;

        private TagOperators(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TagOperators All { get; } = new TagOperators("All");
        public static TagOperators Any { get; } = new TagOperators("Any");

        public static bool operator ==(TagOperators left, TagOperators right) => left.Equals(right);
        public static bool operator !=(TagOperators left, TagOperators right) => !left.Equals(right);

        public static explicit operator string(TagOperators value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TagOperators other && Equals(other);
        public bool Equals(TagOperators other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the visibility of the configuration. The default value is 'Custom'
    /// </summary>
    [EnumType]
    public readonly struct Visibility : IEquatable<Visibility>
    {
        private readonly string _value;

        private Visibility(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Only visible to users with permissions.
        /// </summary>
        public static Visibility Custom { get; } = new Visibility("Custom");
        /// <summary>
        /// Visible to all users.
        /// </summary>
        public static Visibility Public { get; } = new Visibility("Public");

        public static bool operator ==(Visibility left, Visibility right) => left.Equals(right);
        public static bool operator !=(Visibility left, Visibility right) => !left.Equals(right);

        public static explicit operator string(Visibility value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Visibility other && Equals(other);
        public bool Equals(Visibility other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
