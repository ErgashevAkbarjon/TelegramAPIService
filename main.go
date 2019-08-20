package main

import (
	"log"
	"path/filepath"

	"github.com/zelenin/go-tdlib/client"
)

func withLogs() client.Option {
	return func(tdlibClient *client.Client) {
		tdlibClient.SetLogVerbosityLevel(&client.SetLogVerbosityLevelRequest{
			NewVerbosityLevel: 1,
		})
	}
}

func main() {

	const (
		apiID   = 954242
		apiHash = "ed33628ede24f596c53edbecfd40ca0b"
	)

	authorizer := client.ClientAuthorizer()
	go client.CliInteractor(authorizer)

	authorizer.TdlibParameters <- &client.TdlibParameters{
		UseTestDc:              false,
		DatabaseDirectory:      filepath.Join(".tdlib", "database"),
		FilesDirectory:         filepath.Join(".tdlib", "files"),
		UseFileDatabase:        true,
		UseChatInfoDatabase:    true,
		UseMessageDatabase:     true,
		UseSecretChats:         false,
		ApiId:                  apiID,
		ApiHash:                apiHash,
		SystemLanguageCode:     "en",
		DeviceModel:            "Server",
		SystemVersion:          "1.0.0",
		ApplicationVersion:     "1.0.0",
		EnableStorageOptimizer: true,
		IgnoreFileNames:        false,
	}

	tdlibClient, err := client.NewClient(authorizer, withLogs())
	if err != nil {
		log.Fatalf("NewClient error: %s", err)
	}

	listener := tdlibClient.GetListener()
	defer listener.Close()

	for update := range listener.Updates {

		if update.GetClass() == client.ClassUpdate {

			lastMessage, lastMessageExists := update.(*client.UpdateNewMessage)

			if !lastMessageExists {
				continue
			}

			lastMessageText := getMessageText(lastMessage.Message)

			lastMessageChat := getMessageChat(lastMessage.Message, tdlibClient)

			isLastMessageOutgoing := lastMessage.Message.IsOutgoing

			// if !isLastMessageOutgoing {
			// 	sendMessage(lastMessageText, lastMessageChat.Id, tdlibClient)
			// 	sendMessage("asdasdasdasdasd", lastMessageChat.Id, tdlibClient)
			// }

			log.Printf("\nChat: %v\nMessage: %v\nOutgoing?: %v", lastMessageChat.Title, lastMessageText, isLastMessageOutgoing)

		}

	}
}

func getMessageText(message *client.Message) string {
	textContent, hasTextContent := message.Content.(*client.MessageText)

	if !hasTextContent {
		return ""
	}

	return textContent.Text.Text
}

func getMessageChat(message *client.Message, chatClient *client.Client) *client.Chat {
	chatRequest := client.GetChatRequest{
		ChatId: message.ChatId}

	chat, err := chatClient.GetChat(&chatRequest)

	if err != nil {
		log.Printf("Error while getting chat: %v", err)
		return nil
	}

	return chat
}

func createNewMessage(messageText string, chatID int64) *client.SendMessageRequest {

	inputMessage := client.InputMessageText{

		Text: &client.FormattedText{
			Text: messageText}}

	sendMessageRequest := client.SendMessageRequest{
		ChatId:              chatID,
		InputMessageContent: &inputMessage}

	return &sendMessageRequest
}

func sendMessage(messageText string, chatID int64, senderClient *client.Client) {

	messageRequest := createNewMessage(messageText, chatID)

	_, err := senderClient.SendMessage(messageRequest)

	if err != nil {
		log.Printf("An error occured in sending message: %v", err)
	}
}
