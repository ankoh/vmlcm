package controller

import(
  "io/ioutil"
)

func generateIdentifier(address string) string {
  return ""
}

func listDirectory(path string) ([]string, error) {
  // First read the directory
  files, err := ioutil.ReadDir(path)
  if err != nil {
    return nil, err
  }

  // Then store the file names in a result array
  var filePaths = make([]string, len(files))
  for index, file := range files {
      filePaths[index] = file.Name()
  }
  return filePaths, nil
}
