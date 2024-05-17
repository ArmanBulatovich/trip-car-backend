package utils

const perPageDefault = 10

func GetLimitAndOffset(page, perPage int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if perPage <= 0 {
		perPage = perPageDefault
	}

	return perPage, (page - 1) * perPage
}
