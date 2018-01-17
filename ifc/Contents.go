package ifc

import "github.com/intelfike/mulpage/types"

type Content interface {
	Init(*types.Content)
}

type Package interface {
	Init(*types.Package)
}
