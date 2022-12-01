@echo off
setlocal
if '%1' == '' goto :eof
mkdir %1
if %errorlevel% neq 0 (
	echo Catalog %1 exists
	goto :eof
)
copy templ_main.go %1\main.go
if %errorlevel% neq 0 (
	echo Unable to copy template main to %1
)
cd %1
rem go mod init example/davidn/2015day%1
touch README.md
%editor% main.go
