# Email Templates for Documenso Worker

This directory contains email templates used by the documenso-worker to send notifications for lease signing actions.

## Available Templates

1. `sign_request.html` - Email template for requesting a lease signature
2. `signing_complete.html` - Email template for notifying when a lease has been signed
3. `config.json` - Configuration for templates including subjects and filenames

## Template Data Variables

The following variables are available in templates:

- `{{.LogoURL}}` - URL to the RentDaddy logo
- `{{.RecipientName}}` - Name of the recipient
- `{{.DocumentTitle}}` - Title of the document
- `{{.SigningURL}}` - URL for signing the document
- `{{.DownloadURL}}` - URL for downloading the signed document

## How to Add a New Template

1. Create your HTML template file in this directory
2. Add template configuration to `config.json`
3. Update the template manager code in main.go to handle the new template type

## Fallback Behavior

If templates can't be loaded from this directory, the worker will fall back to using built-in templates.