set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0

function Build-Api {
    $foldersApi = Get-ChildItem -Path "lambdas/api/*/" -Directory

    foreach ($folder in $foldersApi) {
        $folderName = $folder.Name
        Push-Location "lambdas/streams/$folderName"
        try {
            go build -o bootstrap -tags lambda.norpc
            Compress-Archive -Path bootstrap -DestinationPath "../../../bin/${folderName}.zip"
            Remove-Item -Path bootstrap
        } catch {
            Write-Error "Failed to build or zip in lambdas/$folderName"
            Pop-Location
            continue
        }
        Pop-Location
    }
}

Build-Api

# $foldersCognito = Get-ChildItem -Path "lambdas/cognito/*/" -Directory

# function Build-Cognito {
#     foreach ($folder in $foldersCognito) {
#         $folderName = $folder.Name
#         Push-Location "lambdas/streams/$folderName"
#         try {
#             go build -o bootstrap -tags lambda.norpc
#             Compress-Archive -Path bootstrap -DestinationPath "../../../bin/${folderName}.zip"
#             Remove-Item -Path bootstrap
#         } catch {
#             Write-Error "Failed to build or zip in lambdas/$folderName"
#             Pop-Location
#             continue
#         }
#         Pop-Location
#     }
# }

# Build-Cognito
