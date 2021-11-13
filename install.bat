@ECHO OFF

ECHO ---------- information ---------
systeminfo | findstr /c:"OS Name"
systeminfo | findstr /c:"OS Version"
systeminfo | findstr /c:"System Type"
ECHO --------------------------------
ECHO Installing......
go get github.com/bndr/gotabulate
ECHO done 
