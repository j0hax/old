# old

A Golang reimplementation of openSUSE's [old](https://github.com/openSUSE/aaa_base/blob/master/files/usr/bin/old) script.

## Usage

```
old [FILES]
```

- The files will be renamed with a datestamp.
- If an old-ified file of the same name and timestamp already exists, an integer starting from 0 will be added.
