package staging

import (
	"image/color"

	"github.com/quasilyte/ge"
	"github.com/quasilyte/roboden-game/viewport"
)

type beamNode struct {
	from   ge.Pos
	to     ge.Pos
	camera *viewport.Camera

	color color.RGBA
	width float64
	line  *ge.Line

	texture *ge.Texture
	texLine *ge.TextureLine
}

var (
	repairBeamColor          = ge.RGB(0x6ac037)
	rechargerBeamColor       = ge.RGB(0x66ced6)
	railgunBeamColor         = ge.RGB(0xbd1844)
	dominatorBeamColorCenter = ge.RGB(0x7a51f2)
	dominatorBeamColorRear   = ge.RGB(0x5433c3)
	builderBeamColor         = color.RGBA{R: 0xae, G: 0x4c, B: 0x78, A: 150}
	destroyerBeamColor       = ge.RGB(0xf58f54)
	courierResourceBeamColor = ge.RGB(0xd2e352)
	prismBeamColor1          = ge.RGB(0x529eb8)
	prismBeamColor2          = ge.RGB(0x61bad8)
	prismBeamColor3          = ge.RGB(0x7bdbfc)
	prismBeamColor4          = ge.RGB(0xccf2ff)
	evoBeamColor             = ge.RGB(0xa641c2)
)

var prismBeamColors = []color.RGBA{
	prismBeamColor1,
	prismBeamColor2,
	prismBeamColor3,
	prismBeamColor4,
}

func newBeamNode(camera *viewport.Camera, from, to ge.Pos, clr color.RGBA) *beamNode {
	return &beamNode{
		camera: camera,
		from:   from,
		to:     to,
		width:  1,
		color:  clr,
	}
}

func newTextureBeamNode(camera *viewport.Camera, from, to ge.Pos, texture *ge.Texture) *beamNode {
	return &beamNode{
		camera:  camera,
		from:    from,
		to:      to,
		texture: texture,
	}
}

func (b *beamNode) Init(scene *ge.Scene) {
	if b.texture == nil {
		b.line = ge.NewLine(b.from, b.to)
		var c ge.ColorScale
		c.SetColor(b.color)
		b.line.SetColorScale(c)
		b.line.Width = b.width
		b.camera.AddGraphicsAbove(b.line)
	} else {
		b.texLine = ge.NewTextureLine(scene.Context(), b.from, b.to)
		b.texLine.SetTexture(b.texture)
		b.camera.AddGraphicsAbove(b.texLine)
	}
}

func (b *beamNode) IsDisposed() bool {
	if b.texture == nil {
		return b.line.IsDisposed()
	}
	return b.texLine.IsDisposed()
}

func (b *beamNode) Update(delta float64) {
	if b.texture == nil {
		if b.line.GetAlpha() < 0.1 {
			b.line.Dispose()
			return
		}
		b.line.SetAlpha(b.line.GetAlpha() - float32(delta*4))
		return
	}

	if b.texLine.GetAlpha() < 0.1 {
		b.texLine.Dispose()
		return
	}
	b.texLine.SetAlpha(b.texLine.GetAlpha() - float32(delta*4))
}
