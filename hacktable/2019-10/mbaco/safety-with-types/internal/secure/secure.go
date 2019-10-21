package secure

type algorithm string

const (
	TypeA algorithm = "TypeA"
	TypeB algorithm = "TypeB"
)

// Compute accepts one of possible Algorithm types and a value and returns a hash.
func Compute(a algorithm, v string) (hash string) {
	return v + "-" + string(a)
}
