## sender mails - небольшой проект, представляющий собой сервис для e- mail рассылок.

## инструкция по запуску:

1) склонируйте репозиторий на локальную машину

2) создайте с аккаунте  google пароль для прилодения. по которому сервис будет иметь доступ к посте с которой будут рассылатться письма

3) откройте проект в ide и заполните email и пароль, и другие поля(помечены комментариями), а так же наполните хранилище информацией с существубщими email адресами

4) из корневой папки проекта в консоли выполните команду go run main.go

5)по эндпоитну http://localhost:5100/ будет доступен список email-адресов вместе с информацией о хозяине адреса (
  по ендпоиту http://localhost:5100/send доступна отправка сообщения на указаный e-mail адресс. В теле Post зпроса нужно передать email-адресс {"email": "example@gmail.com"}
