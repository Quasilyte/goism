package function

type LispFn struct {
	Sym string
}

var (
	StrCast      = &LispFn{Sym: "Go-rt.bytesToStr"}
	ArrayToSlice = &LispFn{Sym: "Go-rt.arrayToSlice"}
)

var (
	CopySequence   = &LispFn{Sym: "copy-sequence"}
	Intern         = &LispFn{Sym: "intern"}
	Gethash        = &LispFn{Sym: "gethash"}
	MakeVector     = &LispFn{Sym: "make-vector"}
	Remhash        = &LispFn{Sym: "remhash"}
	HashTableCount = &LispFn{Sym: "hash-table-count"}
	Lsh            = &LispFn{Sym: "lsh"}
	Logand         = &LispFn{Sym: "logand"}
	Logior         = &LispFn{Sym: "logior"}
	Logxor         = &LispFn{Sym: "logxor"}
	StrGt          = &LispFn{Sym: "string>"}
	Vector         = &LispFn{Sym: "vector"}
)

/*
func AppendOne(typ *types.Slice) *Fn {
	return &Fn{
		name:    "Go--slice-push",
		results: tuple(typ),
	}
}

func MakeMap(typ *types.Map) *Fn {
	return &Fn{
		name:    "Go--make-map",
		results: tuple(typ),
	}
}

func MakeMapCap(typ *types.Map) *Fn {
	return &Fn{
		name:    "Go--make-map-cap",
		results: tuple(typ),
	}
}

func MakeSlice(typ *types.Slice) *Fn {
	return &Fn{
		name:    "Go--make-slice",
		results: tuple(typ),
	}
}

func MakeSliceCap(typ *types.Slice) *Fn {
	return &Fn{
		name:    "Go--make-slice-cap",
		results: tuple(typ),
	}
}
*/
