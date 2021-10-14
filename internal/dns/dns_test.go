package dns

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverser4(t *testing.T) {
	tests := []struct {
		ip   string
		want string
	}{
		{"8.8.4.4", "4.4.8.8.in-addr.arpa."},
		{"8.8.8.8", "8.8.8.8.in-addr.arpa."},
		{"1.2.3.4", "4.3.2.1.in-addr.arpa."},
	}

	for _, test := range tests {
		got, _ := reverse4(test.ip)
		require.Equal(t, test.want, got)
	}
}

func TestReverse6(t *testing.T) {
	tests := []struct {
		ip   string
		want string
	}{
		{"2607:f8b0:4002:c0c::68", "8.6.0.0.0.0.0.0.0.0.0.0.0.0.0.0.c.0.c.0.2.0.0.4.0.b.8.f.7.0.6.2.ip6.arpa."},
	}

	for _, test := range tests {
		got, _ := reverse6(test.ip)
		require.Equal(t, test.want, got)
	}
}

func TestReverseGeneric(t *testing.T) {
	tests := []struct {
		ip   string
		want string
	}{
		{"2607:f8b0:4002:c0c::68", "8.6.0.0.0.0.0.0.0.0.0.0.0.0.0.0.c.0.c.0.2.0.0.4.0.b.8.f.7.0.6.2.ip6.arpa."},
		{"1.2.3.4", "4.3.2.1.in-addr.arpa."},
		{"foo", ""},
	}

	for _, test := range tests {
		got, _ := Reverse(test.ip)
		require.Equal(t, test.want, got)
	}
}
