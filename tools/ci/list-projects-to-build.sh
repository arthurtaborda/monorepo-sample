#!/bin/bash

# Documentation
read -r -d '' USAGE_TEXT <<EOM
Usage:
    list-projects-to-builds.sh <revision range>

    List all projects which had some changes in given commit range.
    Project is identified with relative path to project's root directory from repository root.
    Output list is ordered respecting dependencies between projects (lower projects depends on upper).
    There can be multiple projects (separated by space) on single line which means they can be build on parallel.
   
    If one of commit messages in given commit range contains [rebuild-all] flag then all projects will be listed.

    <revision range>        range of revision hashes where changes will be looked for
                            format is HASH1..HASH2
EOM

set -e

# Capture input parameter and validate it
COMMIT_RANGE=$1

if [[ -z $COMMIT_RANGE ]]; then
  CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
  if [ -z "$GITHUB_SHA" ]; then
    HEAD_HASH=$(git rev-parse HEAD)
  else
    HEAD_HASH=$GITHUB_SHA
  fi

  if [[ $CURRENT_BRANCH == "master" ]]; then
    HASH_TO_COMPARE=$(git rev-parse "$HEAD_HASH"~1) # get previous commit
  else
    HASH_TO_COMPARE=$(git rev-parse origin/master)
  fi

  COMMIT_RANGE=$HASH_TO_COMPARE..$HEAD_HASH
fi

COMMIT_RANGE_FOR_LOG="$(echo $COMMIT_RANGE | sed -e 's/\.\./.../g')"

# Find script directory (no support for symlinks)
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

# Look for changes in given revision range
CHANGED_PATHS=$(git diff $COMMIT_RANGE --name-status)
#echo -e "Changed paths:\n$CHANGED_PATHS"

# Setup variables for output collecting
CHANGED_PROJECTS=""
CHANGED_DEPENDENCIES=""

# If [rebuild-all] command passed it's enough to take all projects and all dependencies as changed
if [[ $(git log "$COMMIT_RANGE_FOR_LOG" | grep "\[rebuild-all\]") ]]; then
  CHANGED_PROJECTS="$(${DIR}/list-projects.sh)"
else
  # For all known projects check if there was a change and look for all dependant projects
  for PROJECT in $(${DIR}/list-projects.sh); do
    PROJECT_NAME=$(basename $PROJECT)
    if [[ $(echo -e "$CHANGED_PATHS" | grep "$PROJECT") ]]; then
      CHANGED_PROJECTS="$CHANGED_PROJECTS\n$PROJECT"
    fi
  done
fi

# Build output
PROJECTS_TO_BUILD=$(echo -e "$CHANGED_DEPENDENCIES" | tsort | tac)
for PROJECT in $(echo -e "$CHANGED_PROJECTS"); do
  if [[ ! $(echo -e "$PROJECTS_TO_BUILD" | grep "$PROJECT") ]]; then
    PROJECTS_TO_BUILD="$PROJECT $PROJECTS_TO_BUILD"
  fi
done

# Print output
echo -e "$PROJECTS_TO_BUILD"
