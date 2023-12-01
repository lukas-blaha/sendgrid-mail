# sendgrid-mail

## Config
1. You can specify the API key using "SENDGRID_KEY" environment var
2. You can specify the API key using config file

### Environment var
`export SENDGRID_KEY=<your-api-key>`

### Config file
```
echo 'apiKey: "<your-api-key>"' >/etc/sendgrid-mail.conf
chmod 644 /etc/sendgrid-mail.conf
```
