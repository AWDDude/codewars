#!/usr/bin/env pwsh

param (
  [Parameter(Mandatory = $true, ValueFromPipeline = $true)][string]$Name
)

$Name = ($Name -replace ' ', '')

[hashtable]$DefaultFiles = @{
	'main.go' = "package main

import (
	`"`"
)

func $Name() {

}
"
	'main_test.go' = "package main

import (
	`"testing`"

	. `"github.com/onsi/gomega`"
)

func Test_$Name(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect($Name()).To(Equal())
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
