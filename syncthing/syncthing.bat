@rem %APPDATA%\Microsoft\Windows\Start Menu\Programs\Startup

@where syncthing.exe

if %errorlevel% NEQ 0 (msg %username% /time:10 "can't find syncthing.exe" && exit 1)

taskkill /im syncthing.exe /f

syncthing.exe --no-console --no-browser --no-default-folder