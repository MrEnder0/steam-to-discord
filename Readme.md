# Steam to Discord

A customizable webhook relay written in Go to forward messages and other updates from your Steam group directly to your Discord server.

## Note

This is the arm branch of the project, this is meant to be run on a Raspberry Pi or other ARM-based devices. Specific things have been changed to make this work on ARM devices such as the database system, because of this, this branch may take longer to get updates than the main branch so if you can run the main branch on your device I would recommend doing so.

## Features

* Automated Posting: Fetches new Steam group messages at a configurable interval.
* TOML Configuration: Easy-to-edit config file for customization.
* Steam Profile Details: Fetches and displays the profile name and picture of the user who posted the message on the steam group.

## Configuration

This is the base config.toml file which is generated on inital runtime, this is meant to be modified to work with your own use case.

```toml
checkfreq = 60 # How often to check for new Steam announcements (in seconds)
groupname = 'steamgroupname' # The name of your Steam group, this is found in the URL of your group
showsteamprefix = true # Adds a "[Steam]" prefix to the webhook username (true/false)
webhookurl = 'https://discord.com/api/webhooks/id' # Your Discord webhook URL, this can be found in the settings of your Discord server
```

## Usage

1. Download the latest release from the [releases page](https://github.com/MrEnder0/steam-to-discord).
2. Run the executable, this will automatically generate a config.toml file and exit.
3. Edit the config.toml file to your liking, make sure to replace the placeholder group name and webhook URL with your own or the program will not work.
4. Run the executable again, this time it will start the relay and begin fetching messages from your Steam group.
