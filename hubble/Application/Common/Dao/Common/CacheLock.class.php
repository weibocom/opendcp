<?php
/*
 *  Copyright 2009-2016 Weibo, Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
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
