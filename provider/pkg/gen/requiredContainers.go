package gen

func additionalRequiredContainers(resourceTok string) requiredContainers {
	switch resourceTok {
	case "azure-native:app:ManagedEnvironment":
		return [][]string{{"properties"}}
	}
	return nil
}
