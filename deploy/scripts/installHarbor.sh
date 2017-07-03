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


if [ $# != 4 ] ; then
   success "USAGE: $0 ip           port    path    cloud"
   success "e.g.:  $0 xx.xx.xx.xx  12380   /data1  aliyun"
exit 1;
fi

info "add config of harbor registry conf/jupiter.conf"

sed -i "s#harbor_registry = .*#harbor_registry = $1:$2#g" ../conf/jupiter.conf
 
info "begin install..."

info "install path:$3"
mkdir -p $3
cd $3

dir="harbor"
if [ -d "$dir" ]; then
    info "remove dir harbor"
    cd harbor
    rm -rf harbor
else
    info "create dir harbor"
    mkdir harbor
    cd harbor
fi

oldTgz=`ls -l|grep harbor|grep .tgz|awk '{print $9}'|head -1`
if [ -f "$oldTgz" ]; then
    info "use exists tgz: $oldTgz"
else
    info "download harbor tgz..."
    wget https://github.com/vmware/harbor/releases/download/0.4.5/harbor-online-installer-0.4.5.tgz
    success "success!"
fi

harborTgz=`ls -l|grep harbor|awk '{print $9}'|head -1`
info "find harbrTar file:${harborTgz}"

info "tar xvf ${harborTgz}..."
tar xvf "${harborTgz}"
success "success!"

cd harbor
info "now at `pwd`"
path=`pwd`

info "run ./prepare..."
./prepare
success "run ./prepare success!"

info "replace harbor ip from reg.mydomain.com to $1..."
sed -i "s/reg.mydomain.com/$1/g" harbor.cfg
success "success!"

if [ "$4" = "aliyun" ]; then
    info "replace aliyun images..."
    image_array=(
        "vmware/harbor-log:0.4.5"
        "library/registry:2.5.0"
        "vmware/harbor-db:0.4.5"
        "vmware/harbor-ui:0.4.5"
        "vmware/harbor-jobservice:0.4.5"
        "nginx:1.9"
    )

    image_r_array=(
        "registry.cn-beijing.aliyuncs.com/opendcp/harbor-log:0.4.5"
        "registry.cn-beijing.aliyuncs.com/opendcp/registry:2.5.0"
        "registry.cn-beijing.aliyuncs.com/opendcp/harbor-db:0.4.5"
        "registry.cn-beijing.aliyuncs.com/opendcp/harbor-ui:0.4.5"
        "registry.cn-beijing.aliyuncs.com/opendcp/harbor-jobservice:0.4.5"
        "registry.cn-beijing.aliyuncs.com/opendcp/nginx:latest"
    )

    for inx in "${!image_array[@]}";
    do
        sed -i "s#${image_array[$inx]}#${image_r_array[$inx]}#g" docker-compose.yml
    done
    success "replace aliyun images success!"
fi

info "run ./install.sh"
./install.sh
success "install and start success!"


info "replace docker-compose.yml prot $2..."
sed -i "s/80:80/$2:80/g" docker-compose.yml
success "success!"

info "replace docker-compose.yml path $3..."
sed -i "s#/data/#$3/#g" docker-compose.yml
success "success!"

info "replace common/config/registry/config.yml port $2..."
sed -i "s/$1/$1:$2/g" common/config/registry/config.yml
success "success!"

docker-compose down
docker-compose up -d



sleep 10
info "create base,default_cluster registry..."
curl -H "content-type: application/json" -d "{\"project_name\":\"default_cluster\",\"public\":1}" -u admin:Harbor12345 "http://$1:$2/api/projects"
curl -H "content-type: application/json" -d "{\"project_name\":\"base\",\"public\":1}" -u admin:Harbor12345 "http://$1:$2/api/projects"
success "create base,default_cluster success!"

success "create success!"

