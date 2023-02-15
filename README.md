<p align="center"><img src="https://avatars.githubusercontent.com/u/91626055?v=4" width="128" /></p>

<div align="center">

# AikoXrayR
XrayR AikoCute Projects

[![](https://img.shields.io/badge/Telegram-group-green?style=flat-square)](https://t.me/AikoXrayR)
[![](https://img.shields.io/badge/Telegram-channel-blue?style=flat-square)](https://t.me/AikoCute_Support)
[![](https://img.shields.io/github/downloads/AikoXrayR-Project/XrayR/total.svg?style=flat-square)](https://github.com/AikoXrayR-Project/XrayR/releases)
[![](https://img.shields.io/github/v/release/AikoXrayR-Project/XrayR?style=flat-square)](https://github.com/AikoXrayR-Project/XrayR/releases)
[![docker](https://img.shields.io/docker/v/aikocute/xrayr?label=Docker%20image&sort=semver)](https://hub.docker.com/r/aikocute/xrayr)
[![Go-Report](https://goreportcard.com/badge/github.com/AikoAikoXrayR-Project/XrayR?style=flat-square)](https://goreportcard.com/report/github.com/AikoXrayR-Project/XrayR)
</div>


# Mô tả về AikoXrayR
XrayR Aiko Hỗ trợ Nhiều Panel khác nhau ( V2board , ProxyPanel, sspanel, Pmpanel ...)

Một khung công tác back-end dựa trên Xray, hỗ trợ các giao thức V2ay, Trojan, Shadowsocks, cực kỳ dễ dàng mở rộng và hỗ trợ kết nối nhiều bảng điều khiển。

Nếu bạn thích dự án này, bạn có thể nhấn vào dấu sao + xem ở góc trên bên phải để theo dõi tiến độ của dự án này.

Sử dụng hướng dẫn：[Hướng dẫn chi tiết](https://xrayr.aikocute.com)
## Tuyên bố từ chối trách nhiệm

Dự án này chỉ mang tính chất học tập, phát triển và bảo trì của cá nhân tôi, tôi không đảm bảo tính khả dụng và tôi không chịu trách nhiệm về bất kỳ hậu quả nào do sử dụng phần mềm này.

## Đặc trưng
* Mã nguồn mở `Bản này tuỳ Tâm trạng vui buồn`
* Hỗ trợ nhiều giao thức V2ray, Trojan, Shadowsocks.
* Hỗ trợ các tính năng mới như Vless và XTLS.
* Hỗ trợ kết nối đơn lẻ với nhiều bảng và nút mà không cần khởi động lại.
* Hỗ trợ IP trực tuyến bị hạn chế
* Hỗ trợ mức cổng nút, giới hạn tốc độ mức người dùng.
* Cấu hình đơn giản và rõ ràng.
* Sửa đổi cấu hình để tự động khởi động lại phiên bản.
* Dễ dàng biên dịch và nâng cấp, có thể nhanh chóng cập nhật phiên bản lõi, hỗ trợ các tính năng mới của Xray-core.
* Hỗ trợ UDP và nhiều chức năng khác

## Đặc trưng

| Đặc trưng                         | v2ray | trojan | shadowsocks |
| ---------------------------       | ----- | ------ | ----------- |
| Nhận thông tin về nút             | √     | √      | √           |
| Nhận thông tin người dùng         | √     | √      | √           |
| Thống kê lưu lượng người dùng     | √     | √      | √           |
| Báo cáo thông tin máy chủ         | √     | √      | √           |
| Tự động đăng ký chứng chỉ TLS     | √     | √      | √           |
| tự động gia hạn chứng chỉ tls     | √     | √      | √           |
| Số người trực tuyến               | √     | √      | √           |
| Hạn chế Người dùng Trực tuyến     | √     | √      | √           |
| quy tắc kiểm toán                 | √     | √      | √           |
| Giới hạn tốc độ cổng nút          | √     | √      | √           |
| Giới hạn tốc độ của người dùng    | √     | √      | √           |
| DNS tùy chỉnh                     | √     | √      | √           |
## Hỗ trợ giao diện người dùng

| đầu trước                                              | v2ray | trojan | shadowsocks                                 |
| ------------------------------------------------------ | ----- | ------ | ------------------------------------------- |
| [sspanel-uim](https://github.com/Anankke/SSPanel-Uim)  | √     | √      | √ (Đa người dùng một cổng và V2ray-Plugin)  |
| [v2board](https://github.com/v2board/v2board)          | √     | √      | √                                           |
| [PMPanel](https://github.com/ByteInternetHK/PMPanel)   | √     | √      | √                                           |
| [ProxyPanel](https://github.com/ProxyPanel/ProxyPanel) | √     | √      | √                                           |
| [V2board] New API ( UniProxy ) - PanelType: AikoPanel  | √     | √      | √                                           |

## Cài đặt phần mềm - release
```
bash <(curl -Ls https://raw.githubusercontent.com/AikoXrayR-Project/AikoXrayR-install/master/install.sh)
```
### Một cài đặt chính - docker
```
docker pull aikocute/xrayr:latest && docker run --restart=always --name xrayr -d -v ${PATH_TO_CONFIG}/aiko.yml:/etc/XrayR/aiko.yml --network=host aikocute/xrayr:latest
```
### Tệp cấu hình và hướng dẫn chi tiết
[Hướng dẫn chi tiết](https://xrayr.aikocute.com)
## Telgram

[AikoTelegram](https://t.me/AikoXrayR)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/AikoCute/XrayR.svg)](https://starchart.cc/AikoCute/XrayR)

