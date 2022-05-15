@ECHO OFF

set argC=0
for %%x in (%*) do Set /A argC+=1

SET comm="%1%"
IF %comm%=="build" goto build
IF %comm%=="install" goto install
IF %comm%=="run" goto run
IF %comm%=="generate-ts" goto generate_ts
IF %comm%=="help" goto show_help

goto show_help

:install
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go install .\coco
goto Exit

:generate_ts
swag init --pd --parseDepth 1 -g .\coco\launch.go
npx swagger-typescript-api -p .\docs\swagger.json -o .\ui\src -n api.ts
goto Exit

:build
swag init --pd --parseDepth 1 -g .\coco\launch.go
goto Exit

:run
swag init --pd --parseDepth 1 -g .\coco\launch.go
go build -o coco-dev.exe .\coco
.\coco-dev.exe
goto Exit

:show_help
echo Usage: ccoco.bat [commands]
echo Commands: 
echo      build             -   To build the Cloud CoCo
echo      install           -   To install the Cloud CoCo dependencies
echo      run               -   To run develop with all generator
echo      generate-ts       -   To generate typescript interface and axios api
echo      help              -   To show all commands
goto Exit

:Exit