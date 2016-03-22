package fakes

import "github.com/pivotal-cf-experimental/bosh-bootloader/boshinit"

type BOSHInitManifestBuilder struct {
	BuildCall struct {
		Receives struct {
			Properties boshinit.ManifestProperties
		}
		Returns struct {
			Manifest   boshinit.Manifest
			Properties boshinit.ManifestProperties
			Error      error
		}
	}
}

func (b *BOSHInitManifestBuilder) Build(properties boshinit.ManifestProperties) (boshinit.Manifest, boshinit.ManifestProperties, error) {
	b.BuildCall.Receives.Properties = properties

	return b.BuildCall.Returns.Manifest, b.BuildCall.Returns.Properties, b.BuildCall.Returns.Error
}