package fileWrite

import "testing"

func TestFileWrite(t *testing.T) {
    result := Hello("kouki")
    want := "Hi, kouki. Welcome!"
    if result != want {
        t.Errorf("fileWrite.Hello() = %q want %q", result, want)
    }
}
