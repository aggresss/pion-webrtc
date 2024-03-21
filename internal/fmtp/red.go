package fmtp

import "strings"

type redFMTP struct {
	parameters []string
}

var _ FMTP = &redFMTP{}

func (r *redFMTP) MimeType() string {
	return "audio/red"
}

func (r *redFMTP) Match(b FMTP) bool {
	c, ok := b.(*redFMTP)
	if !ok {
		return false
	}
	if !strings.EqualFold(r.MimeType(), c.MimeType()) || len(c.parameters) != len(r.parameters) {
		return false
	}
	for i, v := range r.parameters {
		if c.parameters[i] != v {
			return false
		}
	}

	return true
}

func (r *redFMTP) Parameter(key string) (string, bool) {
	for _, p := range r.parameters {
		if p == key {
			return key, true
		}
	}

	return "", false
}
