package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

// -1003212636286
var captMsg string
var channelId string
var adminIDs []int

func PostingSetup(captMsg_, channelId_, adminIDs_ string) {
	captMsg = captMsg_
	channelId = channelId_

	adminIDs, _ = stringToInt(adminIDs_)
	log.Println(adminIDs_, channelId)
}

func PostingContent(bot *tb.Bot, filePath string, fileType string, userID int) error {
	ch, _ := strconv.ParseInt(channelId, 10, 64)
	channelChatID := tb.ChatID(ch)
	caption := fmt.Sprintf(captMsg, userID)

	isAdmin := false
	for _, adminID := range adminIDs {
		if userID == adminID {
			isAdmin = true
			break
		}
	}

	if isAdmin {
		log.Println("Content from admin. Posting... ", userID)
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

func stringToInt(input string) ([]int, error) {

	numberStrings := strings.Split(input, ",")
	numbers := make([]int, 0, len(numberStrings))

	for _, str := range numberStrings {
		trimmedStr := strings.TrimSpace(str)
		if trimmedStr == "" {
			continue
		}
		num, err := strconv.Atoi(trimmedStr)
		if err != nil {
			return nil, fmt.Errorf("ошибка преобразования '%s': %w", trimmedStr, err)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}
