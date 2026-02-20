package brokers

import (
	"context"
	"os"
	"testing"
)

func TestSendEmailOTP(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &SendEmailOTPRequest{
		AccessToken:  "your_access_token_here",
		CredentialId: "asdd122e-asdx-adin-xzca-asd123y1983x",
	}

	resp, err := client.SendEmailOTP(context.Background(), req)
	if err != nil {
		t.Fatalf("SendEmailOTP failed: %v", err)
	}

	t.Logf("SendEmailOTP Response:")
	t.Logf("  Success: %v", resp.Success)
	t.Logf("  Message: %s", resp.Message)
	t.Logf("  OTP: %s", resp.Otp)

	if !resp.Success {
		t.Errorf("Expected success=true, got success=%v, message=%s", resp.Success, resp.Message)
	}
}
