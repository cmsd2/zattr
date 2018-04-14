# Zattr

Zattr is a program to modify the attributes of files in zip files.

## Usecases

### Setting executable bit on windows

As a developer on a windows system I want to be able to create cross-compiled binaries with the correct attributes e.g. for AWS Lambda.

These binaries need to have the executable bits set.

**Unfortunately this cross-platform usecase doesn't work because zipinfo detects the platform on which the zip was created and reinterprets the attribute bits.**

Example:

```
zattr chattr lambda-win.zip -m "0755" -o lambda-linux.zip
```