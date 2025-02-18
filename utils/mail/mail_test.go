package mail

import "testing"

func TestMail(t *testing.T) {
	err := SendMail("asharudheen@omaemirates.com", "test email", "this email for testing email functionalty")
	if err != nil {
		t.Fatalf("error %v", err)
	}
}
