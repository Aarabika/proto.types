package ptypes

import "github.com/Aarabika/nulls"

func NullInt(val *Int) nulls.Int {
	return nulls.Int{
		Int: int(val.GetInt()),
		Valid:  val.GetValid(),
	}
}

func NullIntProto(val nulls.Int) *Int {
	return &Int{
		Int: int32(val.Int),
		Valid:   val.Valid,
	}
}
