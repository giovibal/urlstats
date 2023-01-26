package store

type OrderBy int

const (
	OrderByCreatedAt OrderBy = iota
	OrderByCreatedAtDesc
	OrderByHitCount
	OrderByHitCountDesc
	OrderByAvgBytes
	OrderByAvgBytesDesc
)

func OrderByFromString(s string) OrderBy {
	switch s {
	case "createdAt":
		return OrderByCreatedAt
	case "createdAtDesc":
		return OrderByCreatedAtDesc
	case "hitCount":
		return OrderByHitCount
	case "hitCountDesc":
		return OrderByHitCountDesc
	case "avgBytes":
		return OrderByAvgBytes
	case "avgBytesDesc":
		return OrderByAvgBytesDesc
	}

	// default to the most interesting order
	return OrderByHitCountDesc
}

func OrderByFromStringArray(sarr []string) []OrderBy {
	var ret []OrderBy
	for _, s := range sarr {
		ret = append(ret, OrderByFromString(s))
	}
	return ret
}
