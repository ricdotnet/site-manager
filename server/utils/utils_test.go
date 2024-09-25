package utils

import (
	"testing"
)

func TestIsValidFileName(t *testing.T) {
	for k, v := range map[string]bool{
		"000-default.conf":         true,
		"ssl-default.conf":         true,
		"some-domain.com.conf":     true,
		"ssl-some-domain.com.conf": true,
		"sub.domain.com.conf":      true,
		"ssl-sub.domain.com.conf":  true,
		"../../another/dir":        false,
		"/ && rm -r /":             false,
		".../...//":                false,
		"..":                       false,
		" ":                        false,
	} {
		got := IsValidFilename(k)
		if got != v {
			t.Errorf("expected %t for %s, got %t", v, k, got)
		}
	}
}

//func TestBuildApachePath(t *testing.T) {
//	goenvironmental.ParseEnv("../.env")
//	apachePath, _ := goenvironmental.Get("APACHE_PATH")
//
//	got := BuildApachePath("sites-available/")
//	want := filepath.Join(apachePath, "sites-available")
//
//	if got != want {
//		t.Errorf("expected %s but got %s", want, got)
//	}
//}
