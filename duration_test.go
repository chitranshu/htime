package htime

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	if res := Duration(""); 0 != res {
		t.Errorf("Duration(\"\") returned expected %d actual %d", 0, res)
	}

	expect := time.Minute + (time.Second * 30)
	if res := Duration("1 minute and 30 seconds"); expect != res {
		t.Errorf("Duration(\"1 min and 30 seconds\") expected %d actual %d", expect, res)
	}

	day := time.Hour * 24

	expect = day * 395
	if res := Duration("1 year 30 days"); expect != res {
		t.Errorf("Duration(\"1 year 30 days\") expected %d actual %d", expect, res)
	}

	expect = day * 3
	if res := Duration("3 days 4 burgers"); expect != res {
		t.Errorf("Duration(\"3 days 4 burgers\") expected %d actual %d", expect, res)
	}

	expect = 0
	if res := Duration("- days 4 hours"); expect != res {
		t.Errorf("Duration(\"- days 4 hours\") expected %d actual %d", expect, res)
	}

	expect = -3 * day
	if res := Duration("-3 days 4 burgers"); expect != res {
		t.Errorf("Duration(\"-3 days 4 burgers\") expected %d actual %d", expect, res)
	}

}
