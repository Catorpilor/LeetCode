package mtd

import "testing"

func TestMinDiff(t *testing.T) {
    st := []struct {
        name string
        tps  []string
        exp  int
    }{
        {"just two", []string{"03:00", "00:00"}, 180},
        {"testcase1", []string{"03:15", "00:00", "23:58"}, 2},
        {"failed1", []string{"12:01", "23:59"}, 718},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := minDiff(tt.tps)
            if out != tt.exp {
                t.Fatalf("with time points: %v wanted %d but got %d", tt.tps, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}

func TestSimpler(t *testing.T) {
    st := []struct {
        name string
        tps  []string
        exp  int
    }{
        {"just two", []string{"03:00", "00:00"}, 180},
        {"testcase1", []string{"03:15", "00:00", "23:58"}, 2},
        {"failed1", []string{"12:01", "23:59"}, 718},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := simpler(tt.tps)
            if out != tt.exp {
                t.Fatalf("with time points: %v wanted %d but got %d", tt.tps, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}

func TestBucket(t *testing.T) {
    st := []struct {
        name string
        tps  []string
        exp  int
    }{
        {"just two", []string{"03:00", "00:00"}, 180},
        {"testcase1", []string{"03:15", "00:00", "23:58"}, 2},
        {"failed1", []string{"12:01", "23:59"}, 718},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := bucket(tt.tps)
            if out != tt.exp {
                t.Fatalf("with time points: %v wanted %d but got %d", tt.tps, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
