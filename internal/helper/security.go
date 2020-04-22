package helper

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ContainsPathTraversal checks if the path contains any traversal
func ContainsPathTraversal(path string) bool {
	// Disallow directory traversal for security
	separator := string(os.PathSeparator)
	return strings.HasPrefix(path, ".."+separator) ||
		strings.Contains(path, separator+".."+separator) ||
		strings.HasSuffix(path, separator+"..") ||
		filepath.IsAbs(path) ||
		path == ".."
}

// Pattern taken from Regular Expressions Cookbook, slightly modified though
//                                        |Scheme                |User                         |Named/IPv4 host|IPv6+ host
var hostPattern = regexp.MustCompile(`(?i)([a-z][a-z0-9+\-.]*://)([a-z0-9\-._~%!$&'()*+,;=:]+@)([a-z0-9\-._~%]+|\[[a-z0-9\-._~%!$&'()*+,;=:]+\])`)

// SanitizeString will clean password and tokens from URLs, and replace them
// with [FILTERED].
func SanitizeString(str string) string {
	return hostPattern.ReplaceAllString(str, "$1[FILTERED]@$3$4")
}

// SanitizeError does the same thing as SanitizeString but for error types
func SanitizeError(err error) error {
	return errors.New(SanitizeString(err.Error()))
}
