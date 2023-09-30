set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64

@echo off
if exist main (
  rm main
  echo "rm main"
)
if exist blog.tar.gz (
  rm blog.tar.gz
  echo "rm blog.tar.gz"
)

@echo off
for /f "tokens=2 delims==" %%i in ('wmic path win32_operatingsystem get LocalDateTime /value ^| findstr "="') do (
  set "strDate=%%i"
)
set "Now=%strDate:~0,4%-%strDate:~4,2%-%strDate:~6,2%,%strDate:~8,2%:%strDate:~10,2%:%strDate:~12,2%"

setlocal EnableDelayedExpansion
for /f "delims=" %%i in ('git rev-parse HEAD') do (
  set "commitId=%%i"
)

set ver=%1
if [%ver%] == [] (set ver=unkonw)
@echo on
go build -ldflags "-X main.buildTime=%Now% -X main.buildVersion=%ver% -X main.gitCommitID=%commitId%" -o main
chmod a+x main

@echo off
echo packing...
@echo on
tar -zcvf blog.tar.gz configs main