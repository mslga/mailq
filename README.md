# mailq

Lightweight email-sending utility written in Go

## Usage

### Launch options and defaults

| Опция    | Описание   | Значение по умолчанию   |
|---------------|------------|------------|
|   -h   |  SMTP server FQDN |  "smtp.company.com""  |
|   -p   |  SMTP server port number  |  25  |
|   -i   |  Skip TLS certificate verification (use only for trusted SMTP servers) |  false  |
|   -s   |  Mail sender |  "mailq.sender@company.com"  |
|   -w   |  SMTP password (optional) |  ""  |
|   -r   |  Mail recipients. Separate multiple recipients with spaces or commas |  "recipient@company.com"  |
|   -u   |  Mail subject  |  "Empty subject"  |
|   -m   |  Mail message in HTML format  |  "Empty message"  |

### Launch examples

* Without auth:

```bash
mailq \
  -h mail.company.local \
  -p 25 \
  -i true
  -s noreply@company.local \
  -r "user1@company.local,user2@company.local" \
  -u "Test" \
  -m "<h1>Hello!</h1>"
```

* With auth:

```bash
mailq \
  -h smtp.office365.com \
  -p 587 \
  -s noreply@company.com \
  -w "MySecretPassword" \
  -r user@example.com \
  -u "Test" \
  -m "<h1>Hello!</h1>"
```

### Docker image

```bash
docker pull mslga/mailq:v2.0.1
```
