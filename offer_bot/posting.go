package main

import (
	"fmt"
	"log"

	tb "gopkg.in/tucnak/telebot.v2"
)

// -1003212636286
var captMsg string

var adminIDs []int

func PostingSetup(captMsg_ string) {
	captMsg = captMsg_
	adminIDs = append(adminIDs, 7149937853)
}

func PostingContent(bot *tb.Bot, filePath string, fileType string, userID int) error {

	channelChatID := tb.ChatID(-1003212636286)
	log.Println("get + " + captMsg)
	caption := fmt.Sprintf("Контент от пользователя: %d \n"+captMsg, userID)

	isAdmin := false
	for _, adminID := range adminIDs {
		if userID == adminID {
			isAdmin = true
			break
		}
	}

	if isAdmin {
		log.Printf("Content from admin. Posting... ", userID)
		return publishContent(bot, channelChatID, filePath, fileType, caption)
	} else {
		log.Printf("%d User's contents sent to admins", userID)
		return moderateContent(bot, filePath, fileType, userID)
	}
}

func publishContent(bot *tb.Bot, channelChatID tb.ChatID, filePath string, fileType string, caption string) error {
	var err error

	switch fileType {
	case "photo":
		photo := &tb.Photo{File: tb.FromDisk(filePath), Caption: caption}
		_, err = bot.Send(channelChatID, photo, markdown)

	case "gif":
		animation := &tb.Animation{File: tb.FromDisk(filePath), Caption: caption}
		_, err = bot.Send(channelChatID, animation, markdown)

	case "video":
		video := &tb.Video{File: tb.FromDisk(filePath), Caption: caption}
		_, err = bot.Send(channelChatID, video, markdown)

	default:
		return fmt.Errorf("неизвестный тип файла: %s", fileType)
	}

	if err != nil {
		log.Printf("Ошибка при публикации в канал: %v", err)
		return err
	}

	log.Printf("Контент успешно опубликован в канале: %s", filePath)
	return nil
}

func moderateContent(bot *tb.Bot, filePath string, fileType string, userID int) error {
	// TODO: Реализовать функционал модерации позже
	log.Printf("Content from user %d sent to admins: %s (%s)", userID, filePath, fileType)

	log.Printf("[MODERATION] File: %s, Type: %s, UserID: %d", filePath, fileType, userID)

	return nil
}
