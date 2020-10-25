# learn Slack

## learn webhook.

1. Access [Incoming WebHooks](https://slack.com/services/new/incoming-webhook).
1. Choose a channel.
1. Copy the Webhook URL.
1. Post message.
    `$ curl -d '{"text":"Hello, World!"}' https://hooks.slack.com/services/<YOUR WEBHOOK URL>`
1. Check a message on Slack.
