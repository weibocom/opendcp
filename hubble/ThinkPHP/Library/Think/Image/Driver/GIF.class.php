<?php
// +----------------------------------------------------------------------
// | TOPThink [ WE CAN DO IT JUST THINK ]
// +----------------------------------------------------------------------
// | Copyright (c) 2010 http://topthink.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: 麦当苗儿 <zuojiazi@vip.qq.com> <http://www.zjzit.cn>
// +----------------------------------------------------------------------
// | GIF.class.php 2013-03-09
// +----------------------------------------------------------------------
namespace Think\Image\Driver;
class GIF{
	/**
	 * GIF帧列表
	 * @var array
	 */
	private $frames = array();

	/**
	 * 每帧等待时间列表
	 * @var array
	 */
	private $delays = array();

	/**
	 * 构造方法，用于解码GIF图片
	 * @param string $src GIF图片数据
	 * @param string $mod 图片数据类型
	 */
	public function __construct($src = null, $mod = 'url') {
		if(!is_null($src)){
			if('url' == $mod && is_file($src)){
				$src = file_get_contents($src);
			}
			
			/* 解码GIF图片 */
			try{
				$de = new GIFDecoder($src);
				$this->frames = $de->GIFGetFrames();
				$this->delays = $de->GIFGetDelays();
			} catch(\Exception $e){
				E("解码GIF图片出错");
			}
		}
	}

	/**
	 * 设置或获取当前帧的数据
	 * @param  string $stream 二进制数据流
	 * @return boolean        获取到的数据
	 */
	public function image($stream = null){
		if(is_null($stream)){
			$current = current($this->frames);
			return false === $current ? reset($this->frames) : $current;
		} else {
			$this->frames[key($this->frames)] = $stream;
		}
	}

	/**
	 * 将当前帧移动到下一帧
	 * @return string 当前帧数据
	 */
	public function nextImage(){
		return next($this->frames);
	}

	/**
	 * 编码并保存当前GIF图片
	 * @param  string $gifname 图片名称
	 */
	public function save($gifname){
		$gif = new GIFEncoder($this->frames, $this->delays, 0, 2, 0, 0, 0, 'bin');
		file_put_contents($gifname, $gif->GetAnimation());
	}

}


/*
:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
::
::	GIFEncoder Version 2.0 by László Zsidi, http://gifs.hu
::
::	This class is a rewritten 'GifMerge.class.php' version.
::
::  Modification:
::   - Simplified and easy code,
::   - Ultra fast encoding,
::   - Built-in errors,
::   - Stable working
::
::
::	Updated at 2007. 02. 13. '00.05.AM'
::
::
::
::  Try on-line GIFBuilder Form demo based on GIFEncoder.
::
::  http://gifs.hu/phpclasses/demos/GifBuilder/
::
:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
*/

Class GIFEncoder {
	private $GIF = "GIF89a";		/* GIF header 6 bytes	*/
	private $VER = "GIFEncoder V2.05";	/* Encoder version		*/

	private $BUF = Array ( );
	private $LOP =  0;
	private $DIS =  2;
	private $COL = -1;
	private $IMG = -1;

	private $ERR = Array (
		'ERR00'	=>	"Does not supported function for only one image!",
		'ERR01'	=>	"Source is not a GIF image!",
		'ERR02'	=>	"Unintelligible flag ",
		'ERR03'	=>	"Does not make animation from animated GIF source",
	);

	/*
	:::::::::::::::::::::::::::::::::::::::::::::::::::
	::
	::	GIFEncoder...
	::
	*/
	public function __construct($GIF_src, $GIF_dly, $GIF_lop, $GIF_dis,$GIF_red, $GIF_grn, $GIF_blu, $GIF_mod) {
		if ( ! is_array ( $GIF_src ) && ! is_array ( $GIF_dly ) ) {
			printf	( "%s: %s", $this->VER, $this->ERR [ 'ERR00' ] );
			exit	( 0 );
		}
		$this->LOP = ( $GIF_lop > -1 ) ? $GIF_lop : 0;
		$this->DIS = ( $GIF_dis > -1 ) ? ( ( $GIF_dis < 3 ) ? $GIF_dis : 3 ) : 2;
		$this->COL = ( $GIF_red > -1 && $GIF_grn > -1 && $GIF_blu > -1 ) ?
						( $GIF_red | ( $GIF_grn << 8 ) | ( $GIF_blu << 16 ) ) : -1;

		for ( $i = 0; $i < count ( $GIF_src ); $i++ ) {
			if ( strToLower ( $GIF_mod ) == "url" ) {
				$this->BUF [ ] = fread ( fopen ( $GIF_src [ $i ], "rb" ), filesize ( $GIF_src [ $i ] ) );
			}
			else if ( strToLower ( $GIF_mod ) == "bin" ) {
				$this->BUF [ ] = $GIF_src [ $i ];
			}
			else {
				printf	( "%s: %s ( %s )!", $this->VER, $this->ERR [ 'ERR02' ], $GIF_mod );
				exit	( 0 );
			}
			if ( substr ( $this->BUF [ $i ], 0, 6 ) != "GIF87a" && substr ( $this->BUF [ $i ], 0, 6 ) != "GIF89a" ) {
				printf	( "%s: %d %s", $this->VER, $i, $this->ERR [ 'ERR01' ] );
				exit	( 0 );
			}
			for ( $j = ( 13 + 3 * ( 2 << ( ord ( $this->BUF [ $i ] { 10 } ) & 0x07 ) ) ), $k = TRUE; $k; $j++ ) {
				switch ( $this->BUF [ $i ] { $j } ) {
					case "!":
						if ( ( substr ( $this->BUF [ $i ], ( $j + 3 ), 8 ) ) == "NETSCAPE" ) {
							printf	( "%s: %s ( %s source )!", $this->VER, $this->ERR [ 'ERR03' ], ( $i + 1 ) );
							exit	( 0 );
						}
						break;
					case ";":
						$k = FALSE;
						break;
				}
			}
		}
		$this->GIFAddHeader ( );
		for ( $i = 0; $i < count ( $this->BUF ); $i++ ) {
			$this->GIFAddFrames ( $i, $GIF_dly [ $i ] );
		}
		$this->GIFAddFooter ( );
	}
	/*
	:::::::::::::::::::::::::::::::::::::::::::::::::::
	::
	::	GIFAddHeader...
	::
	*/
	private function GIFAddHeader ( ) {
		$cmap = 0;

		if ( ord ( $this->BUF [ 0 ] { 10 } ) & 0x80 ) {
			$cmap = 3 * ( 2 << ( ord ( $this->BUF [ 0 ] { 10 } ) & 0x07 ) );

			$this->GIF .= substr ( $this->BUF [ 0 ], 6, 7		);
			$this->GIF .= substr ( $this->BUF [ 0 ], 13, $cmap	);
			$this->GIF .= "!\377\13NETSCAPE2.0\3\1" . $this->GIFWord ( $this->LOP ) . "\0";
		}
	}
	/*
	:::::::::::::::::::::::::::::::::::::::::::::::::::
	::
	::	GIFAddFrames...
	::
	*/
	private function GIFAddFrames ( $i, $d ) {

		$Locals_str = 13 + 3 * ( 2 << ( ord ( $this->BUF [ $i ] { 10 } ) & 0x07 ) );

		$Locals_end = strlen ( $this->BUF [ $i ] ) - $Locals_str - 1;
		$Locals_tmp = substr ( $this->BUF [ $i ], $Locals_str, $Locals_end );

		$Global_len = 2 << ( ord ( $this->BUF [ 0  ] { 10 } ) & 0x07 );
		$Locals_len = 2 << ( ord ( $this->BUF [ $i ] { 10 } ) & 0x07 );

		$Global_rgb = substr ( $this->BUF [ 0  ], 13,
							3 * ( 2 << ( ord ( $this->BUF [ 0  ] { 10 } ) & 0x07 ) ) );
		$Locals_rgb = substr ( $this->BUF [ $i ], 13,
							3 * ( 2 << ( ord ( $this->BUF [ $i ] { 10 } ) & 0x07 ) ) );

		$Locals_ext = "!\xF9\x04" . chr ( ( $this->DIS << 2 ) + 0 ) .
						chr ( ( $d >> 0 ) & 0xFF ) . chr ( ( $d >> 8 ) & 0xFF ) . "\x0\x0";

		if ( $this->COL > -1 && ord ( $this->BUF [ $i ] { 10 } ) & 0x80 ) {
			for ( $j = 0; $j < ( 2 << ( ord ( $this->BUF [ $i ] { 10 } ) & 0x07 ) ); $j++ ) {
				if	(
						ord ( $Locals_rgb { 3 * $j + 0 } ) == ( ( $this->COL >> 16 ) & 0xFF ) &&
						ord ( $Locals_rgb { 3 * $j + 1 } ) == ( ( $this->COL >>  8 ) & 0xFF ) &&
						ord ( $Locals_rgb { 3 * $j + 2 } ) == ( ( $this->COL >>  0 ) & 0xFF )
					) {
					$Locals_ext = "!\xF9\x04" . chr ( ( $this->DIS << 2 ) + 1 ) .
									chr ( ( $d >> 0 ) & 0xFF ) . chr ( ( $d >> 8 ) & 0xFF ) . chr ( $j ) . "\x0";
					break;
				}
			}
		}
		switch ( $Locals_tmp { 0 } ) {
			case "!":
				$Locals_img = substr ( $Locals_tmp, 8, 10 );
				$Locals_tmp = substr ( $Locals_tmp, 18, strlen ( $Locals_tmp ) - 18 );
				break;
			case ",":
				$Locals_img = substr ( $Locals_tmp, 0, 10 );
				$Locals_tmp = substr ( $Locals_tmp, 10, strlen ( $Locals_tmp ) - 10 );
				break;
		}
		if ( ord ( $this->BUF [ $i ] { 10 } ) & 0x80 && $this->IMG > -1 ) {
			if ( $Global_len == $Locals_len ) {
				if ( $this->GIFBlockCompare ( $Global_rgb, $Locals_rgb, $Global_len ) ) {
					$this->GIF .= ( $Locals_ext . $Locals_img . $Locals_tmp );
				}
				else {
					$byte  = ord ( $Locals_img { 9 } );
					$byte |= 0x80;
					$byte &= 0xF8;
					$byte |= ( ord ( $this->BUF [ 0 ] { 10 } ) & 0x07 );
					$Locals_img { 9 } = chr ( $byte );
					$this->GIF .= ( $Locals_ext . $Locals_img . $Locals_rgb . $Locals_tmp );
				}
			}
			else {
				$byte  = ord ( $Locals_img { 9 } );
				$byte |= 0x80;
				$byte &= 0xF8;
				$byte |= ( ord ( $this->BUF [ $i ] { 10 } ) & 0x07 );
				$Locals_img { 9 } = chr ( $byte );
				$this->GIF .= ( $Locals_ext . $Locals_img . $Locals_rgb . $Locals_tmp );
			}
		}
		else {
			$this->GIF .= ( $Locals_ext . $Locals_img . $Locals_tmp );
		}
		$this->IMG  = 1;
	}
	/*
	:::::::::::::::::::::::::::::::::::::::::::::::::::
	::
	::	GIFAddFooter...
	::
	*/
	private function GIFAddFooter ( ) {
		$this->GIF .= ";";
	}
	/*
	:::::::::::::::::::::::::::::::::::::::::::::::::::
	::
	::	GIFBlockCompare...
	::
	*/
	private function GIFBlockCompare ( $GlobalBlock, $LocalBlock, $Len ) {

		for ( $i = 0; $i < $Len; $i++ ) {
			if	(
					$GlobalBlock { 3 * $i + 0 } != $LocalBlock { 3 * $i + 0 } ||
					$GlobalBlock { 3 * $i + 1 } != $LocalBlock { 3 * $i + 1 } ||
					$GlobalBlock { 3 * $i + 2 } != $LocalBlock { 3 * $i + 2 }
				) {
					return ( 0 );
			}
		}

		return ( 1 );
	}
	/*
	:::::::::::::::::::::::::::::::::::::::::::::::::::
	::
	::	GIFWord...
	::
	*/
	private function GIFWord ( $int ) {

		return ( chr ( $int & 0xFF ) . chr ( ( $int >> 8 ) & 0xFF ) );
	}
	/*
	:::::::::::::::::::::::::::::::::::::::::::::::::::
	::
	::	GetAnimation...
	::
	*/
	public function GetAnimation ( ) {
		return ( $this->GIF );
	}
}


/*
:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
::
::	GIFDecoder Version 2.0 by László Zsidi, http://gifs.hu
::
::	Created at 2007. 02. 01. '07.47.AM'
::
::
::
::
::  Try on-line GIFBuilder Form demo based on GIFDecoder.
::
::  http://gifs.hu/phpclasses/demos/GifBuilder/
::
:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
*/

Class GIFDecoder {
	private $GIF_buffer = Array ( );
	private $GIF_arrays = Array ( );
	private $GIF_delays = Array ( );
	private $GIF_stream = "";
	private $GIF_string = "";
	private $GIF_bfseek =  0;

	private $GIF_screen = Array ( );
	private $GIF_global = Array ( );
	private $GIF_sorted;
	private $GIF_colorS;
	private $GIF_colorC;
	private $GIF_colorF;

	/*
	:::::::::::::::::::::::::::::::::::::::::::::::::::
	::
	::	GIFDecoder ( $GIF_pointer )
	::
	*/
	public function __construct ( $GIF_pointer ) {
		$this->GIF_stream = $GIF_pointer;

		$this->GIFGetByte ( 6 );	// GIF89a
		$this->GIFGetByte ( 7 );	// Logical Screen Descriptor

		$this->GIF_screen = $this->GIF_buffer;
		$this->GIF_colorF = $this->GIF_buffer [ 4 ] & 0x80 ? 1 : 0;
		$this->GIF_sorted = $this->GIF_buffer [ 4 ] & 0x08 ? 1 : 0;
		$this->GIF_colorC = $this->GIF_buffer [ 4 ] & 0x07;
		$this->GIF_colorS = 2 << $this->GIF_colorC;

		if ( $this->GIF_colorF == 1 ) {
			$this->GIFGetByte ( 3 * $this->GIF_colorS );
			$this->GIF_global = $this->GIF_buffer;
		}
		/*
		 *
		 *  05.06.2007.
		 *  Made a little modification
		 *
		 *
		 -	for ( $cycle = 1; $cycle; ) {
		 +		if ( GIFDecoder::GIFGetByte ( 1 ) ) {
		 -			switch ( $this->GIF_buffer [ 0 ] ) {
		 -				case 0x21:
		 -					GIFDecoder::GIFReadExtensions ( );
		 -					break;
		 -				case 0x2C:
		 -					GIFDecoder::GIFReadDescriptor ( );
		 -					break;
		 -				case 0x3B:
		 -					$cycle = 0;
		 -					break;
		 -		  	}
		 -		}
		 +		else {
		 +			$cycle = 0;
		 +		}
		 -	}
		*/
		for ( $cycle = 1; $cycle; ) {
			if ( $this->GIFGetByte ( 1 ) ) {
				switch ( $this->GIF_buffer [ 0 ] ) {
					case 0x21:
						$this->GIFReadExtensions ( );
						break;
					case 0x2C:
						$this->GIFReadDescriptor ( );
						break;
					case 0x3B:
						$cycle = 0;
						break;
				}
			}
			else {
				$cycle = 0;
			}
		}
	}
	/*
	:::::::::::::::::::::::::::::::::::::::::::::::::::
	::
	::	GIFReadExtension ( )
	::
	*/
	private function GIFReadExtensions ( ) {
		$this->GIFGetByte ( 1 );
		for ( ; ; ) {
			$this->GIFGetByte ( 1 );
			if ( ( $u = $this->GIF_buffer [ 0 ] ) == 0x00 ) {
				break;
			}
			$this->GIFGetByte ( $u );
			/*
			 * 07.05.2007.
			 * Implemented a new line for a new function
			 * to determine the originaly delays between
			 * frames.
			 *
			 */
			if ( $u == 4 ) {
				$this->GIF_delays [ ] = ( $this->GIF_buffer [ 1 ] | $this->GIF_buffer [ 2 ] << 8 );
			}
		}
	}
	/*
	:::::::::::::::::::::::::::::::::::::::::::::::::::
	::
	::	GIFReadExtension ( )
	::
	*/
	private function GIFReadDescriptor ( ) {
		$GIF_screen	= Array ( );

		$this->GIFGetByte ( 9 );
		$GIF_screen = $this->GIF_buffer;
		$GIF_colorF = $this->GIF_buffer [ 8 ] & 0x80 ? 1 : 0;
		if ( $GIF_colorF ) {
			$GIF_code = $this->GIF_buffer [ 8 ] & 0x07;
			$GIF_sort = $this->GIF_buffer [ 8 ] & 0x20 ? 1 : 0;
		}
		else {
			$GIF_code = $this->GIF_colorC;
			$GIF_sort = $this->GIF_sorted;
		}
		$GIF_size = 2 << $GIF_code;
		$this->GIF_screen [ 4 ] &= 0x70;
		$this->GIF_screen [ 4 ] |= 0x80;
		$this->GIF_screen [ 4 ] |= $GIF_code;
		if ( $GIF_sort ) {
			$this->GIF_screen [ 4 ] |= 0x08;
		}
		$this->GIF_string = "GIF87a";
		$this->GIFPutByte ( $this->GIF_screen );
		if ( $GIF_colorF == 1 ) {
			$this->GIFGetByte ( 3 * $GIF_size );
			$this->GIFPutByte ( $this->GIF_buffer );
		}
		else {
			$this->GIFPutByte ( $this->GIF_global );
		}
		$this->GIF_string .= chr ( 0x2C );
		$GIF_screen [ 8 ] &= 0x40;
		$this->GIFPutByte ( $GIF_screen );
		$this->GIFGetByte ( 1 );
		$this->GIFPutByte ( $this->GIF_buffer );
		for ( ; ; ) {
			$this->GIFGetByte ( 1 );
			$this->GIFPutByte ( $this->GIF_buffer );
			if ( ( $u = $this->GIF_buffer [ 0 ] ) == 0x00 ) {
				break;
			}
			$this->GIFGetByte ( $u );
			$this->GIFPutByte ( $this->GIF_buffer );
		}
		$this->GIF_string .= chr ( 0x3B );
		/*
		   Add frames into $GIF_stream array...
		*/
		$this->GIF_arrays [ ] = $this->GIF_string;
	}
	/*
	:::::::::::::::::::::::::::::::::::::::::::::::::::
	::
	::	GIFGetByte ( $len )
	::
	*/

	/*
	 *
	 *  05.06.2007.
	 *  Made a little modification
	 *
	 *
	 -	function GIFGetByte ( $len ) {
	 -		$this->GIF_buffer = Array ( );
	 -
	 -		for ( $i = 0; $i < $len; $i++ ) {
	 +			if ( $this->GIF_bfseek > strlen ( $this->GIF_stream ) ) {
	 +				return 0;
	 +			}
	 -			$this->GIF_buffer [ ] = ord ( $this->GIF_stream { $this->GIF_bfseek++ } );
	 -		}
	 +		return 1;
	 -	}
	 */
	private function GIFGetByte ( $len ) {
		$this->GIF_buffer = Array ( );

		for ( $i = 0; $i < $len; $i++ ) {
			if ( $this->GIF_bfseek > strlen ( $this->GIF_stream ) ) {
				return 0;
			}
			$this->GIF_buffer [ ] = ord ( $this->GIF_stream { $this->GIF_bfseek++ } );
		}
		return 1;
	}
	/*
	:::::::::::::::::::::::::::::::::::::::::::::::::::
	::
	::	GIFPutByte ( $bytes )
	::
	*/
	private function GIFPutByte ( $bytes ) {
		for ( $i = 0; $i < count ( $bytes ); $i++ ) {
			$this->GIF_string .= chr ( $bytes [ $i ] );
		}
	}
	/*
	:::::::::::::::::::::::::::::::::::::::::::::::::::
	::
	::	PUBLIC FUNCTIONS
	::
	::
	::	GIFGetFrames ( )
	::
	*/
	public function GIFGetFrames ( ) {
		return ( $this->GIF_arrays );
	}
	/*
	:::::::::::::::::::::::::::::::::::::::::::::::::::
	::
	::	GIFGetDelays ( )
	::
	*/
	public function GIFGetDelays ( ) {
		return ( $this->GIF_delays );
	}
}
