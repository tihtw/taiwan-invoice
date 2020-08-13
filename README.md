# Taiwan Invoice
Golang SDK for Taiwan E-Invoice API

這個 SDK 是提供開發者串接以及測試台灣財政部的電子發票系統時可以更加快速以及方便所使用的。

如果有任何疑問歡迎開立 ISSUE。


SDK 功能部分分為兩個大類：
1. 電子發票應用API
2. 電子發票營業人應用 API


其中在第二類營業人應用 API 需要向財政部申請並設定 Client 端 IP 才能進行連線，這點請特別注意。

# 實作狀態

* 電子發票營業人應用 API
[v] 手機條碼驗證
[v] 捐贈碼驗證
[ ] 營業人個別化主題
[ ] 信用卡載具卡 BIN 查詢
[ ] 信用卡載具發卡銀行查詢
[ ] 查詢該營業人是否設定接收方式
[v] 查詢該統一編號是否為營業人


# 關於 IP 限制

在營業人應用 API 的部分會限制連線 IP, 可能會影響開發測試，因此開發者可以自行設定 `ConnectionHost` 利用 `socat` 或者是 SSH 反向隧道進行，另外需要設定 `ConnectionHost` 作為連線用的 host.

舉例來說：
在 example.tw 主機設定了 
```
socat tcp-listen:1237,reuseaddr,fork tcp:www-vc.einvoice.nat.gov.tw:443,bind=103.0.0.0
```

如此連線至 example.tw 的主機會利用 103.0.0.0 的 IP 連線到指定主機 (HOST)。

在 SDK 中可以設定 `ConnectionHost` 為 `example.tw:1237` 來讓 SDK 先連線到該主機。


# 參照規格

* 電子發票應用 API 規格 版本: 1.7
* 電子發票營業人應用 API 規格: 1.6


# 其他相關連結

這邊整理了其他相關的電子發票 API, 如果不一定是 Golang 的 SDK, 如果有興趣放在這個列表中請發送 Pull Request 或者是開立 Issue, 感謝。

* [PichuChen/einvoice](https://github.com/PichuChen/einvoice) PHP 版的電子發票 SDK
* [ShenTengTu/node-tw-e-invoice](https://github.com/ShenTengTu/node-tw-e-invoice) node.js 版的電子發票 SDK



