# Go app that converts .md to .html

## How to use?

```
docker build -t md2html .
docker run --rm -v "$PWD":/work md2html -in /work/README.md -out /work/README.html
```