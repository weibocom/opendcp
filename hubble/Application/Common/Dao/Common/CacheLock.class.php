<?php
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