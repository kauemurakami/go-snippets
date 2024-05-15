# Define path to executável gof.exe
$executablePath = "C:\caminho\para\o\seu\executavel\gof.exe"

# Add the directory of the exe to PATH system
$env:Path += ";$executablePath"

# Persist changes in PATH
# (Notes: Required ADMIN mode)
[Environment]::SetEnvironmentVariable("Path", $env:Path, [EnvironmentVariableTarget]::Machine)