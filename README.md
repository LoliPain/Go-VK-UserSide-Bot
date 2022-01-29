# Go-VK-UserSide-Bot
Порт с [Python](https://github.com/P2LOVE/VK-UserSide-Bot "Python") на GoLang бота для удаления соообщений используя триггер.

Скорость работы ~в 5 раз быстрее чем у Python-версии. 

## Функционал
В отличии от [Python бота](https://github.com/P2LOVE/VK-UserSide-Bot "Python бота") реализован только функционал для удаления сообщений используя триггер и их редактирование если после триггера указан `-`

![](https://lolipa.in/static/img/opera_qJohlQoHf4.gif)

## Установка и использование
Главная зависимость - библиотека `vksdk`

##### Для установки VKSDK со стандартным GoPath и GoRoot используйте
```bash
go get github.com/SevereCloud/vksdk/v2@latest
```

##### Клонируйте репозиторий 
```bash
git clone https://github.com/P2LOVE/Go-VK-UserSide-Bot
cd Go-VK-UserSide-Bot
```

###### Далее в файле `del-msg-vk.go` настройте значение `VKToken` в строке №13, заполнив её токеном с доступом к личным сообщениям юзера. Рекоммендую использовать для получения токена сервис [VK Host](https://vkhost.github.io "VK Host")

###### И в этом же файле заполните значение `DeleteTrigger` на строке №17, записав туда триггер которым вы хотите инициировать удаления сообщений.

![Пример настроенного скрипта](http://lolipa.in/static/img/mintty_CAMsyEovLS.png "Пример настроенного скрипта")

##### И для запуска используйте
```bash
go mod init vk-bot
go mod tidy

go run del-msg-vk.go
```

Учтите что при таком запуске скрипт выключится вместе с закрытием терминала.
Для запуска в фоне используйте что-то из `nohup`, `screen`, `disown`

## О найденных багах сообщать в раздел Issue:
https://github.com/P2LOVE/Go-VK-UserSide-Bot/issues
