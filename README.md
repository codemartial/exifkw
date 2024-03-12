# exifkw: Append keywords with exiftool
## Installation
`go get -u github.com/codemartial/exifkw`

👉Please ensure that you have `exiftool` binary installed and accessible on your command-line path. E.g. the following command should work for you:
```
$ exiftool -ver
11.59
```

## Usage
`exifkw "KEYWORD[,KEYWORD]..." EXIFTOOL_OPTS... FILE...`
### Example
```
$ exifkw "New Zealand,South Island,Queenstown,D750,70-200mm f/4G VR" *.jpg
   12 image files updated
```
`exifkw` preserves the file modification timestamps and it does not create the `*_original` files that `exiftool` creates by default. It's better if you have a backup or copy of the files in case you add the wrong keywords, make other mistakes that you want to undo or if `exifkw` doesn't work as intended.
