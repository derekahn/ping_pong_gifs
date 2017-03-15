# pingif-pongif ⚡

A simple program that can take a specified directory full of `.gifs` and encode it with a `pingpong style` concurrently; all the while saving to the output directory `gifs/` and uploading to [gifs.com](https://gifs.com/).

![Turns this:](https://gifs.com/gif/Wn01zn)
![Into this:](http://gifs.com/gif/alice-mathing-bridge-mwgm1n)

## Usage
With my provided `test_gifs` full of awesomeness 🐶 🐕 🐩:
```bash
# Linux peoples
$ 7z x -so test_gifs.tar.7z | tar xf - -C test_gifs/

# Mac peoples
$ brew update && brew doctor
$ brew install P7zip
$ 7za x -so test_gifs.tar.7z | tar xf -

$ go run *.go

> Enter ./path/to/gifs/directory: test_gifs

# Enter your own gifs API key
> Enter your gifs.com API key (ie. gifs58xxce10ad223): gifs56d63999f0f34
# > Key set to: gifs56d63999f0f34
# Else hit enter
> Enter your gifs.com API key (ie. gifs58xxce10ad223):
# > No key set! Won't be uploading to your dashboard.

# Check out localhost:8080 of the encoded pingpong gifs
# while you wait for the uploads! 🍻
```
### Resources
- [P7zip](http://brewformulas.org/P7zip)
- [gifs API](http://docs.gifs.com/v1.0/docs#)

### Terms
- `pingpong style` = It plays both forward and then backward in a loop.
