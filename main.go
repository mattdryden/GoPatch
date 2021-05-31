package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Patch Reminder Utility by Matt Dryden")

	config := LoadConfig()
	checkPendingReminders(config)
}

func checkPendingReminders(config *Config) {
	if time.Now().After(config.Patches.Seed) {
		sendReminder(config.Patches.Seed, config)
	} else {
		fmt.Printf("Not due until %s", config.Patches.Seed)
	}
}

func sendReminder(due time.Time, config *Config) {
	var err error
	var success bool
	quote, err := getQuotes(config)

	if err != nil {
		quote = ""
	}

	for _, recipient := range config.Recipients {
		if recipient.Key != "" {
			message := fmt.Sprintf("Hello %s, your patch is due as of %s. Feel free to ask Matt!\n\n%s", recipient.Name, due.Format(time.ANSIC), quote)
			notification := Notification{user: recipient.Key, token: config.Pushover.Token, message: message}
			success, err = sendNotification(&notification, config)

			if err != nil {
				fmt.Errorf("Error: %q", err)
			}
		}
	}

	if success {
		seed(due, config)
	}
}

func seed(due time.Time, config *Config) {
	nextDue := due.Add(config.Patches.Interval)
	config.Patches.Seed = nextDue

	_, err := SaveConfig(config)

	if err != nil {
		panic(err)
	} else {
		fmt.Print("Next reminder scheduled!")
	}
}
