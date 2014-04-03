Sends a message to Slack for each ticket sale on Amiando.

Needs these environment variables to run:
 * PORT
 * SLACK_URL: something like https://xxxx.slack.com/services/hooks/incoming-webhook?token=yyyyy
 * SLACK_CHANNEL: a channel name to post (like "#tickets")

Procfile is provided for easy deployment on Heroku.

You could then enter "http://xxx.herokuapp.com/amiando-server-call" in Amiando's "Integration" page.

SLACK_URL and SLACK_CHANNEL can also be added as Heroku config variables.