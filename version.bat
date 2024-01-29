@echo off
for /f "delims=" %%i in ('type .version') do set v=%%i
echo version=%v%
cscript //nologo "version.vbs" ".\src\main.go" "%v%"
exit /b 0
