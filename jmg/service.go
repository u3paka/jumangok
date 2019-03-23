package jmg

import (
	"bufio"
	"bytes"
	"io"
	"os/exec"
)

//Service execute the given command
type Service struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
	ch     chan string
}

//NewService creates a new Service
func NewService(command string, options ...string) (*Service, error) {
	cmd := exec.Command(command, options...)
	cl := Service{cmd: cmd}
	var err error
	if cl.stdin, err = cl.cmd.StdinPipe(); err != nil {
		return nil, err
	}
	if cl.stdout, err = cl.cmd.StdoutPipe(); err != nil {
		return nil, err
	}
	err = cmd.Start()
	cl.startParser()
	return &cl, err
}

func (s *Service) RawParse(t string) string {
	s.ch <- t
	return <-s.ch
}

func (s *Service) startParser() {
	s.ch = make(chan string)
	go func() {
		for {
			io.WriteString(s.stdin, <-s.ch)
			io.WriteString(s.stdin, "\n")

			var buf bytes.Buffer
			sc := bufio.NewScanner(s.stdout)
			for sc.Scan() {
				l := sc.Text()
				if l == "EOS" {
					buf.WriteString("EOS")
					break
				}
				buf.WriteString(l)
				buf.WriteRune('\n')
			}
			s.ch <- buf.String()
		}
	}()
	return
}
