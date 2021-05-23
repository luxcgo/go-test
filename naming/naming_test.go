package naming

import "testing"

func TestDog_Bark(t *testing.T) {

}

// Dog Type + Bark Method + muzzled Params
func TestDog_Bark_muzzled(t *testing.T) {

}

func TestDog_Bark_unmuzzled(t *testing.T) {

}

func TestColor(t *testing.T) {
	arg := "blue"
	want := "#00000FF"
	got := Color(arg)
	if got != want {
		t.Errorf("Color(%q) = %q; want %q", arg, got, want)
	}
}
