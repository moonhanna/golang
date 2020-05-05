# GOLANG



### 개발환경설정

1. golang 설치
   - https://golang.org/
2. Atom Editor 설치
   - https://atom.io/
3. git 설치
   - https://git-scm.com/

### 설치 확인

```bash
C:\Users\mhn97>go
Go is a tool for managing Go source code.

Usage:

        go <command> [arguments]

The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

        buildmode   build modes
        c           calling between Go and C
        cache       build and test caching
        environment environment variables
        filetype    file types
        go.mod      the go.mod file
        gopath      GOPATH environment variable
        gopath-get  legacy GOPATH go get
        goproxy     module proxy protocol
        importpath  import path syntax
        modules     modules, module versions, and more
        module-get  module-aware go get
        module-auth module authentication using go.sum
        module-private module configuration for non-public modules
        packages    package lists and patterns
        testflag    testing flags
        testfunc    testing functions

Use "go help <topic>" for more information about that topic.


C:\Users\mhn97>atom //Atom Editor 실행

C:\Users\mhn97>git
usage: git [--version] [--help] [-C <path>] [-c <name>=<value>]
           [--exec-path[=<path>]] [--html-path] [--man-path] [--info-path]
           [-p | --paginate | -P | --no-pager] [--no-replace-objects] [--bare]
           [--git-dir=<path>] [--work-tree=<path>] [--namespace=<name>]
           <command> [<args>]

These are common Git commands used in various situations:

start a working area (see also: git help tutorial)
   clone     Clone a repository into a new directory
   init      Create an empty Git repository or reinitialize an existing one

work on the current change (see also: git help everyday)
   add       Add file contents to the index
   mv        Move or rename a file, a directory, or a symlink
   restore   Restore working tree files
   rm        Remove files from the working tree and from the index

examine the history and state (see also: git help revisions)
   bisect    Use binary search to find the commit that introduced a bug
   diff      Show changes between commits, commit and working tree, etc
   grep      Print lines matching a pattern
   log       Show commit logs
   show      Show various types of objects
   status    Show the working tree status

grow, mark and tweak your common history
   branch    List, create, or delete branches
   commit    Record changes to the repository
   merge     Join two or more development histories together
   rebase    Reapply commits on top of another base tip
   reset     Reset current HEAD to the specified state
   switch    Switch branches
   tag       Create, list, delete or verify a tag object signed with GPG

collaborate (see also: git help workflows)
   fetch     Download objects and refs from another repository
   pull      Fetch from and integrate with another repository or a local branch
   push      Update remote refs along with associated objects

'git help -a' and 'git help -g' list available subcommands and some
concept guides. See 'git help <command>' or 'git help <concept>'
to read about a specific subcommand or concept.
See 'git help git' for an overview of the system.

C:\Users\mhn97>
```

cmd에서 go, atom, git을 각각 입력하여 확인

### workspace

src - 소스코드

pkg - 패키지

bin - 컴파일 후 binary

각각의 폴더 생성

### 환경변수 등록 및 확인

시스템 속성 -> 환경변수 편집

1. GOROOT - go 설치된 경로
2. GOPATH - go workspace 경로
3. GOBIN - go workspace의 bin 경로

```bash
C:\Users\mhn97>go env
set GOBIN=C:\Users\mhn97\Desktop\Study\go\bin
set GOPATH=C:\Users\mhn97\Desktop\Study\go\
set GOROOT=C:\Go
C:\Users\mhn97>
```

### 패키지 다운

1. Atom Editor 실행
2. File - Settings
3. go-plus install
4. script install
5. platformio-ide-terminal install
   - platformio-ide-terminal 설치 후 맨 밑에 +로 터미널 창 열어서 확인
   - 만약 안열린다면 cmd - 속성 - 레거시 콘솔 사용 - Atom Editor 재실행
6. pakages에서 다운받은 패키지 확인

### 프로젝트 파일 추가

File - Add Project Folder - workpace 추가

### 실행 방법

1. run - exe파일 없이 단위 테스트용
   - ctrl + shift + b
   - go run ~~~.go
2. build - 실행파일 생성
   - go build ~~~.go
3. install - 상위 폴더의 이름으로 실행파일 생성(배포의 개념), bin 폴더에 생김
   - go install

