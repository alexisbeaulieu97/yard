#!/bin/bash

# canopy shell wrapper
# Source this script in your shell profile (e.g. .zshrc, .bashrc)
# Usage: source /path/to/canopy.sh

canopy() {
    # Resolve canopy binary
    local CANOPY_BIN="canopy"
    if [[ -x "./canopy" ]]; then
        CANOPY_BIN="./canopy"
    fi

    # Check for "workspace new" or "tui" commands which might need cd
    if [[ "$1" == "workspace" && "$2" == "new" ]]; then
        # Capture output to see if we need to cd
        # We use a temporary file to capture stdout while letting stderr pass through
        tmp_out=$(mktemp)
        
        if [ -f "./canopy" ]; then
            ./canopy "$@" --print-path > "$tmp_out"
        else
            command canopy "$@" --print-path > "$tmp_out"
        fi
        
        exit_code=$?
        output=$(cat "$tmp_out")
        rm "$tmp_out"
        
        if [ $exit_code -eq 0 ] && [ -n "$output" ] && [ -d "$output" ]; then
            echo "Changing directory to: $output"
            cd "$output" || return
        else
            # If not a path or failed, just print output
            echo "$output"
        fi
    elif [[ "$1" == "workspace" && "$2" == "switch" ]]; then
        # Capture output to see if we need to cd
        tmp_out=$(mktemp)
        
        if [ -f "./canopy" ]; then
            ./canopy "$@" > "$tmp_out"
        else
            command canopy "$@" > "$tmp_out"
        fi
        
        exit_code=$?
        output=$(cat "$tmp_out")
        rm "$tmp_out"
        
        if [ $exit_code -eq 0 ] && [ -n "$output" ] && [ -d "$output" ]; then
            echo "Changing directory to: $output"
            cd "$output" || return
        else
            echo "$output"
        fi
    elif [[ "$1" == "tui" ]]; then
        # Similar logic for TUI if it returns a path
        tmp_out=$(mktemp)
        
        if [ -f "./canopy" ]; then
            ./canopy "$@" --print-path > "$tmp_out"
        else
            command canopy "$@" --print-path > "$tmp_out"
        fi
        
        exit_code=$?
        output=$(cat "$tmp_out")
        rm "$tmp_out"
        
        if [ $exit_code -eq 0 ] && [ -n "$output" ] && [ -d "$output" ]; then
             echo "Changing directory to: $output"
             cd "$output" || return
        fi
    else
        # Normal execution
        if [ -f "./canopy" ]; then
            ./canopy "$@"
        else
            command canopy "$@"
        fi
    fi
}
