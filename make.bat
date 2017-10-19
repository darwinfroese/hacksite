@ECHO OFF
PowerShell.exe -Command %rootdir%scripts/build.ps1 %1
EXIT /b %errorlevel%
