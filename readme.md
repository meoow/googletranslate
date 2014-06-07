#Usage
##1. On UNIX(Linux, OS X), open terminal:
```
$ ./gt too much shit on the bed
```
>太多狗屎上了床

```
$ ./gt I don\'t want to set the world on fire.
```
>我不想让世界上的火。

####Default is AUTO -> Chinese, but any two language combinations supported by Google Translate can be used：
* _you need to provide two environment variables: **GTF**(source) and **GTT**(target)_

```
# temperary (Chinese -> Japanese)
$ GTF=zh GTT=ja ./gt 只有三喵神才是宇宙真神
```
>わずか3ニャー神は真の神の世界である

```
$ GTF=zh GTT=en ./gt 微软就是只猪队友
```
>Microsoft is pig teammate


```
# exported (Chinese Simplified -> Polish)
$ export GTF=zh-CN
$ export GTT=pl
$ ./gt 三喵神万岁
```  
>Trzy meow Bóg Viva


##2. On Windows, open CMD:
`> gt.exe too much shit on the bed`
>太多狗屎上了床

```
:: use set to assign GTF and GTT
> set gtf=zh
> set gtt=en
> gt.exe 微软就是只猪队友
```
>Microsoft is pig teammate

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

####you could put them in a batch file
```
@echo off
pushd %~dp0
set gtf=zh
set gtt=ja
gt.exe %*
popd
```  
* _保存为bat文件，例如gt.bat，和gt.exe放在相同目录_

```
> gt.bat 天上天下，唯我独尊
```
> Tenjho、傲慢

---
##3. Google Dictionary

When environment variable **GTDICT** exists and not empty, and you are passing only one word, it wil invoke google dictionay。

```
$ GTDICT=1 ./gt fire
```

> NOUN: 火, 炉火, 炉, 炉子, 燧, 烽火, 炬, 煚, 烽, 爓  
> VERB: 解雇, 发射, 射击, 射, 开炮, 发, 开除, 燃起, 施放, 放枪, 点起

解释的顺序是按照其常用度排序。

---
##4. Customize the address for querying Google Translate
Default address is **translate.google.com**, but you can use a specific ip address instead, due to recent restricted access to google services.

```
# assign GTADDR to your own ip
export GTADDR=111.111.111.111
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
