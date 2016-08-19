# Exercise 8.2
Implement a concurrent File Transfer Protocol(FTP) server. The server should interpret commands from each client such as `cd` to change directory, `ls` to list a directory, `get` to send the contents of a file, and `close` to close the connection. You can use the standard `ftp` command as the client, or write your own.

---
# 練習問題 8.2
並列に動作するFTP(File Transfer Protocol)サーバを実装しなさい。サーバは、ディレクトリを変更する`cd`、ディレクトリを列挙する`ls`、ファイルの内容を取り出す`get`、接続を閉じる`close`といったコマンドをクライアントごとに読み取るべきです。クライアントとして標準の`ftp`コマンドを使えますし、あるいは自分でクライアントを書いてみてください。


# Quote from RFC959

[https://www.ietf.org/rfc/rfc959.txt](https://www.ietf.org/rfc/rfc959.txt)

## 7.  TYPICAL FTP SCENARIO

   User at host U wanting to transfer files to/from host S:

   In general, the user will communicate to the server via a mediating
   user-FTP process.  The following may be a typical scenario.  The
   user-FTP prompts are shown in parentheses, '---->' represents
   commands from host U to host S, and '<----' represents replies from
   host S to host U.

| LOCAL COMMANDS BY USER | ACTION INVOLVED|
|---|---|
|ftp (host) multics`<CR>` | Connect to host S, port L, establishing control connections.|
||<---- 220 Service ready `<CRLF>`.|
|username Doe `<CR>` | USER Doe`<CRLF>`---->|
|| <---- 331 User name ok, need password`<CRLF>`.|
|password mumble `<CR>`| PASS mumble`<CRLF>`---->|
|| <---- 230 User logged in`<CRLF>`.|
|retrieve (local type) ASCII`<CR>`||
|(local pathname) test 1 `<CR>`   |User-FTP opens local file in ASCII.|
|(for. pathname) test.pl1`<CR>`  | RETR test.pl1`<CRLF>` ---->|
||<---- 150 File status okay;  about to open data connection`<CRLF>`.  Server makes data connection to port U.|
||<---- 226 Closing data connection, file transfer successful`<CRLF>`.|
|type Image`<CR>` |TYPE I`<CRLF>` ---->|
||<---- 200 Command OK`<CRLF>`|
|store (local type) image`<CR>`||
|(local pathname) file dump`<CR>` User-FTP opens local file in Image.||
|(for.pathname) >udd>cn>fd`<CR>` |STOR >udd>cn>fd`<CRLF>` ---->|
||<---- 550 Access denied`<CRLF>`|
|terminate|QUIT `<CRLF>` ---->|
||Server closes all connections.|


# Appendix

## Start standard FTP server on Mac OS X

````shell
$ sudo launchctl load -w /System/Library/LaunchDaemons/ftp.plist
````

## Stop standard FTP server on Mac OS X

````shell
$ sudo launchctl unload -w /System/Library/LaunchDaemons/ftp.plist
````
