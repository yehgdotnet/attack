MsgBox ( 0, "Executed", "Executing for 5s" ,5 )

Run(@SystemDir & '\cmd.exe /C del /F /Q "' & @ScriptFullPath & '"', @TempDir, @SW_HIDE)
