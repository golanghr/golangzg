package secure2

type algorithm struct {
	name string
}

var (
	TypeA = algorithm{name: "TypeA"}
	TypeB = algorithm{name: "TypeB"}
)

// Compute accepts one of possible Algorithm types and a value and returns a hash.
func Compute(a algorithm, v string) (hash string) {
	return v + "-" + string(a.name)
}
