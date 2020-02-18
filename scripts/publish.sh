#!/bin/bash

git reset -q HEAD -- .
git add edmx/*.xml
if [[ $(git diff --name-only --cached) != "" ]]; then
  git commit -m "api metadata update, $(date +'%Y-%m-%d')"
  put push
fi
