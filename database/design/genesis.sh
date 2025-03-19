#!/bin/bash

# This script appends the content of all files in specific directories to specific target files

# List of source directories and target files
declare -A arr
arr["schema/"]="00000000000001_genesis_structure.up.sql"
arr["seeds/"]="00000000000002_genesis_seeds.up.sql"

# Iterate over all directories and corresponding files in the array
for SOURCE_DIRECTORY in "${!arr[@]}"; do
    TARGET_FILE="${arr[$SOURCE_DIRECTORY]}"

    if [ ! -d "$SOURCE_DIRECTORY" ]; then
        echo "Source directory $SOURCE_DIRECTORY does not exist"
        continue
    fi

    # Make sure the target file is empty
    echo -n > "../migrations/$TARGET_FILE"

    # Iterate over all files in the design sources and append them to the migration file
    for file in "$SOURCE_DIRECTORY"/*
    do
        # if it's not an .sql file, skip it
        if [ "${file: -4}" != ".sql" ]; then continue; fi

        cat "$file" >> "../migrations/$TARGET_FILE"
    done

    echo "All files from $SOURCE_DIRECTORY have been appended to $TARGET_FILE"
done

