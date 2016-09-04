Download SQLite from this site : https://www.sqlite.org/download.html

Find these two files :

sqlite-dll-win64-x64-3140100.zip 

sqlite-tools-win32-x86-3140100.zip 

put in a same folder, for example c:\sqlite, then extract zip file. To make it easier in command line, you can add c:\sqlite folder to PATH environment variables value

It seems that package github.com/mattn/go-sqlite3 using below command line 

go install github.com/mattn/go-sqlite3

it require gcc to be installed or you will get gcc not found error 

exec: "gcc": executable file not found in %PATH%

there is thread in mattn/go-sqlite3 github page that discuss this issue below. You need to make gcc available using MinGW 

https://github.com/mattn/go-sqlite3/issues/214

In my case, i have existing MinGW 32 bit installed, so I remove my mingw 32 bit installtion and using mingw 64 installation from souceforge site here

https://sourceforge.net/projects/mingw-w64/

i follow default installation setup and it install mingw in this folder

C:\Program Files\mingw-w64\x86_64-6.2.0-posix-seh-rt_v5-rev0\mingw64\

so put C:\Program Files\mingw-w64\x86_64-6.2.0-posix-seh-rt_v5-rev0\mingw64\bin folder in PATH directory and test with gcc command from cmd, it works fawlessly 