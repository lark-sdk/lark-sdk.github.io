package main

import (
  "io/ioutil"
  "os"
  "path"
  "strings"
)

func main() {
  err := build(os.Args[1])
  if err != nil {
    panic(err)
  }
}

func build(dir string) error {
  mdFile := path.Join(dir, "/index.html.md")

  return readRandReplaceMd(mdFile, dir)
}

func readRandReplaceMd(mdFile, codeDir string) error {
  bs, err := ioutil.ReadFile(mdFile)
  if err != nil {
    return err
  }
  res := []string{}
  for _, v := range strings.Split(string(bs), "\n") {
    if strings.HasPrefix(v, "file-embed:") {
      codeFilename := strings.TrimSpace(v[len("file-embed:"):])
      codeFilename = path.Join(codeDir, codeFilename)
      codeBS, err := ioutil.ReadFile(codeFilename)
      if err != nil {
        return err
      }
      for _, vv := range strings.Split(string(codeBS), "\n") {
        res = append(res, vv)
      }
      continue
    }

    res = append(res, v)
  }

  return ioutil.WriteFile(mdFile, []byte(strings.Join(res, "\n")), 0666)
}
