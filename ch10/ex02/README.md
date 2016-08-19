# Exercise 10.2
Define a genetic archive file-reading function capable of reading ZIP files (`archive/zip`) and POSIX tar files(`archive/tar`). Use a registration mechanism similar to the one described above so that support for each file format can be plugged using blank imports.

---
# 練習問題 10.2
ZIPファイル(`archive/zip`)とPOSIX tarファイル(`archive/tar`)を読み込むことができる汎用のアーカイブ読み込み関数を定義しなさい。個々のファイル形式のサポートをブランクインポートを使って組み込むために、前述の方法と同様な登録の仕組みを使いなさい。

# Result

````shell
budougumi0617@~/git/gotraining/ch10/ex02 (remainingwork@GoTraining)
$  tree .
.
├── README.md
├── myunarchive
├── myunarchive.go
├── test.tar
├── test.zip
└── unarchive
    ├── format.go
    ├── tar
    │   └── reader.go
    └── zip
        └── reader.go

3 directories, 8 files
budougumi0617@~/git/gotraining/ch10/ex02 (remainingwork@GoTraining)
$  ./myunarchive test.tar test.zip
Support format is below:
tar
zip
Use tar
start !!!!
2016/08/18 21:54:43 tarfile/ is directory
2016/08/18 21:54:43 tarfile/bar/ is directory
2016/08/18 21:54:43 Write file ./tarfile/foo.txt
2016/08/18 21:54:43 Write file ./tarfile/bar/bar.txt
Use zip
Cannot unzip subfiles in zipfile/ yet
2016/08/18 21:54:43 Try name zipfile/
2016/08/18 21:54:43 open ./zipfile/: is a directory
budougumi0617@~/git/gotraining/ch10/ex02 (remainingwork@GoTraining)
$  tree .
.
├── README.md
├── myunarchive
├── myunarchive.go
├── tarfile
│   ├── bar
│   │   └── bar.txt
│   └── foo.txt
├── test.tar
├── test.zip
├── unarchive
│   ├── format.go
│   ├── tar
│   │   └── reader.go
│   └── zip
│       └── reader.go
└── zipfile

6 directories, 10 files
````
