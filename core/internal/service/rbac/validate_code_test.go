package rbac

import "testing"

func TestGenerateCaptcha(t *testing.T) {
	// Basic functionality test - only checks return value types and non-emptiness
	t.Run("Basic functionality check", func(t *testing.T) {
		gotId, gotB64s, gotAnswer, err := GenerateCaptcha()
		if err != nil {
			t.Errorf("GenerateCaptcha() should not return error, but got: %v", err)
		}
		if gotId == "" {
			t.Error("GenerateCaptcha() returned ID should not be empty")
		}
		if gotB64s == "" {
			t.Error("GenerateCaptcha() returned Base64 string should not be empty")
		}
		if len(gotAnswer) != 4 {
			t.Errorf("GenerateCaptcha() returned answer length should be 4, but got: %d", len(gotAnswer))
		}
	})

	// Verify if the captcha answer can pass verification
	t.Run("Captcha verification test", func(t *testing.T) {
		gotId, _, gotAnswer, err := GenerateCaptcha()
		if err != nil {
			t.Errorf("GenerateCaptcha() should not return error, but got: %v", err)
			return
		}

		// Verify with correct answer
		if !VerifyCaptcha(gotId, gotAnswer) {
			t.Error("VerifyCaptcha() should return true with correct answer")
		}

		// Verify with wrong answer
		if VerifyCaptcha(gotId, "wrong answer") {
			t.Error("VerifyCaptcha() should return false with wrong answer")
		}
	})

	// Verify the format of generated captcha
	t.Run("Captcha format test", func(t *testing.T) {
		_, _, gotAnswer, err := GenerateCaptcha()
		if err != nil {
			t.Errorf("GenerateCaptcha() should not return error, but got: %v", err)
			return
		}

		// According to defaultDriver configuration, the answer should only contain characters from "34578abcdefhjkmnprstuvwxy"
		validChars := "34578abcdefhjkmnprstuvwxy"
		for _, char := range gotAnswer {
			found := false
			for _, validChar := range validChars {
				if char == validChar {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Captcha contains invalid character: %c", char)
			}
		}
	})

	// Test the GetCaptchaAnswer function
	t.Run("Get answer test", func(t *testing.T) {
		gotId, _, gotAnswer, err := GenerateCaptcha()
		if err != nil {
			t.Errorf("GenerateCaptcha() should not return error, but got: %v", err)
			return
		}

		// Get answer via GetCaptchaAnswer
		retrievedAnswer := GetCaptchaAnswer(gotId)
		if retrievedAnswer != gotAnswer {
			t.Errorf("GetCaptchaAnswer() = %v, want %v", retrievedAnswer, gotAnswer)
		}
	})
}

func TestGetCaptcha(t *testing.T) {
	// Test basic functionality
	t.Run("Basic functionality", func(t *testing.T) {
		gotId, gotB64s, err := GetCaptcha()
		if err != nil {
			t.Errorf("GetCaptcha() should not return error, but got: %v", err)
		}
		if gotId == "" {
			t.Error("GetCaptcha() returned ID should not be empty")
		}
		if gotB64s == "" {
			t.Error("GetCaptcha() returned Base64 string should not be empty")
		}
	})

	// Test handling of empty captcha pool
	t.Run("Empty pool handling", func(t *testing.T) {
		// Create a temporary pool for testing
		originalPool := globalPool
		defer func() {
			globalPool = originalPool // Restore original pool after test
		}()

		// Create an empty pool
		emptyPool := NewCodePool(0, DefaultExpiration, DefaultRefreshRate)
		globalPool = emptyPool

		// Force empty the pool
		emptyPool.mutex.Lock()
		emptyPool.pool = make(map[string]string, 0)
		emptyPool.mutex.Unlock()

		// Call GetCaptcha, it should handle empty pool and generate new captcha
		gotId, gotB64s, err := GetCaptcha()
		if err != nil {
			// Even with empty pool, it should auto-refresh and generate captcha
			t.Errorf("GetCaptcha() should auto-refresh with empty pool, but got error: %v", err)
		}
		if gotId == "" || gotB64s == "" {
			t.Error("After auto-refresh, GetCaptcha() results should not be empty")
		}
	})

	// Test verification process
	t.Run("Verification test", func(t *testing.T) {
		// Get a captcha
		gotId, _, err := GetCaptcha()
		if err != nil {
			t.Errorf("GetCaptcha() should not return error, but got: %v", err)
			return
		}

		// Get correct answer
		answer := GetCaptchaAnswer(gotId)
		if answer == "" {
			t.Error("Failed to get captcha answer")
			return
		}

		// Verify with correct answer
		if !VerifyCaptcha(gotId, answer) {
			t.Error("Verification failed with correct answer")
		}

		// Verify with wrong answer
		if VerifyCaptcha(gotId, "wrong answer") {
			t.Error("Verification should not succeed with wrong answer")
		}
	})
}
