package templates

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Email handling functions
var DefaultManager *EmailTemplateManager

// InitializeDefaultManager sets up the global template manager
func InitializeDefaultManager() error {
	// Get logo URL from environment
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		// Try to derive from backend URL
		backendURL := os.Getenv("BACKEND_URL")
		if backendURL != "" {
			// Extract domain from backend URL
			backendDomain := strings.TrimPrefix(strings.TrimPrefix(backendURL, "https://"), "http://")
			if strings.HasPrefix(backendDomain, "api.") {
				backendDomain = backendDomain[4:] // Remove "api." prefix if present
			}
			frontendURL = "https://app." + backendDomain
		} else {
			// Default fallback
			frontendURL = "https://app.curiousdev.net"
		}
	}
	
	logoURL := frontendURL + "/logo.png"
	log.Printf("Email template manager initializing with logo URL: %s", logoURL)
	
	// Look for templates in a standard location
	templatePath := filepath.Join("internal", "templates", "emails")
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		// Try alternate path
		templatePath = filepath.Join(".", "internal", "templates", "emails")
		if _, err := os.Stat(templatePath); os.IsNotExist(err) {
			return fmt.Errorf("template directory not found at %s or %s", 
				filepath.Join("internal", "templates", "emails"),
				filepath.Join(".", "internal", "templates", "emails"))
		}
	}
	
	var err error
	DefaultManager, err = NewEmailTemplateManager(templatePath)
	return err
}

// RenderSignRequestEmail renders the signing request email template
func RenderSignRequestEmail(recipientName, documentTitle, signingURL string) (string, string, error) {
	if DefaultManager == nil {
		if err := InitializeDefaultManager(); err != nil {
			return getFallbackSignRequestEmail(recipientName, documentTitle, signingURL), 
				"Please sign your RentDaddy lease agreement", nil
		}
	}
	
	data := EmailTemplateData{
		RecipientName: recipientName,
		DocumentTitle: documentTitle,
		SigningURL:    signingURL,
	}
	
	html, err := DefaultManager.RenderTemplate("sign_request", data)
	if err != nil {
		return getFallbackSignRequestEmail(recipientName, documentTitle, signingURL), 
			"Please sign your RentDaddy lease agreement", nil
	}
	
	subject := DefaultManager.GetTemplateSubject("sign_request")
	return html, subject, nil
}

// RenderSigningCompleteEmail renders the signing complete email template
func RenderSigningCompleteEmail(recipientName, documentTitle, downloadURL string) (string, string, error) {
	if DefaultManager == nil {
		if err := InitializeDefaultManager(); err != nil {
			return getFallbackSigningCompleteEmail(recipientName, documentTitle, downloadURL), 
				"Your RentDaddy lease agreement has been signed", nil
		}
	}
	
	data := EmailTemplateData{
		RecipientName: recipientName,
		DocumentTitle: documentTitle,
		DownloadURL:   downloadURL,
	}
	
	html, err := DefaultManager.RenderTemplate("signing_complete", data)
	if err != nil {
		return getFallbackSigningCompleteEmail(recipientName, documentTitle, downloadURL), 
			"Your RentDaddy lease agreement has been signed", nil
	}
	
	subject := DefaultManager.GetTemplateSubject("signing_complete")
	return html, subject, nil
}

// Fallback templates
func getFallbackSignRequestEmail(recipientName, documentTitle, signingURL string) string {
	logoURL := os.Getenv("FRONTEND_URL")
	if logoURL == "" {
		logoURL = "https://app.curiousdev.net"
	}
	logoURL += "/logo.png"
	
	return fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
  <style>
    body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
    .container { max-width: 600px; margin: 0 auto; padding: 20px; }
    .header { text-align: center; margin-bottom: 20px; }
    .logo { max-width: 150px; }
    .button { display: inline-block; background-color: #4CAF50; color: white; padding: 12px 24px; 
              text-decoration: none; border-radius: 4px; font-weight: bold; margin-top: 20px; }
    .footer { margin-top: 30px; font-size: 12px; color: #777; text-align: center; }
  </style>
</head>
<body>
  <div class="container">
    <div class="header">
      <img src="%s" alt="RentDaddy Logo" class="logo">
      <h2>Your Lease Agreement is Ready to Sign</h2>
    </div>
    
    <p>Hello %s,</p>
    
    <p>Your lease agreement <strong>"%s"</strong> is ready for your signature.</p>
    
    <p>Please review the document carefully before signing. Once all parties have signed, you'll receive a final copy of the completed document.</p>
    
    <div style="text-align: center;">
      <a href="%s" class="button">Sign Document Now</a>
    </div>
    
    <p>If you have any questions about this document, please contact your property manager.</p>
    
    <p>Thank you,<br>
    The RentDaddy Team</p>
    
    <div class="footer">
      <p>© 2025 RentDaddy - Making property management easy</p>
    </div>
  </div>
</body>
</html>`, logoURL, recipientName, documentTitle, signingURL)
}

func getFallbackSigningCompleteEmail(recipientName, documentTitle, downloadURL string) string {
	logoURL := os.Getenv("FRONTEND_URL")
	if logoURL == "" {
		logoURL = "https://app.curiousdev.net"
	}
	logoURL += "/logo.png"
	
	return fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
  <style>
    body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
    .container { max-width: 600px; margin: 0 auto; padding: 20px; }
    .header { text-align: center; margin-bottom: 20px; }
    .logo { max-width: 150px; }
    .success-icon { font-size: 48px; color: #4CAF50; text-align: center; margin: 20px 0; }
    .button { display: inline-block; background-color: #4CAF50; color: white; padding: 12px 24px; 
              text-decoration: none; border-radius: 4px; font-weight: bold; margin-top: 20px; }
    .footer { margin-top: 30px; font-size: 12px; color: #777; text-align: center; }
  </style>
</head>
<body>
  <div class="container">
    <div class="header">
      <img src="%s" alt="RentDaddy Logo" class="logo">
      <h2>Lease Agreement Completed!</h2>
    </div>
    
    <div class="success-icon">✓</div>
    
    <p>Hello %s,</p>
    
    <p>Great news! Your lease agreement <strong>"%s"</strong> has been signed by all parties and is now complete.</p>
    
    <p>You can access the signed document through your RentDaddy account or by clicking the button below.</p>
    
    <div style="text-align: center;">
      <a href="%s" class="button">Download Signed Document</a>
    </div>
    
    <p>If you have any questions about this document, please contact your property manager.</p>
    
    <p>Thank you,<br>
    The RentDaddy Team</p>
    
    <div class="footer">
      <p>© 2025 RentDaddy - Making property management easy</p>
    </div>
  </div>
</body>
</html>`, logoURL, recipientName, documentTitle, downloadURL)
}