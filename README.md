# Test luồng PDU SESSION ESTABLISHMENT

Sử dụng Wireshark để kiểm tra luồng PDU SESSION ESTABLISHMENT với cú pháp lọc như sau:

```plaintext
(tcp.port == 8080 || tcp.port == 8081 || tcp.port == 8082 || udp.port == 8805) && (pfcp || http)
