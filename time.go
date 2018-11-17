package ptypes

import (
	"github.com/Aarabika/nulls"
	"github.com/golang/protobuf/ptypes"
)

func NullTime(val *Time) (nulls.Time, error) {
	t, err := ptypes.Timestamp(val.Time)

	if err != nil {
		return nulls.Time{}, err
	}

	return nulls.Time{
		Time:  t,
		Valid: val.GetValid(),
	}, nil
}

func NullTimeProto(val nulls.Time) (*Time, error) {

	t, err := ptypes.TimestampProto(val.Time)

	if err != nil {
		return nil, err
	}

	return &Time{
		Time:  t,
		Valid: val.Valid,
	}, nil
}
