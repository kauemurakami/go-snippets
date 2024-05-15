#!/bin/bash

# Define the target to exe gof
executablePath="/caminho/para/o/seu/executavel/gof"

# Add directory of the exe in PATH system
echo "export PATH=\$PATH:$executablePath" >> ~/.bashrc