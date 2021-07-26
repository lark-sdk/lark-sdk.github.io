package main

import (
  "io/ioutil"
  "os"
  "strings"
)

func main() {
  err := build(os.Args[1])
  if err != nil {
    panic(err)
  }
}

func build(dir string) error {
  mdFile := dir + "/index.html.md"
  codeDir := dir + "/codes"

  return readRandReplaceMd(mdFile, codeDir)
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
      codeFilename = codeDir + "/" + codeFilename
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
