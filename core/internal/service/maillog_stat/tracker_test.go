package maillog_stat

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestNewMailTracker(t *testing.T) {
	type args struct {
		mailHTML   string
		campaignID int
		messageID  string
		recipient  string
		baseURL    string
	}
	tests := []struct {
		name string
		args args
		want *MailTracker
	}{
		{
			name: "Basic email tracker creation",
			args: args{
				mailHTML:   "<html><body>Test Mail</body></html>",
				campaignID: 123,
				messageID:  "abc123",
				recipient:  "test@example.com",
				baseURL:    "https://example.com",
			},
			want: &MailTracker{
				originalMailHTML: "<html><body>Test Mail</body></html>",
				modified:         false,
				mailHTML:         "<html><body>Test Mail</body></html>",
				campaignID:       123,
				messageID:        "abc123",
				recipient:        "test@example.com",
				baseURL:          "https://example.com/pmta",
				hrefPattern:      regexp.MustCompile(`href\s*=\s*"([^"]+)"`),
			},
		},
		{
			name: "Email with slash in baseURL",
			args: args{
				mailHTML:   "<html><body>Content</body></html>",
				campaignID: 456,
				messageID:  "def456",
				recipient:  "user@example.org",
				baseURL:    "https://test.com/",
			},
			want: &MailTracker{
				originalMailHTML: "<html><body>Content</body></html>",
				modified:         false,
				mailHTML:         "<html><body>Content</body></html>",
				campaignID:       456,
				messageID:        "def456",
				recipient:        "user@example.org",
				baseURL:          "https://test.com/pmta",
				hrefPattern:      regexp.MustCompile(`href\s*=\s*"([^"]+)"`),
			},
		},
		{
			name: "Email with empty HTML",
			args: args{
				mailHTML:   "",
				campaignID: 789,
				messageID:  "ghi789",
				recipient:  "empty@example.net",
				baseURL:    "https://empty.com",
			},
			want: &MailTracker{
				originalMailHTML: "",
				modified:         false,
				mailHTML:         "",
				campaignID:       789,
				messageID:        "ghi789",
				recipient:        "empty@example.net",
				baseURL:          "https://empty.com/pmta",
				hrefPattern:      regexp.MustCompile(`href\s*=\s*"([^"]+)"`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMailTracker(tt.args.mailHTML, tt.args.campaignID, tt.args.messageID, tt.args.recipient, tt.args.baseURL)

			// Can't compare regexps directly, so we compare their string representations
			if got.hrefPattern.String() != tt.want.hrefPattern.String() {
				t.Errorf("NewMailTracker() hrefPattern = %v, want %v", got.hrefPattern, tt.want.hrefPattern)
			}

			// Set hrefPattern to nil for comparison
			got.hrefPattern = nil
			tt.want.hrefPattern = nil

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMailTracker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMailTracker_TrackLinks(t *testing.T) {
	type fields struct {
		originalMailHTML string
		modified         bool
		mailHTML         string
		campaignID       int
		messageID        string
		recipient        string
		baseURL          string
		hrefPattern      *regexp.Regexp
	}
	type test struct {
		name            string
		fields          fields
		wantModified    bool
		wantContains    string
		wantNotContains string
	}

	tests := []test{
		{
			name: "Track links in HTML with href",
			fields: fields{
				originalMailHTML: `<html><body><a href="https://example.com">Link</a></body></html>`,
				modified:         false,
				mailHTML:         `<html><body><a href="https://example.com">Link</a></body></html>`,
				campaignID:       123,
				messageID:        "abc123",
				recipient:        "test@example.com",
				baseURL:          "https://track.example.com/pmta",
				hrefPattern:      regexp.MustCompile(`href\s*=\s*"([^"]+)"`),
			},
			wantModified:    true,
			wantContains:    "https://track.example.com/pmta/",
			wantNotContains: `href="https://example.com"`,
		},
		{
			name: "No links to track",
			fields: fields{
				originalMailHTML: `<html><body>No links here</body></html>`,
				modified:         false,
				mailHTML:         `<html><body>No links here</body></html>`,
				campaignID:       123,
				messageID:        "abc123",
				recipient:        "test@example.com",
				baseURL:          "https://track.example.com/pmta",
				hrefPattern:      regexp.MustCompile(`href\s*=\s*"([^"]+)"`),
			},
			wantModified:    false,
			wantContains:    `<html><body>No links here</body></html>`,
			wantNotContains: "https://track.example.com/pmta/",
		},
		{
			name: "Invalid URL in href",
			fields: fields{
				originalMailHTML: `<html><body><a href="javascript:void(0)">JS Link</a></body></html>`,
				modified:         false,
				mailHTML:         `<html><body><a href="javascript:void(0)">JS Link</a></body></html>`,
				campaignID:       123,
				messageID:        "abc123",
				recipient:        "test@example.com",
				baseURL:          "https://track.example.com/pmta",
				hrefPattern:      regexp.MustCompile(`href\s*=\s*"([^"]+)"`),
			},
			wantModified:    false,
			wantContains:    `href="javascript:void(0)"`,
			wantNotContains: "https://track.example.com/pmta/",
		},
		{
			name: "Multiple links to track",
			fields: fields{
				originalMailHTML: `<html><body><a href="https://example1.com">Link1</a><a href="https://example2.com">Link2</a></body></html>`,
				modified:         false,
				mailHTML:         `<html><body><a href="https://example1.com">Link1</a><a href="https://example2.com">Link2</a></body></html>`,
				campaignID:       123,
				messageID:        "abc123",
				recipient:        "test@example.com",
				baseURL:          "https://track.example.com/pmta",
				hrefPattern:      regexp.MustCompile(`href\s*=\s*"([^"]+)"`),
			},
			wantModified:    true,
			wantContains:    "https://track.example.com/pmta/",
			wantNotContains: `href="https://example1.com"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tracker := &MailTracker{
				originalMailHTML: tt.fields.originalMailHTML,
				modified:         tt.fields.modified,
				mailHTML:         tt.fields.mailHTML,
				campaignID:       tt.fields.campaignID,
				messageID:        tt.fields.messageID,
				recipient:        tt.fields.recipient,
				baseURL:          tt.fields.baseURL,
				hrefPattern:      tt.fields.hrefPattern,
			}

			tracker.TrackLinks()

			if tracker.IsModified() != tt.wantModified {
				t.Errorf("TrackLinks() modified = %v, want %v", tracker.IsModified(), tt.wantModified)
			}

			if !strings.Contains(tracker.GetHTML(), tt.wantContains) {
				t.Errorf("TrackLinks() HTML does not contain %q", tt.wantContains)
			}

			if tt.wantNotContains != "" && strings.Contains(tracker.GetHTML(), tt.wantNotContains) {
				t.Errorf("TrackLinks() HTML contains %q when it should not", tt.wantNotContains)
			}
		})
	}
}

func TestMailTracker_AppendTrackingPixel(t *testing.T) {
	type fields struct {
		originalMailHTML string
		modified         bool
		mailHTML         string
		campaignID       int
		messageID        string
		recipient        string
		baseURL          string
		hrefPattern      *regexp.Regexp
	}
	type test struct {
		name            string
		fields          fields
		wantModified    bool
		wantContains    string
		wantNotContains string
	}

	tests := []test{
		{
			name: "Append tracking pixel to HTML with body tag",
			fields: fields{
				originalMailHTML: `<html><body>Test content</body></html>`,
				modified:         false,
				mailHTML:         `<html><body>Test content</body></html>`,
				campaignID:       123,
				messageID:        "abc123",
				recipient:        "test@example.com",
				baseURL:          "https://track.example.com/pmta",
				hrefPattern:      regexp.MustCompile(`href\s*=\s*"([^"]+)"`),
			},
			wantModified:    true,
			wantContains:    `<img src="https://track.example.com/pmta/`,
			wantNotContains: `</body><img src=`,
		},
		{
			name: "Append tracking pixel to HTML with html tag but no body tag",
			fields: fields{
				originalMailHTML: `<html>Test content</html>`,
				modified:         false,
				mailHTML:         `<html>Test content</html>`,
				campaignID:       123,
				messageID:        "abc123",
				recipient:        "test@example.com",
				baseURL:          "https://track.example.com/pmta",
				hrefPattern:      regexp.MustCompile(`href\s*=\s*"([^"]+)"`),
			},
			wantModified:    true,
			wantContains:    `<img src="https://track.example.com/pmta/`,
			wantNotContains: `</html><img src=`,
		},
		{
			name: "Append tracking pixel to HTML without body or html tag",
			fields: fields{
				originalMailHTML: `Just plain text content`,
				modified:         false,
				mailHTML:         `Just plain text content`,
				campaignID:       123,
				messageID:        "abc123",
				recipient:        "test@example.com",
				baseURL:          "https://track.example.com/pmta",
				hrefPattern:      regexp.MustCompile(`href\s*=\s*"([^"]+)"`),
			},
			wantModified: true,
			wantContains: `Just plain text content<img src="https://track.example.com/pmta/`,
		},
		{
			name: "Empty HTML content",
			fields: fields{
				originalMailHTML: ``,
				modified:         false,
				mailHTML:         ``,
				campaignID:       123,
				messageID:        "abc123",
				recipient:        "test@example.com",
				baseURL:          "https://track.example.com/pmta",
				hrefPattern:      regexp.MustCompile(`href\s*=\s*"([^"]+)"`),
			},
			wantModified: true,
			wantContains: `<img src="https://track.example.com/pmta/`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tracker := &MailTracker{
				originalMailHTML: tt.fields.originalMailHTML,
				modified:         tt.fields.modified,
				mailHTML:         tt.fields.mailHTML,
				campaignID:       tt.fields.campaignID,
				messageID:        tt.fields.messageID,
				recipient:        tt.fields.recipient,
				baseURL:          tt.fields.baseURL,
				hrefPattern:      tt.fields.hrefPattern,
			}

			tracker.AppendTrackingPixel()

			if tracker.IsModified() != tt.wantModified {
				t.Errorf("AppendTrackingPixel() modified = %v, want %v", tracker.IsModified(), tt.wantModified)
			}

			if !strings.Contains(tracker.GetHTML(), tt.wantContains) {
				t.Errorf("AppendTrackingPixel() HTML does not contain %q", tt.wantContains)
			}

			if tt.wantNotContains != "" && strings.Contains(tracker.GetHTML(), tt.wantNotContains) {
				t.Errorf("AppendTrackingPixel() HTML contains %q when it should not", tt.wantNotContains)
			}

			if strings.Contains(tt.fields.originalMailHTML, "</body>") {
				if !strings.Contains(tracker.GetHTML(), `<img src="https://track.example.com/pmta/`) ||
					!strings.Contains(tracker.GetHTML(), `</body>`) {
					t.Errorf("Tracking pixel not correctly inserted before </body>")
				}
			} else if strings.Contains(tt.fields.originalMailHTML, "</html>") {
				if !strings.Contains(tracker.GetHTML(), `<img src="https://track.example.com/pmta/`) ||
					!strings.Contains(tracker.GetHTML(), `</html>`) {
					t.Errorf("Tracking pixel not correctly inserted before </html>")
				}
			}
		})
	}
}
