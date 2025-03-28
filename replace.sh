#!/bin/bash

# Get the replacement string from command line argument
replace_str="$1"

# Function to replace "gin-template" with the specified string in file contents
replace_in_file_contents() {
  local files=( $(grep -rl 'gin-template' "$1") )

  for file in "${files[@]}"; do
    # Exclude the replace.sh file from replacement
    if [ "$file" != "./replace.sh" ]; then
      sed -i '' "s/gin-template/$replace_str/g" "$file"
    fi
  done
}

# Function to replace "gin-template" with the specified string in directory and file names
replace_in_dir_and_file_names() {
  local dirs=( $(find . -depth -type d -name "*gin-template*") )

  for dir in "${dirs[@]}"; do
    # Exclude the current directory from replacement
    if [ "$dir" != "." ]; then
      newdir="${dir//gin-template/$replace_str}"
      mv "$dir" "$newdir"
    fi
  done

  local files=( $(find . -type f -name "*gin-template*") )

  for file in "${files[@]}"; do
    # Exclude the replace.sh file from replacement
    if [ "$file" != "./replace.sh" ]; then
      newfile="${file//gin-template/$replace_str}"
      mv "$file" "$newfile"
    fi
  done
}

# Function to replace the project root directory name
replace_project_root() {
  local current_dir="$(pwd)"
  local parent_dir="$(dirname "$current_dir")"
  local project_dir="$(basename "$current_dir")"

  # Check if the current directory's name is "gin-template"
  if [ "$project_dir" = "gin-template" ]; then
    local new_project_dir="${parent_dir}/${replace_str}"
    mv "$current_dir" "$new_project_dir"
  fi
}

# Function to remove the first 20 lines of Makefile
remove_first_20_lines_of_makefile() {
  if [ -f "Makefile" ]; then
    sed -i '' '1,20d' "Makefile"
  fi
}

# Remove the first 20 lines of Makefile
remove_first_20_lines_of_makefile

# Replace directory and file names
replace_in_dir_and_file_names

# Replace the project root directory name
replace_project_root

# Replace in file contents
replace_in_file_contents .
