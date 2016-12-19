DIRS="orion jupiter octans db_init proxy hubble dcp_open imagebuild"
REG=<your_repo>
LOC=opendcp
VER=latest

if [ "" != "$1" ] ;then
    VER=$1
fi

for DIR in $DIRS; do
    TAG=${REG}/${LOC}/${DIR}:${VER}
    echo "Pulling $TAG ..."
    docker pull $TAG
done