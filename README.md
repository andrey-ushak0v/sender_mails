# sender mails - небольшой проект, представляющий собой сервис для e-mail рассылок.
если выполнить отправку, то получателю придет сообщение в такого вида "привет <имя> <фамилия>, ты был рожден в <год>

### Инструкция по запуску:

1) Склонируйте репозиторий на локальную машину
   
   ```https://github.com/andrey-ushak0v/sender_mails.git```

2) Создайте с аккаунте google пароль для приложения по которому сервис будет иметь доступ к почте с которой будут рассылатться письма

3) откройте проект в ide и заполните email и пароль, и другие поля(помечены комментариями), а так же наполните хранилище информацией с существующими email адресами

4) из корневой папки проекта в консоли выполните команду go run main.go

5) по эндпоитну  ```http://localhost:5100/```  будет доступен список email-адресов вместе с информацией о хозяине адреса.
  А по ендпоиту ```http://localhost:5100/send``` доступна отправка сообщения на указаный e-mail адресс. В теле Post зпроса нужно передать email-адресс         
   ```{"email": "example@gmail.com"}```
