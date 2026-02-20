package brokers

import (
	"context"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &LoginRequest{
		CredentialId: "asdd122e-asdx-adin-xzca-asd123y1983x",
		Otp:          "148357", // Replace with actual OTP from SendEmailOTP
	}

	resp, err := client.Login(context.Background(), req)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	t.Logf("Login Response:")
	t.Logf("  Success: %v", resp.Success)
	t.Logf("  Message: %s", resp.Message)
	t.Logf("  AccessToken: %s", resp.AccessToken)

	if !resp.Success {
		t.Errorf("Expected success=true, got success=%v, message=%s", resp.Success, resp.Message)
	}
}
