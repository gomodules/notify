package main

import (
	"fmt"
	"gomodules.xyz/notify/unified"
)
/*
	Please make sure to set below environment variables

	TINIYO_AUTH_ID
	- Login to tiniyo.com and get your auth_id

	TINIYO_AUTH_TOKEN
	- Login to Tiniyo.com and get your auth_token

	TINIYO_FROM -
	- In case of TINIYO_NOTIFY_CHANNEL=voice,Please use your verified phone-number
		or purchased phone number from tiniyo
	- In case of TINIYO_NOTIFY_CHANNEL=sms,Please use your approved senderid from tiniyo

	TINIYO_NOTIFY_CHANNEL
	- sms or voice

 */
func main() {
	fmt.Println("We are going to send sms via tiniyo")
	err := unified.NotifyViaDefault("tiniyo", "This is my notification via tiniyo", "", "", "REPLACE_WITH_YOUR_DESTINATION")
	if err != nil {
		fmt.Println("Error is there", err)
	}
}
