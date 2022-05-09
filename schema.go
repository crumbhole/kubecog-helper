package main

type CogValues struct {
	ArgoCD   ArgoCD   `json:"name,omitempty"`
	Platform Platform `json:"platform" validate:"required"`
}

type ArgoCD struct {
	HA bool `json:"ha,omitempty"`
}

type Platform struct {
	Provider string `json:"provider" validate:"required,oneof=rke k3s aks eks"`
}
