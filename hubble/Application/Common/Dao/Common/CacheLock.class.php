<?php
/**
 *    Copyright (C) 2016 Weibo Inc.
 *
 *    This file is part of Opendcp.
 *
 *    Opendcp is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation; version 2 of the License.
 *
 *    Opendcp is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with Opendcp; if not, write to the Free Software
 *    Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA
 */


/**
 * Created by PhpStorm.
 * User: reposkeeper
 * Date: 16/9/4
 * Time: 下午5:36
 */

class CacheLock {

    private $postfix = '_2fb9ad68_9b76_4029_8be6_9024f8c9434f';
    private $lockString = 'LOCK';
    private $unlockString = 'UNLOCK';

    public function lock($filename){

        if(empty($filename))
            return false;

        $state = 1;
        do{
            if(S($filename.$this->postfix) == $this->lockString){
                $state--;
                if($state >= 0) sleep(C('HUBBLE_CACHE_LOCK_WAIT_TIME'));
            }else{
                S($filename.$this->postfix, $this->lockString, C('HUBBLE_CACHE_LOCK_TIME')*60);
                return true;
            }
        }while($state >= 0);

        return false;
    }

    public function unlock($filename){
        S($filename.$this->postfix, $this->unlockString);
        return true;
    }
}
