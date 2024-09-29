# pwnd
パスワードが漏洩したかどうかを確認するサーバーデーモン\
**OpenBSDのみ**

## インストールする方法
```sh
make
doas make install
doas rcctl enable pwnd
doas rcctl start pwnd
```
