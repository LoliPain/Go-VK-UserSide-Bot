package main

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/longpoll-user"
	wrapper "github.com/SevereCloud/vksdk/v2/longpoll-user/v3"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)

const VKToken = ""

// Вставьте ваш токен, имеющий доступ к работе с ЛС (Kate Mobile итд)

var DeleteTrigger = strings.ToLower("")

// Введите свой триггер для удаления сообщений внутри кавычек

var VK = api.NewVK(VKToken)
var UserList, _ = VK.UsersGet(api.Params{})
var MyID = UserList[0].ID

func main() {
	GetLongpoll, _ := longpoll.NewLongPoll(VK, 0)
	EventWrapper := wrapper.NewWrapper(GetLongpoll)
	EventWrapper.OnNewMessage(func(message wrapper.NewMessage) {
		MessageText := strings.ToLower(message.ExtraFields.Text)
		MessageObject, _ := VK.MessagesGetByID(api.Params{"message_ids": message.MessageID})
		if strings.HasPrefix(MessageText, DeleteTrigger) && MessageObject.Items[0].FromID == MyID {
			ToDeleteArgument := string([]rune(MessageText)[utf8.RuneCountInString(DeleteTrigger):])
			if _, err := strconv.Atoi(ToDeleteArgument); err == nil {
				if strings.HasPrefix(ToDeleteArgument, "-") {
					toDeleteArgument, _ := strconv.Atoi(ToDeleteArgument[1:])
					DeleteMsg(toDeleteArgument, true, message.PeerID)
				} else {
					toDeleteArgument, _ := strconv.Atoi(ToDeleteArgument)
					DeleteMsg(toDeleteArgument, false, message.PeerID)
				}
			}
		}
		if MessageText == DeleteTrigger && MessageObject.Items[0].FromID == MyID {
			DeleteMsg(1, false, message.PeerID)
		}
		if MessageText == DeleteTrigger+"-" && MessageObject.Items[0].FromID == MyID {
			DeleteMsg(1, true, message.PeerID)
		}
	})
	if err := GetLongpoll.Run(); err != nil {
		log.Fatal(err)
	}
}

func DeleteMsg(Count int, Redact bool, PeerId int) {
	var MessageIDs []string
	MessageHistory, _ := VK.MessagesGetHistory(api.Params{"peer_id": PeerId})
	for _, Message := range MessageHistory.Items {
		if Message.FromID == MyID {
			MessageIDs = append(MessageIDs, strconv.Itoa(Message.ID))
		}
		if len(MessageIDs) >= Count+1 {
			break
		}
	}
	if Redact {
		for _, MessageID := range MessageIDs {
			_, err := VK.MessagesEdit(api.Params{"peer_id": PeerId, "message_id": MessageID, "message": "ᅠ"})
			if err != nil {
				break
			}
		}
	}
	_, _ = VK.MessagesDelete(api.Params{"message_ids": strings.Join(MessageIDs, ","), "delete_for_all": 1})
}
