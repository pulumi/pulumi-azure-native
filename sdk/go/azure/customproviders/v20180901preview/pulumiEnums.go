


package v20180901preview

type ActionRouting string

const (
	ActionRoutingProxy = ActionRouting("Proxy")
)

type ResourceTypeRouting string

const (
	ResourceTypeRoutingProxy        = ResourceTypeRouting("Proxy")
	ResourceTypeRouting_Proxy_Cache = ResourceTypeRouting("Proxy,Cache")
)

type ValidationType string

const (
	ValidationTypeSwagger = ValidationType("Swagger")
)

func init() {
}
