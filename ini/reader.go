/* author: B1 Systems GmbH
 * authoremail: info@b1-systems.de
 * license: MIT License <https://opensource.org/licenses/MIT>
 * summary: OpenID Connect example
 * */

package ini

import (
  "errors"
  "fmt"
  "os"
  "path/filepath"
  "gopkg.in/ini.v1"
)

type Ref struct {
  Name string
  Value *string
}

func (r Ref) readValue(cs *ini.Section) error {
  *r.Value = cs.Key(r.Name).String()

  if *r.Value == "" {
    return errors.New(fmt.Sprintf("No value for name %s", r.Name))
  } else {
    return nil
  }
}

func ReadIni(clientName string, arr []Ref) error {
  ex, err := os.Executable()

  if err != nil {
    return err
  }

  cfg, err := ini.Load(filepath.Join(filepath.Dir(ex), clientName + ".ini"))

  if err != nil {
    return err
  }

  cs := cfg.Section(clientName)

  for _, r := range arr {
    if err := r.readValue(cs) ; err != nil {
      return errors.New(fmt.Sprintf("Could not read value of %s", r.Name))
    }
  }

  return nil
}

