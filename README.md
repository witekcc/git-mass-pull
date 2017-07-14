# git utility

## Prerequisites
install libgit2 v0.22
https://github.com/libgit2/libgit2
    - checkout old v0.22 branch and build it:
        mkdir build && cd build
        cmake ..
        cmake --build .
        cmake --build . --target install
use `git "gopkg.in/libgit2/git2go.v22" in the app`


## Build
export LD_LIBRARY_PATH=/lib:/usr/lib:/usr/local/lib - !!!!!
go build 

### Run
export LD_LIBRARY_PATH=/lib:/usr/lib:/usr/local/lib - ???
export GITLAB_TOKEN=XXX
GODEBUG=cgocheck=0 ./git-mass-pull
GODEBUG=cgocheck=0 ./git-mass-pull -t [token] -registryType gitlab -d /home/witek/_backups/gitlab-pd
GODEBUG=cgocheck=0 ./git-mass-pull -u [username] -p [password] -registryType gitlab -d /home/witek/_backups/gitlab-pd
!!!
set gitlab token in GITLAB_TOKEN env var
user name and pass don't work