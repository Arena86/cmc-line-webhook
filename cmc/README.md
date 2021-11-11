# Webhook receiver to receive LINE event

## Introduction
- Write a webhook receiver to receive LINE event then reply a message back (e.g. “hello {full name} : you said {event message}”)
- persist log events into MongoDB or MySQL

## Documentation

- Import cmc_line.sql file to Mysql
- Change mysql connect information in .env file
- Change line secrect, token in .env file
- Signup line account: https://developers.line.biz/en/
- Signin line account
- Create a provider
- Create a channel
- Get channel secrect: choose channel > Basic settings tab > Channel secret
- Get channel access token: choose channel > Messaging API tab > Channel access token (long-lived) > issue button
- Setting webhook: choose channel > Messaging API tab > Webhook URL > add webhook url > verify
- Use webhook: choose channel > Messaging API tab > Use webhook > enable

## Testing

- I used https://ngrok.com/ for localhost https
- User testing: choose channel > statistics tab > click link "LINE Official Account Manager" > Home tab > Gain friends > scan QR code to chat

## Requirements

This library requires Go 1.11 or later.


