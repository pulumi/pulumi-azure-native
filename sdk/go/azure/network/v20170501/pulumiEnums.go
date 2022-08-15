


package v20170501

type EndpointMonitorStatus string

const (
	EndpointMonitorStatusCheckingEndpoint = EndpointMonitorStatus("CheckingEndpoint")
	EndpointMonitorStatusOnline           = EndpointMonitorStatus("Online")
	EndpointMonitorStatusDegraded         = EndpointMonitorStatus("Degraded")
	EndpointMonitorStatusDisabled         = EndpointMonitorStatus("Disabled")
	EndpointMonitorStatusInactive         = EndpointMonitorStatus("Inactive")
	EndpointMonitorStatusStopped          = EndpointMonitorStatus("Stopped")
)

type EndpointStatus string

const (
	EndpointStatusEnabled  = EndpointStatus("Enabled")
	EndpointStatusDisabled = EndpointStatus("Disabled")
)

type MonitorProtocol string

const (
	MonitorProtocolHTTP  = MonitorProtocol("HTTP")
	MonitorProtocolHTTPS = MonitorProtocol("HTTPS")
	MonitorProtocolTCP   = MonitorProtocol("TCP")
)

type ProfileMonitorStatus string

const (
	ProfileMonitorStatusCheckingEndpoints = ProfileMonitorStatus("CheckingEndpoints")
	ProfileMonitorStatusOnline            = ProfileMonitorStatus("Online")
	ProfileMonitorStatusDegraded          = ProfileMonitorStatus("Degraded")
	ProfileMonitorStatusDisabled          = ProfileMonitorStatus("Disabled")
	ProfileMonitorStatusInactive          = ProfileMonitorStatus("Inactive")
)

type ProfileStatus string

const (
	ProfileStatusEnabled  = ProfileStatus("Enabled")
	ProfileStatusDisabled = ProfileStatus("Disabled")
)

type TrafficRoutingMethod string

const (
	TrafficRoutingMethodPerformance = TrafficRoutingMethod("Performance")
	TrafficRoutingMethodPriority    = TrafficRoutingMethod("Priority")
	TrafficRoutingMethodWeighted    = TrafficRoutingMethod("Weighted")
	TrafficRoutingMethodGeographic  = TrafficRoutingMethod("Geographic")
)

func init() {
}
