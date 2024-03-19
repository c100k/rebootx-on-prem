package main

import "strconv"

func parseInt(raw *string) *int32 {
	if raw == nil {
		return nil
	}

	v, err := strconv.ParseInt(*raw, 0, 32)
	if err != nil {
		return nil
	}

	asInt32 := int32(v)

	return &asInt32
}
