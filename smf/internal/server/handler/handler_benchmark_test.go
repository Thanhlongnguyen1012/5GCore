package handler

import (
	"os"
	"testing"
)

func BenchmarkGetSessionManagementSubscription(b *testing.B) {

	UdmBaseURL = os.Getenv("UDM_BASE_URL")

	payload := Payload{
		Supi: "452040989692072",
	}

	for i := 0; i < b.N; i++ {
		resp, err := payload.GetSessionManagementSubscription()
		if err != nil {
			b.Fatalf("Failed to get session data: %v", err)
		}
		resp.Body.Close()
	}
}
