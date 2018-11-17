package ptypes

import "github.com/Aarabika/nulls"

func NullBool(val *Bool) nulls.Bool {
	return nulls.Bool{
		Bool:  val.GetBool(),
		Valid: val.GetValid(),
	}
}

func NullBoolProto(val nulls.Bool) *Bool {
	return &Bool{
		Bool:  val.Bool,
		Valid: val.Valid,
	}
}
