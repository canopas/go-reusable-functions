# Email

Email is important when you want to improve more user engagement on your application. This functionality will keep up-to-date your users about your application.

Amazon SES is a cloud email service provider that can integrate into any application for bulk email sending. Whether you use an email software to send transactional emails, marketing emails, or newsletter emails, you pay only for what you use.

## Steps to use email-reusables
 
- We have used go.embed to parse email templates for emails

## Prerequisites

- Email template in **templates** directory of your **project's root**. 
- AWS credentials
  - Access key id
  - Secret Access key
  - Region
  - AWS session token (If MFA enabled for AWS account)

## Example

Find working example of Email-reusable at [Email-Example](https://github.com/canopas/go-reusables/blob/main/examples/email.go).