package dns

import (
	"testing"
	"github.com/miekg/dns"
)

func TestRadixName(t *testing.T) {
	tests := map[string]string{".": ".",
		"www.miek.nl.": ".nl.miek.www",
		"miek.nl.":     ".nl.miek",
		"mi\\.ek.nl.":  ".nl.mi\\.ek",
		`mi\\.ek.nl.`:  `.nl.ek.mi\\`,
		"":             "."}
	for i, o := range tests {
		t.Logf("%s %v\n", i, SplitLabels(i))
		if x := toRadixName(i); x != o {
			t.Logf("%s should convert to %s, not %s\n", i, o, x)
			t.Fail()
		}
	}
}

func TestInsert(t *testing.T) {
}
func TestRemove(t *testing.T) {
	z := dns.NewZone("miek.nl.")
	mx, _ := dns.NewRR("foo.miek.nl. MX 10 mx.miek.nl.")
	z.Insert(mx)
	zd, exact := z.Find("foo.miek.nl.")
	if exact != true {
		t.Fail() // insert broken?
	}
	z.Remove(zd.RR[dns.TypeMX][0])
	zd, exact = z.Find("foo.miek.nl.")
	if exact != false {
		t.Errorf("zd(%s) exact(%s) still exists", zd, exact) // it should no longer be in the zone
	}
}
