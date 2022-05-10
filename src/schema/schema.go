package schema

// CogValues is the top level structure for cogvalues.yaml as used by
// cog-plugin. Add json: and validate: tags as appropriate for validation.
type CogValues struct {
	ArgoCD   ArgoCD   `json:"name,omitempty"`
	Platform Platform `json:"platform" validate:"required"`
}

// ArgoCD is a member of CogValues for controlling the installation of ArgoCD
type ArgoCD struct {
	HA bool `json:"ha,omitempty"`
}

// Platform is a member of CogValues for telling cog where it is being installed
type Platform struct {
	Provider string `json:"provider" validate:"required,oneof=rke k3s aks eks"`
}
