package vcs

import (
        "log"
        "os/exec"
        "path/filepath"
        "strings"
        "time"
       )

func init() {
    RegisterVCS("cvs", &CvsDriver{})
}

type CvsDriver struct{}

func (g *CvsDriver) HeadHash(dir string) (string, error) {
     t := time.Now()
     return strings.TrimSpace(t.Format("20060102150405")), nil
}

func (g *CvsDriver) Pull(dir string) (string, error) {
cmd := exec.Command("cvs", "up")
         cmd.Dir = dir
         out, err := cmd.CombinedOutput()
         if err != nil {
             log.Printf("Failed to cvs up %s, see output below\n%sContinuing...", dir, out)
                 return "", err
         }

     return g.HeadHash(dir)
}

func (g *CvsDriver) Clone(dir, url string) (string, error) {
    par, rep := filepath.Split(dir)
        cmd := exec.Command(
                "cvs",
                "-f",
                "-d",
                url,
                "checkout",
                "-r",
                "HEAD",
                "-P",
                rep)
        cmd.Dir = par
        out, err := cmd.CombinedOutput()
        if err != nil {
            log.Printf("Failed to checkout %s, see output below\n%sContinuing...", url, out)
                return "", err
        }

    return g.HeadHash(dir)
}
