// +build !linux,!plan9,!freebsd,!solaris,!darwin

package lsof

func Scan(chkFn StringChecker) (map[string][]int, error) {
	return nil, ErrUnsupportedOS
}
