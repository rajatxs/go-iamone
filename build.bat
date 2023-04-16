@ECHO OFF
SET SOURCE_FILE=cmd\main.go
SET EXEC_FILE=bin\iamone-core.exe

IF NOT EXIST "bin" (
   MKDIR bin
)

IF EXIST "env.bat" (
   CALL "env.bat"
)

go build -o %EXEC_FILE% %SOURCE_FILE%

@ECHO ON
