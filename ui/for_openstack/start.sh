while true
do
#if test $( pgrep -f "php for_openstack/doInit.php" | wc -l ) -eq 0
#then
        cd /data1/web && nohup php for_openstack/doInit.php >> /tmp/init.log &
#else
#        echo 'hasmake'
#fi

sleep 10
done
