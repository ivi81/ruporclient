**Клиентская библиотека на golang для RUPOR (R-Vision)**

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

