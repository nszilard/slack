# Slack

[![CI](https://github.com/nszilard/slack/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/nszilard/slack/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/nszilard/slack)](https://goreportcard.com/report/github.com/nszilard/slack)
[![GoDoc](https://godoc.org/github.com/nszilard/slack?status.svg)](https://godoc.org/github.com/nszilard/slack)

---

## About

`slack` is a simple CLI utility to send Slack messages based on a pre-defined templates, but it can also be used as a library.  
Inspired by the [slack-go/slack](https://github.com/slack-go/slack) project, but wanted something simpler for personal use.

## Installation

``` shell
go get -u github.com/nszilard/slack
```

## Usage

It supports both environment variables as well as flags to set values, where the latter one takes priority if the values differ.

### Main command

| Environment variable  | Flag                  | Description                                                                                   |
| :---------------------|:--------------------- |:----------------------------------------------------------------------------------------------|
| `SLACK_ORG_ID`        | `slack-org-id`        | Slack Organization ID. From the webhhok token **`Txxxxxx`**`/Bxxxxxx/xxxxxxx` the first group |
| `SLACK_WEBHOOK_ID`    | `slack-webhook-id`    | Slack Webhhok ID. From the webhhok token `Txxxxxx/`**`Bxxxxxx`**/`xxxxxxx` the second group   |
| `SLACK_WEBHOOK_TOKEN` | `slack-webhook-token` | Slack Webhhok token. From the webhhok token `Txxxxxx/Bxxxxxx/`**`xxxxxxx`** the third group   |
| `SLACK_CHANNEL`       | `slack-channel`       | Slack channel name (or personal ID) to send the message to                                    |
| `SLACK_USER`          | `slack-user`          | To specify the username for the published message.                                            |
| `SLACK_USER_IMAGE`    | `slack-user-image`    | To specify a URL to an image to use as the profile photo alongside the message                |

### Changelog command

| Environment variable  | Flag                | Description                                             |
| :---------------------|:------------------- |:--------------------------------------------------------|
| `CHANGELOG_SERVICE`   | `changelog-service` | Name of the Git repositry or preferred service name     |
| `CHANGELOG_VERSION`   | `changelog-version` | The new version                                         |
| `CHANGELOG_COMMITS`   |                     | Commits which have been released with the given version |

### See generated docs

* [slack](docs/slack.md) - CLI to send Slack messages programmatically.
