#!/bin/sh

echo "Enter the IP of the this machine [ENTER]: "
read host_ip
echo "Enter the IP of harbor [ENTER]: "
read harbor_ip
echo "Enter your Aliyun Key [ENTER]: "
read ali_key
echo "Enter your Aliyun Secret [ENTER]: "
read ali_secret
echo "Do you have OpenStack account [y/n]"
read hasAccount
if [ $hasAccount = "y" ]; then
    echo "Enter your OpenStack Platform IP"
    read op_ip
    echo "Enter your OpenStack Platform Port"
    read op_port
    echo "Enter your OpenStack UserName"
    read op_user_name
    echo "Enter your OpenStack PassWord"
    read op_password
fi

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

sed -i "s/your_aliyun_key_id/$ali_key/g" ${DEPLOY_DIR}/${ALIYUN_FILE}
sed -i "s/your_aliyun_secret/$ali_secret/g" ${DEPLOY_DIR}/${ALIYUN_FILE}

sed -i "s/host_ip/$host_ip/g" ${DEPLOY_DIR}/${ALIYUN_FILE}


if [ $hasAccount = "y" ] ;then
sed -i "s/your_op_ip/$op_ip/g" ${DEPLOY_DIR}/${ALIYUN_FILE}
sed -i "s/your_op_port/$op_port/g" ${DEPLOY_DIR}/${ALIYUN_FILE}
sed -i "s/your_op_username/$op_user_name/g" ${DEPLOY_DIR}/${ALIYUN_FILE}
sed -i "s/your_op_password/$op_password/g" ${DEPLOY_DIR}/${ALIYUN_FILE}
sed -i "s/127.0.0.1/$op_ip/g" ${DEPLOY_DIR}/../jupiter/run.sh
fi
