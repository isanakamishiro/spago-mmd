package components

import (
	"github.com/nobonobo/spago"
)

//go:generate spago generate -c RendererInfoPane -p components render_info_pane.html

// RendererInfoPane  ...
type RendererInfoPane struct {
	spago.Core
}
