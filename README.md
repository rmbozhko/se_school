# se_school

### Currency exchange third-party service
Please, visit https://app.exchangerate-api.com/sign-up to create an account which comes with an API key that should be provided to `API_TOKEN` environment variable.
If you are running the service with `docker-compose.yaml`, then specify the value in api service's environment section. 

### SMTP
Specify the settings of your SMTP server setup in `docker-compose.yaml`. Required properties to provide are:
```
SMTP_HOST - SMTP server host, e.g. smtp.gmail.com
SMTP_PORT - SMTP server port, e.g. 587
SMTP_USER - SMTP server user, e.g. <your Gmail account>
SMTP_PASSWORD - SMTP server user's password <your Gmail account password>
```

In case you experience the error related to Gmail SMTP server:
```
Error: Invalid login: 534-5.7.9 Application-specific password required
```
then follow these steps: https://stackoverflow.com/a/60718806

