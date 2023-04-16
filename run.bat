@ECHO OFF
SET SOURCE_FILE=cmd\main.go

IF EXIST "env.bat" (
   CALL "env.bat"
)

go run %SOURCE_FILE%

@ECHO ON
