package manager

func Destroy(projectName, backedURL string, sp Provider) error {
	// this will return a provider which implements the api.Provider interface
	p, err := GetProvider(sp)
	if err != nil {
		return err
	}
	return destroyStack(stack{
		name:           stackCreate,
		backedURL:      backedURL,
		providerPlugin: *p.GetPlugin()})
}
