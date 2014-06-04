#使用说明


##使用方法
##1. 在类UNIX系统下(Linux, Mac OSX, BSD)，打开终端:
```
$ ./gt too much shit on the bed
```
>太多狗屎上了床

```
$ ./gt I don\'t want to set the world on fire.
```
>我不想让世界上的火。

####支持其他任意两种语言翻译：
* _需要给定环境变量**GTF**(源语言)和**GTT**(目标语言)_  
* _例如: GTF=zh-CN (zh-CN是汉语简体所写代码)_

```
$ GTF=zh GTT=ja ./gt 只有三喵神才是宇宙真神
```
>わずか3ニャー神は真の神の世界である

```
$ GTF=zh GTT=en ./gt 微软就是只猪队友
```
>Microsoft is pig teammate

####环境变量也可以单独设置，该设置会一直在窗口关闭前一直生效，而前一种方法只单次命令有效:
```
$ export GTF=zh-CN
$ export GTT=pl
$ ./gt 三喵神万岁
```  
>Trzy meow Bóg Viva


##2. 在windows系统下，打开命令提示符:
`> gt.exe too much shit on the bed`
>太多狗屎上了床

####汉译英(windows的cmd不支持临时环境变量，只能预先单独设置)
```
> set gtf=zh
> set gtt=en
> gt.exe 微软就是只猪队友
```
>Microsoft is pig teammate

####汉译日
```
> set gtf=zh
> set gtt=ja
> gt.exe 只有三喵神才是宇宙真神
```
>わずか3ニャー神は真の神の世界である

```
> gt.exe 吃瘪
```
>打たれた

####由于windows单独设置单独设置变量比较麻烦，可以使用批处理
```
@echo off
pushd %~dp0
set gtf=zh
set gtt=ja
gt.exe %*
popd
```  
保存为bat文件，例如gt.bat，和gt.exe放在相同目录

```
> gt.bat 天上天下，唯我独尊
```
> Tenjho、傲慢

---
##3. 使用Google词典
如果GTDICT环境变量存在并且是非空值，并且只提供一个单词的情况下，会使用谷歌词典，而非句子翻译。

```
$ GTDICT=1 ./gt fire
```

> NOUN: 火, 炉火, 炉, 炉子, 燧, 烽火, 炬, 煚, 烽, 爓  
> VERB: 解雇, 发射, 射击, 射, 开炮, 发, 开除, 燃起, 施放, 放枪, 点起

解释的顺序是按照其常用度排序。

---
##4. 支持自定义Google Translate的地址
默认是translate.google.com
但是鉴于最近GƒW的问题，所以可以通过环境变量GTADDR自定义地址（IP）

```
export GTADDR=111.111.111.111 #自定义的地址
got I am the God of War
我是战神

```
具体用什么地址，可以自行搜索通过HOSTS文件访问谷歌服务的相关内容，里面的IP地址可以用在这里

##多国语言缩写代码： 


|Code      | Language |
|:---------|:--------:|
|af          |Afrikaans|
|ak          |Akan|
|sq          |Albanian|
|am          |Amharic|
|ar          |Arabic|
|hy          |Armenian|
|az          |Azerbaijani|
|eu          |Basque|
|be          |Belarusian|
|bem         |Bemba|
|bn          |Bengali|
|bh          |Bihari|
|xx-bork     |Bork, bork, bork!|
|bs          |Bosnian|
|br          |Breton|
|bg          |Bulgarian|
|km          |Cambodian|
|ca          |Catalan|
|chr         |Cherokee|
|ny          |Chichewa|
|**zh-CN**   |**Chinese (Simplified)**|
|zh-TW       |Chinese (Traditional)|
|co          |Corsican|
|hr          |Croatian|
|cs          |Czech|
|da          |Danish|
|nl          |Dutch|
|xx-elmer    |Elmer Fudd|
|**en**      |**English**|
|eo          |Esperanto|
|et          |Estonian|
|ee          |Ewe|
|fo          |Faroese|
|tl          |Filipino|
|fi          |Finnish|
|fr          |French|
|fy          |Frisian|
|gaa         |Ga|
|gl          |Galician|
|ka          |Georgian|
|de          |German|
|el          |Greek|
|gn          |Guarani|
|gu          |Gujarati|
|xx-hacker   |Hacker|
|ht          |Haitian Creole|
|ha          |Hausa|
|haw         |Hawaiian|
|iw          |Hebrew|
|hi          |Hindi|
|hu          |Hungarian|
|is          |Icelandic|
|ig          |Igbo|
|id          |Indonesian|
|ia          |Interlingua|
|ga          |Irish|
|it          |Italian|
|ja          |Japanese|
|jw          |Javanese|
|kn          |Kannada|
|kk          |Kazakh|
|rw          |Kinyarwanda|
|rn          |Kirundi|
|xx-klingon  |Klingon|
|kg          |Kongo|
|ko          |Korean|
|kri         |Krio (Sierra Leone)|
|ku          |Kurdish|
|ckb         |Kurdish (Soranî)|
|ky          |Kyrgyz|
|lo          |Laothian|
|la          |Latin|
|lv          |Latvian|
|ln          |Lingala|
|lt          |Lithuanian|
|loz         |Lozi|
|lg          |Luganda|
|ach         |Luo|
|mk          |Macedonian|
|mg          |Malagasy|
|ms          |Malay|
|ml          |Malayalam|
|mt          |Maltese|
|mi          |Maori|
|mr          |Marathi|
|mfe         |Mauritian Creole|
|mo          |Moldavian|
|mn          |Mongolian|
|sr-ME       |Montenegrin|
|ne          |Nepali|
|pcm         |Nigerian Pidgin|
|nso         |Northern Sotho|
|no          |Norwegian|
|nn          |Norwegian (Nynorsk)|
|oc          |Occitan|
|or          |Oriya|
|om          |Oromo|
|ps          |Pashto|
|fa          |Persian|
|xx-pirate   |Pirate|
|pl          |Polish|
|pt-BR       |Portuguese (Brazil)|
|pt-PT       |Portuguese (Portugal)|
|pa          |Punjabi|
|qu          |Quechua|
|ro          |Romanian|
|rm          |Romansh|
|nyn         |Runyakitara|
|ru          |Russian|
|gd          |Scots Gaelic|
|sr          |Serbian|
|sh          |Serbo-Croatian|
|st          |Sesotho|
|tn          |Setswana|
|crs         |Seychellois Creole|
|sn          |Shona|
|sd          |Sindhi|
|si          |Sinhalese|
|sk          |Slovak|
|sl          |Slovenian|
|so          |Somali|
|es          |Spanish|
|es-419      |Spanish (Latin American)|
|su          |Sundanese|
|sw          |Swahili|
|sv          |Swedish|
|tg          |Tajik|
|ta          |Tamil|
|tt          |Tatar|
|te          |Telugu|
|th          |Thai|
|ti          |Tigrinya|
|to          |Tonga|
|lua         |Tshiluba|
|tum         |Tumbuka|
|tr          |Turkish|
|tk          |Turkmen|
|tw          |Twi|
|ug          |Uighur|
|uk          |Ukrainian|
|ur          |Urdu|
|uz          |Uzbek|
|vi          |Vietnamese|
|cy          |Welsh|
|wo          |Wolof|
|xh          |Xhosa|
|yi          |Yiddish|
|yo          |Yoruba|
|zu          |Zul|
