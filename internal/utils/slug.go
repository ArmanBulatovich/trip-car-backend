package utils

import "regexp"

func IsValidSlug(slug string) bool {
	if len(slug) < 3 {
		return false
	}
	re := regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)
	return re.MatchString(slug)
}
