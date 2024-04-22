package main

import (
    "fmt"
    "os"
    "path/filepath"
    "sort"
)

type FileInfo struct {
  size int
  filename string
}

func main() {
  mb := int64(1024*1024)
  list := []FileInfo{}

  // Get current directory
  dir, err := os.Getwd()
  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  // Walk through all directories and files
  err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
    if err != nil {
      fmt.Println("Error:", err)
      return err
    }

    // Check if it's a regular file and its size is greater than 100MB
    if info.Mode().IsRegular() && info.Size() > 100*mb {
      file := FileInfo{int(info.Size() / mb), path}
      list = append(list, file)
    }
    return nil
  })

  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  // Sort the list by filesize, desc
  sort.Slice(list, func(i, j int) bool {
    return list[i].size > list[j].size
  })

  for _, file := range list {
    fmt.Printf("[%d MB] %s\n", file.size, file.filename)
  }
}
