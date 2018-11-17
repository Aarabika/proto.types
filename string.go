package ptypes

import "github.com/Aarabika/nulls"

func NullString(val *String) nulls.String {
	return nulls.String{
		String: val.GetString_(),
		Valid:  val.GetValid(),
	}
}

func NullStringProto(val nulls.String) *String {
	return &String{
		String_: val.String,
		Valid:   val.Valid,
	}
}
