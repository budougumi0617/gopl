# Exercise 4.10
Modify `issues` to report the results in age categories, say less than a month old, less than a year old, and more than a year old.

---
# 練習問題 4.10
一ヶ月未満、一年未満、一年以上の期間で分類された結果を報告するように`issues`を修正しなさい。

# Result

````sh
$  go run issues.go github.go golang/go is:open json decoder

------------------------
Within a single month
------------------------
3 issues:
----------------------------------------
Number: 14640
User:   AntiPaste
Title:  encoding/json: decoding a null value does not replace previous v
Age:    3 days
----------------------------------------
Number: 49
User:   vetinari
Title:  add basic ldif reader
Age:    22 days
----------------------------------------
Number: 1838
User:   Spritekin
Title:  Crash using 1.1.1
Age:    18 days

------------------------
Within a year
------------------------
17 issues:
----------------------------------------
Number: 1666
User:   carlanton
Title:  Add decoder for Docker Fluentd [wip]
Age:    214 days
----------------------------------------
Number: 271
User:   fatih
Title:  Encoding to JSON produces null values
Age:    276 days
----------------------------------------
Number: 39
User:   GoogleCodeExporter
Title:  An array of floats represented as strings causes json decoding t
Age:    339 days
----------------------------------------
Number: 39
User:   GoogleCodeExporter
Title:  An array of floats represented as strings causes json decoding t
Age:    345 days
----------------------------------------
Number: 1
User:   gilliek
Title:  Add support for compressed JSON (gz or bz2) in input
Age:    344 days
----------------------------------------
Number: 16964
User:   liggitt
Title:  Preserve int data when unmarshaling
Age:    122 days
----------------------------------------
Number: 114
User:   ugorji
Title:  XML support
Age:    130 days
----------------------------------------
Number: 1488
User:   jaredgisin
Title:  [TCPInput] net = unix* parameters do not work
Age:    326 days
----------------------------------------
Number: 5004
User:   AxelVoitier
Title:  [0.9.5.1] SELECT returns both int and float for the same field
Age:    93 days
----------------------------------------
Number: 246
User:   Dieterbe
Title:  nilpointer when adding new collector
Age:    258 days
----------------------------------------
Number: 3046
User:   DriveU
Title:  Communicator:None
Age:    59 days
----------------------------------------
Number: 437
User:   workingenius
Title:  server blocks and can not get any http response
Age:    172 days
----------------------------------------
Number: 3384
User:   vrparmar
Title:  Host/Container management UI unable to fetch CPU/Memory/Network/
Age:    45 days
----------------------------------------
Number: 1
User:   sdboyer
Title:  uint64 use
Age:    83 days
----------------------------------------
Number: 37
User:   lcg635
Title:  请问可以和protobuf结合使用吗？
Age:    194 days
----------------------------------------
Number: 3441
User:   wint00
Title:  crash during google container
Age:    152 days
----------------------------------------
Number: 3
User:   theGeekPirate
Title:  Failed to load systray.dll (not a valid Win32 application)
Age:    304 days

------------------------
Over a year ago
------------------------
6 issues:
----------------------------------------
Number: 6716
User:   gopherbot
Title:  encoding/json: include field name in unmarshal error messages
Age:    854 days
----------------------------------------
Number: 7872
User:   extemporalgenome
Title:  encoding/json: Encoder internally buffers full output
Age:    681 days
----------------------------------------
Number: 6384
User:   joeshaw
Title:  encoding/json: encode precise floating point integers using %.0f
Age:    906 days
----------------------------------------
Number: 6
User:   cmars
Title:  Support some playground features by using play.golang.org
Age:    920 days
----------------------------------------
Number: 4237
User:   gjemiller
Title:  encoding/base64: URLEncoding padding is optional
Age:    1243 days
----------------------------------------
Number: 73
User:   tve
Title:  handling many request types
Age:    465 days
````
