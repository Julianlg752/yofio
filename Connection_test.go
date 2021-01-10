// vim: sw=4 ts=4 expandtab
package main

import (
    "os"
    "testing"
)

func TestNoEnvConnection(t *testing.T) {
    _, db_err := connection()
    if db_err != nil {
        t.Errorf("Error connecting to the Database: %s", db_err.Error())
    }
}

func TestConnection(t *testing.T) {
    os.Setenv("DB_USER", "yofio_user")
    os.Setenv("DB_PASSWORD", "2021Yofio.5347")
    os.Setenv("DB_NAME", "yofio")
    _, db_err := connection()
    if db_err != nil {
        t.Errorf("Error connecting to the Database: %s", db_err.Error())
    }
}

