# clip

A simple and robust clipboard manager

# Usage

- Copying

```sh
echo "foo" | clip copy
cat ./foo.bar | clip copy
```

- Pasting

```sh
clip paste >&1
clip paste > ./test
```

- Posting (posts stdin to paste.rs)

```sh
echo "foo" | clip poste
cat ./foo.bar | clip post
```

# Installation

## Github

- Grab a release directly from [Github Releases](https://github.com/saucesteals/clip/releases)
- Move it to a proper location (ex. `/usr/local/bin`)
- If your shell cant find `clip`, add the location to $PATH

## Manual

---

_Prerequisites_

- https://go.dev/doc/install

---

- Pull clip from the repo

```sh
go get -d github.com/saucesteals/clip
```

- Build the clip binary

```sh
go install github.com/saucesteals/clip/cmd/clip
```