package assets

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/ge"
)

func registerShaderResources(ctx *ge.Context) {
	// Associate shader resources.
	shaderResources := map[resource.ShaderID]resource.ShaderInfo{
		ShaderDissolve:    {Path: "shader/dissolve.go"},
		ShaderColonyBuild: {Path: "shader/colony_build.go"},
	}
	for id, res := range shaderResources {
		ctx.Loader.ShaderRegistry.Set(id, res)
		ctx.Loader.LoadShader(id)
	}
}

const (
	ShaderDissolve resource.ShaderID = iota
	ShaderColonyBuild
)
