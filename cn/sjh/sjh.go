package sjh

import "regexp"

const sjhPattern = `^1[34578][0-9]{9}$`

var regSjh = regexp.MustCompile(sjhPattern)

func Valid(phone string) bool { return regSjh.MatchString(phone) }

func ValidAndMask(phone string) string {
	if Valid(phone) {
		return phone[:3] + "****" + phone[7:]
	}
	return phone
}
