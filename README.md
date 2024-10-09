**Клиентская библиотека на golang для RUPOR (R-Vision)**

**Установка данного модуля**:
  * Установка последней версии
      *go get gitlab.cloud.gcm/i.ippolitov/go-ruporclient*
  * Установка конкретной версии 
      *go get gitlab.cloud.gcm/i.ippolitov/go-ruporclient@tag_num*
  * Установка вместе с последним коммитом из main 
     *go get gitlab.cloud.gcm/i.ippolitov/go-ruporclient@main*
  * Обновление до последней минорной версии
     *go get -u gitlab.cloud.gcm/i.ippolitov/go-ruporclient*

Структура проекта:
```
├── config
│   ├── config.go
│   └── const.go
├── go.mod
├── go.sum
├── README.md
└── rupor
    ├── api
    │   ├── apiClient.go
    │   ├── cons
    │   │   ├── activitystatus_enummethods.go
    │   │   ├── activityStatus.go
    │   │   ├── const.go
    │   │   ├── constType.go
    │   │   ├── doccategory_enummethods.go
    │   │   ├── docCategory.go
    │   │   ├── endpoint.go
    │   │   ├── filterCommentField.go
    │   │   ├── filtercommentfields_enummethods.go
    │   │   ├── filternotiffield_enummethods.go
    │   │   ├── filterNotifField.go
    │   │   ├── malwareFunc.go
    │   │   ├── messageStatus.go
    │   │   ├── msgstatus_enummethods.go
    │   │   ├── notificationStatus.go
    │   │   └── notifstatus_enummethods.go
    │   ├── export
    │   │   ├── observ.go
    │   │   ├── simpletype.go
    │   │   └── type.go
    │   ├── param
    │   │   ├── field.go
    │   │   ├── filter.go
    │   │   ├── param.go
    │   │   └── sort.go
    │   ├── resp
    │   │   ├── handler.go
    │   │   └── type.go
    │   └── type.go
    ├── config
    │   ├── config.go
    │   └── const.go
    └── logger
        └── loggerinterface.go
```
* **config** - пакет содержащий структуры данных в которых хранятся настройки конфигураци api\-клиента\, такие как url\, credentials и т\.д\.
* **rupor\api\cons** - пакет хранящий описание констант и перчислений используемых в качаестве значений полей документов.

* **logger** - описание интерфеса логирования. Любое програмное средство логирования котрое может применятся в коде должно соответствовать фкнциональному поведению описанному данным интерфейсом.

Примеры использования:

