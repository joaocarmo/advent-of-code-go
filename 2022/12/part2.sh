#!/opt/homebrew/bin/bash
# Path: 2022/12/part2.sh
# This script is part of the Advent of Code 2020 challenge.
# It will be used to automate the process of solving the puzzle.

# Definitions
BIN_NAME="day12"
BIN_PATH="${PWD}/${BIN_NAME}"
INPUT_FILE="${PWD}/${1}"

execute_cmd() {
  if [[ ! -x "${BIN_PATH}" ]]; then
    echo "Binary is not executable!"
    exit 1
  fi

  "${BIN_PATH}" "${INPUT_FILE}" "${1}"
}

# Check if the input file exists.
if [[ ! -f "${INPUT_FILE}" ]]; then
  echo "Input file does not exist!"
  exit 1
fi

# Create a temporary file to store the input data.
TMP_FILE_POINTS=$(mktemp)
TMP_FILE_STEPS=$(mktemp)

# Print the current directory.
echo "Current directory: $PWD"

# Build the binary.
echo "Building the binary..."
go build -o "${BIN_PATH}" main.go || exit 1

# Execute the script to get the input data.
echo "Executing the script to get the input data..."
execute_cmd all | tee "${TMP_FILE_POINTS}"
echo ""

# For each line in the points input file, execute the script to get the steps.
echo "Executing the script to get the steps..."
while read -r LINE || [ -n "$LINE" ]; do
  # The reasoning for ignoring the stderr is that a starting point which causes
  # the script to fail is probably not an optimal starting point.
  execute_cmd "${LINE}" 2>/dev/null | tee -a "${TMP_FILE_STEPS}"
done < "${TMP_FILE_POINTS}"

# Sort the steps file and get the first line.
RESULT=$(sort -n "${TMP_FILE_STEPS}" | head -n 1)

# Print the result.
echo "The result is: ${RESULT}"

# Remove the temporary files.
rm -f "${TMP_FILE_POINTS}"
rm -f "${TMP_FILE_STEPS}"

# Remove the binary.
rm -f "${BIN_NAME}"
