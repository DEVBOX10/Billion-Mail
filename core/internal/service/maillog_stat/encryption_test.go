package maillog_stat

import (
	"encoding/json"
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Encrypt String",
			args: args{
				data: "Test String 123",
			},
		},
		{
			name: "Test Encrypt Numbers",
			args: args{
				data: 123456,
			},
		},
		{
			name: "Test Encrypt Map",
			args: args{
				data: map[string]interface{}{
					"id":         1001,
					"email":      "test@example.com",
					"event_type": "click",
					"timestamp":  1623456789,
					"metadata": map[string]interface{}{
						"campaign_id": 5678,
						"url":         "https://example.com/page",
						"ip":          "192.168.1.1",
					},
					"tags": []string{"marketing", "newsletter"},
				},
			},
		},
		{
			name: "Test Encrypt Empty String",
			args: args{
				data: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encrypted := Encrypt(tt.args.data)
			if encrypted == "" {
				t.Errorf("Encrypt() returned empty string")
				return
			}

			if len(encrypted) < 10 {
				t.Errorf("Invalid encrypt length: %v", encrypted)
			}

			// Test Decryption
			var decrypted interface{}
			err := Decrypt(encrypted, &decrypted)
			if err != nil {
				t.Errorf("Decrypt failed: %v", err)
				return
			}

			// Marshal both original and decrypted data to JSON for comparison
			originalJSON, _ := json.Marshal(tt.args.data)
			decryptedJSON, _ := json.Marshal(decrypted)

			if string(originalJSON) != string(decryptedJSON) {
				t.Errorf("Decrypted data not match: origin=%v, decrypted=%v", string(originalJSON), string(decryptedJSON))
			}
		})
	}
}
