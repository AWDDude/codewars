#!/usr/bin/env pwsh

param (
  [Parameter(Mandatory = $true, ValueFromPipeline = $true)][string]$Name
)

$Name = ($Name -replace ' ', '')

[hashtable]$DefaultFiles = @{
  'main.go' = "package main

import (
  `"fmt`"
)

func main() {
  fmt.Println($Name())
}

func $Name() {

}
"
}

cd $PSScriptRoot

if ((ls -Directory).Name -icontains $Name) {
  Write-Error "A directory by the name of '$Name' already exists"
  return
}

Write-Host "Creating '$Name'" -ForegroundColor Green

mkdir $Name | Out-Null
cd $Name
go mod init $Name

foreach($key in $DefaultFiles.Keys) {
  Out-File -FilePath $key -InputObject $DefaultFiles[$key]
}

cd ..
go work use "./$Name"
