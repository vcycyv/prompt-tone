@echo off
echo Building Wooden Fish Prompt Tone Generator...
go build -o ding.exe main.go
if %errorlevel% neq 0 (
    echo Build failed!
    pause
    exit /b 1
)

echo.
echo Running Wooden Fish Prompt Tone Generator...
echo Usage examples:
echo   Default (90 minutes): ding.exe
echo   Custom duration: ding.exe -duration 60
echo   Custom output: ding.exe -output meditation.mp3
echo   Both: ding.exe -duration 120 -output session.mp3
echo.
echo Press any key to run with default settings (90 minutes)...
pause >nul

ding.exe
pause