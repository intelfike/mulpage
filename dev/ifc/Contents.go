package ifc

import "github.com/intelfike/mulpage/dev/types"

type Content interface {
	Init(types.Content)
}

type Package interface {
	Init(types.Package)
}
