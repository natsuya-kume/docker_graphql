package utils

import (
	"math"
)

// 現在のページ番号取得
func Page(limit int, offset *int) int {
	if limit > 0 {
		page := int(math.Floor(float64(*offset)/float64(limit)) + 1)
		return page
	}
	return 1
}

// 最大ページ数の取得
func PaginationLength(totalCount int, limit int) int {
	if limit > 0 {
		page := int(math.Ceil(float64(totalCount) / float64(limit)))
		return page
	}
	return 0
}

// 前のページがあるか
func HasPreviousPage(offset *int) bool {
	return *offset > 0
}

// 次のページがあるか
func HasNextPage(totalCount int, limit int, offset *int) bool {
	return *offset+limit < totalCount
}
