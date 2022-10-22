package client

func emitEvents(update map[string]interface{}, offset *int, bot Bot) {
	for key, value := range update {
		// setting offset to only react on new updates
		if key == "update_id" {
			*offset = int(value.(float64)) + 1
			continue
		}
		//emitting general events like message, channel post, callback query, etc
		bot.emit(key)

		// variables to safely emit event
		var text string
		var ofst, length int

		//checking additional events like text, location, dice, photo, etc
		for key, value = range value.(map[string]interface{}) {
			bot.emit(key)
			//text value for text messages, commands, etc
			if key == "text" {
				text = value.(string)
			}
			//enitity type and text value for message entities (commands, links, etc)
			if key == "entities" {
				for _, rawEntity := range value.([]interface{}) {
					entity := rawEntity.(map[string]interface{})
					ofst = int(entity["offset"].(float64))
					length = int(entity["length"].(float64))

					bot.emit(entity["type"].(string))
				}
			}
		}
		//emitting text event for more flexibility

		// if length variable was assigned just emitting part that triggered enitity
		if length != 0 {
			bot.emit(text[ofst : length+ofst])
			return
		}
		// otherwise emitting text from the message
		bot.emit(text)
	}
}
