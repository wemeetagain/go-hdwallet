package hdwalletutil

import (
    "testing"
    )

func TestPow(t *testing.T) {
    a, b, c := 3, 5, 243
    ans := pow(a,b)
    if ans != c {
        t.Errorf("3^5=%i, should equal 243",ans)
    }
}

