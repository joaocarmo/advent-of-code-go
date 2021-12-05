#!/usr/bin/env bash

set -e

NEW_PROJECT_NAME="$1"
WORKDIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
TEMPLATEDIR="$WORKDIR/template"
NEW_PROJECT_DIR="$WORKDIR/$NEW_PROJECT_NAME"

# stop if no project name is given
if [ -z "$NEW_PROJECT_NAME" ]; then
  >&2 echo "ERROR: No project name given"
  exit 1
fi

# stop if project already exists
if [ -d "$NEW_PROJECT_DIR" ]; then
  >&2 echo "ERROR: Project already exists"
  exit 1
fi

echo "Creating new project '$NEW_PROJECT_NAME' from template..."

# copy the template to the new project
cp -r "$TEMPLATEDIR" "$NEW_PROJECT_DIR"

# replace the project name in the template
sed -i "" "s/template/$NEW_PROJECT_NAME/g" "$NEW_PROJECT_DIR/go.mod"

echo "Done"
