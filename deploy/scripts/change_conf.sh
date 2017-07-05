#!/bin/sh

echo "Enter the IP of the this machine [ENTER]: "
read host_ip
echo "Enter the IP of harbor [ENTER]: "
read harbor_ip
echo "Enter your Aliyun Key [ENTER]: "
read ali_key
echo "Enter your Aliyun Secret [ENTER]: "
read ali_secret



# 定位到脚本所在目录
cd `dirname $0`

# 定位到上层目录也就是deploy目录
DEPLOY_DIR=$(dirname $(pwd))
cd $DEPLOY_DIR

echo "current pwd:"`pwd`
echo "Changing config..."
FILES="conf/hubble_conf.php conf/web_conf.php docker-compose.yml conf/octans_roles/init/files/docker ../db_init/sql/orion.sql conf/jupiter.conf ../ui/for_openstack/doInit.php"
for FILE in $FILES; do
    echo "  changing config in $FILE"

    if [ ! -f ${DEPLOY_DIR}/${FILE}_orig ]; then
          #echo "cp ${DEPLOY_DIR}/${FILE} ${DEPLOY_DIR}/${FILE}_orig"
          cp ${DEPLOY_DIR}/${FILE} ${DEPLOY_DIR}/${FILE}_orig

    else
          #echo "cp ${DEPLOY_DIR}/${FILE}_orig ${DEPLOY_DIR}/${FILE}"
          cp ${DEPLOY_DIR}/${FILE}_orig ${DEPLOY_DIR}/${FILE}
    fi


    sed -i "s/harbor_ip/$harbor_ip/g" $FILE
    sed -i "s/host_ip/$host_ip/g" $FILE
done
echo "Changing config...[DONE]"



ALIYUN_FILE="conf/jupiter.json"
if [ ! -f ${DEPLOY_DIR}/${ALIYUN_FILE}_orig ]; then
      #echo "cp ${DEPLOY_DIR}/${ALIYUN_FILE} ${DEPLOY_DIR}/${ALIYUN_FILE}_orig"
      cp ${DEPLOY_DIR}/${ALIYUN_FILE} ${DEPLOY_DIR}/${ALIYUN_FILE}_orig

else
      #echo "cp ${DEPLOY_DIR}/${ALIYUN_FILE}_orig ${DEPLOY_DIR}/${ALIYUN_FILE}"
      cp ${DEPLOY_DIR}/${ALIYUN_FILE}_orig ${DEPLOY_DIR}/${ALIYUN_FILE}
fi

sed -i "s/your_key_id/$ali_key/g" ${DEPLOY_DIR}/${ALIYUN_FILE}
sed -i "s/your_secret/$ali_secret/g" ${DEPLOY_DIR}/${ALIYUN_FILE}
