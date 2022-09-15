package main

// sqtdq is a small program that change a single quote to double quote

// idea of this program, i needed to create a slice of string and the data
// i will be populating it with is data from a python program i wrote some
// time ago, in the program i make use of python list, enclosing each element(string) of the list in
// single quote which means something else in golang(a byte or a rune) not string
// so i decided to automate the boring stuff(not a robot), instead of manually changing
// single quote to double quote which would take tremendous amount of time base on how vast the data is
// so i decided to write this short script so save me time and stress.
// if you found yourself in the same situation you could write yours, in just short lines of code
// or you can just use this and save yourself more time
//

import (
	"fmt"
	"strings"
)

func main() {
	data := `'.txt','.exe','.php','.pl','.7z','.rar','.m4a','.wma',
	'.avi','.wmv','.csv','.d3dbsp','.sc2save','.sie','.sum',
	'.ibank','.t13','.t12','.qdf','.gdb','.tax','.pkpass','.bc6',
	'.bc7','.bkp','.qic','.bkf','.sidn','.sidd','.mddata','.itl',
	'.itdb','.icxs','.hvpl','.hplg','.hkdb','.mdbackup','.syncdb',
	'.gho','.cas','.svg','.map','.wmo','.itm','.sb','.fos','.mcgame',
	'.vdf','.ztmp','.sis','.sid','.ncf','.menu','.layout','.dmp','.blob',
	'.esm','.001','.vtf','.dazip','.fpk','.mlx','.kf','.iwd','.vpk','.tor',
	'.psk','.rim','.w3x','.fsh','.ntl','.arch00','.lvl','.snx','.cfr','.ff',
	'.vpp_pc','.lrf','.m2','.mcmeta','.vfs0','.mpqge','.kdb','.db0','.mp3',
	'.upx','.rofl','.hkx','.bar','.upk','.das','.iwi','.litemod','.asset',
	'.forge','.ltx','.bsa','.apk','.re4','.sav','.lbf','.slm','.bik','.epk',
	'.rgss3a','.pak','.big','.unity3d','.wotreplay','.xxx','.desc','.py',
	'.m3u','.flv','.js','.css','.rb','.png','.jpeg','.p7c','.p7b','.p12',
	'.pfx','.pem','.crt','.cer','.der','.x3f','.srw','.pef','.ptx','.r3d',
	'.rw2','.rwl','.raw','.raf','.orf','.nrw','.mrwref','.mef','.erf','.kdc',
	'.dcr','.cr2','.crw','.bay','.sr2','.srf','.arw','.3fr','.dng','.jpeg',
	'.jpg','.cdr','.indd','.ai','.eps','.pdf','.pdd','.psd','.dbfv','.mdf',
	'.wb2','.rtf','.wpd','.dxg','.xf','.dwg','.pst','.accdb','.mdb','.pptm',
	'.pptx','.ppt','.xlk','.xlsb','.xlsm','.xlsx','.xls','.wps','.docm',
	'.docx','.doc','.odb','.odc','.odm','.odp','.ods','.odt','.sql','.zip',
	'.tar','.tar.gz','.tgz','.biz','.ocx','.html','.htm','.3gp','.srt','.cpp',
	'.mid','.mkv','.mov','.asf','.mpeg','.vob','.mpg','.fla','.swf','.wav',
	'.qcow2','.vdi','.vmdk','.vmx','.gpg','.aes','.ARC','.PAQ','.tar.bz2','.tbk',
	'.bak','.djv','.djvu','.bmp','.cgm','.tif','.tiff','.NEF','.cmd','.class',
	'.jar','.java','.asp','.brd','.sch','.dch','.dip','.vbs','.asm',
	'.pas','.ldf','.ibd','.MYI','.MYD','.frm','.dbf','.SQLITEDB','.SQLITE3',
	'.asc','.lay6','.lay','.ms11(Securitycopy)','.sldm','.sldx','.ppsm',
	'.ppsx','.ppam','.docb','.mml','.sxm','.otg','.slk','.xlw','.xlt','.xlm',
	'.xlc','.dif','.stc','.sxc','.ots','.ods','.hwp','.dotm','.dotx','.docm',
	'.DOT','.max','.xml','.uot','.stw','.sxw','.ott','.csr','.key',
	'wallet.dat', 'pdf',`

	r := TurnToSlice(data)
	fmt.Println(r)
}

func SqtDq(data string) string {
	res := strings.ReplaceAll(data, "'", `"`)
	return res
}

func TurnToSlice(data string) []string {
	dat := SqtDq(data)
	res := strings.Split(dat, " ,")
	return res
}
