# Exercise 8.2
Implement a concurrent File Transfer Protocol(FTP) server. The server should interpret commands from each client such as `cd` to change directory, `ls` to list a directory, `get` to send the contents of a file, and `close` to close the connection. You can use the standard `ftp` command as the client, or write your own.

---
# 練習問題 8.2
並列に動作するFTP(File Transfer Protocol)サーバを実装しなさい。サーバは、ディレクトリを変更する`cd`、ディレクトリを列挙する`ls`、ファイルの内容を取り出す`get`、接続を閉じる`close`といったコマンドをクライアントごとに読み取るべきです。クライアントとして標準の`ftp`コマンドを使えますし、あるいは自分でクライアントを書いてみてください。

# Result

## Client side Log
````shel
budougumi0617@~/tmp
$  ftp localhost 8000
Trying ::1...
ftp: Can't connect to `::1': Connection refused
Trying 127.0.0.1...
Connected to localhost.
220 Service ready for new user
Name (localhost:budougumi0617): yshimizu
230 User logged in, proceed.
ftp> ls
502 Command not implemented.
200 PORT command successful.
150 Here comes the directory listing.
azure.png
cmd.go
main.go
README.md
rfc.txt
subdir
226 Closing data connection. List successful.
ftp> get rfc.txt
local: rfc.txt remote: rfc.txt
200 PORT command successful.
150 File status okay; about to open data connection.
   136 KiB    7.54 MiB/s
250 Requested file action okay, completed.
WARNING! 3933 bare linefeeds received in ASCII mode.
File may not have transferred correctly.
139979 bytes received in 00:00 (5.13 MiB/s)
ftp> cd su
502 Command not implemented.
ftp> cd subdir
250 Requested file action okay, completed. Move to subdir
ftp> pwd
Remote directory: subdir
ftp> put azure.png
local: azure.png remote: azure.png
200 PORT command successful.
150 Ok to send data.
100% |*****************************************************************************************************************| 42222        5.68 MiB/s    --:-- ETA
226 Transfer complete. Created subdir/azure.png
42222 bytes sent in 00:00 (2.56 MiB/s)
ftp> ls
200 PORT command successful.
150 Here comes the directory listing.
azure.png
226 Closing data connection. List successful.
ftp> by
221 Service closing control connection. Logged out if appropriate.
````


## Server side log

```shel
budougumi0617@~/git/gotraining/ch08/ex02 (remainingwork@GoTraining)
$  go run main.go cmd.go &
[1] 7239
2016/08/20 21:57:34 main.go:32: Connect from 127.0.0.1:64539
2016/08/20 21:57:34 cmd.go:90: Response: 220 Service ready for new user
2016/08/20 21:57:36 cmd.go:48: Parse text : "USER yshimizu"
2016/08/20 21:57:36 cmd.go:65: Search command : USER
2016/08/20 21:57:36 cmd.go:90: Response: 230 User logged in, proceed.
2016/08/20 21:57:36 cmd.go:48: Parse text : "SYST"
2016/08/20 21:57:36 cmd.go:65: Search command : SYST
2016/08/20 21:57:36 cmd.go:90: Response: 502 Command not implemented.
2016/08/20 21:57:36 cmd.go:48: Parse text : "FEAT"
2016/08/20 21:57:36 cmd.go:65: Search command : FEAT
2016/08/20 21:57:36 cmd.go:90: Response: 502 Command not implemented.
2016/08/20 21:57:36 cmd.go:48: Parse text : "PWD"
2016/08/20 21:57:36 cmd.go:65: Search command : PWD
2016/08/20 21:57:36 cmd.go:136: Execute PWD "127.0.0.1:64539"
2016/08/20 21:57:36 cmd.go:90: Response: 257 "." is current directory.
2016/08/20 21:57:39 cmd.go:48: Parse text : "EPSV"
2016/08/20 21:57:39 cmd.go:65: Search command : EPSV
2016/08/20 21:57:39 cmd.go:90: Response: 502 Command not implemented.
2016/08/20 21:57:39 cmd.go:48: Parse text : "PASV"
2016/08/20 21:57:39 cmd.go:65: Search command : PASV
2016/08/20 21:57:39 cmd.go:90: Response: 502 Command not implemented.
2016/08/20 21:57:39 cmd.go:48: Parse text : "PORT 127,0,0,1,252,28"
2016/08/20 21:57:39 cmd.go:65: Search command : PORT
2016/08/20 21:57:39 cmd.go:90: Response: 200 PORT command successful.
2016/08/20 21:57:39 cmd.go:48: Parse text : "LIST"
2016/08/20 21:57:39 cmd.go:65: Search command : LIST
2016/08/20 21:57:39 cmd.go:90: Response: 150 Here comes the directory listing.
2016/08/20 21:57:39 cmd.go:90: Response: 226 Closing data connection. List successful.
2016/08/20 21:57:45 cmd.go:48: Parse text : "SIZE rfc.txt"
2016/08/20 21:57:45 cmd.go:65: Search command : SIZE
2016/08/20 21:57:45 cmd.go:90: Response: 502 Command not implemented.
2016/08/20 21:57:45 cmd.go:48: Parse text : "PORT 127,0,0,1,252,30"
2016/08/20 21:57:45 cmd.go:65: Search command : PORT
2016/08/20 21:57:45 cmd.go:90: Response: 200 PORT command successful.
2016/08/20 21:57:45 cmd.go:48: Parse text : "RETR rfc.txt"
2016/08/20 21:57:45 cmd.go:65: Search command : RETR
2016/08/20 21:57:45 cmd.go:90: Response: 150 File status okay; about to open data connection.
2016/08/20 21:57:45 cmd.go:90: Response: 250 Requested file action okay, completed.
2016/08/20 21:57:46 cmd.go:48: Parse text : "MDTM rfc.txt"
2016/08/20 21:57:46 cmd.go:65: Search command : MDTM
2016/08/20 21:57:46 cmd.go:90: Response: 502 Command not implemented.
2016/08/20 21:57:49 cmd.go:48: Parse text : "PORT 127,0,0,1,252,32"
2016/08/20 21:57:49 cmd.go:65: Search command : PORT
2016/08/20 21:57:49 cmd.go:90: Response: 200 PORT command successful.
2016/08/20 21:57:49 cmd.go:48: Parse text : "NLST"
2016/08/20 21:57:49 cmd.go:65: Search command : NLST
2016/08/20 21:57:49 cmd.go:90: Response: 502 Command not implemented.
2016/08/20 21:57:51 cmd.go:48: Parse text : "CWD subdir"
2016/08/20 21:57:51 cmd.go:65: Search command : CWD
2016/08/20 21:57:51 cmd.go:90: Response: 250 Requested file action okay, completed. Move to subdir
2016/08/20 21:57:51 cmd.go:48: Parse text : "PWD"
2016/08/20 21:57:51 cmd.go:65: Search command : PWD
2016/08/20 21:57:51 cmd.go:136: Execute PWD "127.0.0.1:64539"
2016/08/20 21:57:51 cmd.go:90: Response: 257 "subdir" is current directory.
2016/08/20 21:57:59 cmd.go:48: Parse text : "PORT 127,0,0,1,252,33"
2016/08/20 21:57:59 cmd.go:65: Search command : PORT
2016/08/20 21:57:59 cmd.go:90: Response: 200 PORT command successful.
2016/08/20 21:57:59 cmd.go:48: Parse text : "STOR azure.png"
2016/08/20 21:57:59 cmd.go:65: Search command : STOR
2016/08/20 21:57:59 cmd.go:90: Response: 150 Ok to send data.
2016/08/20 21:57:59 cmd.go:90: Response: 226 Transfer complete. Created subdir/azure.png
2016/08/20 21:58:05 cmd.go:48: Parse text : "PORT 127,0,0,1,252,35"
2016/08/20 21:58:05 cmd.go:65: Search command : PORT
2016/08/20 21:58:05 cmd.go:90: Response: 200 PORT command successful.
2016/08/20 21:58:05 cmd.go:48: Parse text : "LIST"
2016/08/20 21:58:05 cmd.go:65: Search command : LIST
2016/08/20 21:58:05 cmd.go:90: Response: 150 Here comes the directory listing.
2016/08/20 21:58:05 cmd.go:90: Response: 226 Closing data connection. List successful.
2016/08/20 21:58:12 cmd.go:48: Parse text : "QUIT"
2016/08/20 21:58:12 cmd.go:65: Search command : QUIT
2016/08/20 21:58:12 cmd.go:90: Response: 221 Service closing control connection. Logged out if appropriate.
budougumi0617@~/git/gotraining/ch08/ex02 (remainingwork@GoTraining)
$  killall go run
[1]  + 7239 terminated  go run main.go cmd.go
budougumi0617@~/git/gotraining/ch08/ex02 (remainingwork@GoTraining)
$  tree ./
./
├── README.md
├── azure.png
├── cmd.go
├── main.go
├── rfc.txt
└── subdir
    └── azure.png

1 directory, 6 files
```


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
