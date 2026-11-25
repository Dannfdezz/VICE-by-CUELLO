# VICE by CUELLO

**Unofficial GTA VI PC Launcher — Manage, install and update experimental compatibility engines in one place. Developed by CUELLO.**

Independent experimental Windows x64 project. Not an official GTA VI PC release and not affiliated with Rockstar Games, Take-Two, Microsoft or third-party engine developers. No game files are included. Use only files you have the right to use. Compatibility and performance are not guaranteed.

## Features

- Manage and link locally installed compatibility engines.
- Check for updates on GitHub and confirm before installing new engine versions.
- Preserve existing installations, game paths, settings, logs and shader caches.
- Dark interface and a standalone installer.

## Downloads

The installer is intended to be the primary download; a standalone launcher is optional. See [Releases](https://github.com/Dannfdezz/VICE-by-CUELLO/releases) for any reviewed releases. Binary distribution remains pending security review.

## Source code

Version: v14 R2. The maintainer will upload the `launcher-source/` and `setup-source/` directories, including graphical assets and tests. Do not commit local game files, keys, private configurations or engine binaries.

## Building

With Go 1.23.2 on Windows x64, from the repository root:

```powershell
$env:GO111MODULE = 'off'
$env:GOOS = 'windows'
$env:GOARCH = 'amd64'
$env:CGO_ENABLED = '0'
New-Item -ItemType Directory -Force dist | Out-Null
Push-Location launcher-source
go test .
go build -ldflags='-H windowsgui -s -w' -o ../dist/VICE-by-CUELLO.exe .
Pop-Location
Copy-Item dist/VICE-by-CUELLO.exe setup-source/payload.exe
Push-Location setup-source
go test .
go build -ldflags='-H windowsgui -s -w' -o ../dist/VICE-by-CUELLO-Setup.exe .
Pop-Location
Get-FileHash dist/*.exe -Algorithm SHA256
```

The setup embeds the locally built launcher. Byte-for-byte reproducibility has not been independently verified.

## Security

Previously reported VirusTotal results: source archive 0/65, launcher 4/69, installer 6/70. These are snapshots, not security certifications. The detections are unresolved. Windows runtime testing, independent review and signing remain pending. Do not disable antivirus protections. See [SECURITY.md](SECURITY.md).

## Licensing

Source visibility does not grant redistribution rights. Project and third-party asset licensing remain pending review. All trademarks belong to their respective owners.
