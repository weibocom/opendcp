<?php
@session_start();

Class openstack
{
    static $cur_region = 'region1';
    static $arr_region_config = array();
    static $arr_region_config_orig = array(
        'region1' => array(
            'name' => 'Region1',
            'authorization_domain' => 'http://{{controller_ip}}:5000',
            'controller_domain' => 'http://{{controller_ip}}:5000',
            'image_domain' => 'http://{{controller_ip}}:9292',
            'server_domain' => 'http://{{controller_ip}}:8774',
            'network_domain' => 'http://{{controller_ip}}:9696',
            'storage_domain' => 'http://{{controller_ip}}:8776',
            'admin_user' => 'admin',
            'admin_pass' => 'root',
        ),
    );

    static $token = '';
    static $user_id = '';
    static $arr_project = array();
    static $cur_project_id = '';
    static $arr_httpinfo = array();
    static $arr_error = array();
    static $needadmin = false;
    static $arr_serverstatus = array(
        'ACTIVE' => '运行中',
        'BUILD' => '创建中',
        'DELETED' => '已删除',
        'HARD_REBOOT' => '硬重启中',
        'REBOOT' => '软重启中',
        'STOPPED' => '已关机',
        'ERROR' => '错误',
        'SUSPENDED' => '已挂起',
        'BUILDING' => '创建中',
        'MIGRATING' => '迁移中',
    );

    static function getControllerIp()
    {
        require_once('keydata.php');
        $ip = keydata::getContentByKey('controller_ip');
        if (empty($ip)) {
            return false;
        }
        foreach (self::$arr_region_config_orig as $rn => $rc) {
            foreach ($rc as $k => $v) {
                self::$arr_region_config[$rn][$k] = str_replace('{{controller_ip}}', $ip, $v);
            }
        }
        return true;
    }

    static function needOpenstackLogin()
    {
        if (!self::getControllerIp()) return false;
        if (!empty($_SESSION['openstack_token_' . self::$cur_region])) {
            //self::$token = $_SESSION['openstack_token_'.self::$cur_region];
            //return true;
        }

        $username = self::$arr_region_config[self::$cur_region]['admin_user'];
        $password = self::$arr_region_config[self::$cur_region]['admin_pass'];
        self::getAdminToken();
        self::getToken($username, $password, '');
        if (!empty(self::$token)) {
            self::getProject();
            return true;
        }
        return false;
    }

    static function setUserProject($project_id)
    {
        if (!self::getControllerIp()) return false;
        self::$cur_project_id = $project_id;
        $_SESSION['openstack_cur_project_id'] = self::$cur_project_id;
        $username = self::$arr_region_config[self::$cur_region]['admin_user'];
        $password = self::$arr_region_config[self::$cur_region]['admin_pass'];
        return self::getToken($username, $password, $project_id);
    }

    static function setRegion($region_id)
    {
        if (!self::getControllerIp()) return false;
        self::$cur_region = $region_id;
        $_SESSION['openstack_region'] = $region_id;
        $username = self::$arr_region_config[self::$cur_region]['admin_user'];
        $password = self::$arr_region_config[self::$cur_region]['admin_pass'];
        return self::getToken($username, $password, $project_id);
    }

    static function getAdminToken($project_id = '5e140c8f1f38414ab9160a4164a7ca93')
    {
        if (!self::getControllerIp()) return false;
        $url = self::$arr_region_config[self::$cur_region]['authorization_domain'] . '/v3/auth/tokens';
        $method = 'POST';
        $username = self::$arr_region_config[self::$cur_region]['admin_user'];
        $password = self::$arr_region_config[self::$cur_region]['admin_pass'];
        $arr = array(
            'auth' => array(
                'identity' => array(
                    'methods' => array('password'),
                    'password' => array(
                        'user' => array(
                            'name' => $username,
                            'domain' => array('name' => 'default'),
                            'password' => $password,
                        ),
                    ),
                ),
            ),
        );
        if (!empty($project_id)) {
            $arr['auth']['scope'] = array(
                'project' => array(
                    'domain' => array('name' => 'default'),
                    //'id'=>$project_id,
                    'name' => 'admin',
                ),
            );
        }
        $param = json_encode($arr);

        $arr_header = array(
            'Content-Type: application/json',
        );

        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_FOLLOWLOCATION, true);
        curl_setopt($ch, CURLOPT_CUSTOMREQUEST, $method);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $param);
        curl_setopt($ch, CURLOPT_HTTPHEADER, $arr_header);
        curl_setopt($ch, CURLOPT_HEADER, 1);
        $rs = curl_exec($ch);
        $arr_match = array();
        if (preg_match('/X-Subject-Token: (.+?)\\r\\n/si', $rs, $arr_match)) {
            $_SESSION['openstack_admin_token'] = $arr_match[1];
            $arr_match = array();
            if (preg_match('/\\r\\n\\r\\n(.+)/si', $rs, $arr_match)) {
                $rs = $arr_match[1];
                $arr_rs = @json_decode($rs, true);
                if (!empty($arr_rs['token']['user']['id'])) {
                    $_SESSION['openstack_admin_user_id'] = $arr_rs['token']['user']['id'];
                }
                return true;
            }
        }
        return false;
    }

    static function getToken($username, $password, $project_id = '')
    {
        if (!self::getControllerIp()) return false;
        $url = self::$arr_region_config[self::$cur_region]['authorization_domain'] . '/v3/auth/tokens';
        $method = 'POST';
        $arr = array(
            'auth' => array(
                'identity' => array(
                    'methods' => array('password'),
                    'password' => array(
                        'user' => array(
                            'name' => $username,
                            'domain' => array('name' => 'default'),
                            'password' => $password,
                        ),
                    ),
                ),
            ),
        );
        if (!empty($project_id)) {
            $arr['auth']['scope'] = array(
                'project' => array(
                    'domain' => array('name' => 'default'),
                    'id' => $project_id,
                ),
            );
        }
        $param = json_encode($arr);

        $arr_header = array(
            'Content-Type: application/json',
        );

        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_FOLLOWLOCATION, true);
        curl_setopt($ch, CURLOPT_CUSTOMREQUEST, $method);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $param);
        curl_setopt($ch, CURLOPT_HTTPHEADER, $arr_header);
        curl_setopt($ch, CURLOPT_HEADER, 1);
        $rs = curl_exec($ch);
        $arr_match = array();
        if (preg_match('/X-Subject-Token: (.+?)\\r\\n/si', $rs, $arr_match)) {
            self::$token = $arr_match[1];
            $_SESSION['openstack_token_' . self::$cur_region] = self::$token;
            $arr_match = array();
            if (preg_match('/\\r\\n\\r\\n(.+)/si', $rs, $arr_match)) {
                $rs = $arr_match[1];
                $arr_rs = @json_decode($rs, true);
                if (!empty($arr_rs['token']['user']['id'])) {
                    self::$user_id = $arr_rs['token']['user']['id'];
                    $_SESSION['openstack_user_id'] = self::$user_id;
                }
                return true;
            }
        }
        return false;
    }

    static function getProject()
    {
        if (!self::getControllerIp()) return false;
        $arr_project = self::getUserProjectList(self::$user_id);
        if (!empty($arr_project['projects'])) {
            $_SESSION['openstack_arr_project'] = $arr_project['projects'];
            foreach ($arr_project['projects'] as $k => $oneproject) {
                self::$cur_project_id = $oneproject['id'];
                $_SESSION['openstack_cur_project_id'] = self::$cur_project_id;
                self::setUserProject(self::$cur_project_id);
                break;
            }
        }
    }

    static function getUserProjectList($id)
    {
        if (!self::getControllerIp()) return false;
        $arr = array();
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['controller_domain'] . '/v3/users/' . $id . '/projects',
            $arr
        );
    }

    static function getImageList($name = '', $marker = 0, $limit = 20)
    {
        if (!self::getControllerIp()) return false;
        $arr = array(
            'marker' => $marker,
            'limit' => $limit,
        );
        if (!empty($name)) {
            $arr['name'] = $name;
        }
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['image_domain'] . '/v2/images',
            $arr
        );
    }

    static function getKeypairList($arr_option = array(), $marker = 0, $limit = 20)
    {
        if (!self::getControllerIp()) return false;
        $arr = array(
            'marker' => $marker,
            'limit' => $limit,
        );
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['server_domain'] . '/v2.1/os-keypairs',
            $arr
        );
    }

    static function getOneNetwork($network_id)
    {
        if (!self::getControllerIp()) return false;
        $arr = array();
        $networkinfo = self::send_http(
            self::$arr_region_config[self::$cur_region]['network_domain'] . '/v2.0/networks/' . $network_id,
            $arr
        );
        $networkinfo = empty($networkinfo['network']) ? array() : $networkinfo['network'];
        $subnetlist = self::getNetworkSubnetList($network_id);
        $networkinfo['subnets'] = empty($subnetlist['subnets']) ? array() : $subnetlist['subnets'];
        $portlist = self::getNetworkPortList($network_id);
        $networkinfo['ports'] = empty($portlist['ports']) ? array() : $portlist['ports'];
        return $networkinfo;
    }

    static function getNetworkPortList($network_id)
    {
        if (!self::getControllerIp()) return false;
        $arr = array(
            'network_id' => $network_id,
        );
        self::$needadmin = true;
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['network_domain'] . '/v2.0/ports',
            $arr
        );
    }

    static function getNetworkIpStatus($network_id)
    {
        if (!self::getControllerIp()) return false;
        $arr = array();
        self::$needadmin = true;
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['network_domain'] . '/v2.0/network-ip-availabilities/' . $network_id,
            $arr
        );
    }

    static function getNetworkSubnetList($network_id, $marker = 0, $limit = 20)
    {
        if (!self::getControllerIp()) return false;
        $arr = array(
            'network_id' => $network_id,
            'marker' => $marker,
            'limit' => $limit,
        );
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['network_domain'] . '/v2.0/subnets',
            $arr
        );
    }

    static function getNetworkList($arr_option = array(), $marker = 0, $limit = 20)
    {
        if (!self::getControllerIp()) return false;
        $arr = array(
            'marker' => $marker,
            'limit' => $limit,
        );
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['network_domain'] . '/v2.0/networks',
            $arr
        );
    }

    static function getFlavorList($arr_option = array(), $marker = 0, $limit = 20)
    {
        if (!self::getControllerIp()) return false;
        $arr = array(
            //'marker'=>$marker,
            //'limit'=>$limit,
        );
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['server_domain'] . '/v2.1/flavors/detail',
            $arr
        );
    }

    static function getHypervisorServerList($id)
    {
        if (!self::getControllerIp()) return false;
        self::$needadmin = true;
        $arr_server = self::send_http(
            self::$arr_region_config[self::$cur_region]['server_domain'] . '/v2.1/os-hypervisors/' . $id . '/servers'

        );
        return $arr_server;
    }

    static function getHypervisorList($arr_option = array(), $marker = 0, $limit = 20)
    {
        if (!self::getControllerIp()) return false;
        $arr = array(
            'marker' => $marker,
            'limit' => $limit,
        );

        if (isset($arr_option['region']) && !empty(self::$arr_region_config[$arr_option['region']])) {
            $currRegion = $arr_option['region'];
        } else {
            $currRegion = self::$cur_region;
        }

        self::$needadmin = true;
        $arr_server = self::send_http(
            self::$arr_region_config[$currRegion]['server_domain'] . '/v2.1/os-hypervisors/detail',
            $arr
        );

        return $arr_server;
    }

    static function getStorageHostList($arr_option = array(), $marker = 0, $limit = 20)
    {
        if (!self::getControllerIp()) return false;
        $arr = array(
            'marker' => $marker,
            'limit' => $limit,
        );

        if (isset($arr_option['region']) && !empty(self::$arr_region_config[$arr_option['region']])) {
            $currRegion = $arr_option['region'];
        } else {
            $currRegion = self::$cur_region;
        }

        self::$needadmin = true;
        $arr_server = self::send_http(
            self::$arr_region_config[$currRegion]['storage_domain'] . '/v3/' . $_SESSION['openstack_cur_project_id'] .'/os-hosts'
        );

        $detail_list = array();
        foreach ($arr_server['hosts'] as $one_host){
            $host_name = $one_host['host_name'];
            $detail = self::getStorageHostDetail($host_name);
            if($one_host['service-status'] != 'available'){
                continue;
            }
            $detail_list[] = $detail;
        }

        return $detail_list;
    }

    static function getStorageHostDetail($host_name = '')
    {
        $curRegion = self::$cur_region;
        self::$needadmin = true;
        $host_detail = self::send_http(
            self::$arr_region_config[$curRegion]['storage_domain'] . '/v3/' . $_SESSION['openstack_cur_project_id'] .'/os-hosts/' . $host_name
        );
        return $host_detail;
    }

    static function getOneHypervisor($id)
    {
        if (!self::getControllerIp()) return false;
        self::$needadmin = true;
        $hinfo = self::send_http(
            self::$arr_region_config[self::$cur_region]['server_domain'] . '/v2.1/os-hypervisors/' . $id

        );
        $arr_server = self::getHypervisorServerList($hinfo['hypervisor']['hypervisor_hostname']);
        $hinfo['hypervisor']['arr_server'] = empty($arr_server['hypervisors'][0]['servers']) ? array() : $arr_server['hypervisors'][0]['servers'];
        foreach ($hinfo['hypervisor']['arr_server'] as $k => $v) {
            $serverid = $v['uuid'];
            $sinfo = self::getOneServer($serverid);
            $hinfo['hypervisor']['arr_server'][$k]['server'] = $sinfo;
        }
        return $hinfo['hypervisor'];
    }

    static function getServerList($arr_option = array(), $marker = 0, $limit = 20)
    {
        if (!self::getControllerIp()) return false;
        $arr = array(
            'marker' => $marker,
            'limit' => $limit,
        );
        if (!empty($arr_option['name'])) {
            $arr['name'] = $arr_option['name'];
        }
        if (!empty($arr_option['status'])) {
            $arr['status'] = $arr_option['status'];
        }
        $arr_server = self::send_http(
            self::$arr_region_config[self::$cur_region]['server_domain'] . '/v2.1/servers/detail',
            $arr
        );

        foreach ($arr_server['servers'] as $k => $v) {
            $arr_server['servers'][$k]['status_str'] = empty(self::$arr_serverstatus[$v['status']]) ? $v['status'] : self::$arr_serverstatus[$v['status']];
        }
        return $arr_server;
    }

    static function getOneServer($id)
    {
        if (!self::getControllerIp()) return false;
        $arr = array();
        $sinfo = self::send_http(
            self::$arr_region_config[self::$cur_region]['server_domain'] . '/v2.1/servers/' . $id,
            $arr
        );

        $sinfo['server']['status_str'] = empty(self::$arr_serverstatus[$sinfo['server']['status']]) ? $sinfo['server']['status'] : self::$arr_serverstatus[$sinfo['server']['status']];
        $sinfo['server']['flavorinfo'] = self::getOneFlavor($sinfo['server']['flavor']['id']);
        $sinfo['server']['imageinfo'] = self::getOneImage($sinfo['server']['image']['id']);
        return $sinfo['server'];
    }

    static function getOneFlavor($id)
    {
        if (!self::getControllerIp()) return false;
        $arr = array();
        $sinfo = self::send_http(
            self::$arr_region_config[self::$cur_region]['server_domain'] . '/v2.1/flavors/' . $id,
            $arr
        );
        return $sinfo['flavor'];
    }

    static function getOneImage($id)
    {
        if (!self::getControllerIp()) return false;
        $arr = array();
        $sinfo = self::send_http(
            self::$arr_region_config[self::$cur_region]['image_domain'] . '/v2/images/' . $id,
            $arr
        );
        return $sinfo;
    }


    static function saveImage($id, $name)
    {
        if (!self::getControllerIp()) return false;
        $arr = array(
            'createImage' => array(
                'name' => $name,
            ),
        );
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['server_domain'] . '/v2.1/servers/' . $id . '/action',
            $arr,
            array(),
            'POST'
        );
    }

    static function deleteImage($id)
    {
        if (!self::getControllerIp()) return false;
        $arr = array();
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['image_domain'] . '/v2/images/' . $id,
            $arr,
            array(),
            'DELETE'
        );
    }

    static function rebootServerHard($id)
    {
        if (!self::getControllerIp()) return false;
        return self::rebootServer($id, 'HARD');
    }

    static function rebootServer($id, $type = 'SOFT')
    {
        if (!self::getControllerIp()) return false;
        $arr = array(
            'reboot' => array(
                'type' => $type,
            ),
        );
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['server_domain'] . '/v2.1/servers/' . $id . '/action',
            $arr,
            array(),
            'POST'
        );
    }

    static function deleteServer($id)
    {
        if (!self::getControllerIp()) return false;
        $arr = array();
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['server_domain'] . '/v2.1/servers/' . $id,
            $arr,
            array(),
            'DELETE'
        );
    }

    static function createKeypair($name)
    {
        if (!self::getControllerIp()) return false;
        $arr = array();
        $arr['keypair'] = array('name' => $name);
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['server_domain'] . '/v2.1/os-keypairs',
            $arr,
            array(),
            'POST'
        );
    }

    static function createServer($arr_data)
    {
        if (!self::getControllerIp()) return false;
        $arr = array();
        $arr['name'] = $arr_data['name'];
        $arr['imageRef'] = $arr_data['image'];
        $arr['flavorRef'] = $arr_data['flavor'];
        $arr['availability_zone'] = 'nova';
        //$arr['availability_zone'] = 'nova:75-29-211-yf-core.jpool.sinaimg.cn';
        $arr['networks'] = array(
            array(
                'uuid' => $arr_data['network'],
            )
        );
        $arr['security_groups'] = array(
            array('name' => 'default'),
        );
        //$arr['key_name'] = $arr_data['keypair'];
        $arr['adminPass'] = 'ASDqwe123';
        $arr_param = array('server' => $arr);
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['server_domain'] . '/v2.1/servers',
            $arr_param,
            array(),
            'POST'
        );
    }

    static function createFlavor($arr_data)
    {
        if (!self::getControllerIp()) return false;
        $arr = array();
        $arr['name'] = $arr_data['flavor_name'];
        $arr['ram'] = $arr_data['flavor_ram'];
        $arr['disk'] = $arr_data['flavor_disk'];
        $arr['vcpus'] = $arr_data['flavor_vcpus'];
        self::$needadmin = true;

        $arr_param = array('flavor' => $arr);
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['server_domain'] . '/v2.1/flavors',
            $arr_param,
            array(),
            'POST',
            true,
            true
        );
    }

    static function createNetwork($arr_data)
    {
        if (!self::getControllerIp()) return false;
        $arr = array();
        $arr['name'] = $arr_data['network_name'];
        $arr['provider:network_type'] = 'flat';
        $arr['provider:physical_network'] = $arr_data['network_providername'];
        $arr['admin_state_up'] = true;

        $arr_param = array('network' => $arr);
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['network_domain'] . '/v2.0/networks',
            $arr_param,
            array(),
            'POST',
            false,
            true
        );
    }

    static function createSubnet($arr_data)
    {
        if (!self::getControllerIp()) return false;
        $arr = array();
        $arr['name'] = $arr_data['network_subnet_name'];
        $arr['network_id'] = $arr_data['network_id'];
        $arr['enable_dhcp'] = true;
        $ap = array();
        $arr_ip = explode(',', $arr_data['network_subnet_ip']);
        $ap[] = array(
            'start' => $arr_ip[0],
            'end' => $arr_ip[1],
        );
        $arr['allocation_pools'] = $ap;
        $arr['gateway_ip'] = $arr_data['network_subnet_gateway'];
        $arr['cidr'] = $arr_data['network_subnet_range'];
        $arr['ip_version'] = 4;

        $arr_param = array('subnet' => $arr);
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['network_domain'] . '/v2.0/subnets',
            $arr_param,
            array(),
            'POST',
            false,
            true
        );
    }

    static function getOnePort($port_id)
    {

        if (!self::getControllerIp()) return false;
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['network_domain'] . '/v2.0/ports/' . $port_id,

            array(),
            'GET',
            true,
            true
        );
    }

    static function updatePorts($port_id, $arr_data)
    {

        if (!self::getControllerIp()) return false;
        $arr_param = array('port' => $arr_data);
        print_r($arr_param);
        return self::send_http(
            self::$arr_region_config[self::$cur_region]['network_domain'] . '/v2.0/ports/' . $port_id,
            $arr_param,
            array(),
            'PUT',
            true,
            true
        );
    }

    static function send_http($url, $param = array(), $header = array(), $method = 'GET', $admin = false, $retmsg = false)
    {
        if (empty($url)) return false;

        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_FOLLOWLOCATION, true);
        curl_setopt($ch, CURLOPT_CUSTOMREQUEST, $method);
        $str_param = empty($param) ? '' : json_encode($param);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $str_param);
        if (self::$needadmin) {
            $header[] = 'X-Auth-Token: ' . $_SESSION['openstack_admin_token'];
        } else {
            $header[] = 'X-Auth-Token: ' . $_SESSION['openstack_token_' . self::$cur_region];
        }
        $header[] = 'Content-Type: application/json';
        curl_setopt($ch, CURLOPT_HTTPHEADER, $header);
        $rs = curl_exec($ch);
        self::$arr_httpinfo = curl_getinfo($ch);

        if (self::$arr_httpinfo['http_code'] == 401) {
            $_SESSION['openstack_token_' . self::$cur_region] = '';
            self::$token = '';
            return false;
        }
        if (self::$arr_httpinfo['http_code'] == 201) {
            if ($retmsg) {
                $arr_rs = @json_decode($rs, true);
                return $arr_rs;
            }
            return true;
        }
        if (self::$arr_httpinfo['http_code'] == 204) {
            if ($retmsg) {
                $arr_rs = @json_decode($rs, true);
                return $arr_rs;
            }
            return true;
        }
        if (self::$arr_httpinfo['http_code'] == 202) {
            if ($retmsg) {
                $arr_rs = @json_decode($rs, true);
                return $arr_rs;
            }
            return true;
        }
        if (self::$arr_httpinfo['http_code'] != 200) {
            self::$arr_error[] = curl_error($ch);
            curl_close($ch);
            return false;
        }
        curl_close($ch);
        $arr_rs = @json_decode($rs, true);
        return $arr_rs;
    }

    public static function clearSession()
    {
        unset($_SESSION['openstack_token']);
        foreach (self::$arr_region_config as $region_id => $r) {
            unset($_SESSION['openstack_token_' . $region_id]);
        }
        unset($_SESSION['openstack_region']);
        unset($_SESSION['openstack_admin_token']);
        unset($_SESSION['openstack_admin_user_id']);
        unset($_SESSION['openstack_cur_user_id']);
        unset($_SESSION['openstack_cur_project_id']);
        unset($_SESSION['openstack_arr_project']);
    }
}

if (!empty($_SESSION['openstack_region'])) {
    openstack::$cur_region = $_SESSION['openstack_region'];
}


?>
