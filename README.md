# Easypdf
A simple CLI application to convert multiple markdown files to pdf using wkhtmltopdf.

# Prerequisites
- go 1.19 
- wkhtmltopdf version 0.12.6

# Installation
 Clone the repositorty and build a binary.
 ```
git clone https://github.com/blackfly19/Easypdf
cd Easypdf
make cmd-first-build
```
# Usage
You can either give specific files or mention the directory of files you want to convert.
```
easypdf convert -f testdoc1.md -f testdoc2.md -c cover.md -o test.pdf
```
```
easypdf convert -d basic_example -o test.pdf
```

Create table of contents with a single flag.
```
easypdf convert -d basic_example -o test.pdf --toc
```
Use watch mode to see changes in pdf without having to run the command again
```
easypdf convert -d basic_example -o test.pdf -w
```
Other options:
```
easypdf convert -h
Converts single or multiple markdown files to pdf

Usage:
  easypdf convert [flags]

Flags:
  -c, --cover string        Cover page of the document
      --css string          css file name
      --ddl                 Disable dotted lines in toc
  -d, --dir string          Directory name containing markdown files
      --dtl                 Disable toc links
  -f, --files stringArray   Single file or multiple files
  -h, --help                help for convert
  -o, --output string       Output file name (default "2c88f9cf-c24e-41b7-8e40-f7a5f725cb0b")
      --tht string          Toc header text (default "Table of Contents")
      --tli uint            Toc level indentation (default 1)
  -t, --toc                 Add table of contents
      --ttss float          Font scaling for each level of heading (default 0.8)
  -w, --watch               Enable watch mode to see changes in real time

```

# License
[MIT](https://github.com/blackfly19/Easypdf/blob/master/LICENSE)
