#!/bin/bash

git reset -q HEAD -- .

git add meta/*.xml
if [[ $(git diff --name-only --cached) != "" ]]; then
  git commit -m "api metadata update, $(date +'%Y-%m-%d')"
  git push
fi

git add docs/*.md
if [[ $(git diff --name-only --cached) != "" ]]; then
  git commit -m "comparison update, $(date +'%Y-%m-%d')"
  git push
fi
