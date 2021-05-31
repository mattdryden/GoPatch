package main

import (
	"fmt"
	"time"
)

type Reminder struct {
	Message *string
}

func main() {
	fmt.Println("Patch Reminder Utility")

	config := LoadConfig()

	fmt.Println(config)

	checkPendingReminders(config)
}

func checkPendingReminders(config *Config) {

	SeedDate, err := time.Parse(time.RFC3339, config.Patches.Seed)
	if err != nil {
		fmt.Errorf("error: %q", err)
	}

	if time.Now().After(SeedDate) {
		sendReminder(SeedDate, config)
	} else {
		fmt.Printf("Not due until %s", SeedDate)
	}

}

func sendReminder(due time.Time, config *Config) {
	fmt.Printf("Reminder due %s", due)
	notification := Notification{}
	notification.user = config.Recipients.One.Key
	notification.message = "Patch Reminder"
	sendNotification(&notification)
}

// func seed(seed time.Time, interval string) {
// 	var err error
// 	interval, err = time.ParseDuration(fmt.Sprintf("%s", interval))
// 	if err != nil {
// 		fmt.Errorf("error: %q", err)
// 	}
// 	fmt.Printf("Seeding... %s", interval)

// 	fmt.Println(seed.Add(interval))
// }
