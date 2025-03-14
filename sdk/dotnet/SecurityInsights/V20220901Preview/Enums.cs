// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.SecurityInsights.V20220901Preview
{
    /// <summary>
    /// The entity query kind
    /// </summary>
    [EnumType]
    public readonly struct EntityTimelineKind : IEquatable<EntityTimelineKind>
    {
        private readonly string _value;

        private EntityTimelineKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// activity
        /// </summary>
        public static EntityTimelineKind Activity { get; } = new EntityTimelineKind("Activity");
        /// <summary>
        /// bookmarks
        /// </summary>
        public static EntityTimelineKind Bookmark { get; } = new EntityTimelineKind("Bookmark");
        /// <summary>
        /// security alerts
        /// </summary>
        public static EntityTimelineKind SecurityAlert { get; } = new EntityTimelineKind("SecurityAlert");
        /// <summary>
        /// anomaly
        /// </summary>
        public static EntityTimelineKind Anomaly { get; } = new EntityTimelineKind("Anomaly");

        public static bool operator ==(EntityTimelineKind left, EntityTimelineKind right) => left.Equals(right);
        public static bool operator !=(EntityTimelineKind left, EntityTimelineKind right) => !left.Equals(right);

        public static explicit operator string(EntityTimelineKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EntityTimelineKind other && Equals(other);
        public bool Equals(EntityTimelineKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
