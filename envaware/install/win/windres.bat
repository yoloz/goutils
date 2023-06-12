@ECHO OFF

SET PATH=C:\TDM-GCC-64\bin;%PATH%
REM SET PATH=C:\mingw64\bin;%PATH%

windres -i C:\envaware\install\win\versioninfo.rc -O coff -o C:\envaware\install\win\versioninfo.syso

PAUSE