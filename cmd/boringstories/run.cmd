@echo off
setlocal
set "APP_NAME=boringstories.exe"

if not exist %APP_NAME% (
  echo %APP_NAME% not found
  call build.cmd
)

start "" /B %APP_NAME%
