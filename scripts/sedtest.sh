# !/usr/bin/bash

# This script is used to test the sed command
file="test.md"

# d  -- delete pattern space
sed -i '/<!-- Folder Structure:START -->/,/<!-- Folder Structure:END -->/d' $file
echo '<!-- Folder Structure:START -->' >> $file
echo '## Folder Structure' >> $file
echo '```' >> $file
(tree -I "documentation*") >> $file
echo '```' >> $file
echo '<!-- Folder Structure:END -->\n' >> $file