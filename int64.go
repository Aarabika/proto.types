package ptypes

import "github.com/Aarabika/nulls"

func NullInt64(val *Int64) nulls.Int64 {
	return nulls.Int64{
		Int64: val.GetInt64(),
		Valid:  val.GetValid(),
	}
}

func NullInt64Proto(val nulls.Int64) *Int64 {
	return &Int64{
		Int64: val.Int64,
		Valid:   val.Valid,
	}
}
