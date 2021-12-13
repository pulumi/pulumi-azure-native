


package v20200201preview

type MediaGraphRtspTransport string

const (
	// HTTP/HTTPS transport. This should be used when HTTP tunneling is desired.
	MediaGraphRtspTransportHttp = MediaGraphRtspTransport("Http")
	// TCP transport. This should be used when HTTP tunneling is not desired.
	MediaGraphRtspTransportTcp = MediaGraphRtspTransport("Tcp")
)

func init() {
}
