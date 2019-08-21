package comm

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/sebastianwebber/cmdr"
)

func ExecOSCmdNew(command string, killInMilliSeconds time.Duration) (okResult bool, stdout, stderr string) {
	str := strings.Split(command, " ")
	cmd := exec.Command(str[0], str[1:]...)
	//cmd := exec.Command("/usr/sbin/nginx", "-t")
	var outbuf, errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	okResult = true

	err := cmd.Start()
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	isKilled := false
	select {
	case <-time.After(killInMilliSeconds * time.Millisecond):
		if err := cmd.Process.Kill(); err != nil {
			okResult = false
		} else {
			isKilled = true
		}
	case err := <-done:

		if err != nil {
			//	log.Printf("process done with error = %v", err)
			okResult = false
		}
	}
	if err != nil {
		log.Fatal(err)
		okResult = false
	}

	if isKilled && errbuf.String() == "" {
		return false, outbuf.String(), "执行超时"
	}
	return okResult, outbuf.String(), errbuf.String()
}

//执行系统命令，errstr为空时则返回正常结果，当timeout =true时，timeoutNum值生效,
func ExecOSCmd(cmdstr string, timeout bool, timeoutNum int) (res string, errstr string) {
	str := strings.Split(cmdstr, " ")
	cmd := cmdr.New(true, str[0], str[1:]...)
	if timeout {
		cmd.Options.Timeout = timeoutNum
	}

	if cmd.IsValid() {
		out, err := cmd.Run()
		if err != nil {
			return "", string(out)
			//fmt.Println("eeee:", string(comm.CustomeError(err)))
		}

		//fmt.Println(string(out))
		return string(out), ""
	}

	return "", "Command Invalid"
}

//执行系统命令，errstr为空时则返回正常结果，当timeoutNum =0时，默认12000秒，即不超时
func ExecOSCmd2(cmdstr string, timeoutNum int) (res string, errstr string) {
	str := strings.Split(cmdstr, " ")
	cmd := cmdr.New(true, str[0], str[1:]...)
	if timeoutNum == 0 {
		cmd.Options.Timeout = 12000
	} else {
		cmd.Options.Timeout = timeoutNum
	}

	if cmd.IsValid() {
		out, err := cmd.Run()
		if err != nil {
			return "", string(out)
			//fmt.Println("eeee:", string(comm.CustomeError(err)))
		}

		//fmt.Println(string(out))
		return string(out), ""
	}

	return "", "Command Invalid"
}

func ExecOSCmdForBash(cmdName string) (okstr, errstr string) {
	var stdoutBuf, stderrBuf bytes.Buffer
	//cmdArgs := strings.Fields(cmdName)
	cmd := exec.Command("/usr/bin/sh", "-c", cmdName)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		//log.Fatalf("cmd.Start() failed with '%s'\n", err)
		return "", CustomeError(err)
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		//log.Fatalf("cmd.Run() failed with %s\n", err)
		return "", CustomeError(err)
	}
	if errStdout != nil || errStderr != nil {
		//log.Fatal("failed to capture stdout or stderr\n")
		return "", "failed to capture stdout or stderr"
	}
	outStr, errStr := stdoutBuf.Bytes(), stderrBuf.Bytes()
	//	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
	if len(errStr) > 0 {
		errstr = string(errStr)
	}
	return string(outStr), errstr
}
