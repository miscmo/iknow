package utils

func GetLimitOffset(page, pageSize int32) (int32, int32) {
	var (
		limit int32
		offset int32
	)
	if pageSize < 0 {
		limit = 0
	}

	if page <= 0 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	return limit, offset
}
