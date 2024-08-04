import json
import re

# Step 1: Read the version from ../typescript/package.json
with open('../typescript/package.json', 'r') as file:
    package_json = json.load(file)
    version = package_json['version']

# Step 2: Update the version in setup.py
setup_file = './setup.py'
with open(setup_file, 'r') as file:
    setup_content = file.read()

# Using regular expression to replace the version string
setup_content = re.sub(
    r"version='[^']+'",
    f"version='{version}'",
    setup_content
)

with open(setup_file, 'w') as file:
    file.write(setup_content)