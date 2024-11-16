package version

import (
  "runtime"
)

var (
  version = "v0.1"
)

func Get() string {
  return version
}

func GetAll() string {
  v := Get() + "," + runtime.Version()
   return v
 }

