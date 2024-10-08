# uservice-bot
Ready-made bot to build your own Digital Business in the Utopia Ecosystem

TODO: оглавление.
TODO: сделать README на русском и английском.

## как на основе этого построить свой бизнес

Допустим, у тебя есть своя услуга, которую ты готов оказывать пользователям. Это решение поможет тебе автоматизировать:
- Общение с пользователями
- Прием заявок
- Оплату твоих услуг
- Сбор клиентов в твой канал

Клиенты будут производить оплату и записываться на твою услугу. Тебе будут приходить оповещения о новых заказах на uMail, далее тебе остается только связаться с пользователем для осуществления услуги. Бот будет автоматически выводить тебе средства, полученные от клиентов.

Все что тебе остается - это приводить клиентов к боту.

Далее я привожу инструкцию как начать.

## как настроить и запустить бота

Рекомендую использовать систему Linux. Потребуется установленный docker.

Далее скачиваем или клонируем репозиторий:

```bash
git clone https://github.com/Sagleft/uservice-bot
```

Далее необходимо создать файл аккаунта Utopia `account.db` и поместить его в папку `data`. Аккаунту надо назначить пароль `password`.

Можно в аккаунте указать определенное имя, аватар для бота.

Далее копируем файл настроек. Например, командой:

```bash
cp data/config_example.yml data/config.yml
```

Далее вносим данные в файл `config.yml`. Подробнее про параметры можно прочесть в документации ниже.

И запускаем систему в фоновом режиме:

```bash
docker-compose up -d
```

проверить, что всё ОК можно через

```bash
docker ps
```

у контейнеров, связанных с приложением, не должно быть статуса `Restarting`.

## Документация

Тут будет разбор настроек.

## Разработка

Помощь проекту приветствуется!

Чат в Utopia по разработке: https://utopia.im/RUTECH

1. сделай форк репозитория
2. создай свою ветку
3. внеси изменения
4. открой Pull Request к моему репозиторию

После изменений пересобери приложение через

```bash
docker-compose build
```

запуск в фоновом режиме:

```bash
docker-compose up -d
```
