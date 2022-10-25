# Mirai-Backup-GroupFriends

鉴于某开发群被和谐很多群友都失去了联系，虽然创建了新的群。但是很多群友联系方式都没有了，所以有了这个项目。用于 QQ 群备份群友联系方式。

## 使用 🛴

在项目根目录创建 `device.json` 文件。输入如下内容（或者通过 MiraiAndroid 项目获取 device.json）

> 注: 该设备信息文件由开发者随机生成，有可能具有风险请酌情使用。

```
{
	"display": "MIRAI.233092.001",
	"product": "mirai",
	"device": "mirai",
	"board": "mirai",
	"brand": "mamoe",
	"model": "mirai",
	"bootloader": "unknown",
	"fingerprint": "mamoe/mirai/mirai:10/MIRAI.200122.001/6812682:user/release-keys",
	"boot_id": "CAFBD15D-C451-B990-69F3-DB480F8648C6",
	"proc_version": "Linux version 3.0.31-Ix9Kw9Eo (android-build@xxx.xxx.xxx.xxx.com)",
	"base_band": "",
	"version": { "incremental": "5891938", "release": "10", "codename": "REL" },
	"sim_info": "T-Mobile",
	"os_type": "android",
	"mac_address": "02:00:00:00:00:00",
	"wifi_bssid": "02:00:00:00:00:00",
	"wifi_ssid": "<unknown ssid>",
	"imsi_md5": "",
	"imei": "695468231243493",
	"apn": "wifi"
}

```

运行
```
go run .
```

扫描二维码登录后，程序将会备份登录的账号全部QQ群群友列表到 `group_users` 目录下，以 `群号_群名.txt` 文件格式保存联系方式。


## 鸣谢 🦾
 - [MiraiGo-Template](https://github.com/Logiase/MiraiGo-Template)
    MiraiGo 脚手架
 - [MiraiGo](https://github.com/Mrs4s/MiraiGo)
    核心协议库
 - [viper](https://github.com/spf13/viper)
    用于解析配置文件，同时可监听配置文件的修改
 - [logrus](https://github.com/sirupsen/logrus)
    功能丰富的Logger
 - [asciiart](https://github.com/yinghau76/go-ascii-art)
    用于在console显示图形验证码
