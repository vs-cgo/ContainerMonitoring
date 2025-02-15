package db

import (
  "log"
)

func Get(db *pg.DB, c Container) error {
  

} 

func Set(db *pg.DB, c Container) error {
  _, err := db.Model(c).Insert()
  if err != nil {
    log.Printf("failed insert container: %v\n", err)
    return err
  }
  return nil
}
