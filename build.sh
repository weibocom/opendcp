#!/bin/sh

reset=$(tput sgr0)
green=$(tput setaf 76)
blue=$(tput setaf 25)
red=$(tput setaf 1)

info() {
    printf "${blue}➜ %s${reset}\n" "$@"
}
error() {
    printf "${red}➜ %s${reset}\n" "$@"
}
success() {
    printf "${green}✔ %s${reset}\n" "$@"
}

info "welcome! "
info "attention: docker 1.10.0+ and docker-compose 1.6.0+"
info "attention: install path Disk capacity is large enough"
echo "";

###check docker and docker-compose
command -v docker >/dev/null 2>&1 || { error "Opendcp require docker but it's not installed. "; error "Please refer to https://docs.docker.com/ or https://docs.docker.com/engine/installation/linux/"; error "Aborting."; exit 1; }
command -v docker-compose >/dev/null 2>&1 || { error "Opendcp require docker-compose but it's not installed.  Aborting."; error "Please refer to https://docs.docker.com/compose/install/";  error "Aborting."; exit 1; }


if [[ $# < 1 ]] ; then
   success "USAGE: $0 CLOUD   TAG "
   success "e.g.:  $0 aliyun"
   success "e.g.:  $0 aliyun  latest"
   exit 1;
fi

if [[ "aliyun" != "$1" && "dockerio" != "$1" ]] ;then
    success "SUPPORT CLOUD: dockerio , aliyun !"
    exit 1;
fi



DIRS="orion jupiter octans db_init proxy hubble ui imagebuild"
REG=weibo.com
LOC=opendcp
VER=latest
CLOUD=$1

if [ "" != "$2" ] ;then
    VER=$2
fi

fail=0
for DIR in $DIRS; do
    cd $DIR
    TAG=${REG}/${LOC}/${DIR}:${VER}

    info "build docker image for $DIR ..."
    info "tag is $TAG"

    if [ "$CLOUD" = "aliyun" ]; then
        \cp -f Dockerfile_Aliyun  Dockerfile
    fi

    if [ "$CLOUD" = "dockerio" ]; then
        \cp -f Dockerfile_Dockerio  Dockerfile
    fi

    if [ "$DIR" = "imagebuild" ]; then
        ./build.sh $CLOUD $TAG
    else
        ./build.sh $TAG
    fi


    if [[ 0 != $? ]] ; then
        info "FAIL"
        fail=1
        break
    fi

    #echo "push docker image for $DIR ..."
    #docker push $TAG
    if [[ 0 != $? ]] ; then
        echo "FAIL"
        fail=1
        break
    fi
    
    cd ..
    success "$DIR OK"
done

if [[ $fail == 1 ]]; then
    exit 1
fi

echo "ALL DONE"
