package msg

import (
	"fmt"
	"testing"
)

func TestMsgListManual(t *testing.T) {
	ml := NewMsgList(3)

	u1 := &Msg{User: "u1"}
	u2 := &Msg{User: "u2"}
	u3 := &Msg{User: "u3"}
	u4 := &Msg{User: "u4"}

	ml.Add(u1)
	if ml.pos != 1 {
		t.Errorf("Error: ml.pos should be 1 not %v", ml.pos)
	}
	if ml.Len() != 1 {
		t.Errorf("Error: ml.Len() should be 1 not %v", ml.Len())
	}
	if ml.msgs[0] != u1 {
		t.Error("Error: msgs[0] != u1")
	}
	if ml.Get(0) != u1 {
		t.Error("Error: ml.Get(0) != u1")
	}

	ml.Add(u2)
	if ml.pos != 2 {
		t.Errorf("Error: ml.pos should be 2 not %v", ml.pos)
	}
	if ml.Len() != 2 {
		t.Errorf("Error: ml.Len() should be 2 not %v", ml.Len())
	}
	if ml.msgs[1] != u2 {
		t.Error("Error: msgs[0] != u2")
	}
	if ml.Get(1) != u2 {
		t.Error("Error: ml.Get(1) != u2")
	}

	ml.Add(u3)
	if ml.pos != 0 {
		t.Errorf("Error: ml.pos should be 0 not %v", ml.pos)
	}
	if ml.Len() != 3 {
		t.Errorf("Error: ml.Len() should be 3 not %v", ml.Len())
	}
	if ml.msgs[2] != u3 {
		t.Error("Error: msgs[2] != u3")
	}
	if ml.Get(2) != u3 {
		t.Error("Error: ml.Get(2) != u3")
	}

	ml.Add(u4)
	if ml.pos != 1 {
		t.Errorf("Error: ml.pos should be 1 not %v", ml.pos)
	}
	if ml.Len() != 3 {
		t.Errorf("Error: ml.Len() should be 3 not %v", ml.Len())
	}
	if ml.msgs[0] != u4 {
		t.Error("Error: msgs[0] != u4")
	}
	if ml.Get(2) != u4 {
		t.Error("Error: ml.Get(2) != u4")
	}
}

func TestMsgListAuto(t *testing.T) {
	testMsgListAutoHelper(t, 3, 30)
	testMsgListAutoHelper(t, 4, 40)
	testMsgListAutoHelper(t, 1000, 10000)
}

func testMsgListAutoHelper(t *testing.T, cap, iters int) {
	ml := NewMsgList(cap)

	for i := 0; i <= iters; i++ {
		usr := fmt.Sprintf("User %v", i)
		m := &Msg{User: usr}
		ml.Add(m)
		if i < cap {
			if ml.Len() != i+1 {
				t.Fatalf("Error: ml.Len() should be %v not %v", i+1, ml.Len())
			}
			if ml.Get(i) != m {
				t.Fatalf("Error: ml.Get(%v) should be %v not %v", i, m.User, ml.Get(i).User)
			}
		} else {
			if ml.Len() != cap {
				t.Fatalf("Error: ml.Len() should be %v not %v", cap, ml.Len())
			}
			if ml.Get(cap-1) != m {
				t.Fatalf("Error: ml.Get(%v) should be %v not %v", cap-1, m.User, ml.Get(cap-1).User)
			}
		}

	}
}
