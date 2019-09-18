package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type selpg_args struct {
	start_page  int
	end_page    int
	in_filename string
	dest        string
	page_len    int
	page_type   int
}

var sa selpg_args
var argc int

func usage() {
	fmt.Println("\nHow to use selpg:")
	fmt.Println("\tselpg -sNumber -eNumber [options] [filename]")
	fmt.Println("\noptions:")
	fmt.Println("\t-lNumber: 每页包含Number行，如果既没有给出“-lNumber”也没有给出“-f”选项，则 selpg 会理解为页有固定的长度（每页 72 行）。")
	fmt.Println("\t-f: selpg在输入中寻找换页符，并将其作为页定界符处理。“-lNumber”和“-f”选项是互斥的。")
	fmt.Println("\t-dDestination: 将选定的页直接发送至打印机。")
	fmt.Println("\t[filename]: 输入文件名。")
}

func process_args(argv []string) {
	if len(argv) < 3 {
		fmt.Fprintf(os.Stderr, "error: not enough arguments\n")
		usage()
		os.Exit(1)
	}

	if argv[1][0] != '-' || argv[1][1] != 's' {
		fmt.Fprintf(os.Stderr, "error: the 1st arg should be -sNumber\n")
		usage()
		os.Exit(1)
	}

	sp, _ := strconv.Atoi(argv[1][2:])
	if sp < 1 {
		fmt.Fprintf(os.Stderr, "error: invalid start page number %d\n", sp)
		usage()
		os.Exit(1)
	}
	sa.start_page = sp

	if argv[2][0] != '-' || argv[2][1] != 'e' {
		fmt.Fprintf(os.Stderr, "error: the 2nd arg should be -eNumebr\n")
		usage()
		os.Exit(1)
	}

	ep, _ := strconv.Atoi(argv[2][2:])
	if ep < 1 || ep < sp {
		fmt.Fprintf(os.Stderr, "error: invalid end page %d\n", ep)
		usage()
		os.Exit(1)
	}
	sa.end_page = ep

	argindex := 3
	for {
		if argindex > argc-1 || argv[argindex][0] != '-' {
			break
		}
		switch argv[argindex][1] {
		case 'l':

			pl, _ := strconv.Atoi(argv[argindex][2:])
			if pl < 1 {
				fmt.Fprintf(os.Stderr, "error: invalid page length %d\n", pl)
				usage()
				os.Exit(1)
			}
			sa.page_len = pl
			argindex++
		case 'f':
			if len(argv[argindex]) > 2 {
				fmt.Fprintf(os.Stderr, "error: option should be \"-f\"\n")
				usage()
				os.Exit(1)
			}
			sa.page_type = 'f'
			argindex++
		case 'd':
			if len(argv[argindex]) <= 2 {
				fmt.Fprintf(os.Stderr, "error: -d option requires a printer destination\n")
				usage()
				os.Exit(1)
			}
			sa.dest = argv[argindex][2:]
			argindex++
		default:
			fmt.Fprintf(os.Stderr, "error: unknown option")
			usage()
			os.Exit(1)
		}
	}

	if argindex <= argc-1 {
		sa.in_filename = argv[argindex]
	}
}

func process_input() {
	var cmd *exec.Cmd
	var cmd_in io.WriteCloser
	var cmd_out io.ReadCloser
	if sa.dest != "" {
		cmd = exec.Command("bash", "-c", sa.dest)
		cmd_in, _ = cmd.StdinPipe()
		cmd_out, _ = cmd.StdoutPipe()
		//执行设定的命令
		cmd.Start()
	}
	if sa.in_filename != "" {
		inf, err := os.Open(sa.in_filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		line_count := 1
		page_count := 1
		fin := bufio.NewReader(inf)
		for {
			line, _, err := fin.ReadLine()
			if err != io.EOF && err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if err == io.EOF {
				break
			}
			if page_count >= sa.start_page && page_count <= sa.end_page {
				if sa.dest == "" {
					fmt.Println(string(line))
				} else {
					fmt.Fprintln(cmd_in, string(line))
				}
			}
			line_count++
			if sa.page_type == 'l' {
				if line_count > sa.page_len {
					line_count = 1
					page_count++
				}
			} else {
				if string(line) == "\f" {
					page_count++
				}
			}
		}
		if sa.dest != "" {
			cmd_in.Close()
			cmdBytes, err := ioutil.ReadAll(cmd_out)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print(string(cmdBytes))
			cmd.Wait()
		}
	} else {
		sc := bufio.NewScanner(os.Stdin)
		line_count := 1
		page_count := 1
		out := ""
		for sc.Scan() {
			line := sc.Text()
			line += "\n"
			if page_count >= sa.start_page && page_count <= sa.end_page {
				out += line
			}
			line_count++
			if sa.page_type == 'l' {
				if line_count > sa.page_len {
					line_count = 1
					page_count++
				}
			} else {
				if string(line) == "\f" {
					page_count++
				}
			}
		}
		if sa.dest == "" {
			fmt.Print(out)
		} else {
			fmt.Fprint(cmd_in, out)
			cmd_in.Close()
			cmdBytes, err := ioutil.ReadAll(cmd_out)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print(string(cmdBytes))
			cmd.Wait()
		}
	}
}

func main() {
	argv := os.Args
	sa.start_page = 1
	sa.end_page = 1
	sa.in_filename = ""
	sa.dest = ""
	sa.page_len = 10
	sa.page_type = 'l'
	argc = len(argv)
	process_args(argv)
	process_input()
}
