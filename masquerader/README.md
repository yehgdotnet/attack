# Objective
To simulate execution of non-signed executables under the name of windows binaries in C:\Windows directory

# Steps

1. Generate a list of binaries using catalog.cmd
```
 dir /b /s c:\windows\*.exe > winbins.txt
```

This will generate winbins.txt in the current directory.

2. go run main.go


Every 4 seconds, the program will copy source.exe to c:\windows\temp directory with each name listed in winbins.txt.  
The source.exe contains auto-exit and auto-delete functionality which cleans up itself. 
 
Refer to masqueradingexecution.log for detailed execution log. 	
