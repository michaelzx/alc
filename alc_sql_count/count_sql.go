package alc_sql_count

import (
	"strings"
)

const (
	from1          = 'f'
	from2          = 'r'
	from3          = 'o'
	from4          = 'm'
	countSqlPrefix = "select count(*) "
)

func getFromIdx(sub []rune) int {
	if len(sub) < 4 {
		return -1
	}
	for i := 0; i <= len(sub)-4; i++ {
		if sub[i] != from1 {
			continue
		}
		if sub[i+1] != from2 {
			continue
		}
		if sub[i+2] != from3 {
			continue
		}
		if sub[i+3] != from4 {
			continue
		}
		return i
	}
	return -1
}
func Convert(originSql string) string {
	letters := []rune(strings.ToLower(originSql))
	levelNow := 0
	rootGroups := make([][2]int, 0)
	for i, letter := range letters {
		if letter == '(' {
			levelNext := levelNow + 1
			if levelNext == 1 {
				rootGroups = append(rootGroups, [2]int{i, 0})
			}
			levelNow = levelNext
		}
		if letter == ')' {
			if levelNow == 1 {
				rootGroups[len(rootGroups)-1][1] = i
			}
			levelNow = levelNow - 1
		}
	}
	iStart := 0
	fromPos := 0
	if len(rootGroups) > 0 {
		for gi, gv := range rootGroups {
			sub := letters[iStart:gv[0]]
			if fromIdx := getFromIdx(sub); fromIdx >= 0 {
				fromPos = iStart + fromIdx
				break
			}
			if gi == len(rootGroups)-1 && gv[1] < len(letters)-1 {
				sub := letters[gv[1]+1:]
				if fromIdx := getFromIdx(sub); fromIdx >= 0 {
					fromPos = gv[1] + fromIdx
					break
				}
			}
			iStart = gv[1] + 1
		}
	} else {
		if fromIdx := getFromIdx(letters); fromIdx >= 0 {
			fromPos = fromIdx
		}
	}
	countSql := countSqlPrefix + string(letters[fromPos:])
	return countSql
}
