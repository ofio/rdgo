release guide:
  following instructions here: https://goreleaser.com/quick-start/
  export GITHUB_TOKEN=
  git tag -a v0.1.1 -m "structs update"
  git push origin v0.1.1
  goreleaser release --rm-dist

usage:
