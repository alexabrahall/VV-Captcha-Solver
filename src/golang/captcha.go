package captcha

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func SolveVVCaptcha(targetURL string) (string, error) {
	// Parse URL to get domain
	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %v", err)
	}
	domain := parsedURL.Hostname()

	// Create HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/handler/vv_captcha/code", domain), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36")
	req.Header.Set("Referer", targetURL)

	// Make request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// Status code handling
	switch resp.StatusCode {
	case 403:
		return "", fmt.Errorf("invalid URL - Domain verification failed (this domain is likely not a vv captcha enabled domain)")
	case 500:
		return "", fmt.Errorf("failed to verify captcha - likely error in domain handling")
	case 200:
		// Continue processing
	default:
		return "", fmt.Errorf("unexpected error [status: %d] - Please check your domain settings", resp.StatusCode)
	}

	// Read response body
	body := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			body = append(body, buf[:n]...)
		}
		if err != nil {
			break
		}
	}
	captchaData := string(body)

	if strings.TrimSpace(captchaData) == "" {
		return "", fmt.Errorf("failed to fetch captcha data - likely error in domain handling")
	}

	// Regex pattern to extract the captcha id from the response
	// Go regex doesn't support backreferences, so we'll match both quote types separately
	re := regexp.MustCompile(`var\s+token\s*=\s*['"](.*?)['"]\s*;`)
	matches := re.FindStringSubmatch(captchaData)

	if len(matches) < 2 {
		return "", fmt.Errorf("failed to extract captcha id from response")
	}

	captchaID := matches[1]
	return captchaID, nil
}
